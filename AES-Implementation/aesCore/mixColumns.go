package aesCore

var mixColumnsTable [16]byte = [16]byte{
	0x02, 0x03, 0x01, 0x01,
	0x01, 0x02, 0x03, 0x01,
	0x01, 0x01, 0x02, 0x03,
	0x03, 0x01, 0x01, 0x02,
}

var mixColumnsInverseTable [16]byte = [16]byte{
	0x0e, 0x0b, 0x0d, 0x09,
	0x09, 0x0e, 0x0b, 0x0d,
	0x0d, 0x09, 0x0e, 0x0b,
	0x0b, 0x0d, 0x09, 0x0e,
}

// gfMul performs multiplication in Galois Field GF(2^8) for AES
func gfMul(a, b byte) byte {
	var res byte = 0
	for i := 0; i < 8; i++ {
		if b&1 != 0 {
			res ^= a
		}
		hiBitSet := (a & 0x80) != 0
		a <<= 1
		if hiBitSet {
			a ^= 0x1b
		}
		b >>= 1
	}
	return res
}

func MixColumns(bytesArray []byte, decrypt bool) []byte {
	var mixTable [16]byte
	if decrypt {
		mixTable = mixColumnsInverseTable
	} else {
		mixTable = mixColumnsTable
	}

	mixedBytes := make([]byte, 16)

	// Process each of the 4 columns independently
	for c := 0; c < 4; c++ {
		colStart := c * 4
		s0 := bytesArray[colStart+0]
		s1 := bytesArray[colStart+1]
		s2 := bytesArray[colStart+2]
		s3 := bytesArray[colStart+3]

		// Matrix multiplication row by column using GF arithmetic and XOR (^)
		mixedBytes[colStart+0] = gfMul(mixTable[0], s0) ^ gfMul(mixTable[1], s1) ^ gfMul(mixTable[2], s2) ^ gfMul(mixTable[3], s3)
		mixedBytes[colStart+1] = gfMul(mixTable[4], s0) ^ gfMul(mixTable[5], s1) ^ gfMul(mixTable[6], s2) ^ gfMul(mixTable[7], s3)
		mixedBytes[colStart+2] = gfMul(mixTable[8], s0) ^ gfMul(mixTable[9], s1) ^ gfMul(mixTable[10], s2) ^ gfMul(mixTable[11], s3)
		mixedBytes[colStart+3] = gfMul(mixTable[12], s0) ^ gfMul(mixTable[13], s1) ^ gfMul(mixTable[14], s2) ^ gfMul(mixTable[15], s3)
	}

	return mixedBytes
}
