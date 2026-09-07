package myDES

import (
	"fmt"
	"strconv"
	"strings"
	"des-lab/utils"
)

var IPMap = [64]int{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7}

var IPInvMap = [64]int{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25}

func EncryptBlock64(block64Bit, key64Bit string) string {
	permutation := PermuteString(block64Bit, IPMap[:])
	L := permutation[:32]
	R := permutation[32:]
	keyState := InitialKeyPermutation(key64Bit)
	for i := 1; i <= 16; i++ {
		var subKey string
		subKey, keyState = TransformKey(keyState, i)
		L, R = EncryptionRound(L, R, subKey)
	}
	finalRoundResult := R + L
	return PermuteString(finalRoundResult, IPInvMap[:])
}

func DecryptBlock64(cipherBlock64Bit, key64Bit string) string {
	permutation := PermuteString(cipherBlock64Bit, IPMap[:])
	L := permutation[:32]
	R := permutation[32:]
	var subKeys [16]string
	keyState := InitialKeyPermutation(key64Bit)
	for i := 1; i <= 16; i++ {
		subKeys[i-1], keyState = TransformKey(keyState, i)
	}
	for i := 15; i >= 0; i-- {
		L, R = EncryptionRound(L, R, subKeys[i])
	}
	finalRoundResult := R + L
	return PermuteString(finalRoundResult, IPInvMap[:])
}
