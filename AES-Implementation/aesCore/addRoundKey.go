package aesCore

func AddKey(stateBytes [16]byte, key [16]byte) []byte {
	var newState []byte = make([]byte, len(stateBytes), len(stateBytes))
		newState = XorBytes(stateBytes, key)
	return newState
}
