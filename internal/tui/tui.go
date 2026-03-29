package tui

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	tea "charm.land/bubbletea/v2"
	"github.com/TyostoKarry/sleepycli/internal/render"
	"github.com/TyostoKarry/sleepycli/internal/styles"
	"github.com/TyostoKarry/sleepycli/internal/validate"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle         = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("75"))
	selectedStyle      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
	dimStyle           = styles.Dim
	errorStyle         = styles.Error
	labelStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	sectionHeaderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("67"))
	cursorStyle        = lipgloss.NewStyle().Reverse(true)
	inputStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	placeholderStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))
)

type inputKind int

const (
	kindTime inputKind = iota
	kindNum
)

type textInput struct {
	value       string
	placeholder string
	editing     bool
	maxLength   int
	kind        inputKind
}

func newTimeInput(placeholder string) textInput {
	return textInput{placeholder: placeholder, maxLength: 5, kind: kindTime}
}

func newNumInput(placeholder string) textInput {
	return textInput{placeholder: placeholder, maxLength: 3, kind: kindNum}
}

func (t *textInput) handleKey(msg tea.KeyPressMsg) {
	key := msg.String()
	switch key {
	case "backspace":
		if len(t.value) > 0 {
			r := []rune(t.value)
			t.value = string(r[:len(r)-1])
		}
	default:
		if len(key) != 1 {
			return
		}
		ch := rune(key[0])
		if !unicode.IsDigit(ch) {
			return
		}
		if len([]rune(t.value)) >= t.maxLength {
			return
		}
		t.value += string(ch)
		if t.kind == kindTime && len([]rune(t.value)) == 2 {
			t.value += ":"
		}
	}
}

func (t textInput) view() string {
	if !t.editing && t.value == "" {
		return placeholderStyle.Render(t.placeholder)
	}
	display := inputStyle.Render(t.value)
	if t.editing {
		display += cursorStyle.Render(" ")
	}
	return display
}

// Menu items
const (
	rowNow      = 0
	rowWake     = 1
	rowSleep    = 2
	rowWindow   = 3
	rowBuffer   = 4
	rowCycleMin = 5
	rowCycleMax = 6
	menuRows    = 7
)

type Model struct {
	// menuCursor is the highlighted row (0–6)
	menuCursor int
	// selectedMode is the active sleep mode (rows 0–3)
	selectedMode int
	// editing is true when the cursor is on a settings/time row and the user pressed enter to type into it
	editing bool

	// time inputs. Only used when a mode that needs input is selected
	inputPrimary   textInput // wake / sleep / from
	inputSecondary textInput // to (window only)
	// settings inputs
	inputBuffer    textInput
	inputCyclesMin textInput
	inputCyclesMax textInput

	// timeField tracks which time input is active (0 = primary, 1 = secondary)
	timeField int

	// PrintResult holds the result to print to stdout after the TUI exits (empty = don't print)
	PrintResult string
}

func InitialModel() Model {
	return Model{
		menuCursor:     rowNow,
		selectedMode:   rowNow,
		inputPrimary:   newTimeInput("HH:MM"),
		inputSecondary: newTimeInput("HH:MM"),
		inputBuffer:    newNumInput("15"),
		inputCyclesMin: newNumInput("4"),
		inputCyclesMax: newNumInput("6"),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m.handleKey(msg)
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyPressMsg) (Model, tea.Cmd) {
	key := msg.String()

	if key == "ctrl+c" || key == "q" {
		return m, tea.Quit
	}

	// Editing time input
	if m.editing && m.menuCursor < rowBuffer {
		switch key {
		case "esc":
			m.editing = false
			m.inputPrimary.editing = false
			m.inputSecondary.editing = false
		case "tab", "enter":
			if m.selectedMode == rowWindow && m.timeField == 0 {
				m.inputPrimary.editing = false
				m.timeField = 1
				m.inputSecondary.editing = true
			} else {
				m.editing = false
				m.inputPrimary.editing = false
				m.inputSecondary.editing = false
				m.timeField = 0
				hasInput := m.inputPrimary.value != ""
				if m.selectedMode == rowWindow {
					hasInput = hasInput && m.inputSecondary.value != ""
				}
				if hasInput {
					m.PrintResult = m.computeResult()
					return m, tea.Quit
				}
			}
		default:
			if m.timeField == 1 {
				m.inputSecondary.handleKey(msg)
			} else {
				m.inputPrimary.handleKey(msg)
			}
		}
		return m, nil
	}

	// Editing settings input
	if m.editing && m.menuCursor >= rowBuffer {
		switch key {
		case "esc", "enter":
			m.editing = false
			activeSettingsInput(&m).editing = false
		default:
			activeSettingsInput(&m).handleKey(msg)
		}
		return m, nil
	}

	// Menu navigation
	switch key {
	case "up":
		m.menuCursor--
		if m.menuCursor < 0 {
			m.menuCursor = menuRows - 1
		}
	case "down":
		m.menuCursor++
		if m.menuCursor >= menuRows {
			m.menuCursor = 0
		}
	case "enter":
		if m.menuCursor <= rowWindow {
			if m.menuCursor == rowNow {
				m.selectedMode = rowNow
				m.PrintResult = m.computeResult()
				return m, tea.Quit
			}
			// Mode selection
			if m.selectedMode != m.menuCursor {
				m.selectedMode = m.menuCursor
				// reset inputs when switching modes
				m.inputPrimary.value = ""
				m.inputSecondary.value = ""
				m.timeField = 0
			}
			m.editing = true
			m.inputPrimary.editing = true
		} else {
			// Settings editing
			m.editing = true
			activeSettingsInput(&m).editing = true
		}
	}

	return m, nil
}

func activeSettingsInput(m *Model) *textInput {
	switch m.menuCursor {
	case rowBuffer:
		return &m.inputBuffer
	case rowCycleMin:
		return &m.inputCyclesMin
	case rowCycleMax:
		return &m.inputCyclesMax
	default:
		return nil
	}
}

func (m Model) bufferMinutes() int {
	if m.inputBuffer.value == "" {
		return 15
	}
	var n int
	fmt.Sscanf(m.inputBuffer.value, "%d", &n)
	return n
}

func (m Model) cyclesMin() int {
	if m.inputCyclesMin.value == "" {
		return 4
	}
	var n int
	fmt.Sscanf(m.inputCyclesMin.value, "%d", &n)
	return n
}

func (m Model) cyclesMax() int {
	if m.inputCyclesMax.value == "" {
		return 6
	}
	var n int
	fmt.Sscanf(m.inputCyclesMax.value, "%d", &n)
	return n
}

func (m Model) computeResult() string {
	buffer := time.Duration(m.bufferMinutes()) * time.Minute
	minCycles := m.cyclesMin()
	maxCycles := m.cyclesMax()

	if minCycles < 0 || maxCycles < 0 || maxCycles < minCycles || buffer < 0 {
		return errorStyle.Render("Invalid Settings")
	}

	switch m.selectedMode {
	case rowNow:
		now := time.Now()
		return render.WakeTimes(now, buffer, minCycles, maxCycles,
			fmt.Sprintf("Sleeping now at %s", now.Format("15:04")))

	case rowWake:
		raw := m.inputPrimary.value
		if raw == "" {
			return dimStyle.Render("Enter a wake time to see results")
		}
		wakeTime, err := time.Parse("15:04", validate.NormalizeHour(raw))
		if err != nil {
			return errorStyle.Render("Invalid time, use HH:MM")
		}
		return render.Bedtimes(wakeTime, buffer, minCycles, maxCycles,
			fmt.Sprintf("To wake up at %s", raw))

	case rowSleep:
		raw := m.inputPrimary.value
		if raw == "" {
			return dimStyle.Render("Enter a sleep time to see results")
		}
		sleepTime, err := time.Parse("15:04", validate.NormalizeHour(raw))
		if err != nil {
			return errorStyle.Render("Invalid time, use HH:MM")
		}
		return render.WakeTimes(sleepTime, buffer, minCycles, maxCycles,
			fmt.Sprintf("Sleeping at %s", raw))

	case rowWindow:
		rawFrom := m.inputPrimary.value
		rawTo := m.inputSecondary.value
		if rawFrom == "" || rawTo == "" {
			return dimStyle.Render("Enter both sleep and wake times to see results")
		}
		fromTime, err := time.Parse("15:04", validate.NormalizeHour(rawFrom))
		if err != nil {
			return errorStyle.Render("from: Invalid time, use HH:MM")
		}
		toTime, err := time.Parse("15:04", validate.NormalizeHour(rawTo))
		if err != nil {
			return errorStyle.Render("to: Invalid time, use HH:MM")
		}
		return render.Window(rawFrom, rawTo, fromTime, toTime, m.bufferMinutes())
	}
	return ""
}

func (m Model) View() tea.View {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("SleepyCLI") + "\n\n")

	// Mode rows (0-3)
	modeLabels := []string{"now", "wake", "sleep", "from/to"}
	modeDescriptions := []string{
		"wake times starting from right now",
		"bedtimes for a target wake time",
		"wake times for a target sleep time",
		"cycles that fit in a sleep window",
	}
	for i, label := range modeLabels {
		cursor := "  "
		if m.menuCursor == i {
			cursor = selectedStyle.Render("▶ ")
		}
		text := fmt.Sprintf("%-8s", label)
		desc := " " + modeDescriptions[i]
		if m.selectedMode == i {
			sb.WriteString(cursor + selectedStyle.Render(text) + dimStyle.Render(desc) + "\n")
		} else {
			sb.WriteString(cursor + dimStyle.Render(text+desc) + "\n")
		}
	}

	// Blank separator and section header between modes and settings
	sb.WriteString("\n")
	sb.WriteString(sectionHeaderStyle.Render("  settings") + "\n")

	// Settings rows (4–6)
	settingsRows := []struct {
		label string
		input textInput
		row   int
	}{
		{"  buffer (min):  ", m.inputBuffer, rowBuffer},
		{"  cycles min:    ", m.inputCyclesMin, rowCycleMin},
		{"  cycles max:    ", m.inputCyclesMax, rowCycleMax},
	}
	for _, s := range settingsRows {
		cursor := "  "
		label := dimStyle.Render(s.label)
		if m.menuCursor == s.row {
			cursor = selectedStyle.Render("▶ ")
			label = selectedStyle.Render(s.label)
		}
		sb.WriteString(cursor + label + s.input.view() + "\n")
	}

	// Time input area
	if m.selectedMode != rowNow {
		sb.WriteString("\n")
		switch m.selectedMode {
		case rowWake:
			sb.WriteString(labelStyle.Render("  Wake time:   ") + m.inputPrimary.view() + "\n")
		case rowSleep:
			sb.WriteString(labelStyle.Render("  Sleep time:  ") + m.inputPrimary.view() + "\n")
		case rowWindow:
			sb.WriteString(labelStyle.Render("  From:  ") + m.inputPrimary.view() + "\n")
			sb.WriteString(labelStyle.Render("  To:    ") + m.inputSecondary.view() + "\n")
		}
	}

	// Result
	sb.WriteString("\n")
	sb.WriteString(m.computeResult())
	sb.WriteString("\n\n")

	// Help bar
	var helpItems []string
	if m.editing {
		helpItems = []string{"type value", "enter confirm", "esc cancel"}
	} else {
		helpItems = []string{"↑↓ navigate", "enter select/edit", "q quit"}
	}
	sb.WriteString(dimStyle.Render(strings.Join(helpItems, "  ·  ")))

	v := tea.NewView(sb.String())
	v.AltScreen = true
	return v
}
