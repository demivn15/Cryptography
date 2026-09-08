package padding

import "strconv"

func Pkcs7Pad(bitString string) string {
	var bitStringLength int = len(bitString)
	if bitStringLength < 8 {
		for i := 0; i < 8 - bitStringLength; i++ {
			bitString = "0" + bitString
		}
	}
	var numberOfBytes int = (bitStringLength + 7) / 8
	var missingBytes = 8 - (numberOfBytes % 8)
	if missingBytes == 0 {
		missingBytes = 8
	}
	for bitPosition := 0; bitPosition < missingBytes; bitPosition++ {
		bitString += fmt.Sprintf("%08b", missingBytes)
	}
	return bitString
}

func Pkcs7Unpad(bitString string) string {
	bitStringLastBit := len(bitString)
	paddingInt, _ := strconv.ParseInt(bitString[bitStringLastBit - 8:bitStringLastBit], 2, 64)
	bitString = bitString[0:bitStringLastBit - 8 * int(paddingInt)]
	return bitString
}
