package aesCore

func xorBytes(byteArray_a []byte, byteArray_b []byte) []byte {
	var xorResult []byte = make([]byte, len(byteArray_a), len(byteArray_a))
	for index := 0; index < len(byteArray_a); index++ {
		xorResult[index] = byteArray_a[index] ^ byteArray_b[index]
	}
	return xorResult
}
