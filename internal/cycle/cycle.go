package cycle

import (
	"time"
)

const CycleDuration = 90 * time.Minute

// CalculateBedtimes calculates potential sleep times based on a desired wake time.
// It subtracts cycles + buffer time from the wake time.
func CalculateBedtimes(wakeTime time.Time, buffer time.Duration, minCycles int, maxCycles int) []time.Time {
	var bedtimes []time.Time

	for cycleCount := minCycles; cycleCount <= maxCycles; cycleCount++ {
		totalDuration := time.Duration(cycleCount)*CycleDuration + buffer
		bedtimes = append(bedtimes, wakeTime.Add(-totalDuration))
	}
	return bedtimes
}

// CalculateWakeTimes calculates potential wake times based on a sleep time.
// It adds cycles + buffer time to the sleep time.
func CalculateWakeTimes(sleepTime time.Time, buffer time.Duration, minCycles int, maxCycles int) []time.Time {
	var wakeTimes []time.Time

	for cycleCount := minCycles; cycleCount <= maxCycles; cycleCount++ {
		totalDuration := time.Duration(cycleCount)*CycleDuration + buffer
		wakeTimes = append(wakeTimes, sleepTime.Add(totalDuration))
	}
	return wakeTimes
}
