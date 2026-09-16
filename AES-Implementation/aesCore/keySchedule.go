package aesCore

var roundConstantMap128 = [10]byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1b, 0x36}
var roundConstantMap192 = [8]byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}
var roundConstantMap256 = [7]byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40}

func KeySchedule(keySize uint16, prevKeyRound []byte, key []byte, round uint8) [][]byte {
	if keySize == 128 {
		keyRoundArray := make([][]byte, 0, 4)
		nextKeyRoundArray := make([][]byte, 4, 4)
		
		for index := 0; index < len(prevKeyRound) && index < 16; index += 4 {
			keyRoundArray = append(keyRoundArray, prevKeyRound[index : index + 4])
		}
		if len(keyRoundArray) < 4 {
			return nextKeyRoundArray
		}
		
		nextKeyRoundArray[0] = xorBytes(keyRoundArray[0], gFunction(keySize, [4]byte(keyRoundArray[3]), round))
		nextKeyRoundArray[1] = xorBytes(nextKeyRoundArray[0], keyRoundArray[1])
		nextKeyRoundArray[2] = xorBytes(nextKeyRoundArray[1], keyRoundArray[2])
		nextKeyRoundArray[3] = xorBytes(nextKeyRoundArray[2], keyRoundArray[3])
		return nextKeyRoundArray

	} else if keySize == 192 {
		keyRoundArray := make([][]byte, 0, 6)
		nextKeyRoundArray := make([][]byte, 6, 6)
		
		for index := 0; index < len(prevKeyRound) && index < 24; index += 4 {
			keyRoundArray = append(keyRoundArray, prevKeyRound[index : index + 4])
		}
		if len(keyRoundArray) < 6 {
			// Fallback/pad if length is short
			for len(keyRoundArray) < 6 {
				keyRoundArray = append(keyRoundArray, []byte{0, 0, 0, 0})
			}
		}
		
		nextKeyRoundArray[0] = xorBytes(keyRoundArray[0], gFunction(keySize, [4]byte(keyRoundArray[5]), round))
		nextKeyRoundArray[1] = xorBytes(nextKeyRoundArray[0], keyRoundArray[1])
		nextKeyRoundArray[2] = xorBytes(nextKeyRoundArray[1], keyRoundArray[2])
		nextKeyRoundArray[3] = xorBytes(nextKeyRoundArray[2], keyRoundArray[3])
		nextKeyRoundArray[4] = xorBytes(nextKeyRoundArray[3], keyRoundArray[4])
		nextKeyRoundArray[5] = xorBytes(nextKeyRoundArray[4], keyRoundArray[5])
		return nextKeyRoundArray

	} else { // AES-256
		keyRoundArray := make([][]byte, 0, 8)
		nextKeyRoundArray := make([][]byte, 8, 8)
		
		for index := 0; index < len(prevKeyRound) && index < 32; index += 4 {
			keyRoundArray = append(keyRoundArray, prevKeyRound[index : index + 4])
		}
		if len(keyRoundArray) < 8 {
			for len(keyRoundArray) < 8 {
				keyRoundArray = append(keyRoundArray, []byte{0, 0, 0, 0})
			}
		}
		
		if round%2 != 0 {
			nextKeyRoundArray[0] = xorBytes(keyRoundArray[0], gFunction(keySize, [4]byte(keyRoundArray[7]), (round+1)/2))
		} else {
			nextKeyRoundArray[0] = xorBytes(keyRoundArray[0], hFunction([4]byte(keyRoundArray[7])))
		}
		
		nextKeyRoundArray[1] = xorBytes(nextKeyRoundArray[0], keyRoundArray[1])
		nextKeyRoundArray[2] = xorBytes(nextKeyRoundArray[1], keyRoundArray[2])
		nextKeyRoundArray[3] = xorBytes(nextKeyRoundArray[2], keyRoundArray[3])
		
		if round%2 != 0 {
			nextKeyRoundArray[4] = xorBytes(nextKeyRoundArray[3], hFunction([4]byte(keyRoundArray[4])))
		} else {
			nextKeyRoundArray[4] = xorBytes(nextKeyRoundArray[3], keyRoundArray[4])
		}
		nextKeyRoundArray[5] = xorBytes(nextKeyRoundArray[4], keyRoundArray[5])
		nextKeyRoundArray[6] = xorBytes(nextKeyRoundArray[5], keyRoundArray[6])
		nextKeyRoundArray[7] = xorBytes(nextKeyRoundArray[6], keyRoundArray[7])
		
		return nextKeyRoundArray
	}
}

func gFunction(keySize uint16, word [4]byte, round uint8) []byte {
	var roundConstantMap []byte = make([]byte, 4, 4)
	if keySize == 128 {
		idx := int(round - 1) % len(roundConstantMap128)
		roundConstantMap[0] = roundConstantMap128[idx]
	} else if keySize == 192 {
		idx := int(round - 1) % len(roundConstantMap192)
		roundConstantMap[0] = roundConstantMap192[idx]
	} else {
		idx := int(round - 1) % len(roundConstantMap256)
		roundConstantMap[0] = roundConstantMap256[idx]
	}
	var wordVal uint32 = uint32(word[0]) << 24 | uint32(word[1]) << 16 | uint32(word[2]) << 8 | uint32(word[3])
	var shiftedwordVal uint32 = (wordVal << 8) | (wordVal >> (32 - 8))
	var rotWord []byte = []byte{byte(shiftedwordVal >> 24), byte(shiftedwordVal >> 16 & 0xFF), byte(shiftedwordVal >> 8 & 0xFF), byte(shiftedwordVal & 0xFF)}
	var subWord []byte = SBoxChunckMapping([4]byte(rotWord), false)
	var rcon []byte = xorBytes(subWord, roundConstantMap)
	return rcon
}

func hFunction(word [4]byte) []byte {
	var subWord []byte = SBoxChunckMapping(word, false)
	return subWord
}
