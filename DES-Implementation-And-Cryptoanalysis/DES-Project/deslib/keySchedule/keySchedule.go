package keySchedule

import (
	pmt "desProject/deslib/permutation"
	bo "desProject/utils/binaryOperations"
)

var pc1 = [56]uint8{
	57, 49, 41, 33, 25, 17, 9, 1, 
	58, 50, 42, 34, 26, 18,	10, 2, 
	59, 51, 43, 35, 27, 19, 11, 3, 
	60, 52, 44, 36, 63, 55, 47, 39, 
	31, 23, 15, 7, 62, 54, 46, 38, 
	30, 22, 14, 6, 61, 53, 45, 37, 
	29, 21, 13, 5, 28, 20, 12, 4}

var pc2 = [48]uint8{
	14, 17, 11, 24, 1, 5, 3, 28, 
	15, 6, 21, 10, 23, 19, 12, 4, 
	26, 8, 16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55, 30, 40, 
	51, 45, 33, 48, 44, 49, 39, 56, 
	34, 53, 46, 42, 50, 36, 29, 32}

func initialKeyPermutation(key string) string {
	return pmt.Permute(key, pc1[:])
}

func transformKey(permutedKey string, round int) (string, string) {
	C_0 := permutedKey[:28]
	D_0 := permutedKey[28:]
	var numOfRotations int
	switch round {
	case 1, 2, 9, 16:
		numOfRotations = 1
	default:
		numOfRotations = 2
	}
	rotatedC := bo.RotateLeftString(C_0, numOfRotations)
	rotatedD := bo.RotateLeftString(D_0, numOfRotations)
	nextKeyBits := rotatedC + rotatedD
	subKey := pmt.Permute(nextKeyBits, pc2[:])
	return subKey, nextKeyBits
}

func DesKeySchedule(key string) []string {
	var subKeyList [16]string
	var permutedKey string = initialKeyPermutation(key)
	for round := 0; round == 15; round++ {
		subKeyList[round] = transformKey(permutedKey, round)
	}
	return subKeyList
}
