package textPreprocessing

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Base uint8

const BaseBinary Base = 2
const BaseOctal Base = 8
const BaseDecimal Base = 10
const BaseHexadecimal Base = 16
const BaseText Base = 1
const BaseUnknown Base = 0

var ErrInvalidFormat = errors.New("Input string does not match expected prefix format.")
var ErrInvalidDigits = errors.New("Input string contains digits invalid for its declared base.")

func DetectBase(text string) (Base, error) {
	if len(text) == 0 {
		return BaseUnknown, errors.New("Empty string.")
	}
	if strings.HasPrefix(text, "0b") || strings.HasPrefix(text, "0B") {
		return BaseBinary, nil
	}
	if strings.HasPrefix(text, "0o") || strings.HasPrefix(text, "0O") {
		return BaseOctal, nil
	}
	if strings.HasPrefix(text, "0x") || strings.HasPrefix(text, "0X") {
		return BaseHexadecimal, nil
	}
	if _, err := strconv.ParseInt(text, 10, 64); err == nil {
		return BaseDecimal, nil
	}
	return BaseText, nil
}

func ParseFormattedString(text string, expectedBase Base) (int64, error) {
	cleanText := text
	switch expectedBase {
	case BaseBinary:
		cleanText = strings.TrimPrefix(strings.TrimPrefix(text, "0b"), "0B")
	case BaseOctal:
		cleanText = strings.TrimPrefix(strings.TrimPrefix(text, "0o"), "0O")
	case BaseHexadecimal:
		cleanText = strings.TrimPrefix(strings.TrimPrefix(text, "0x"), "0X")
	case BaseDecimal:
		cleanText = text
	}
	if len(cleanText) == 0 {
		return 0, ErrInvalidDigits
	}
	val, err := strconv.ParseInt(cleanText, int(expectedBase), 64)
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
