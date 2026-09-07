package permutation

func Permute(text string, table []int) string {
	var permutation strings.Builder
	for _, value := range table {
		permutation.WriteByte(text[value - 1])
	}
	return permutation.String()
}
