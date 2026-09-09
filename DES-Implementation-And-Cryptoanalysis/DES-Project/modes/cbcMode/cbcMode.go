package cbcMode

import (
	"desProject/deslib/desCore"
	bo "desProject/utils/binaryOperations"
	pd "desProject/utils/padding"
)

func DESEncryptionCBC(plaintext string, encryptionKey string, initializationVector string) string {
	if len(plaintext) == 0 {
		return ""
	}

	paddedPlaintext := pd.Pkcs7Pad(plaintext)
	var completeEncryption string
	currentIV := initializationVector

	for i := 0; i < len(paddedPlaintext); i += 64 {
		block := paddedPlaintext[i : i+64]
		xorBlock := bo.XorBitStrings(currentIV, block)
		encryptedBlock := desCore.EncryptBlock(xorBlock, encryptionKey)
		completeEncryption += encryptedBlock
		currentIV = encryptedBlock
	}
	return completeEncryption
}

func DESDecryptionCBC(encryptedText string, encryptionKey string, initializationVector string) string {
	// Guard against empty input or misaligned ciphertext blocks
	if len(encryptedText) == 0 || len(encryptedText)%64 != 0 {
		return ""
	}

	var completeDecryption string
	previousBlock := initializationVector

	for i := 0; i < len(encryptedText); i += 64 {
		block := encryptedText[i : i+64]
		decryptedBlock := desCore.DecryptBlock(block, encryptionKey)
		plaintextBlock := bo.XorBitStrings(previousBlock, decryptedBlock)
		completeDecryption += plaintextBlock
		previousBlock = block
	}

	return pd.Pkcs7Unpad(completeDecryption)
}
