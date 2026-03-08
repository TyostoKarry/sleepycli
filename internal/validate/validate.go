package validate

import (
	"fmt"
	"time"
)

type Config struct {
	WakeTime  string
	SleepTime string
	Buffer    int
	MinCycles int
	MaxCycles int
}

func (c *Config) Validate() error {
	if c.WakeTime != "" && c.SleepTime != "" {
		return fmt.Errorf("cannot specify both wake time and sleep time")
	}
	if c.WakeTime == "" && c.SleepTime == "" {
		return fmt.Errorf("must specify either wake time or sleep time")
	}
	if c.WakeTime != "" {
		if _, err := time.Parse("15:04", c.WakeTime); err != nil {
			return fmt.Errorf("invalid wake time format: %v", err)
		}
	}
	if c.SleepTime != "" {
		if _, err := time.Parse("15:04", c.SleepTime); err != nil {
			return fmt.Errorf("invalid sleep time format: %v", err)
		}
	}
	if c.Buffer < 0 {
		return fmt.Errorf("buffer time cannot be negative")
	}
	if c.MinCycles < 0 || c.MaxCycles < 0 {
		return fmt.Errorf("cycle counts cannot be negative")
	}
	if c.MinCycles > c.MaxCycles {
		return fmt.Errorf("min cycles cannot be greater than max cycles")
	}
	return nil
}
