package binaryOperations

func RotateLeftString(text string, numberOfRotations int) string {
	textLength := len(text)
	if textLength == 0 {
		return text
	}
	numberOfRotations = numberOfRotations % textLength
	return text[numberOfRotations:] + text[:numberOfRotations]
}

func XorBitStrings(a, b string) string {
	var sb strings.Builder
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			sb.WriteByte('0')
		} else {
			sb.WriteByte('1')
		}
	}
	return sb.String()
}
