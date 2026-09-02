package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Base int

const (
	BaseBinary Base = 2
	BaseOctal Base = 8
	BaseDecimal Base = 10
	BaseHexadecimal Base = 16
	BaseText Base = 1
	BaseUnknown Base = 0
)

var (
	ErrInvalidFormat = errors.New("input string does not match expected prefix format")
	ErrInvalidDigits = errors.New("input string contains digits invalid for its declared base")
)

func RotateLeftString(s string, k int) string {
	n := len(s)
	if n == 0 {
		return s
	}
	k = k % n
	return s[k:] + s[:k]
}

func DetectBase(s string) (Base, error) {
	if len(s) == 0 {
		return BaseUnknown, errors.New("empty string")
	}
	if strings.HasPrefix(s, "0b") || strings.HasPrefix(s, "0B") {
		return BaseBinary, nil
	}
	if strings.HasPrefix(s, "0o") || strings.HasPrefix(s, "0O") {
		return BaseOctal, nil
	}
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return BaseHexadecimal, nil
	}
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		return BaseDecimal, nil
	}
	return BaseText, nil
}

func ParseFormattedString(s string, expectedBase Base) (int64, error) {
	cleanStr := s
	switch expectedBase {
	case BaseBinary:
		cleanStr = strings.TrimPrefix(strings.TrimPrefix(s, "0b"), "0B")
	case BaseOctal:
		cleanStr = strings.TrimPrefix(strings.TrimPrefix(s, "0o"), "0O")
	case BaseHexadecimal:
		cleanStr = strings.TrimPrefix(strings.TrimPrefix(s, "0x"), "0X")
	case BaseDecimal:
		cleanStr = s
	}
	if len(cleanStr) == 0 {
		return 0, ErrInvalidDigits
	}
	val, err := strconv.ParseInt(cleanStr, int(expectedBase), 64)
	if err != nil {
		return 0, ErrInvalidDigits
	}
	return val, nil
}

func ConvertToBinaryString(input string) (string, error) {
	base, err := DetectBase(input)
	if err != nil {
		return "", err
	}

	if base == BaseText {
		var sb strings.Builder
		for i := 0; i < len(input); i++ {
			sb.WriteString(fmt.Sprintf("%08b", input[i]))
		}
		return sb.String(), nil
	}

	val, err := ParseFormattedString(input, base)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%b", val), nil
}

func ApplyPKCS7PaddingBits(bitStr string) string {
	byteLen := (len(bitStr) + 7) / 8
	padBytes := 8 - (byteLen % 8)
	if padBytes == 0 {
		padBytes = 8
	}
	
	for len(bitStr)%8 != 0 {
		bitStr = "0" + bitStr
	}

	for i := 0; i < padBytes; i++ {
		bitStr += fmt.Sprintf("%08b", padBytes)
	}
	return bitStr
}

func BinaryBitsToText(bitStr string) string {
	var sb strings.Builder
	for i := 0; i < len(bitStr); i += 8 {
		if i+8 > len(bitStr) {
			break
		}
		val, _ := strconv.ParseInt(bitStr[i:i+8], 2, 64)
		sb.WriteByte(byte(val))
	}
	return sb.String()
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter plaintext: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
