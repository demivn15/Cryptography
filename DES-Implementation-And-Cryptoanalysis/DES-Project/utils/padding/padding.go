package padding

import (
	"strconv"
	"fmt"
)

func Pkcs7Pad(bitString string) string {
	var bitStringLength int = len(bitString)
	var totalBytes int = bitStringLength / 8
	var missingBytes = 8 - (totalBytes % 8)
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
	if bitStringLastBit < 8 {
		return bitString
	}
	paddingInt, _ := strconv.ParseInt(bitString[bitStringLastBit - 8:], 2, 64)
	var paddingBits int = int(paddingInt) * 8
	if paddingBits > bitStringLastBit {
		return bitString
	}
	return bitString[:bitStringLastBit - paddingBits]
}
