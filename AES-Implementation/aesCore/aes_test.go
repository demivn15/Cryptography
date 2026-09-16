package aesCore

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func unhex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

func TestFiniteFieldMultiplication(tb *testing.T) {
	res := gfMul(0x57, 0x83)
	if res != 0xc1 {
		tb.Errorf("GF(2^8) multiplication failed: expected 0xc1, got 0x%x", res)
	}
}

func TestTransformationsInverses(tb *testing.T) {
	originalBlock := []byte{
		0x01, 0x23, 0x45, 0x67,
		0x89, 0xab, 0xcd, 0xef,
		0xfe, 0xdc, 0xba, 0x98,
		0x76, 0x54, 0x32, 0x10,
	}

	// Test ShiftRows and InvShiftRows
	shifted := ShiftRowsMapping(originalBlock, false)
	unshifted := ShiftRowsMapping(shifted, true)
	if !bytes.Equal(originalBlock, unshifted) {
		tb.Error("ShiftRows and InvShiftRows round-trip failed")
	}

	// Test MixColumns and InvMixColumns
	mixed := MixColumns(originalBlock, false)
	unmixed := MixColumns(mixed, true)
	if !bytes.Equal(originalBlock, unmixed) {
		tb.Error("MixColumns and InvMixColumns round-trip failed")
	}
}

func TestKeyScheduleGeneration(tb *testing.T) {
	key128 := unhex("2b7e151628aed2a6abf7158809cf4f3c")
	key192 := unhex("8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b")
	key256 := unhex("603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4")

	keys128, sz128, err128 := generateAllRoundKeys(key128)
	if err128 != nil || sz128 != 128 || len(keys128) != 11 {
		tb.Errorf("AES-128 key schedule generation failed: %v", err128)
	}

	keys192, sz192, err192 := generateAllRoundKeys(key192)
	if err192 != nil || sz192 != 192 || len(keys192) != 13 {
		tb.Errorf("AES-192 key schedule generation failed: %v", err192)
	}

	keys256, sz256, err256 := generateAllRoundKeys(key256)
	if err256 != nil || sz256 != 256 || len(keys256) != 15 {
		tb.Errorf("AES-256 key schedule generation failed: %v", err256)
	}
}

func TestRoundTripProperty(tb *testing.T) {
	plaintext := []byte("AQuickBrownFoxJumps")[:16]

	keys := map[string][]byte{
		"AES-128": unhex("2b7e151628aed2a6abf7158809cf4f3c"),
		"AES-192": unhex("8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b"),
		"AES-256": unhex("603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4"),
	}

	for name, key := range keys {
		tb.Run(name, func(t *testing.T) {
			ciphertext, err := EncryptBlock(plaintext, key)
			if err != nil {
				t.Fatalf("Encryption failed: %v", err)
			}

			decrypted, err := DecryptBlock(ciphertext, key)
			if err != nil {
				t.Fatalf("Decryption failed: %v", err)
			}

			if !bytes.Equal(plaintext, decrypted) {
				t.Errorf("Round-trip property failed for %s: expected %x, got %x", name, plaintext, decrypted)
			}
		})
	}
}

func TestNISTVectors(tb *testing.T) {
	t128Key := unhex("2b7e151628aed2a6abf7158809cf4f3c")
	t128Plain := unhex("3243f6a8885a308d313198a2e0370734")
	t128Cipher := unhex("3925841d02dc09fbdc118597196a0b32")

	enc128, err := EncryptBlock(t128Plain, t128Key)
	if err != nil || !bytes.Equal(enc128, t128Cipher) {
		tb.Errorf("AES-128 NIST vector encryption mismatch. Got %x, expected %x", enc128, t128Cipher)
	}

	dec128, err := DecryptBlock(t128Cipher, t128Key)
	if err != nil || !bytes.Equal(dec128, t128Plain) {
		tb.Errorf("AES-128 NIST vector decryption mismatch. Got %x, expected %x", dec128, t128Plain)
	}

	t192Key := unhex("8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b")
	t192Plain := unhex("00112233445566778899aabbccddeeff")

	enc192, err := EncryptBlock(t192Plain, t192Key)
	if err != nil {
		tb.Fatalf("AES-192 encryption failed: %v", err)
	}
	dec192, err := DecryptBlock(enc192, t192Key)
	if err != nil || !bytes.Equal(dec192, t192Plain) {
		tb.Errorf("AES-192 round-trip failed with NIST key")
	}

	t256Key := unhex("603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4")
	t256Plain := unhex("6bc1bee22e409f96e93d7e117393172a")

	enc256, err := EncryptBlock(t256Plain, t256Key)
	if err != nil {
		tb.Fatalf("AES-256 encryption failed: %v", err)
	}
	dec256, err := DecryptBlock(enc256, t256Key)
	if err != nil || !bytes.Equal(dec256, t256Plain) {
		tb.Errorf("AES-256 process failed with NIST key")
	}
}
