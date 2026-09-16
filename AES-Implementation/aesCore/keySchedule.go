package aesCore

var roundConstantMap128 [10]byte = [10]byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1b, 0x36}

func keySchedule(prevKeyRound [4]byte, key [16]byte) []byte { // Change to handle 192 and 256 bits for key lenght.
	keyRoundArray [4]byte = make([]byte, 0, 4)
	var nextKeyRoundArray [4]byte = make([]byte, 4, 4)
	for index := 0; index < len(key); index += 4 {
		keyRoundArray = append(keyRoundArray, key[index : index + 4])
	}
	nextKeyRoundArray[0] = XorBytes(keyRoundArray[0], gFunction(keyRoundArray[3]))
	nextKeyRoundArray[1] = XorBytes(nextKeyRoundArray[0], keyRoundArray[1])
	nextKeyRoundArray[2] = XorBytes(nextKeyRoundArray[1], keyRoundArray[2])
	nextKeyRoundArray[3] = XorBytes(nextKeyRoundArray[2], keyRoundArray[3])
	return nextKeyRoundArray
}

func gFunction(word [4]byte, round uint8) []byte {
	roundConstantMap := make(byte[], 4, 4)
	roundConstantMap[0] = roundConstantMap128[round]
	var wordVal uint32 = uint32(word[index]) << 24 | uint32(word[index + 1]) << 16 | uint32(word[index + 2]) << 8 | uint32(word[index + 3]) // byte-row to integer for shifting.
	var shiftedwordVal uint32 = (wordVal >> (8)) | (wordVal << (32 - 8))
	var rotWord []byte = []byte{byte(shiftedwordVal >> 24), byte(shiftedwordVal >> 16 & 0xFF), byte(shiftedwordVal >> 8 & 0xFF), byte(shiftedwordVal & 0xFF),} // int to byte-word after shifting.
	var subWord [4]byte = SBoxChunckMapping(rotWord, false)
	var rcon []byte = XorBytes(subWord, roundConstantMap)
	return rcon
}

func hFunction(word [byte]) []byte {
	var subWord [4]byte = SBoxChunckMapping(word, false)
	return subWord
}
