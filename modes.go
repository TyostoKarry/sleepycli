package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/TyostoKarry/sleepycli/internal/cycle"
	"github.com/TyostoKarry/sleepycli/internal/styles"
	"github.com/TyostoKarry/sleepycli/internal/validate"
)

func validateAndSelectMode(
	nowFlag bool,
	wakeFlag, sleepFlag, fromFlag, toFlag string,
	bufferFlag, cyclesMinFlag, cyclesMaxFlag int,
) error {
	cfg := validate.Config{
		Now:       nowFlag,
		WakeTime:  wakeFlag,
		SleepTime: sleepFlag,
		FromTime:  fromFlag,
		ToTime:    toFlag,
		Buffer:    bufferFlag,
		MinCycles: cyclesMinFlag,
		MaxCycles: cyclesMaxFlag,
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	buffer := time.Duration(bufferFlag) * time.Minute

	if nowFlag {
		return runNowMode(buffer, cyclesMinFlag, cyclesMaxFlag)
	}
	if fromFlag != "" && toFlag != "" {
		return runWindowMode(fromFlag, toFlag, buffer)
	}
	if wakeFlag != "" {
		return runWakeMode(wakeFlag, buffer, cyclesMinFlag, cyclesMaxFlag)
	}
	if sleepFlag != "" {
		return runSleepMode(sleepFlag, buffer, cyclesMinFlag, cyclesMaxFlag)
	}
	return fmt.Errorf("no valid mode selected")
}

func runNowMode(buffer time.Duration, minCycles, maxCycles int) error {
	now := time.Now()
	wakeTimes := cycle.CalculateWakeTimes(now, buffer, minCycles, maxCycles)

	fmt.Println(styles.Result.Render(fmt.Sprintf("Sleeping now at %s", now.Format("15:04"))))
	fmt.Println(styles.Separator.Render(strings.Repeat("─", 30)))
	fmt.Println(styles.Dim.Render(fmt.Sprintf("Assuming %d min to fall asleep", int(buffer.Minutes()))))
	fmt.Println()
	for i := len(wakeTimes) - 1; i >= 0; i-- {
		cycleCount := minCycles + i
		fmt.Printf("  %s cycles  →  wake at %s  %s\n",
			styles.Result.Render(fmt.Sprintf("%d", cycleCount)),
			styles.Result.Render(wakeTimes[i].Format("15:04")),
			styles.Dim.Render("("+formatDuration(cycleCount)+")"))
	}
	return nil
}

func runWindowMode(from, to string, buffer time.Duration) error {
	fromTime, err := time.Parse("15:04", validate.NormalizeHour(from))
	if err != nil {
		return err
	}
	toTime, err := time.Parse("15:04", validate.NormalizeHour(to))
	if err != nil {
		return err
	}

	cycles, remainder := cycle.CalculateCyclesInWindow(fromTime, toTime, buffer)
	fmt.Println(styles.Result.Render(fmt.Sprintf("Between %s and %s", from, to)))
	fmt.Println(styles.Separator.Render(strings.Repeat("─", 30)))
	fmt.Println(styles.Dim.Render(fmt.Sprintf("Assuming %d min to fall asleep", int(buffer.Minutes()))))
	fmt.Println()
	fmt.Printf("  %s complete cycles  %s\n",
		styles.Result.Render(fmt.Sprintf("%d", cycles)),
		styles.Dim.Render("("+formatDuration(cycles)+")"))
	fmt.Printf("  %s minutes remaining\n",
		styles.Result.Render(fmt.Sprintf("%d", int(remainder.Minutes()))))
	return nil
}

func runWakeMode(wake string, buffer time.Duration, minCycles, maxCycles int) error {
	wakeTime, err := time.Parse("15:04", validate.NormalizeHour(wake))
	if err != nil {
		return err
	}
	bedtimes := cycle.CalculateBedtimes(wakeTime, buffer, minCycles, maxCycles)

	fmt.Println(styles.Result.Render(fmt.Sprintf("To wake up at %s", wake)))
	fmt.Println(styles.Separator.Render(strings.Repeat("─", 30)))
	fmt.Println(styles.Dim.Render(fmt.Sprintf("Assuming %d min to fall asleep", int(buffer.Minutes()))))
	fmt.Println()
	for i := len(bedtimes) - 1; i >= 0; i-- {
		cycleCount := minCycles + i
		fmt.Printf("  %s cycles  →  sleep at %s  %s\n",
			styles.Result.Render(fmt.Sprintf("%d", cycleCount)),
			styles.Result.Render(bedtimes[i].Format("15:04")),
			styles.Dim.Render("("+formatDuration(cycleCount)+")"))
	}
	return nil
}

func runSleepMode(sleep string, buffer time.Duration, minCycles, maxCycles int) error {
	sleepTime, err := time.Parse("15:04", validate.NormalizeHour(sleep))
	if err != nil {
		return err
	}
	wakeTimes := cycle.CalculateWakeTimes(sleepTime, buffer, minCycles, maxCycles)

	fmt.Println(styles.Result.Render(fmt.Sprintf("Sleeping at %s", sleep)))
	fmt.Println(styles.Separator.Render(strings.Repeat("─", 30)))
	fmt.Println(styles.Dim.Render(fmt.Sprintf("Assuming %d min to fall asleep", int(buffer.Minutes()))))
	fmt.Println()
	for i := len(wakeTimes) - 1; i >= 0; i-- {
		cycleCount := minCycles + i
		fmt.Printf("  %s cycles  →  wake at %s  %s\n",
			styles.Result.Render(fmt.Sprintf("%d", cycleCount)),
			styles.Result.Render(wakeTimes[i].Format("15:04")),
			styles.Dim.Render("("+formatDuration(cycleCount)+")"))
	}
	return nil
}

func formatDuration(cycleCount int) string {
	sleepDuration := time.Duration(cycleCount) * cycle.CycleDuration
	hours := int(sleepDuration.Hours())
	minutes := int(sleepDuration.Minutes()) % 60
	return fmt.Sprintf("%dh %02dm", hours, minutes)
}
