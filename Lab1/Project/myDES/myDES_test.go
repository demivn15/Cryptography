package myDES_test

import (
	"crypto/des"
	"fmt"
	"strconv"
	"testing"

	"des-lab/myDES"
)

// Helper to convert byte slice to a 64-bit binary string
func bytesToBitString(b []byte) string {
	var sb string
	for _, byteVal := range b {
		sb += fmt.Sprintf("%08b", byteVal)
	}
	return sb
}

// Helper to convert a 64-bit binary string back into a byte slice
func bitStringToBytes(bitStr string) []byte {
	bytes := make([]byte, len(bitStr)/8)
	for i := 0; i < len(bitStr); i += 8 {
		val, _ := strconv.ParseUint(bitStr[i:i+8], 2, 8)
		bytes[i/8] = byte(val)
	}
	return bytes
}

func TestDESAgainstStdLib(t *testing.T) {
	// Standard 64-bit NIST DES test key and plaintext
	keyBytes := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	plainBytes := []byte{0x4E, 0x6F, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74} // "Now is t"

	// 1. Run Standard Library DES
	block, err := des.NewCipher(keyBytes)
	if err != nil {
		t.Fatalf("Failed to initialize stdlib DES: %v", err)
	}
	stdCipher := make([]byte, 8)
	block.Encrypt(stdCipher, plainBytes)

	// 2. Run Custom Bit-String DES
	keyBits := bytesToBitString(keyBytes)
	plainBits := bytesToBitString(plainBytes)
	customCipherBits := myDES.EncryptBlock64(plainBits, keyBits)

	// 3. Verify Encryption Output
	expectedCipherBits := bytesToBitString(stdCipher)
	if customCipherBits != expectedCipherBits {
		t.Errorf("Encryption Mismatch!\nGot:      %s\nExpected: %s", customCipherBits, expectedCipherBits)
	} else {
		t.Logf("Encryption Verification Passed!")
	}

	// 4. Verify Decryption Output
	decryptedBits := myDES.DecryptBlock64(customCipherBits, keyBits)
	if decryptedBits != plainBits {
		t.Errorf("Decryption Mismatch!\nGot:      %s\nExpected: %s", decryptedBits, plainBits)
	} else {
		t.Logf("Decryption Verification Passed!")
	}
}
