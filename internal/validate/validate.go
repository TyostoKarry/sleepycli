package validate

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	WakeTime  string
	SleepTime string
	FromTime  string
	ToTime    string
	Buffer    int
	MinCycles int
	MaxCycles int
}

func (c *Config) Validate() error {
	if err := validateModes(c.WakeTime, c.SleepTime, c.FromTime, c.ToTime); err != nil {
		return err
	}
	if c.FromTime != "" || c.ToTime != "" {
		if err := validateWindow(c.FromTime, c.ToTime); err != nil {
			return err
		}
	} else {
		if err := validateTimeFlags(c.WakeTime, c.SleepTime); err != nil {
			return err
		}
	}
	if err := validateBuffer(c.Buffer); err != nil {
		return err
	}
	return validateCycles(c.MinCycles, c.MaxCycles)
}

func validateModes(wake, sleep, from, to string) error {
	windowSet := from != "" || to != ""
	wakeSleepSet := wake != "" || sleep != ""
	if windowSet && wakeSleepSet {
		return fmt.Errorf("cannot use --from/--to with --wake or --sleep")
	}
	return nil
}

func validateWindow(from, to string) error {
	if from == "" {
		return fmt.Errorf("--from is required when using --to")
	}
	if to == "" {
		return fmt.Errorf("--to is required when using --from")
	}
	if err := validateTimeFormat(from); err != nil {
		return fmt.Errorf("invalid --from value: %w", err)
	}
	if err := validateTimeFormat(to); err != nil {
		return fmt.Errorf("invalid --to value: %w", err)
	}
	return nil
}

func validateTimeFlags(wake, sleep string) error {
	if wake != "" && sleep != "" {
		return fmt.Errorf("cannot specify both --wake and --sleep")
	}
	if wake == "" && sleep == "" {
		return fmt.Errorf("must specify either --wake or --sleep")
	}
	if wake != "" {
		if err := validateTimeFormat(wake); err != nil {
			return fmt.Errorf("invalid --wake value: %w", err)
		}
	}
	if sleep != "" {
		if err := validateTimeFormat(sleep); err != nil {
			return fmt.Errorf("invalid --sleep value: %w", err)
		}
	}
	return nil
}

func validateTimeFormat(s string) error {
	_, err := time.Parse("15:04", NormalizeHour(s))
	return err
}

func validateBuffer(buffer int) error {
	if buffer < 0 {
		return fmt.Errorf("--buffer cannot be negative")
	}
	return nil
}

func validateCycles(min, max int) error {
	if min < 0 || max < 0 {
		return fmt.Errorf("--cycles-min and --cycles-max cannot be negative")
	}
	if min > max {
		return fmt.Errorf("--cycles-min cannot be greater than --cycles-max")
	}
	return nil
}

func NormalizeHour(s string) string {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) == 2 && len(parts[0]) == 1 {
		return "0" + s
	}
	return s
}
