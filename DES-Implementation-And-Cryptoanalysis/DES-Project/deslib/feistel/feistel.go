package feistel

import (
	bo "desProject/utils/binaryOperations"
	pmt "desProject/deslib/permutation"
	sb "desProject/deslib/sBoxes"
	"strings"
)

var expansionTable = [48]int{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1}

var feistelPTable = [32]int{
	16, 7, 20, 21, 29, 12, 28, 17,
	1, 15, 23, 26, 5, 18, 31, 10,
	2, 8, 24, 14, 32, 27, 3, 9,
	19, 13, 30, 6, 22, 11, 4, 25}

func feistel(subKey string, R_n string) string {
	expansion := pmt.Permute(R_n, expansionTable[:])
	xorResult := bo.XorBitStrings(expansion, subKey)
	var feistelXor strings.Builder
	for boxIndex := 0; boxIndex < 8; boxIndex++ {
		slice := xorResult[boxIndex*6 : (boxIndex+1)*6]
		feistelXor.WriteString(sb.Substitute(slice, boxIndex))
	}
	return pmt.Permute(feistelXor.String(), feistelPTable[:])
}

func EncryptionRound(L_n, R_n, subKey string) (string, string) {
	feistelResult := feistel(subKey, R_n)
	nextL := R_n
	nextR := bo.XorBitStrings(L_n, feistelResult)
	return nextL, nextR
}
