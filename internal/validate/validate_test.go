package validate

import "testing"

func TestValidateWakeTime(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{"valid wake time", Config{WakeTime: "07:00"}, false},
		{"valid wake time format", Config{WakeTime: "7:00"}, false},
		{"invalid wake time value", Config{WakeTime: "25:00"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
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
		{"invalid sleep time format", Config{SleepTime: "10:00 PM"}, true},
		{"invalid sleep time value", Config{SleepTime: "24:00"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
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
		{"negative buffer", Config{WakeTime: "07:00", Buffer: -5}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
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
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
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
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
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
		{"min cycles greater than max cycles", Config{WakeTime: "07:00", MinCycles: 7, MaxCycles: 6}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
