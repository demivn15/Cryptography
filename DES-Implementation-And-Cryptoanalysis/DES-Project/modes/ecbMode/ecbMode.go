package ecbMode

import (
	"desProject/deslib/desCore"
	pd "desProject/utils/padding"
)


func DESEncryptionECB(plaintext string, encryptionKey string) string {
	var paddedPlaintext string = pd.Pkcs7Pad(plaintext)
	var completeEncryption string
	for i := 0; i < len(paddedPlaintext); i += 64 {
		var block string = paddedPlaintext[i : i+64]
		completeEncryption += desCore.EncryptBlock(block, encryptionKey)
	}
	return completeEncryption
}

func DESDecryptionECB(encryptedText string, encryptionKey string) string {
	var completeDecryption string
	for i := 0; i < len(encryptedText); i += 64 {
		var block string = completeDecryption[i:i+64]
		completeDecryption += desCore.DecryptBlock(block, encryptionKey)
	}
	var unpaddedPlaintext string = pd.Pkcs7Unpad(completeDecryption)
	return unpaddedPlaintext
}
