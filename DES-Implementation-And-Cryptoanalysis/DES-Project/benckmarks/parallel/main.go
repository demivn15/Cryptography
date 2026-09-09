package main

import (
	"fmt"
	"runtime"
	"time"

	"desProject/attacks/bruteforce"
	"desProject/attacks/keyspace"
	"desProject/deslib/desCore"
)

func main() {
	// Sample 64-bit binary plaintext block ("ABCDEFGH")
	plaintext := "0100000101000010010000110100010001000101010001100100011101001000"
	fixedPrefix := "00000000000000000000000000000000000000000000000000000000"

	nBits := 20 // 2^20 search space
	totalKeys := uint64(1) << nBits

	// Position key at 90% of search space to test parallel coverage
	targetIndex := uint64(float64(totalKeys) * 0.90)
	secretKey := keyspace.BuildCandidateKey(targetIndex, nBits, fixedPrefix)
	ciphertext := desCore.EncryptBlock(plaintext, secretKey)

	workerCounts := []int{1, 2, 4, 8}
	maxCores := runtime.NumCPU()

	fmt.Println("=== Exercise 10: Parallel Performance Benchmark ===")
	fmt.Printf("System CPU Cores Available: %d\n", maxCores)
	fmt.Printf("Search Space Size: 2^%d (%d keys)\n", nBits, totalKeys)
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Printf("%-8s | %-10s | %-12s | %-15s | %-10s | %-12s\n",
		"Workers", "Time [s]", "Keys Tested", "Keys/s", "Speedup", "Efficiency")
	fmt.Println("------------------------------------------------------------------------------------------------")

	var sequentialTime float64

	for _, p := range workerCounts {
		if p > maxCores && p > 8 {
			continue
		}

		startTime := time.Now()
		recoveredKey, keysTested, found := bruteforce.ParallelBruteForceDES(
			plaintext, ciphertext, totalKeys, p, nBits, fixedPrefix,
		)
		elapsedSeconds := time.Since(startTime).Seconds()

		if !found || recoveredKey != secretKey {
			fmt.Printf("Error: Parallel search failed for p=%d\n", p)
			continue
		}

		keysPerSec := float64(keysTested) / elapsedSeconds

		if p == 1 {
			sequentialTime = elapsedSeconds
		}

		speedup := sequentialTime / elapsedSeconds
		efficiency := speedup / float64(p)

		fmt.Printf("%-8d | %-10.4f | %-12d | %-15.2f | %-10.2f | %-12.2f%%\n",
			p, elapsedSeconds, keysTested, keysPerSec, speedup, efficiency*100)
	}
	fmt.Println("------------------------------------------------------------------------------------------------")
}
