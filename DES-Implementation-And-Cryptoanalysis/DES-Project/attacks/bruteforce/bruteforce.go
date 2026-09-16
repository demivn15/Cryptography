package bruteforce

import (
	"desProject/attacks/keyspace"
	"desProject/deslib/desCore"
)

func BruteForceDES(plaintext string, ciphertext string, start uint64, end uint64, numUnknownBits int, fixedPrefixBits string) (string, uint64, bool) {
	var candidatesTested uint64 = 0
	for candidateIdx := start; candidateIdx < end; candidateIdx++ {
		candidatesTested++
		candidateKey := keyspace.BuildCandidateKey(candidateIdx, numUnknownBits, fixedPrefixBits)
		testedCipher := desCore.EncryptBlock(plaintext, candidateKey)
		if testedCipher == ciphertext {
			return candidateKey, candidatesTested, true
		}
	}
	return "", candidatesTested, false
}
