// TODO: Implement InvShiftRows

package aesCore

func ShiftRowsMapping(byteArray []byte, decrypt bool) []byte {
	if decrypt == true {
		continue
	} else {
		continue
	}
	var stateBytes []byte = make([]byte, 16, 16)
	var shiftedBytes []byte = make([]byte, 0, 16)
	for index, _byte := range byteArray {
		stateBytes[((index % 4) * 4) + (index / 4)] = _byte // Transposed matrix.
	}	
	for index := 0; index <= 3; index++ { // Appending first row as no shifts are needed.
		shiftedBytes = append(shiftedBytes, stateBytes[index])
	}
	var shifts uint8 = 3
	for index := 4; index < len(stateBytes); index += 4  {
		var rowVal uint32 = uint32(stateBytes[index]) << 24 | uint32(stateBytes[index + 1]) << 16 | uint32(stateBytes[index + 2]) << 8 | uint32(stateBytes[index + 3]) // byte-row to integer for shifting.
		var shiftedRowVal = (rowVal >> (shifts * 8)) | (rowVal << (32 - shifts * 8)) // Shifting according to row value.
		var shiftedRow []byte = []byte{byte(shiftedRowVal >> 24), byte(shiftedRowVal >> 16 & 0xFF), byte(shiftedRowVal >> 8 & 0xFF), byte(shiftedRowVal & 0xFF),} // int to byte-row after shifting.
		for _, _byte := range shiftedRow {
			shiftedBytes = append(shiftedBytes, _byte)
		}
		shifts = shifts - 1
	}
	return shiftedBytes
}
