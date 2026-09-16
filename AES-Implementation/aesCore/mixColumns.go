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

func MixColumns(bytesArray []byte, decrypt bool) []byte {
	var mixTable [16]byte
	if decrypt == true {
		mixTable = mixColumnsInverseTable
	} else {
		mixTable = mixColumnsTable
	}
	var mixedBytes []byte = make([]byte, 16, 16)
	var sum uint16 = 0
	var count uint8 = 0
	for index := 0; index < len(mixedBytes); index++ {
		sum += mixTable[((index / 4) * 4) + (index % 4)] * bytesArray[((index % 4) * 4) + (index / 4)] 
		count += 1
		if count >= 4 {
			mixedBytes[((index / 4) * 4) + (index % 4)] = sum
			sum = 0
			count = 0
		}
	}
	return mixedBytes
}
