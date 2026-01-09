package b58uuid

import (
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		uuid     string
		expected string
		wantErr  bool
	}{
		{
			name:     "standard UUID with hyphens",
			uuid:     "550e8400-e29b-41d4-a716-446655440000",
			expected: "BWBeN28Vb7cMEx7Ym8AUzs",
			wantErr:  false,
		},
		{
			name:     "UUID without hyphens",
			uuid:     "550e8400e29b41d4a716446655440000",
			expected: "BWBeN28Vb7cMEx7Ym8AUzs",
			wantErr:  false,
		},
		{
			name:     "nil UUID (all zeros)",
			uuid:     "00000000-0000-0000-0000-000000000000",
			expected: "1111111111111111111111",
			wantErr:  false,
		},
		{
			name:     "max UUID (all Fs)",
			uuid:     "ffffffff-ffff-ffff-ffff-ffffffffffff",
			expected: "YcVfxkQb6JRzqk5kF2tNLv",
			wantErr:  false,
		},
		{
			name:     "invalid UUID (too short)",
			uuid:     "550e8400-e29b-41d4-a716",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "invalid UUID (non-hex characters)",
			uuid:     "550e8400-e29b-41d4-a716-44665544000g",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "empty string",
			uuid:     "",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Encode(tt.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("Encode() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		b58      string
		expected string
		wantErr  bool
	}{
		{
			name:     "standard b58uuid",
			b58:      "BWBeN28Vb7cMEx7Ym8AUzs",
			expected: "550e8400-e29b-41d4-a716-446655440000",
			wantErr:  false,
		},
		{
			name:     "nil UUID encoded",
			b58:      "1111111111111111111111",
			expected: "00000000-0000-0000-0000-000000000000",
			wantErr:  false,
		},
		{
			name:     "max UUID encoded",
			b58:      "YcVfxkQb6JRzqk5kF2tNLv",
			expected: "ffffffff-ffff-ffff-ffff-ffffffffffff",
			wantErr:  false,
		},
		{
			name:     "invalid character (0)",
			b58:      "0000000000000000",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "invalid character (O)",
			b58:      "OOOOOOOOOOOOOOOO",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "empty string",
			b58:      "",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Decode(tt.b58)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("Decode() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	tests := []string{
		"00000000-0000-0000-0000-000000000000",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
		"550e8400-e29b-41d4-a716-446655440000",
		"deadbeef-cafe-babe-0123-456789abcdef",
		"123e4567-e89b-12d3-a456-426614174000",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			encoded, err := Encode(tt)
			if err != nil {
				t.Fatalf("Encode() error = %v", err)
			}
			decoded, err := Decode(encoded)
			if err != nil {
				t.Fatalf("Decode() error = %v", err)
			}
			if decoded != tt {
				t.Errorf("Round trip failed: input = %q, encoded = %q, decoded = %q", tt, encoded, decoded)
			}
		})
	}
}

func TestNew(t *testing.T) {
	// Generate 100 UUIDs and ensure they are all unique and valid.
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		b58, err := New()
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		if seen[b58] {
			t.Errorf("New() generated duplicate: %q", b58)
		}
		seen[b58] = true

		// Decode and verify it's a valid UUID.
		uuid, err := Decode(b58)
		if err != nil {
			t.Errorf("New() generated invalid b58uuid %q: %v", b58, err)
		}

		// Verify UUID format.
		if len(uuid) != 36 {
			t.Errorf("New() generated UUID with incorrect length: %q", uuid)
		}
		if strings.Count(uuid, "-") != 4 {
			t.Errorf("New() generated UUID with incorrect hyphen count: %q", uuid)
		}
	}
}

func TestMustEncode(t *testing.T) {
	result := MustEncode("550e8400-e29b-41d4-a716-446655440000")
	expected := "BWBeN28Vb7cMEx7Ym8AUzs"
	if result != expected {
		t.Errorf("MustEncode() = %q, want %q", result, expected)
	}
}

func TestMustEncodePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustEncode() did not panic on invalid input")
		}
	}()
	MustEncode("invalid-uuid")
}

func TestMustDecode(t *testing.T) {
	result := MustDecode("BWBeN28Vb7cMEx7Ym8AUzs")
	expected := "550e8400-e29b-41d4-a716-446655440000"
	if result != expected {
		t.Errorf("MustDecode() = %q, want %q", result, expected)
	}
}

func TestMustDecodePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustDecode() did not panic on invalid input")
		}
	}()
	MustDecode("invalid0b58")
}

func BenchmarkEncode(b *testing.B) {
	uuid := "550e8400-e29b-41d4-a716-446655440000"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Encode(uuid)
	}
}

func BenchmarkDecode(b *testing.B) {
	b58 := "BWBeN28Vb7cMEx7Ym8AUzs"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(b58)
	}
}

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = New()
	}
}

func TestOutputLength(t *testing.T) {
	// All encodings must produce exactly 22 characters
	tests := []string{
		"00000000-0000-0000-0000-000000000000",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
		"550e8400-e29b-41d4-a716-446655440000",
		"deadbeef-cafe-babe-0123-456789abcdef",
		"00000000-0000-0000-0000-000000000001",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			encoded, err := Encode(tt)
			if err != nil {
				t.Fatalf("Encode() error = %v", err)
			}
			if len(encoded) != 22 {
				t.Errorf("Encode() length = %d, want 22 for UUID %s", len(encoded), tt)
			}
		})
	}
}

func TestOverflowDetection(t *testing.T) {
	// Create a Base58 string that would overflow u128
	overflowStr := "zzzzzzzzzzzzzzzzzzzzzz" // 22 'z' characters
	_, err := Decode(overflowStr)
	if err == nil {
		t.Error("Decode() should return error for overflow value")
	}
	if !strings.Contains(err.Error(), "overflow") {
		t.Errorf("Decode() error should mention overflow, got: %v", err)
	}
}

func TestUUIDVersionAndVariant(t *testing.T) {
	// Generate multiple UUIDs and verify version and variant bits
	for i := 0; i < 100; i++ {
		b58, err := New()
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}

		bytes, err := DecodeBytes(b58)
		if err != nil {
			t.Fatalf("DecodeBytes() error = %v", err)
		}

		// Check version 4 (bits 0100 in byte 6)
		version := (bytes[6] & 0xF0) >> 4
		if version != 4 {
			t.Errorf("UUID should be version 4, got %d", version)
		}

		// Check variant (bits 10 in byte 8)
		variant := (bytes[8] & 0xC0) >> 6
		if variant != 2 {
			t.Errorf("UUID should have variant 10 (RFC 4122), got %d", variant)
		}
	}
}
