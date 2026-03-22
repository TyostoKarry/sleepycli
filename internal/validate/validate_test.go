package validate

import "testing"

func TestValidateModes(t *testing.T) {
	tests := []struct {
		name    string
		now     bool
		wake    string
		sleep   string
		from    string
		to      string
		wantErr bool
	}{
		{"now only", true, "", "", "", "", false},
		{"wake only", false, "07:00", "", "", "", false},
		{"sleep only", false, "", "22:00", "", "", false},
		{"window only", false, "", "", "22:00", "07:00", false},
		{"now with wake", true, "07:00", "", "", "", true},
		{"now with sleep", true, "", "22:00", "", "", true},
		{"now with window", true, "", "", "22:00", "07:00", true},
		{"wake with window", false, "07:00", "", "22:00", "07:00", true},
		{"sleep with window", false, "", "22:00", "22:00", "07:00", true},
		{"no mode", false, "", "", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateModes(tt.now, tt.wake, tt.sleep, tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateModes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateWakeTime(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"valid wake time", Config{WakeTime: "07:00"}, false},
		{"valid wake time short hour", Config{WakeTime: "7:00"}, false},
		{"invalid wake time value", Config{WakeTime: "25:00"}, true},
		{"invalid wake time format", Config{WakeTime: "7am"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSleepTime(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"valid sleep time", Config{SleepTime: "22:00"}, false},
		{"valid sleep time short hour", Config{SleepTime: "9:00"}, false},
		{"invalid sleep time format", Config{SleepTime: "10:00 PM"}, true},
		{"invalid sleep time value", Config{SleepTime: "24:00"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTimeFlagsMutualExclusion(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"both wake and sleep set", Config{WakeTime: "07:00", SleepTime: "22:00"}, true},
		{"neither wake nor sleep set", Config{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateBuffer(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"valid buffer", Config{WakeTime: "07:00", Buffer: 15}, false},
		{"zero buffer", Config{WakeTime: "07:00", Buffer: 0}, false},
		{"negative buffer", Config{WakeTime: "07:00", Buffer: -5}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMinCycles(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"valid min cycles", Config{WakeTime: "07:00", MinCycles: 4, MaxCycles: 6}, false},
		{"negative min cycles", Config{WakeTime: "07:00", MinCycles: -1, MaxCycles: 6}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMaxCycles(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"valid max cycles", Config{WakeTime: "07:00", MinCycles: 4, MaxCycles: 6}, false},
		{"negative max cycles", Config{WakeTime: "07:00", MinCycles: 4, MaxCycles: -1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMinMaxCycles(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"valid min and max cycles", Config{WakeTime: "07:00", MinCycles: 4, MaxCycles: 6}, false},
		{"equal min and max cycles", Config{WakeTime: "07:00", MinCycles: 6, MaxCycles: 6}, false},
		{"min cycles greater than max cycles", Config{WakeTime: "07:00", MinCycles: 7, MaxCycles: 6}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateWindow(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"valid window", Config{FromTime: "22:00", ToTime: "07:00"}, false},
		{"valid window short hour", Config{FromTime: "9:00", ToTime: "07:00"}, false},
		{"missing --to", Config{FromTime: "22:00"}, true},
		{"missing --from", Config{ToTime: "07:00"}, true},
		{"invalid --from format", Config{FromTime: "10:00 PM", ToTime: "07:00"}, true},
		{"invalid --to format", Config{FromTime: "22:00", ToTime: "7am"}, true},
		{"window mixed with --wake", Config{FromTime: "22:00", ToTime: "07:00", WakeTime: "07:00"}, true},
		{"window mixed with --sleep", Config{FromTime: "22:00", ToTime: "07:00", SleepTime: "22:00"}, true},
		{"window mixed with --now", Config{FromTime: "22:00", ToTime: "07:00", Now: true}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
