package destd

const (
	PLAIN_TEXT_SIZE = 64
	PERM_MATRIX_SIZE int = 64
	EXPANSION_MATRIX_SIZE int = 48
	SBOX_SIZE int = 64
	FEISTEL_PERM_MATRIX_SIZE int = 32
	INITIAL_KEY_PERM_MATRIX_SIZE = 56
	FINAL_KEY_PERM_MATRIX_SIZE = 48
)

var initialPermutationMatrix = [PERM_MATRIX_SIZE]int{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 28, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

var finalPermutationMatrix = [PERM_MATRIX_SIZE]int{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 49, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25,
}

var expansionMatrix = [EXPANSION_MATRIX_SIZE]int{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1,
}

var sBox = [8][SBOX_SIZE]int{
	// sBox_1
	{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7,
	0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8,
	4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0,
	15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
	// sBox_2
	{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10,
	3, 13, 4, 7, 16, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5,
	0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15,
	13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
	// sBox_3
	{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8,
	13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1,
	13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7,
	1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
	// sBox_4
	{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15,
	13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9,
	10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4,
	3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
	// sBox_5
	{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9,
	14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6,
	4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14,
	11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
	// sBox_6
	{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11,
	10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8,
	9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6,
	4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
	// sBox_7
	{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1,
	13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6,
	1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2,
	6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 13, 2, 3, 12},
	// sBox_8
	{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7,
	1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2,
	7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8,
	2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 8, 0, 3, 5, 6, 11}
}

var feistelPermutationMatrix = [FEISTEL_PERM_MATRIX_SIZE]int{
	16, 7, 20, 21, 29, 12, 28, 17,
	1, 15, 23, 26, 5, 18, 31, 10,
	2, 8, 24, 14, 32, 27, 3, 9,
	19, 13, 30, 6, 22, 11, 4, 25,
}

var initialKeyPermutationMatrix = [INITIAL_KEY_PERM_MATRIX_SIZE]int{
	57, 49, 41, 33, 25, 17, 9, 1,
	58, 50, 42, 34, 26, 18, 10, 2,
	59, 51, 43, 35, 27, 19, 11, 3,
	60, 52, 44, 36, 63, 55, 47, 39,
	31, 23, 15, 7, 62, 54, 46, 38,
	30, 22, 14, 6, 61, 53, 45, 37,
	29, 21, 13, 5, 28, 20, 12, 4,
}

var finalKeyPermutationMatrix = [FINAL_KEY_PERM_MATRIX_SIZE]int{
	14, 17, 11, 24, 1, 5, 3, 28,
	15, 6, 21, 10, 23, 19, 12, 4,
	26, 8, 16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55, 30, 40,
	51, 45, 33, 48, 44, 49, 39, 56,
	34, 53, 46, 50, 36, 29, 32,
}

func Permutation(plaintext string, permutationMatrix [PERM_MATRIX_SIZE]int) []byte {
	permutedPlaintext := make([]byte, PLAIN_TEXT_LENGTH)
	for index, val := range permutationMatrix {
		permutedPlaintext[index] = plaintext[val]
	}
	return permutedPlaintext
}

func InitialKeyPermutation(key string, keyPermutationMatrix [INITIAL_KEY_PERM_MATRIX_SIZE]int) []byte {
	permutedKey := make([]byte, INITIAL_KEY_PERM_MATRIX_SIZE)
	for index, val := range keyPermutationMatrix {
		permutedKey[index] = plaintext[val]
	}
	return permutedKey
}

func TransformKey(C_0 byte[], D_0 byte[], i int) string {
	switch i {
	case 1, 2, 9, 16:
		numOfRotations := 1
	default:
		numOfRotations := 2
	}
	

}

func sBox(slice_n byte[], sBox_n [SBOX_SIZE]int) string {
	colNumber := 16
	rowBinary := string(slice_n[0]) + string(slice_n[-1])
	colBinary := string(slice_n[1:-2])
	row, _ := strconv.ParseInt(rowBinary, 2, 64)
	col, _ := strconv.ParseInt(colBinary, 2, 64)
	sBoxOutput := strconv.FormatInt(sBox_n[row * colNumber + col], 2)
	return sBoxOutput
}

func Feistel(transformKey string, R_n string) string {
	expansion := make([]byte, EXPANSION_MATRIX_SIZE)
	for index, val := range expansionMatrix {
		expansion[index] = R_n[val]
	}
	expansionByte := []byte(transformKey)
	xorResult := make([]byte, EXPANSION_MATRIX_SIZE)
	for i := 0; i < EXPANSION_MATRIX_SIZE; i++ {
		xorResult[i] = expansionByte[i] ^ transformKey[i]
	}
	feistelXor := ""
	minSlice = 0
	maxSlice = 5
	sliceSize = 6
	for i := 0; i < 7 {
		slice = xorResult[minSlice:maxSlice]
		feistelXor += sBox(slice, sBox[i])
		minSlice += sliceSize
		maxSlice += sliceSize
	}
	feistelResult := make([]byte, FEISTEL_PERM_MATRIX_SIZE)
	for index, val := range feistelPermutationMatrix {
		permutedPlaintext[index] = plaintext[val]
	}
	return feistelResult
}

func EncryptionRound(permutedPlaintext []byte) string {
	L_n := permutedPlaintext[:32]
	R_n := permutedPlaintext[32:]
	feistelResult := Feistel(transformKey, R_n)
	xorResult := make([]byte, len(L_n))
	for i := 0; i < len(L_n); i++ {
		xorResult[i] = L_n[i] ^ fesitelResult[i]
	}
	L_m := R_n
	R_m := xorResult
	return string(L_m) + string(R_m)
}

func DES(plaintext string) string {
	permutation := Permutation(plaintext, initialPermutationMatrix)
	for i := 0; i < 15; i++ {
		encryption_i := EncryptionRound(permutation)
		permutation := encryption_i
	}
	L_16 := permutation[:32]
	R_16 := permutation[32:]
	finalRoundResult := string(R_16) + string(L_16) 
	encryptedText := Permutation(finalRoundResult, finalPermutationMatrix)
	return encryptedText
}
