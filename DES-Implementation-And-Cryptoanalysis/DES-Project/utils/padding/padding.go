package padding

func Pkcs7Pad(bitString string) string {
	var stringLengthBytes = (len(bitString) + 7) / 8
	var padBytes = 8 - (stringLengthBytes % 8)
	if padBytes == 0 {
		padBytes = 8
	}
	for len(bitString)%8 != 0 {
		bitString = "0" + bitString
	}
	for i := 0; i < padBytes; i++ {
		bitString += fmt.Sprintf("%08b", padBytes)
	}
	return bitString
}

func Pkcs7Unpad(bitString string) string {
	return ""
}
