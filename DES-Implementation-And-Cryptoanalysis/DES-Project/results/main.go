package main

import (
	"fmt"
	"time"
	"desProject/deslib/desCore"
	pd "desProject/utils/padding"
	tp "desProject/utils/textProcessing"
)

var KeyBytes []byte = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
var PlainBytes []byte = []byte{0x4E, 0x6F, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74}

func main() {
	var plainTextInput string = tp.ReadInput("Enter the plaintext: ")
	var encyptionKeyInput string = tp.ReadInput("Enter the encyption key: ")
	if plainTextInput == "" {
		plainTextInput = PlainBytes
	}
	if encryptionKeyInput == "" {
		encryptionKeyInput = KeyBytes
	}
	bitStr, err := tp.ConvertToBinaryString(input)
	if err != nil {
		fmt.Printf("Conversion error: %v\n", err)
		return
	}
	paddedBitStr := pd.ApplyPKCS7PaddingBits(bitStr)
	fmt.Printf("Binary Representation: %s\n", paddedBitStr)
	fmt.Println("")
	startEnc := time.Now()
	var encryptedBlocks string
	for i := 0; i < len(paddedBitStr); i += 64 {
		block := paddedBitStr[i : i+64]
		encryptedBlocks += desCore.EncryptBlock(block, encyptionKeyInput)
	}
	encDuration := time.Since(startEnc)

	startDec := time.Now()
	var decryptedBlocks string
	for i := 0; i < len(encryptedBlocks); i += 64 {
		block := encryptedBlocks[i : i+64]
		decryptedBlocks += desCore.DecryptBlock(block, encryptionKeyInput)
	}
	decDuration := time.Since(startDec)

	textStr := tp.BinaryBitsToText(decryptedBlocks)

	fmt.Printf("Encrypted Output (Bits): %s\n", encryptedBlocks)
	fmt.Printf("Encryption Execution Time: %v\n", encDuration)
	fmt.Println("")
	fmt.Printf("Decrypted Output (Bits): %s\n", decryptedBlocks)
	fmt.Printf("Decryption Execution Time: %v\n", decDuration)
	fmt.Printf("Decryption Output (String): %s\n", textStr)
}
