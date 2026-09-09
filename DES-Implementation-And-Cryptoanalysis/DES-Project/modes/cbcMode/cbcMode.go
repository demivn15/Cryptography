package cbcMode

import (
	"desProject/deslib/desCore"
	pd "desProject/utils/padding"
	bo "desProject/utils/binaryOperations"
)

func DESEncryptionCBC(plaintext string, encryptionKey string, initializationVector string) string {
	var paddedPlaintext string = pd.Pkcs7Pad(plaintext)
	var completeEncryption string
	for i := 0; i < len(paddedPlaintext); i += 64 {
		var block string = paddedPlaintext[i:i+64]
		var xorBlock string = bo.XorBitStrings(initializationVector, block)
		completeEncryption += desCore.EncryptBlock(xorBlock, encryptionKey)
		initializationVector = completeEncryption[i:i+64]
	}
	return completeEncryption
}

func DESDecryptionCBC(encryptedText string, encryptionKey string, initializationVector string) string {
	var completeDecryption string
	var previousBlock string = initializationVector
	for i := 0; i < len(encryptedText); i += 64 {
		var block string = completeDecryption[i:i+64]
		var decryptedBlock string = desCore.DecryptBlock(block, encryptionKey)
		var xorBlock string = bo.XorBitStrings(previousBlock, decryptedBlock)
		var completeDecryption += xorBlock
		previousBlock = block
	}
	var unpaddedPlaintext string = pd.Pkcs7Unpad(completeDecryption)
	return unpaddedPlaintext
}
