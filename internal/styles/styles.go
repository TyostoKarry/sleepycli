package styles

import "github.com/charmbracelet/lipgloss"

var (
	Result    = lipgloss.NewStyle().Foreground(lipgloss.Color("120"))
	Dim       = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	Separator = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	Error     = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
)
