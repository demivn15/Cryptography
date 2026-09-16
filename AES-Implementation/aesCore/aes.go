package aesCore

import (
	"fmt"
	"sync"
)

var (
	keyCacheMu sync.Mutex
	keyCache   = make(map[string][][]byte)
)

func generateAllRoundKeys(key []byte) ([][]byte, uint16, error) {
	keyStr := string(key)

	// Check cache first to avoid redundant key expansions
	keyCacheMu.Lock()
	if rk, exists := keyCache[keyStr]; exists {
		keyCacheMu.Unlock()
		var keySize uint16
		switch len(key) {
		case 16:
			keySize = 128
		case 24:
			keySize = 192
		case 32:
			keySize = 256
		}
		return rk, keySize, nil
	}
	keyCacheMu.Unlock()

	var keySize uint16
	var Nr int
	switch len(key) {
	case 16:
		keySize, Nr = 128, 10
	case 24:
		keySize, Nr = 192, 12
	case 32:
		keySize, Nr = 256, 14
	default:
		return nil, 0, fmt.Errorf("invalid key size: %d bytes (must be 16, 24, or 32)", len(key))
	}

	// Collect all expanded words, starting with the initial key words
	var allWords [][]byte
	for i := 0; i < len(key); i += 4 {
		w := make([]byte, 4)
		copy(w, key[i:i+4])
		allWords = append(allWords, w)
	}

	prevKey := make([]byte, len(key))
	copy(prevKey, key)

	totalWordsNeeded := (Nr + 1) * 4
	round := uint8(1)
	for len(allWords) < totalWordsNeeded {
		rkWords := KeySchedule(keySize, prevKey, key, round)
		for _, w := range rkWords {
			allWords = append(allWords, w)
		}
		
		var fullFlat []byte
		for _, w := range rkWords {
			fullFlat = append(fullFlat, w...)
		}
		copy(prevKey, fullFlat)
		round++
	}

	// Group all words into Nr + 1 round keys (16 bytes / 4 words each)
	roundKeys := make([][]byte, Nr+1)
	for r := 0; r <= Nr; r++ {
		var rkBytes []byte
		for w := 0; w < 4; w++ {
			rkBytes = append(rkBytes, allWords[r*4+w]...)
		}
		roundKeys[r] = rkBytes
	}

	// Store in cache
	keyCacheMu.Lock()
	keyCache[keyStr] = roundKeys
	keyCacheMu.Unlock()

	return roundKeys, keySize, nil
}

func EncryptBlock(plaintext []byte, key []byte) ([]byte, error) {
	if len(plaintext) != 16 {
		return nil, fmt.Errorf("plaintext must be exactly 16 bytes")
	}

	roundKeys, _, err := generateAllRoundKeys(key)
	if err != nil {
		return nil, err
	}
	Nr := len(roundKeys) - 1

	state := make([]byte, 16)
	copy(state, plaintext)

	state = AddKey(state, roundKeys[0])

	for round := 1; round < Nr; round++ {
		subbed := make([]byte, 16)
		for i := 0; i < 16; i += 4 {
			chunk := [4]byte{state[i], state[i+1], state[i+2], state[i+3]}
			res := SBoxChunckMapping(chunk, false)
			copy(subbed[i:], res[:])
		}
		state = subbed

		state = ShiftRowsMapping(state, false)
		state = MixColumns(state, false)
		state = AddKey(state, roundKeys[round])
	}

	subbed := make([]byte, 16)
	for i := 0; i < 16; i += 4 {
		chunk := [4]byte{state[i], state[i+1], state[i+2], state[i+3]}
		res := SBoxChunckMapping(chunk, false)
		copy(subbed[i:], res[:])
	}
	state = subbed
	state = ShiftRowsMapping(state, false)
	state = AddKey(state, roundKeys[Nr])

	return state, nil
}

func DecryptBlock(ciphertext []byte, key []byte) ([]byte, error) {
	if len(ciphertext) != 16 {
		return nil, fmt.Errorf("ciphertext must be exactly 16 bytes")
	}

	roundKeys, _, err := generateAllRoundKeys(key)
	if err != nil {
		return nil, err
	}
	Nr := len(roundKeys) - 1

	state := make([]byte, 16)
	copy(state, ciphertext)

	// 1. Undo the final round (Nr): AddRoundKey(Nr) -> InvShiftRows -> InvSubBytes
	state = AddKey(state, roundKeys[Nr])
	state = ShiftRowsMapping(state, true)
	
	subbed := make([]byte, 16)
	for i := 0; i < 16; i += 4 {
		chunk := [4]byte{state[i], state[i+1], state[i+2], state[i+3]}
		res := SBoxChunckMapping(chunk, true)
		copy(subbed[i:], res[:])
	}
	state = subbed

	// 2. Undo the middle rounds (Nr-1 down to 1): AddRoundKey -> InvMixColumns -> InvShiftRows -> InvSubBytes
	for round := Nr - 1; round >= 1; round-- {
		state = AddKey(state, roundKeys[round])
		state = MixColumns(state, true)
		state = ShiftRowsMapping(state, true)

		subbed = make([]byte, 16)
		for i := 0; i < 16; i += 4 {
			chunk := [4]byte{state[i], state[i+1], state[i+2], state[i+3]}
			res := SBoxChunckMapping(chunk, true)
			copy(subbed[i:], res[:])
		}
		state = subbed
	}

	// 3. Undo the initial AddRoundKey (roundKeys[0])
	state = AddKey(state, roundKeys[0])

	return state, nil
}
