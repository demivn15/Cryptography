package ecbMode

import (
	"desProject/deslib/desCore"
	pd "desProject/utils/padding"
)

func DESEncryptionECB(plaintext string, encryptionKey string) string {
	if len(plaintext) == 0 {
		return ""
	}
	paddedPlaintext := pd.Pkcs7Pad(plaintext)
	var completeEncryption string

	for i := 0; i < len(paddedPlaintext); i += 64 {
		block := paddedPlaintext[i : i+64]
		completeEncryption += desCore.EncryptBlock(block, encryptionKey)
	}
	return completeEncryption
}

func DESDecryptionECB(encryptedText string, encryptionKey string) string {
	// Guard against empty input or misaligned ciphertext blocks
	if len(encryptedText) == 0 || len(encryptedText)%64 != 0 {
		return ""
	}

	var completeDecryption string
	for i := 0; i < len(encryptedText); i += 64 {
		block := encryptedText[i : i+64]
		completeDecryption += desCore.DecryptBlock(block, encryptionKey)
	}
	
	return pd.Pkcs7Unpad(completeDecryption)
}
