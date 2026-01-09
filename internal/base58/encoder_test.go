package base58

import (
	"testing"
)

var testVectors = []struct {
	name string
	uuid [16]byte
	b58  string
}{
	{
		name: "Nil UUID (All Zeros)",
		uuid: [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		b58:  "1111111111111111111111",
	},
	{
		name: "Max UUID (All Fs)",
		uuid: [16]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		b58:  "YcVfxkQb6JRzqk5kF2tNLv",
	},
	{
		name: "Standard UUIDv4 Example 1",
		uuid: [16]byte{0x55, 0x0e, 0x84, 0x00, 0xe2, 0x9b, 0x41, 0xd4, 0xa7, 0x16, 0x44, 0x66, 0x55, 0x44, 0x00, 0x00},
		b58:  "BWBeN28Vb7cMEx7Ym8AUzs",
	},
	{
		name: "Standard UUIDv4 Example 2",
		uuid: [16]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0xba, 0xbe, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		b58:  "UVqy39vS4tbfPzthw5VEKg",
	},
	{
		name: "Sequential Low Value 1",
		uuid: [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		b58:  "1111111111111111111112",
	},
	{
		name: "Sequential Low Value 2",
		uuid: [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16},
		b58:  "111111111111111111111H",
	},
}

func TestEncode(t *testing.T) {
	for _, tv := range testVectors {
		t.Run(tv.name, func(t *testing.T) {
			result := Encode(tv.uuid)
			if result != tv.b58 {
				t.Errorf("Expected %s, got %s", tv.b58, result)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for _, tv := range testVectors {
		t.Run(tv.name, func(t *testing.T) {
			result, err := Decode(tv.b58)
			if err != nil {
				t.Errorf("Decode failed: %v", err)
			}
			if result != tv.uuid {
				t.Errorf("Expected %v, got %v", tv.uuid, result)
			}
		})
	}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	for _, tv := range testVectors {
		t.Run(tv.name, func(t *testing.T) {
			encoded := Encode(tv.uuid)
			decoded, err := Decode(encoded)
			if err != nil {
				t.Errorf("Decode failed: %v", err)
			}
			if decoded != tv.uuid {
				t.Errorf("Round trip failed: %v -> %s -> %v", tv.uuid, encoded, decoded)
			}
		})
	}
}

func TestDecodeInvalidCharacters(t *testing.T) {
	testCases := []string{
		"0000000000000000", // Contains 0
		"OOOOOOOOOOOOOOOO", // Contains O
		"IIIIIIIIIIIIIIII", // Contains I
		"llllllllllllllll", // Contains l
		"invalid",          // Invalid characters
	}

	for _, tc := range testCases {
		_, err := Decode(tc)
		if err == nil {
			t.Errorf("Expected error for invalid b58: %s", tc)
		}
	}
}

func TestDecodeEmptyString(t *testing.T) {
	_, err := Decode("")
	if err == nil {
		t.Errorf("Expected error for empty string")
	}
}

func BenchmarkEncode(b *testing.B) {
	uuid := [16]byte{0x55, 0x0e, 0x84, 0x00, 0xe2, 0x9b, 0x41, 0xd4, 0xa7, 0x16, 0x44, 0x66, 0x55, 0x44, 0x00, 0x00}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(uuid)
	}
}

func BenchmarkDecode(b *testing.B) {
	b58 := "BWBeN28Vb7cMEx7Ym8AUzs"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Decode(b58)
	}
}
