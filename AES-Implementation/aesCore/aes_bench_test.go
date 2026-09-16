package aesCore

import (
	"crypto/rand"
	"testing"
	"time"
)

func TestAESBenchmark(t *testing.T) {
	keys := map[string][]byte{
		"AES-128": unhex("2b7e151628aed2a6abf7158809cf4f3c"),
		"AES-192": unhex("8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b"),
		"AES-256": unhex("603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4"),
	}

	sizes := map[string]int{
		"1 MB":   1 * 1024 * 1024,
		"10 MB":  10 * 1024 * 1024,
		"100 MB": 100 * 1024 * 1024,
	}

	t.Logf("=======================================================================================")
	t.Logf("AES Core Benchmark Report (Averaged over 3 runs)")
	t.Logf("=======================================================================================")

	for name, key := range keys {
		for sizeName, sizeBytes := range sizes {
			// Pad data size to a multiple of 16 bytes for block processing
			numBlocks := (sizeBytes + 15) / 16
			actualBytes := numBlocks * 16
			data := make([]byte, actualBytes)
			rand.Read(data)

			// Pre-generate ciphertext blocks for decryption benchmarking
			ciphertext := make([]byte, actualBytes)
			for i := 0; i < actualBytes; i += 16 {
				enc, err := EncryptBlock(data[i:i+16], key)
				if err != nil {
					t.Fatalf("setup encryption failed: %v", err)
				}
				copy(ciphertext[i:i+16], enc)
			}

			var encTimes []float64
			var decTimes []float64

			// Repeat experiment at least three times
			for run := 0; run < 3; run++ {
				// Measure Encryption Time
				startEnc := time.Now()
				for i := 0; i < actualBytes; i += 16 {
					_, _ = EncryptBlock(data[i:i+16], key)
				}
				encTimes = append(encTimes, time.Since(startEnc).Seconds())

				// Measure Decryption Time
				startDec := time.Now()
				for i := 0; i < actualBytes; i += 16 {
					_, _ = DecryptBlock(ciphertext[i:i+16], key)
				}
				decTimes = append(decTimes, time.Since(startDec).Seconds())
			}

			// Calculate averages
			avgEncTime := (encTimes[0] + encTimes[1] + encTimes[2]) / 3.0
			avgDecTime := (decTimes[0] + decTimes[1] + decTimes[2]) / 3.0

			// Processed data size in MB (binary definition: 1 MB = 1024 * 1024 bytes)
			dMB := float64(actualBytes) / (1024.0 * 1024.0)

			// Calculate throughput R = D / T (MB/s)
			encThroughput := dMB / avgEncTime
			decThroughput := dMB / avgDecTime

			t.Logf("[%s | %6s] Enc Time: %7.4f s | Enc Throughput: %8.2f MB/s || Dec Time: %7.4f s | Dec Throughput: %8.2f MB/s",
				name, sizeName, avgEncTime, encThroughput, avgDecTime, decThroughput)
		}
	}
	t.Logf("=======================================================================================")
}
