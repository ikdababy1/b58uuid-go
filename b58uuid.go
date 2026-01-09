// Package b58uuid provides concise, unambiguous, and URL-safe UUIDs using Base58 encoding.
// It converts standard 36-character UUIDs into approximately 22-character Base58 strings,
// while maintaining full reversibility.
//
// The encoding uses the Bitcoin Base58 alphabet, which excludes visually ambiguous
// characters (0, O, I, l), making it ideal for user-facing applications, URLs,
// and database keys.
//
// Example usage:
//
//	// Encode a UUID
//	b58, err := b58uuid.Encode("550e8400-e29b-41d4-a716-446655440000")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(b58) // Output: BWBeN28Vb7cMEx7Ym8AUzs
//
//	// Decode back to UUID
//	uuid, err := b58uuid.Decode("BWBeN28Vb7cMEx7Ym8AUzs")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(uuid) // Output: 550e8400-e29b-41d4-a716-446655440000
//
//	// Generate a new random b58uuid
//	b58, err := b58uuid.New()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(b58)
package b58uuid

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/b58uuid/b58uuid-go/internal/base58"
)

var (
	// ErrInvalidUUID is returned when the input UUID format is invalid.
	ErrInvalidUUID = errors.New("invalid UUID format")

	// ErrInvalidB58UUID is returned when the input b58uuid cannot be decoded.
	ErrInvalidB58UUID = errors.New("invalid b58uuid format")

	// ErrOverflow is returned when arithmetic overflow occurs during conversion.
	ErrOverflow = errors.New("arithmetic overflow")
)

// Encode converts a standard UUID string (with or without hyphens) into a Base58-encoded string.
// The input UUID must be a valid 128-bit UUID in the format "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
// or "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" (32 hexadecimal characters).
func Encode(uuidStr string) (string, error) {
	// Parse and validate UUID
	parsedUUID, err := parseUUID(uuidStr)
	if err != nil {
		return "", err
	}

	// Use optimized base58 encoder
	return base58.Encode(parsedUUID), nil
}

// EncodeBytes converts a 16-byte UUID to a Base58-encoded string.
func EncodeBytes(uuidBytes [16]byte) string {
	return base58.Encode(uuidBytes)
}

// Decode converts a Base58-encoded string back to a standard UUID string.
// The output UUID will be in the canonical format with hyphens.
func Decode(b58 string) (string, error) {
	if b58 == "" {
		return "", ErrInvalidB58UUID
	}

	// Decode base58 to UUID bytes
	uuidBytes, err := base58.Decode(b58)
	if err != nil {
		// Check if it's an overflow error
		if strings.Contains(err.Error(), "overflow") {
			return "", ErrOverflow
		}
		return "", ErrInvalidB58UUID
	}

	// Format as standard UUID
	return formatUUID(uuidBytes), nil
}

// DecodeBytes converts a Base58-encoded string to a 16-byte UUID.
func DecodeBytes(b58 string) ([16]byte, error) {
	if b58 == "" {
		return [16]byte{}, ErrInvalidB58UUID
	}

	return base58.Decode(b58)
}

// New generates a new random UUID v4 and returns its Base58-encoded representation.
func New() (string, error) {
	uuidBytes, err := generateUUIDv4()
	if err != nil {
		return "", err
	}
	return base58.Encode(uuidBytes), nil
}

// generateUUIDv4 generates a random UUID v4 using crypto/rand.
func generateUUIDv4() ([16]byte, error) {
	var uuid [16]byte

	// Fill with random bytes
	_, err := rand.Read(uuid[:])
	if err != nil {
		return uuid, fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Set version 4 (bits 12-15 of time_hi_and_version)
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// Set variant (bits 6-7 of clock_seq_hi_and_reserved)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return uuid, nil
}

// MustEncode is like Encode but panics if there is an error.
func MustEncode(uuidStr string) string {
	b58, err := Encode(uuidStr)
	if err != nil {
		panic(err)
	}
	return b58
}

// MustDecode is like Decode but panics if there is an error.
func MustDecode(b58 string) string {
	uuid, err := Decode(b58)
	if err != nil {
		panic(err)
	}
	return uuid
}

// parseUUID parses a UUID string with or without hyphens.
func parseUUID(uuidStr string) ([16]byte, error) {
	var uuid [16]byte

	// Remove hyphens if present
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")

	// Validate length
	if len(uuidStr) != 32 {
		return uuid, ErrInvalidUUID
	}

	// Convert hex string to bytes
	bytes, err := hex.DecodeString(uuidStr)
	if err != nil {
		return uuid, ErrInvalidUUID
	}

	// Validate length after decoding
	if len(bytes) != 16 {
		return uuid, ErrInvalidUUID
	}

	// Copy to fixed-size array
	copy(uuid[:], bytes)
	return uuid, nil
}

// formatUUID formats UUID bytes as standard UUID string with hyphens.
func formatUUID(u [16]byte) string {
	return fmt.Sprintf(
		"%08x-%04x-%04x-%04x-%012x",
		u[0:4],
		u[4:6],
		u[6:8],
		u[8:10],
		u[10:16],
	)
}
