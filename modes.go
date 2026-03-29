package main

import (
	"fmt"
	"time"

	"github.com/TyostoKarry/sleepycli/internal/render"
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
	fmt.Print(render.WakeTimes(now, buffer, minCycles, maxCycles,
		fmt.Sprintf("Sleeping now at %s", now.Format("15:04"))))
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

	fmt.Print(render.Window(from, to, fromTime, toTime, int(buffer.Minutes())))
	return nil
}

func runWakeMode(wake string, buffer time.Duration, minCycles, maxCycles int) error {
	wakeTime, err := time.Parse("15:04", validate.NormalizeHour(wake))
	if err != nil {
		return err
	}
	fmt.Print(render.Bedtimes(wakeTime, buffer, minCycles, maxCycles,
		fmt.Sprintf("To wake up at %s", wake)))
	return nil
}

func runSleepMode(sleep string, buffer time.Duration, minCycles, maxCycles int) error {
	sleepTime, err := time.Parse("15:04", validate.NormalizeHour(sleep))
	if err != nil {
		return err
	}
	fmt.Print(render.WakeTimes(sleepTime, buffer, minCycles, maxCycles,
		fmt.Sprintf("Sleeping at %s", sleep)))
	return nil
}
