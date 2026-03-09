package cycle

import (
	"testing"
	"time"
)

func TestCalculateBedtimes(t *testing.T) {
	wakeTime := time.Date(2026, time.March, 8, 7, 0, 0, 0, time.UTC)
	buffer := 15 * time.Minute
	minCycles := 4
	maxCycles := 6

	result := CalculateBedtimes(wakeTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 8, 0, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 23, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 21, 45, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d bedtimes, got %d", len(expected), len(result))
	}
	for i, bedtime := range result {
		if !bedtime.Equal(expected[i]) {
			t.Errorf("expected bedtime %d to be %v, got %v", i, expected[i], bedtime)
		}
	}
}

func TestCalculateWakeTimes(t *testing.T) {
	sleepTime := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	buffer := 15 * time.Minute
	minCycles := 4
	maxCycles := 6

	result := CalculateWakeTimes(sleepTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 8, 4, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 5, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 7, 15, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d wake times, got %d", len(expected), len(result))
	}
	for i, wakeTime := range result {
		if !wakeTime.Equal(expected[i]) {
			t.Errorf("expected wake time %d to be %v, got %v", i, expected[i], wakeTime)
		}
	}
}

func TestCalculateBedtimes_zeroBuffer(t *testing.T) {
	wakeTime := time.Date(2026, time.March, 8, 7, 0, 0, 0, time.UTC)
	buffer := 0 * time.Minute
	minCycles := 4
	maxCycles := 6

	result := CalculateBedtimes(wakeTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 8, 1, 00, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 23, 30, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 22, 00, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d bedtimes, got %d", len(expected), len(result))
	}
	for i, bedtime := range result {
		if !bedtime.Equal(expected[i]) {
			t.Errorf("expected bedtime %d to be %v, got %v", i, expected[i], bedtime)
		}
	}
}

func TestCalculateWakeTimes_zeroBuffer(t *testing.T) {
	sleepTime := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	buffer := 0 * time.Minute
	minCycles := 4
	maxCycles := 6

	result := CalculateWakeTimes(sleepTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 8, 4, 00, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 5, 30, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 7, 00, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d wake times, got %d", len(expected), len(result))
	}
	for i, wakeTime := range result {
		if !wakeTime.Equal(expected[i]) {
			t.Errorf("expected wake time %d to be %v, got %v", i, expected[i], wakeTime)
		}
	}
}

func TestCalculateBedtimes_largeBuffer(t *testing.T) {
	wakeTime := time.Date(2026, time.March, 8, 7, 0, 0, 0, time.UTC)
	buffer := 120 * time.Minute
	minCycles := 4
	maxCycles := 6

	result := CalculateBedtimes(wakeTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 7, 23, 00, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 21, 30, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 20, 00, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d bedtimes, got %d", len(expected), len(result))
	}
	for i, bedtime := range result {
		if !bedtime.Equal(expected[i]) {
			t.Errorf("expected bedtime %d to be %v, got %v", i, expected[i], bedtime)
		}
	}
}

func TestCalculateWakeTimes_largeBuffer(t *testing.T) {
	sleepTime := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	buffer := 120 * time.Minute
	minCycles := 4
	maxCycles := 6

	result := CalculateWakeTimes(sleepTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 8, 6, 00, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 7, 30, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 9, 00, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d wake times, got %d", len(expected), len(result))
	}
	for i, wakeTime := range result {
		if !wakeTime.Equal(expected[i]) {
			t.Errorf("expected wake time %d to be %v, got %v", i, expected[i], wakeTime)
		}
	}
}

func TestCalculateBedtimes_oneCycle(t *testing.T) {
	wakeTime := time.Date(2026, time.March, 8, 7, 0, 0, 0, time.UTC)
	buffer := 15 * time.Minute
	minCycles := 6
	maxCycles := 6

	result := CalculateBedtimes(wakeTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 7, 21, 45, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d bedtimes, got %d", len(expected), len(result))
	}
	for i, bedtime := range result {
		if !bedtime.Equal(expected[i]) {
			t.Errorf("expected bedtime %d to be %v, got %v", i, expected[i], bedtime)
		}
	}
}

func TestCalculateWakeTimes_oneCycle(t *testing.T) {
	sleepTime := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	buffer := 15 * time.Minute
	minCycles := 6
	maxCycles := 6

	result := CalculateWakeTimes(sleepTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 8, 7, 15, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d wake times, got %d", len(expected), len(result))
	}
	for i, wakeTime := range result {
		if !wakeTime.Equal(expected[i]) {
			t.Errorf("expected wake time %d to be %v, got %v", i, expected[i], wakeTime)
		}
	}
}

func TestCalculateBedtimes_tenCycles(t *testing.T) {
	wakeTime := time.Date(2026, time.March, 8, 7, 0, 0, 0, time.UTC)
	buffer := 15 * time.Minute
	minCycles := 1
	maxCycles := 10

	result := CalculateBedtimes(wakeTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 8, 5, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 3, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 2, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 0, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 23, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 21, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 20, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 18, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 17, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 7, 15, 45, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d bedtimes, got %d", len(expected), len(result))
	}
	for i, bedtime := range result {
		if !bedtime.Equal(expected[i]) {
			t.Errorf("expected bedtime %d to be %v, got %v", i, expected[i], bedtime)
		}
	}
}

func TestCalculateWakeTimes_tenCycles(t *testing.T) {
	sleepTime := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	buffer := 15 * time.Minute
	minCycles := 1
	maxCycles := 10

	result := CalculateWakeTimes(sleepTime, buffer, minCycles, maxCycles)

	expected := []time.Time{
		time.Date(2026, time.March, 7, 23, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 1, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 2, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 4, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 5, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 7, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 8, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 10, 15, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 11, 45, 0, 0, time.UTC),
		time.Date(2026, time.March, 8, 13, 15, 0, 0, time.UTC),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d wake times, got %d", len(expected), len(result))
	}
	for i, wakeTime := range result {
		if !wakeTime.Equal(expected[i]) {
			t.Errorf("expected wake time %d to be %v, got %v", i, expected[i], wakeTime)
		}
	}
}

func TestCalculateCyclesInWindow(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 8, 7, 0, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 0)

	if cycles != 6 {
		t.Errorf("expected 6 cycles, got %d", cycles)
	}
	if overflow != 0 {
		t.Errorf("expected 0 overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_withBuffer(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 8, 7, 15, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 15*time.Minute)

	if cycles != 6 {
		t.Errorf("expected 6 cycles, got %d", cycles)
	}
	if overflow != 0 {
		t.Errorf("expected 0 overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_nextDayWake(t *testing.T) {
	from := time.Date(2026, time.March, 7, 23, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 7, 7, 0, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 0)

	if cycles != 5 {
		t.Errorf("expected 5 cycles, got %d", cycles)
	}
	if overflow != 30*time.Minute {
		t.Errorf("expected 30m overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_withOverflow(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 8, 7, 44, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 0)

	if cycles != 6 {
		t.Errorf("expected 6 cycles, got %d", cycles)
	}
	if overflow != 44*time.Minute {
		t.Errorf("expected 44m overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_bufferLargerThanWindow(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 7, 22, 10, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 15*time.Minute)

	if cycles != 0 {
		t.Errorf("expected 0 cycles, got %d", cycles)
	}
	if overflow != 0 {
		t.Errorf("expected 0 overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_overflowOnlyNoCycles(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 7, 22, 45, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 0)

	if cycles != 0 {
		t.Errorf("expected 0 cycles, got %d", cycles)
	}
	if overflow != 45*time.Minute {
		t.Errorf("expected 45m overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_exactlyOneCycle(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 7, 23, 30, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 0)

	if cycles != 1 {
		t.Errorf("expected 1 cycle, got %d", cycles)
	}
	if overflow != 0 {
		t.Errorf("expected 0 overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_exactlyOneCycleWithBuffer(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 7, 23, 45, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 15*time.Minute)

	if cycles != 1 {
		t.Errorf("expected 1 cycle, got %d", cycles)
	}
	if overflow != 0 {
		t.Errorf("expected 0 overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_zeroBuffer(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 8, 5, 30, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 0)

	if cycles != 5 {
		t.Errorf("expected 5 cycles, got %d", cycles)
	}
	if overflow != 0 {
		t.Errorf("expected 0 overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_largeBuffer(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 8, 9, 0, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 120*time.Minute)

	if cycles != 6 {
		t.Errorf("expected 6 cycles, got %d", cycles)
	}
	if overflow != 0 {
		t.Errorf("expected 0 overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_midnightBoundary(t *testing.T) {
	from := time.Date(2026, time.March, 7, 23, 30, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 8, 0, 0, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 0)

	if cycles != 0 {
		t.Errorf("expected 0 cycles, got %d", cycles)
	}
	if overflow != 30*time.Minute {
		t.Errorf("expected 30m overflow, got %v", overflow)
	}
}

func TestCalculateCyclesInWindow_sameFromAndTo(t *testing.T) {
	from := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)
	to := time.Date(2026, time.March, 7, 22, 0, 0, 0, time.UTC)

	cycles, overflow := CalculateCyclesInWindow(from, to, 0)

	if cycles != 16 {
		t.Errorf("expected 16 cycles, got %d", cycles)
	}
	if overflow != 0 {
		t.Errorf("expected 0 overflow, got %v", overflow)
	}
}
