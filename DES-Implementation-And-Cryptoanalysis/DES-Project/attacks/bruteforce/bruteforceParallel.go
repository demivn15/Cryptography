package bruteforce

import (
	"context"
	"sync/atomic"
	"desProject/attacks/keyspace"
	"desProject/deslib/desCore"
)

type SearchResult struct { // SearchResult holds the output from a parallel key search worker
	FoundKey string
	Found    bool
}

// ParallelBruteForceDES partitions the key space [0, totalKeys) into p equal intervals
// and executes searching workers concurrently across available CPU cores.
func ParallelBruteForceDES(plaintext, ciphertext string, totalKeys uint64, numWorkers int, numUnknownBits int, fixedPrefix string) (string, uint64, bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resultChan := make(chan SearchResult, numWorkers)
	var totalCandidatesTested uint64
	intervalSize := totalKeys / uint64(numWorkers)
	for w := 0; w < numWorkers; w++ { // Launch p workers
		start := uint64(w) * intervalSize
		end := start + intervalSize
		if w == numWorkers-1 { // The last worker handles any remaining key space due to integer division remainder
			end = totalKeys
		}
		go func(wStart, wEnd uint64) {
			for candidateIdx := wStart; candidateIdx < wEnd; candidateIdx++ {
				select { // Check for early termination request from another worker
				case <-ctx.Done():
					return
				default:
				}
				atomic.AddUint64(&totalCandidatesTested, 1)
				candidateKey := keyspace.BuildCandidateKey(candidateIdx, numUnknownBits, fixedPrefix) // Candidate mapping: Candidate integer -> 64-bit DES key
				if desCore.EncryptBlock(plaintext, candidateKey) == ciphertext { // Test key against known ciphertext
					cancel() // Signal all other workers to terminate immediately
					resultChan <- SearchResult{FoundKey: candidateKey, Found: true}
					return
				}
			}
		}(start, end)
	}
	select { // Wait for first worker to find key or all workers to exhaust search space
	case res := <-resultChan:
		return res.FoundKey, atomic.LoadUint64(&totalCandidatesTested), res.Found
	case <-ctx.Done():
		select { // Secondary check if result arrived right at cancellation
		case res := <-resultChan:
			return res.FoundKey, atomic.LoadUint64(&totalCandidatesTested), res.Found
		default:
			return "", atomic.LoadUint64(&totalCandidatesTested), false
		}
	}
}
