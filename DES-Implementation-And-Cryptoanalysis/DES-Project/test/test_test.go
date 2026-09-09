package test

import (
	"testing"

	"desProject/attacks/bruteforce"
	"desProject/attacks/keyspace"
	"desProject/deslib/desCore"
	"desProject/modes/cbcMode"
	"desProject/modes/ecbMode"
	"desProject/utils/padding"
)

func TestPaddingRoundTrip(t *testing.T) {
	testCases := []string{
		"01000001",
		"010000010100001001000011",
		"0100000101000010010000110100010001000101010001100100011101001000",
	}

	for _, originalBits := range testCases {
		padded := padding.Pkcs7Pad(originalBits)
		unpadded := padding.Pkcs7Unpad(padded)

		if unpadded != originalBits {
			t.Errorf("Padding round-trip failed! Expected %s, got %s", originalBits, unpadded)
		}
	}
}

func TestInvalidPadding(t *testing.T) {
	invalidPaddedBlock := "0100000101000010010000110100010001000101010001100100011100001010"
	unpadded := padding.Pkcs7Unpad(invalidPaddedBlock)

	if unpadded != invalidPaddedBlock {
		t.Errorf("Invalid padding was not handled correctly!")
	}
}

func TestECBRoundTrip(t *testing.T) {
	key := "0001001100110100010101110111100110011011101111001101111011110001" // 64 bits
	plaintext := "01000001010000100100001101000100" // 32 bits

	ciphertext := ecbMode.DESEncryptionECB(plaintext, key)
	if len(ciphertext) == 0 {
		t.Fatalf("ECB Encryption returned an empty string!")
	}

	decrypted := ecbMode.DESDecryptionECB(ciphertext, key)

	if decrypted != plaintext {
		t.Errorf("ECB round trip failed! Expected %s, got %s", plaintext, decrypted)
	}
}

func TestCBCRoundTrip(t *testing.T) {
	key := "0001001100110100010101110111100110011011101111001101111011110001"
	iv := "0001001000110100010101100111100010010000101010111100110111101111"
	plaintext := "01000001010000100100001101000100"

	ciphertext := cbcMode.DESEncryptionCBC(plaintext, key, iv)
	if len(ciphertext) == 0 {
		t.Fatalf("CBC Encryption returned an empty string!")
	}

	decrypted := cbcMode.DESDecryptionCBC(ciphertext, key, iv)

	if decrypted != plaintext {
		t.Errorf("CBC round trip failed! Expected %s, got %s", plaintext, decrypted)
	}
}

func TestDifferentIVsProduceDifferentCiphertexts(t *testing.T) {
	key := "0001001100110100010101110111100110011011101111001101111011110001"
	iv1 := "0001001000110100010101100111100010010000101010111100110111101111"
	iv2 := "1111111100000000111111110000000011111111000000001111111100000000"
	plaintext := "01000001010000100100001101000100"

	ciphertext1 := cbcMode.DESEncryptionCBC(plaintext, key, iv1)
	ciphertext2 := cbcMode.DESEncryptionCBC(plaintext, key, iv2)

	if ciphertext1 == ciphertext2 {
		t.Errorf("CBC mode with different IVs produced identical ciphertexts!")
	}
}

func TestBruteForceRecovery(t *testing.T) {
	plaintext := "0100000101000010010000110100010001000101010001100100011101001000"
	fixedPrefix := "00000000000000000000000000000000000000000000000000000000"
	numUnknownBits := 8
	targetIndex := uint64(150)

	secretKey := keyspace.BuildCandidateKey(targetIndex, numUnknownBits, fixedPrefix)
	ciphertext := desCore.EncryptBlock(plaintext, secretKey)

	recoveredKey, _, found := bruteforce.BruteForceDES(plaintext, ciphertext, 0, 1<<numUnknownBits, numUnknownBits, fixedPrefix)

	if !found || recoveredKey != secretKey {
		t.Errorf("Brute-force key recovery failed! Expected key %s, got %s", secretKey, recoveredKey)
	}
}
