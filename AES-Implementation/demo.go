package main

import "fmt"
import "aesProject/aesCore"
import "encoding/hex"

func unhex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

var plaintext = unhex("3243f6a8885a308d313198a2e0370736") // 128 bits fixed size
var key = unhex("2b7e151628aed2a6abf7158809cf4f3c") // 128 bits

func main() {
	encrypted, _:= aesCore.EncryptBlock(plaintext, key)
	decrypted , _:= aesCore.DecryptBlock(encrypted, key)
	fmt.Printf("plaintext: %x\n", plaintext)
	fmt.Printf("encrypted: %x\n", encrypted)
	fmt.Printf("decrypted: %x\n", decrypted)
}
