package main

import (
	"fmt"
	"strings"

	"desProject/modes/cbcMode"
	"desProject/modes/ecbMode"
	tp "desProject/utils/textProcessing"
)

func main() {
	// Standard 64-bit DES key and IV in binary
	key := "0001001100110100010101110111100110011011101111001101111111110001"
	iv := "0000000000000000000000000000000000000000000000000000000000000000"

	// Repeated plaintext: "ABCDEFGHABCDEFGHABCDEFGHABCDEFGH" (4 identical 64-bit blocks)
	blockASCII := "ABCDEFGH"
	blockBinary := tp.StringToBinary(blockASCII)
	plaintextBinary := strings.Repeat(blockBinary, 4)

	fmt.Println("==================================================")
	fmt.Println("    SECTION 3: ECB VS CBC REPEATED BLOCK EXPERIMENT")
	fmt.Println("==================================================")
	fmt.Printf("Plaintext ASCII: %s\n", strings.Repeat(blockASCII, 4))
	fmt.Printf("Total Bits: %d (4 blocks of 64 bits)\n", len(plaintextBinary))
	fmt.Println("--------------------------------------------------")

	// 1. ECB Encryption
	ecbEncrypted := ecbMode.DESEncryptionECB(plaintextBinary, key)
	fmt.Println("\n--- [ECB MODE ENCRYPTION] ---")
	for i := 0; i < len(ecbEncrypted)/64; i++ {
		block := ecbEncrypted[i*64 : (i+1)*64]
		fmt.Printf("C_%d: %s\n", i+1, block)
	}

	// Verify identity in ECB
	ecbBlock1 := ecbEncrypted[0:64]
	ecbBlock2 := ecbEncrypted[64:128]
	if ecbBlock1 == ecbBlock2 {
		fmt.Println("Verdict: Identical plaintext blocks produce IDENTICAL ciphertext blocks in ECB.")
	} else {
		fmt.Println("Verdict: Plaintext blocks produced different ciphertexts in ECB.")
	}

	// 2. CBC Encryption
	cbcEncrypted := cbcMode.DESEncryptionCBC(plaintextBinary, key, iv)
	fmt.Println("\n--- [CBC MODE ENCRYPTION] ---")
	for i := 0; i < len(cbcEncrypted)/64; i++ {
		block := cbcEncrypted[i*64 : (i+1)*64]
		fmt.Printf("C_%d: %s\n", i+1, block)
	}

	// Verify uniqueness in CBC
	cbcBlock1 := cbcEncrypted[0:64]
	cbcBlock2 := cbcEncrypted[64:128]
	if cbcBlock1 != cbcBlock2 {
		fmt.Println("Verdict: Identical plaintext blocks produce DISTINCT ciphertext blocks in CBC.")
	} else {
		fmt.Println("Verdict: Plaintext blocks produced identical ciphertexts in CBC.")
	}
	fmt.Println("==================================================")
}
