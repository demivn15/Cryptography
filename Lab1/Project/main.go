package main

import (
	"fmt"
	"time"
	"des-lab/myDES"
	"des-lab/utils"
)

func main() {
	input := utils.ReadInput()
	key64Bit := "0001001100110100010101110111100110011011101111001101111111110001" // Standard test key

	bitStr, err := utils.ConvertToBinaryString(input)
	if err != nil {
		fmt.Printf("Conversion error: %v\n", err)
		return
	}

	paddedBitStr := utils.ApplyPKCS7PaddingBits(bitStr)

	fmt.Printf("Binary Representation: %s\n", paddedBitStr)
	fmt.Println("")

	startEnc := time.Now()
	var encryptedBlocks string
	for i := 0; i < len(paddedBitStr); i += 64 {
		block := paddedBitStr[i : i+64]
		encryptedBlocks += myDES.EncryptBlock64(block, key64Bit)
	}
	encDuration := time.Since(startEnc)

	startDec := time.Now()
	var decryptedBlocks string
	for i := 0; i < len(encryptedBlocks); i += 64 {
		block := encryptedBlocks[i : i+64]
		decryptedBlocks += myDES.DecryptBlock64(block, key64Bit)
	}
	decDuration := time.Since(startDec)

	textStr := utils.BinaryBitsToText(decryptedBlocks)

	fmt.Printf("Encrypted Output (Bits): %s\n", encryptedBlocks)
	fmt.Printf("Encryption Execution Time: %v\n", encDuration)
	fmt.Println("")
	fmt.Printf("Decrypted Output (Bits): %s\n", decryptedBlocks)
	fmt.Printf("Decryption Execution Time: %v\n", decDuration)
	fmt.Printf("Decryption Output (String): %s\n", textStr)
}
