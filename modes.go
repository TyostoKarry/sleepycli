package main

import (
	"fmt"
	"time"

	"github.com/TyostoKarry/sleepycli/internal/cycle"
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

	fmt.Printf("If you go to sleep now at %s:\n", now.Format("15:04"))
	for i, wakeTime := range wakeTimes {
		cycleCount := minCycles + i
		fmt.Printf("  - For %d cycles, wake up at %s\n", cycleCount, wakeTime.Format("15:04"))
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
	fmt.Printf("Between %s and %s:\n", from, to)
	fmt.Printf("You can fit %d complete sleep cycles (%d minutes remaining)\n", cycles, int(remainder.Minutes()))
	return nil
}

func runWakeMode(wake string, buffer time.Duration, minCycles, maxCycles int) error {
	wakeTime, err := time.Parse("15:04", validate.NormalizeHour(wake))
	if err != nil {
		return err
	}
	bedtimes := cycle.CalculateBedtimes(wakeTime, buffer, minCycles, maxCycles)

	fmt.Printf("To wake up at %s:\n", wake)
	for i, bedtime := range bedtimes {
		cycleCount := minCycles + i
		fmt.Printf("  - For %d cycles, go to sleep at %s\n", cycleCount, bedtime.Format("15:04"))
	}
	return nil
}

func runSleepMode(sleep string, buffer time.Duration, minCycles, maxCycles int) error {
	sleepTime, err := time.Parse("15:04", validate.NormalizeHour(sleep))
	if err != nil {
		return err
	}
	wakeTimes := cycle.CalculateWakeTimes(sleepTime, buffer, minCycles, maxCycles)

	fmt.Printf("If you go to sleep at %s:\n", sleep)
	for i, wakeTime := range wakeTimes {
		cycleCount := minCycles + i
		fmt.Printf("  - For %d cycles, wake up at %s\n", cycleCount, wakeTime.Format("15:04"))
	}
	return nil
}
