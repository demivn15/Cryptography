package aesCore

func ShiftRowsMapping(byteArray []byte, decrypt bool) []byte {
	// Step 1: Convert input from column-major to row-major so rows are contiguous blocks of 4 bytes
	rowMajorInput := make([]byte, 16)
	for index, _byte := range byteArray {
		rowMajorInput[((index % 4) * 4) + (index / 4)] = _byte
	}

	shiftedBytes := make([]byte, 0, 16)

	// Row 0: no shifts needed
	for index := 0; index <= 3; index++ {
		shiftedBytes = append(shiftedBytes, rowMajorInput[index])
	}

	// Rows 1, 2, and 3: shift operations
	var shifts uint8 = 1
	for index := 4; index < len(rowMajorInput); index += 4 {
		var rowVal uint32 = uint32(rowMajorInput[index]) << 24 | 
			uint32(rowMajorInput[index+1]) << 16 | 
			uint32(rowMajorInput[index+2]) << 8 | 
			uint32(rowMajorInput[index+3])

		var shiftedRowVal uint32
		if decrypt {
			// Decryption: Shifting right according to row number
			shiftedRowVal = (rowVal >> (shifts * 8)) | (rowVal << (32 - shifts * 8))
		} else {
			// Encryption: Shifting left according to row number
			shiftedRowVal = (rowVal << (shifts * 8)) | (rowVal >> (32 - shifts * 8))
		}

		shiftedRow := []byte{
			byte(shiftedRowVal >> 24), 
			byte(shiftedRowVal >> 16 & 0xFF), 
			byte(shiftedRowVal >> 8 & 0xFF), 
			byte(shiftedRowVal & 0xFF),
		}
		for _, _byte := range shiftedRow {
			shiftedBytes = append(shiftedBytes, _byte)
		}
		shifts++
	}

	// Step 2: Convert back from row-major to column-major so MixColumns and subsequent rounds function correctly
	colMajorOutput := make([]byte, 16)
	for index, _byte := range shiftedBytes {
		colMajorOutput[((index % 4) * 4) + (index / 4)] = _byte
	}

	return colMajorOutput
}
