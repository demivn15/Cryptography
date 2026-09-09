package padding

import (
	"fmt"
	"strconv"
)

func Pkcs7Pad(bitString string) string {
	// Total bits must be padded up to a multiple of 64 bits (8 bytes)
	currentBits := len(bitString)
	remainderBits := currentBits % 64
	missingBits := 64 - remainderBits
	
	if missingBits == 0 {
		missingBits = 64
	}

	missingBytes := missingBits / 8
	paddingByteBinary := fmt.Sprintf("%08b", missingBytes)

	padded := bitString
	for i := 0; i < missingBytes; i++ {
		padded += paddingByteBinary
	}
	return padded
}

func Pkcs7Unpad(bitString string) string {
	bitStringLength := len(bitString)
	// Must be at least one 64-bit block and aligned to 64 bits
	if bitStringLength < 64 || bitStringLength%64 != 0 {
		return bitString
	}

	// Read last byte (8 bits) for padding count
	lastByteBits := bitString[bitStringLength-8:]
	paddingInt, err := strconv.ParseInt(lastByteBits, 2, 64)
	if err != nil {
		return bitString
	}

	paddingBytes := int(paddingInt)
	if paddingBytes <= 0 || paddingBytes > 8 {
		return bitString
	}

	paddingBits := paddingBytes * 8
	if paddingBits > bitStringLength {
		return bitString
	}

	// Verify all padding bytes match
	expectedPaddingByte := fmt.Sprintf("%08b", paddingBytes)
	for i := bitStringLength - paddingBits; i < bitStringLength; i += 8 {
		if bitString[i:i+8] != expectedPaddingByte {
			return bitString
		}
	}

	return bitString[:bitStringLength-paddingBits]
}
