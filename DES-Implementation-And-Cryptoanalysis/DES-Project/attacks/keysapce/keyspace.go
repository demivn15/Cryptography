package keyspace

import (
	"fmt"
	"strings"
)

func GenerateDESKey64(effective56Bits string) string {
	var key64 strings.Builder
	for i := 0; i < 56; i += 7 {
		chunk := effective56Bits[i : i+7]
		onesCount := 0
		for j := 0; j < 7; j++ {
			if chunk[j] == '1' {
				onesCount++
			}
		}
		key64.WriteString(chunk)
		if onesCount%2 == 0 {
			key64.WriteByte('1')
		} else {
			key64.WriteByte('0')
		}
	}
	return key64.String()
}

func BuildCandidateKey(candidateIndex uint64, numUnknownBits int, fixedPrefixBits string) string {
	formatSpec := fmt.Sprintf("%%0%db", numUnknownBits)
	variableBits := fmt.Sprintf(formatSpec, candidateIndex)
	effective56Bits := fixedPrefixBits[:56-numUnknownBits] + variableBits
	return GenerateDESKey64(effective56Bits)
}
