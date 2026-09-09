package main

import (
	"fmt"
	"strings"

	"desProject/modes/cbcMode"
	"desProject/modes/ecbMode"
)

// Helper to flip a single bit at a given bit index in a binary string ('0' <-> '1')
func flipBitInBitString(bitStr string, bitIndex int) string {
	if bitIndex < 0 || bitIndex >= len(bitStr) {
		return bitStr
	}
	runes := []rune(bitStr)
	if runes[bitIndex] == '0' {
		runes[bitIndex] = '1'
	} else {
		runes[bitIndex] = '0'
	}
	return string(runes)
}

// Helper to calculate Hamming distance (number of differing bits) between two bit strings
func hammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		minLen := len(s1)
		if len(s2) < minLen {
			minLen = len(s2)
		}
		diff := 0
		for i := 0; i < minLen; i++ {
			if s1[i] != s2[i] {
				diff++
			}
		}
		return diff + abs(len(s1)-len(s2))
	}
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Key: 64-bit binary string
	key := "0001001100110100010101110111100110011011101111001101111111110001"
	// IV: 64-bit binary string for CBC
	iv := "0000000000000000000000000000000000000000000000000000000000000000"

	// Plaintext: Three identical 64-bit blocks (192 bits total)
	singleBlock := "0100000101000010010000110100010001000101010001100100011101001000" // "ABCDEFGH"
	plaintext := strings.Repeat(singleBlock, 3)

	fmt.Println("==================================================")
	fmt.Println("      SECTION 4: ERROR PROPAGATION EXPERIMENT      ")
	fmt.Println("==================================================")

	// 1. ECB Error Propagation
	ecbEncrypted := ecbMode.DESEncryptionECB(plaintext, key)
	
	// Flip 1 bit in Block 2 (bit index 68, which is the 5th bit of Block 2)
	targetBitIndex := 68
	ecbCorrupted := flipBitInBitString(ecbEncrypted, targetBitIndex)
	ecbDecrypted := ecbMode.DESDecryptionECB(ecbCorrupted, key)

	fmt.Println("\n--- [ECB MODE] ---")
	fmt.Printf("Flipped Bit Index: %d (Block 2, Bit 5)\n", targetBitIndex)
	fmt.Println("Decrypted Blocks Analysis:")
	for blockIdx := 0; blockIdx < 3; blockIdx++ {
		start := blockIdx * 64
		end := start + 64
		origBlock := plaintext[start:end]
		decBlock := ecbDecrypted[start:end]
		diff := hammingDistance(origBlock, decBlock)
		fmt.Printf("  Block %d: %d / 64 bits corrupted (%.2f%% error rate)\n", blockIdx+1, diff, (float64(diff)/64.0)*100)
	}

	// 2. CBC Error Propagation
	cbcEncrypted := cbcMode.DESEncryptionCBC(plaintext, key, iv)
	cbcCorrupted := flipBitInBitString(cbcEncrypted, targetBitIndex)
	cbcDecrypted := cbcMode.DESDecryptionCBC(cbcCorrupted, key, iv)

	fmt.Println("\n--- [CBC MODE] ---")
	fmt.Printf("Flipped Bit Index: %d (Block 2, Bit 5)\n", targetBitIndex)
	fmt.Println("Decrypted Blocks Analysis:")
	for blockIdx := 0; blockIdx < 3; blockIdx++ {
		start := blockIdx * 64
		end := start + 64
		origBlock := plaintext[start:end]
		decBlock := cbcDecrypted[start:end]
		diff := hammingDistance(origBlock, decBlock)
		fmt.Printf("  Block %d: %d / 64 bits corrupted (%.2f%% error rate)\n", blockIdx+1, diff, (float64(diff)/64.0)*100)
	}
	fmt.Println("==================================================")
}
