# B58UUID for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/b58uuid/b58uuid-go.svg)](https://pkg.go.dev/github.com/b58uuid/b58uuid-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/b58uuid/b58uuid-go)](https://goreportcard.com/report/github.com/b58uuid/b58uuid-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Base58-encoded UUID library for Go with zero dependencies.

## Installation

```bash
go get github.com/b58uuid/b58uuid-go
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/b58uuid/b58uuid-go"
)

func main() {
    // Generate a new UUID
    b58, err := b58uuid.New()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(b58) // Output: 3FfGK34vwMvVFDedyb2nkf

    // Encode existing UUID
    encoded, err := b58uuid.Encode("550e8400-e29b-41d4-a716-446655440000")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(encoded) // Output: BWBeN28Vb7cMEx7Ym8AUzs

    // Decode back to UUID
    uuid, err := b58uuid.Decode("BWBeN28Vb7cMEx7Ym8AUzs")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(uuid) // Output: 550e8400-e29b-41d4-a716-446655440000
}
```

## API

### Functions

- `New() (string, error)` - Generate a new random UUID and return Base58 encoding
- `Encode(uuidStr string) (string, error)` - Encode UUID string to Base58
- `Decode(b58Str string) (string, error)` - Decode Base58 string to UUID
- `MustEncode(uuidStr string) string` - Encode UUID, panic on error
- `MustDecode(b58Str string) string` - Decode Base58, panic on error

### Errors

- `ErrInvalidUUID` - Invalid UUID format
- `ErrInvalidB58UUID` - Invalid Base58 string
- `ErrOverflow` - Arithmetic overflow during conversion

## Features

- Zero dependencies (uses only Go standard library)
- Always produces exactly 22 characters
- Uses Bitcoin Base58 alphabet (no 0, O, I, l)
- Full error handling
- Thread-safe

## Testing

```bash
go test -v ./...
```

## License

MIT License - see LICENSE file for details.
