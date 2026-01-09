# Contributing to b58uuid-go

Thank you for your interest in contributing to b58uuid-go! We welcome contributions from the community.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

## How to Contribute

### Reporting Bugs

Before creating a bug report, please check existing issues to avoid duplicates. When creating a bug report, include:

- A clear and descriptive title
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Go version and OS
- Code example demonstrating the issue

### Suggesting Features

Feature requests are welcome! Please provide:

- A clear description of the feature
- Use case and motivation
- Proposed API or implementation (if applicable)
- Any alternatives you've considered

### Pull Requests

1. **Fork the repository** and create your branch from `main`
2. **Make your changes** following our coding standards
3. **Add tests** for any new functionality
4. **Ensure all tests pass**: `go test -v -race ./...`
5. **Run code quality checks**:
   ```bash
   go vet ./...
   gofmt -w .
   ```
6. **Update documentation** if needed
7. **Submit a pull request** with a clear description

## Development Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/b58uuid-go.git
cd b58uuid-go

# Run tests
go test -v ./...

# Run tests with race detector
go test -v -race ./...

# Check test coverage
go test -cover ./...

# Run code quality checks
go vet ./...
gofmt -l .
```

## Coding Standards

### Go Style

- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` to format your code
- Write clear, self-documenting code
- Add comments for exported functions and types
- Keep functions small and focused

### Testing

- Write tests for all new functionality
- Maintain or improve test coverage
- Use table-driven tests where appropriate
- Test edge cases and error conditions
- Ensure tests are deterministic

### Commit Messages

- Use clear and descriptive commit messages
- Start with a verb in present tense (e.g., "Add", "Fix", "Update")
- Keep the first line under 72 characters
- Add detailed description if needed

Example:
```
Add support for custom Base58 alphabets

- Implement configurable alphabet option
- Add tests for custom alphabets
- Update documentation with examples
```

## Project Structure

```
b58uuid-go/
├── b58uuid.go           # Main API
├── b58uuid_test.go      # Main tests
├── internal/
│   └── base58/
│       ├── encoder.go       # Base58 encoding implementation
│       └── encoder_test.go  # Encoding tests
├── .github/
│   ├── workflows/       # CI/CD workflows
│   └── ISSUE_TEMPLATE/  # Issue templates
├── LICENSE
└── README.md
```

## Testing Guidelines

### Unit Tests

- Test all public functions
- Test error conditions
- Test edge cases (empty strings, nil values, overflow)
- Use meaningful test names

### Benchmarks

When adding performance-critical code, include benchmarks:

```go
func BenchmarkEncode(b *testing.B) {
    uuid := "550e8400-e29b-41d4-a716-446655440000"
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Encode(uuid)
    }
}
```

## Documentation

- Update README.md for user-facing changes
- Add godoc comments for exported functions
- Include code examples in documentation
- Keep documentation clear and concise

## Release Process

Releases are managed by maintainers:

1. Update version in documentation
2. Create a git tag: `git tag v1.x.x`
3. Push tag: `git push origin v1.x.x`
4. Create GitHub release with changelog
5. pkg.go.dev will automatically index the new version

## Questions?

If you have questions about contributing, feel free to:

- Open a discussion on GitHub
- Create an issue with the "question" label
- Check existing issues and discussions

## License

By contributing to b58uuid-go, you agree that your contributions will be licensed under the MIT License.

Thank you for contributing! 🎉
