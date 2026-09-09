package main

import (
	"fmt"
	"time"
	"desProject/attacks/bruteforce"
	"desProject/attacks/keyspace"
	"desProject/deslib/desCore"
)

func main() {
	plaintext := "0100000101000010010000110100010001000101010001100100011101001000" // "ABCDEFGH"
	fixedPrefix := "00000000000000000000000000000000000000000000000000000000"     // 56-bit template
	searchSpaceSizes := []int{16, 18, 20}
	runsPerExperiment := 3
	fmt.Println("=== Exercise 8: Sequential Exhaustive Search Benchmark ===")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-6s | %-12s | %-10s | %-12s | %-15s\n", "n-bits", "Key Position", "Tested Keys", "Avg Time (s)", "Throughput (keys/s)")
	fmt.Println("--------------------------------------------------------------------------------")
	for _, n := range searchSpaceSizes {
		totalCandidatesInSpace := uint64(1) << n
		targetIndex := uint64(float64(totalCandidatesInSpace) * 0.70)
		secretKey := keyspace.BuildCandidateKey(targetIndex, n, fixedPrefix)
		ciphertext := desCore.EncryptBlock(plaintext, secretKey)
		var totalDuration time.Duration
		var candidatesTested uint64
		var recoveredKey string
		var found bool
		for run := 0; run < runsPerExperiment; run++ {
			startTime := time.Now()
			recoveredKey, candidatesTested, found = bruteforce.BruteForceDES(
				plaintext, ciphertext, 0, totalCandidatesInSpace, n, fixedPrefix,
			)
			totalDuration += time.Since(startTime)
		}
		if !found || recoveredKey != secretKey {
			fmt.Printf("Error: Key not recovered for n=%d!\n", n)
			continue
		}
		avgTimeSeconds := totalDuration.Seconds() / float64(runsPerExperiment)
		throughput := float64(candidatesTested) / avgTimeSeconds

		fmt.Printf("%-6d | %-12d | %-10d | %-12.4f | %-15.2f\n",
			n, targetIndex, candidatesTested, avgTimeSeconds, throughput)
	}
	fmt.Println("--------------------------------------------------------------------------------")
}
