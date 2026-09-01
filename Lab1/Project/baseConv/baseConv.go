package baseConv

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Base int // Base types represented by integer constants

const (
	BaseBinary      Base = 2
	BaseOctal       Base = 8
	BaseDecimal     Base = 10
	BaseHexadecimal Base = 16
	BaseUnknown     Base = 0
)

var (
	ErrInvalidFormat = errors.New("input string does not match expected prefix format")
	ErrInvalidDigits = errors.New("input string contains digits invalid for its declared base")
)

func DetectBase(s string) (Base, error) {
	if len(s) == 0 {
		return BaseUnknown, errors.New("empty string")
	}
	if strings.HasPrefix(s, "0b") || strings.HasPrefix(s, "0B") { // Check prefixes for explicit base formats
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
	return BaseUnknown, ErrInvalidFormat
}

func ConvertToBinary(input string) (string, error) { // ConvertToBinary detects the base format of the input string and converts it to formatted binary ("0b...").
	base, err := DetectBase(input)
	if err != nil {
		return "", err
	}
	val, err := ParseFormattedString(input, base) // Validate and parse the string to integer
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0b%s", strconv.FormatInt(val, 2)), nil // Format to binary string with "0b" prefix
}

func ConvertFromBinary(binaryInput string, targetBase Base) (string, error) {
	base, err := DetectBase(binaryInput)
	if err != nil || base != BaseBinary { // Ensure the input is valid binary with "0b" prefix
		return "", fmt.Errorf("invalid binary input: %w", ErrInvalidFormat)
	}
	val, err := ParseFormattedString(binaryInput, BaseBinary)
	if err != nil {
		return "", err
	}
	switch targetBase { // Format output with the correct prefix based on target Base
	case BaseBinary:
		return fmt.Sprintf("0b%s", strconv.FormatInt(val, 2)), nil
	case BaseOctal:
		return fmt.Sprintf("0o%s", strconv.FormatInt(val, 8)), nil
	case BaseDecimal:
		return strconv.FormatInt(val, 10), nil // Decimal uses no prefix
	case BaseHexadecimal:
		return fmt.Sprintf("0x%s", strconv.FormatInt(val, 16)), nil
	default:
		return "", fmt.Errorf("unsupported target base: %d", targetBase)
	}
}

func ParseFormattedString(s string, expectedBase Base) (int64, error) { // Strips prefix formatting and parses the digits strictly according to the base.
	cleanStr := s
	switch expectedBase {
	case BaseBinary:
		if !strings.HasPrefix(s, "0b") && !strings.HasPrefix(s, "0B") {
			return 0, ErrInvalidFormat
		}
		cleanStr = s[2:]
	case BaseOctal:
		if !strings.HasPrefix(s, "0o") && !strings.HasPrefix(s, "0O") {
			return 0, ErrInvalidFormat
		}
		cleanStr = s[2:]
	case BaseHexadecimal:
		if !strings.HasPrefix(s, "0x") && !strings.HasPrefix(s, "0X") {
			return 0, ErrInvalidFormat
		}
		cleanStr = s[2:]
	case BaseDecimal:
		cleanStr = s
	}
	if len(cleanStr) == 0 {
		return 0, ErrInvalidDigits
	}
	val, err := strconv.ParseInt(cleanStr, int(expectedBase), 64) // Parse digits strictly using the target base
	if err != nil {
		return 0, ErrInvalidDigits
	}
	return val, nil
}
