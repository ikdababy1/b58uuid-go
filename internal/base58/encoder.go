package base58

import (
	"fmt"
	"math/big"
)

// Alphabet is the Bitcoin Base58 alphabet
const Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// AlphabetBytes is the Bitcoin Base58 alphabet as a byte array
var AlphabetBytes = []byte(Alphabet)

// ReverseAlphabet is a lookup table for decoding Base58 characters
var ReverseAlphabet = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 0, 1, 2, 3, 4, 5, 6, 7, 8, 255, 255, 255, 255, 255, 255,
	255, 9, 10, 11, 12, 13, 14, 15, 16, 255, 17, 18, 19, 20, 21, 255,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 255, 255, 255, 255, 255,
	255, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 255, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// Encode encodes 16 bytes to a Base58 string
// Always returns exactly 22 characters
func Encode(data [16]byte) string {
	// Convert bytes to big.Int
	num := new(big.Int).SetBytes(data[:])
	base := big.NewInt(58)
	zero := big.NewInt(0)

	// Build result in reverse
	var result []byte
	for num.Cmp(zero) > 0 {
		remainder := new(big.Int).Mod(num, base)
		result = append(result, AlphabetBytes[remainder.Int64()])
		num.Div(num, base)
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// Pad with leading '1' characters to ensure exactly 22 characters
	for len(result) < 22 {
		result = append([]byte{AlphabetBytes[0]}, result...)
	}

	return string(result)
}

// Decode decodes a Base58 string to 16 bytes
// Returns error if value exceeds maximum UUID value
func Decode(s string) ([16]byte, error) {
	var result [16]byte

	if s == "" {
		return result, fmt.Errorf("empty string")
	}

	// Maximum UUID value (2^128 - 1)
	maxUUID := new(big.Int)
	maxUUID.SetString("ffffffffffffffffffffffffffffffff", 16)

	// Convert string to big.Int
	num := big.NewInt(0)
	base := big.NewInt(58)

	for _, c := range s {
		if c > 255 {
			return result, fmt.Errorf("invalid character: %c", c)
		}

		ind := ReverseAlphabet[c]
		if ind == 255 {
			return result, fmt.Errorf("invalid character: %c", c)
		}

		num.Mul(num, base)
		num.Add(num, big.NewInt(int64(ind)))

		// Check for overflow
		if num.Cmp(maxUUID) > 0 {
			return result, fmt.Errorf("overflow: value exceeds maximum UUID value")
		}
	}

	// Convert big.Int to bytes
	bytes := num.Bytes()

	// Pad with leading zeros if necessary
	if len(bytes) < 16 {
		bytes = append(make([]byte, 16-len(bytes)), bytes...)
	}

	// Check for overflow (should not happen due to check above, but defensive)
	if len(bytes) > 16 {
		return result, fmt.Errorf("value too large for UUID")
	}

	copy(result[:], bytes)
	return result, nil
}
