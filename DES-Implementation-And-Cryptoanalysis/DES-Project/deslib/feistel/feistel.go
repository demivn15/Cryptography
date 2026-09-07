package fesitel

var expansionMap = [48]int{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1}

var feistelPMap = [32]int{
	16, 7, 20, 21, 29, 12, 28, 17,
	1, 15, 23, 26, 5, 18, 31, 10,
	2, 8, 24, 14, 32, 27, 3, 9,
	19, 13, 30, 6, 22, 11, 4, 25}

func Feistel(subKey string, R_n string) string {
	expansion := PermuteString(R_n, expansionMap[:])
	xorResult := XorBitStrings(expansion, subKey)
	var feistelXor strings.Builder
	for i := 0; i < 8; i++ {
		slice := xorResult[i*6 : (i+1)*6]
		feistelXor.WriteString(sBoxLookup(slice, i))
	}
	return PermuteString(feistelXor.String(), feistelPMap[:])
}

func EncryptionRound(L_n, R_n, subKey string) (string, string) {
	feistelResult := Feistel(subKey, R_n)
	nextL := R_n
	nextR := XorBitStrings(L_n, feistelResult)
	return nextL, nextR
}


