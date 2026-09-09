package main

import (
	"fmt"
	"math"
)

// TimeBreakdown converts seconds into seconds, hours, days, and years
type TimeBreakdown struct {
	Seconds float64
	Hours   float64
	Days    float64
	Years   float64
}

func calculateTimeBreakdown(totalSeconds float64) TimeBreakdown {
	return TimeBreakdown{
		Seconds: totalSeconds,
		Hours:   totalSeconds / 3600.0,
		Days:    totalSeconds / 86400.0,
		Years:   totalSeconds / (86400.0 * 365.25), // Accounting for leap years
	}
}

// Exercise 11: Extrapolation to Full DES (2^56 Effective Key Space)
func PrintDESExtrapolation(bestThroughputKeysPerSec float64) {
	// Full DES effective key space sizes
	totalKeysMax := math.Pow(2, 56) // 2^56 = 72,057,594,037,927,936
	totalKeysAvg := math.Pow(2, 55) // 2^55 = 36,028,797,018,963,968

	tMaxSeconds := totalKeysMax / bestThroughputKeysPerSec
	tAvgSeconds := totalKeysAvg / bestThroughputKeysPerSec

	tMax := calculateTimeBreakdown(tMaxSeconds)
	tAvg := calculateTimeBreakdown(tAvgSeconds)

	fmt.Println("=== Exercise 11: Extrapolation to Full DES (2^56 Key Space) ===")
	fmt.Printf("Best Measured Throughput (R): %.2f keys/second\n", bestThroughputKeysPerSec)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-25s | %-15s | %-15s\n", "Metric Unit", "Average Time (T_avg)", "Maximum Time (T_max)")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-25s | %-15.2f | %-15.2f\n", "Seconds", tAvg.Seconds, tMax.Seconds)
	fmt.Printf("%-25s | %-15.2f | %-15.2f\n", "Hours", tAvg.Hours, tMax.Hours)
	fmt.Printf("%-25s | %-15.2f | %-15.2f\n", "Days", tAvg.Days, tMax.Days)
	fmt.Printf("%-25s | %-15.4f | %-15.4f\n", "Years", tAvg.Years, tMax.Years)
	fmt.Println("--------------------------------------------------------------------------------")
}

func main() {
	// Example: Assuming measured peak throughput R = 500,000 keys/sec
	// In practice, this value will be passed dynamically from Exercise 10 results.
	samplePeakThroughput := 500000.0
	PrintDESExtrapolation(samplePeakThroughput)
}
