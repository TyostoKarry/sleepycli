package render

import (
	"fmt"
	"strings"
	"time"

	"github.com/TyostoKarry/sleepycli/internal/cycle"
	"github.com/TyostoKarry/sleepycli/internal/styles"
)

func FormatDuration(cycleCount int) string {
	duration := time.Duration(cycleCount) * cycle.CycleDuration
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return fmt.Sprintf("%dh %02dm", hours, minutes)
}

func WakeTimes(base time.Time, buffer time.Duration, minCycles, maxCycles int, header string) string {
	wakeTimes := cycle.CalculateWakeTimes(base, buffer, minCycles, maxCycles)
	var sb strings.Builder
	sb.WriteString(styles.Result.Render(header) + "\n")
	sb.WriteString(styles.Separator.Render(strings.Repeat("─", 30)) + "\n")
	sb.WriteString(styles.Dim.Render(fmt.Sprintf("Assuming %d min to fall asleep", int(buffer.Minutes()))) + "\n\n")
	for i := len(wakeTimes) - 1; i >= 0; i-- {
		c := minCycles + i
		fmt.Fprintf(&sb, "  %s cycles  →  wake at %s  %s\n",
			styles.Result.Render(fmt.Sprintf("%d", c)),
			styles.Result.Render(wakeTimes[i].Format("15:04")),
			styles.Dim.Render("("+FormatDuration(c)+")"))
	}
	return sb.String()
}

func Bedtimes(base time.Time, buffer time.Duration, minCycles, maxCycles int, header string) string {
	bedTimes := cycle.CalculateBedtimes(base, buffer, minCycles, maxCycles)
	var sb strings.Builder
	sb.WriteString(styles.Result.Render(header) + "\n")
	sb.WriteString(styles.Separator.Render(strings.Repeat("─", 30)) + "\n")
	sb.WriteString(styles.Dim.Render(fmt.Sprintf("Assuming %d min to fall asleep", int(buffer.Minutes()))) + "\n\n")
	for i := len(bedTimes) - 1; i >= 0; i-- {
		c := minCycles + i
		fmt.Fprintf(&sb, "  %s cycles  →  sleep at %s  %s\n",
			styles.Result.Render(fmt.Sprintf("%d", c)),
			styles.Result.Render(bedTimes[i].Format("15:04")),
			styles.Dim.Render("("+FormatDuration(c)+")"))
	}
	return sb.String()
}

func Window(from, to string, fromTime, toTime time.Time, bufferMinutes int) string {
	buffer := time.Duration(bufferMinutes) * time.Minute
	cycles, remainder := cycle.CalculateCyclesInWindow(fromTime, toTime, buffer)
	var sb strings.Builder
	sb.WriteString(styles.Result.Render(fmt.Sprintf("Between %s and %s", from, to)) + "\n")
	sb.WriteString(styles.Separator.Render(strings.Repeat("─", 30)) + "\n")
	sb.WriteString(styles.Dim.Render(fmt.Sprintf("Assuming %d min to fall asleep", bufferMinutes)) + "\n\n")
	fmt.Fprintf(&sb, "  %s complete cycles  %s\n",
		styles.Result.Render(fmt.Sprintf("%d", cycles)),
		styles.Dim.Render("("+FormatDuration(cycles)+")"))
	fmt.Fprintf(&sb, "  %s minutes remaining\n",
		styles.Result.Render(fmt.Sprintf("%d", int(remainder.Minutes()))))
	return sb.String()
}
