package main

import (
	ecb "desProject/modes/ecbMode"
	cbc "desProject/modes/cbcMode"
	tp "desProject/utils/textProcessing"
	"fmt"
)

var encryptionKey string = tp.ConvertToBinaryString(string([]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}))
var plaintext string = tp.ConvertToBinaryString(string([]byte{0x4E, 0x6F, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74, 0x4E, 0x6F, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74}))
var initializationVector string = tp.ConvertToBinaryString(string([]byte{0x4E, 0x6F, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74}))

func main() {
	var ecbEncryption string = ecb.DESEncryptionECB(plaintext, encryptionKey)
	var ecbDecryption string = ecb.DESDecryptionECB(plaintext, encryptionKey)
	var cbcEncryption string = cbc.DESEncryptionCBC(plaintext, encryptionKey, initializationVector)
	var cbcDecryption string = cbc.DESDecryptionCBC(plaintext, encryptionKey, initializationVector)
	fmt.Printf("%s/n, %s/n, %s/n, %s/n", ecbEncryption, ecbDecryption, cbcEncryption, cbcDecryption)
	
}
