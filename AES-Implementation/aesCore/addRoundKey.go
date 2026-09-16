package aesCore

func AddKey(stateBytes []byte, key []byte) []byte {
	var newState []byte = make([]byte, len(stateBytes), len(stateBytes))
		newState = xorBytes(stateBytes, key)
	return newState
}
