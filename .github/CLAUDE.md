# CLAUDE.md

Quick reference for Claude Code when working with go-validate - a Go struct field validation library with tag-based validations and utility functions.

## 🎯 Project Overview

**go-validate** is a lightweight Go validation library providing:
- Tag-based struct field validation (`validation:"rule=value"`)
- Built-in validators (string length, numeric ranges, email, regex)
- Extra validation utilities (phone, SSN, IP, DNS, enum)
- Thread-safe validation maps using reflection
- Custom validation interface support

## 🏗️ Core Architecture

### Key Types
```go
// Main validation interface - implement this for custom validators
type Interface interface {
    SetFieldIndex(int)
    FieldIndex() int
    SetFieldName(string)
    FieldName() string
    Validate(value interface{}, obj reflect.Value) *ValidationError
}

// Thread-safe validation registry
type Map struct {
    validator sync.Map               // map[reflect.Type][]Interface
    validationNameToBuilder sync.Map // map[string]func(string, reflect.Kind) (Interface, error)
}

// Validation error with field context
type ValidationError struct {
    Key     string // Field name
    Message string // Error message
}
```

### Global Registry
- `DefaultMap` - shared validation map
- Auto-registers string/numeric validations via `InitValidations()`

## 🚀 Quick Commands

```bash
# Development setup
magex update:install    # Install MAGE-X tools
magex help             # View all build commands

# Testing
magex test             # Run tests (fast)
magex test:race        # Run with race detector (slower)
magex bench            # Run benchmarks

# Quality checks
magex lint             # Run linting
magex format           # Format code
magex deps:update      # Update dependencies
```

## 📋 Validation System

### Built-in Tag Validations

**String validators:**
```go
type User struct {
    Name     string `validation:"min_length=2 max_length=50"`
    Email    string `validation:"format=email"`
    Password string `validation:"format=regexp:^[A-Za-z0-9!@#$%]{8,}$"`
    Confirm  string `validation:"compare=Password"`
}
```

**Numeric validators:**
```go
type Product struct {
    Price    float64 `validation:"min=0.01 max=9999.99"`
    Quantity int     `validation:"min=1 max=100"`
    Rating   uint8   `validation:"min=1 max=5"`
}
```

### Extra Validation Functions
```go
// Standalone validation functions (not tag-based)
IsValidEmail(email, mxCheck)          // Email with optional MX check
IsValidPhoneNumber(phone, countryCode) // US/Canada/Mexico phone numbers
IsValidSocial(ssn)                    // US Social Security Numbers
IsValidEnum(value, allowed, emptyOK)   // Enum validation
IsValidIP(ip) / IsValidIPv4/v6(ip)    // IP address validation
IsValidDNSName(dns) / IsValidHost(host) // DNS/Host validation
```

## 🔧 Adding New Validations

1. **Create validation struct** implementing `Interface`:
```go
type myValidation struct {
    Validation // Embed base struct
    customField string
}

func (m *myValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
    // Implementation
    return nil // or &ValidationError{...}
}
```

2. **Create builder function**:
```go
func myValidationBuilder(param string, kind reflect.Kind) (Interface, error) {
    return &myValidation{customField: param}, nil
}
```

3. **Register validation**:
```go
AddValidation("my_rule", myValidationBuilder)
```

## 🧪 Testing Requirements

### Test Structure
- `*_test.go` - unit tests
- `*_fuzz_test.go` - fuzz tests
- Use testify/assert for assertions

### Required Tests Before Commits
```bash
magex test        # All unit tests must pass
magex test:race   # Race condition detection
magex lint        # Code quality checks
```

### Example Test Pattern
```go
func TestMyValidation(t *testing.T) {
    tests := []struct {
        name    string
        input   interface{}
        wantErr bool
        errMsg  string
    }{
        // Test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## 📝 Code Patterns

### Error Handling
- Use sentinel errors from `errors.go` (e.g., `ErrEmailFormatInvalid`)
- Return `*ValidationError` with field context
- Use `fmt.Errorf` with `%w` for error wrapping

### Validation Registration
- Use `sync.Once` for one-time registration
- Group related validations in same file
- Follow naming: `RegisterXxxValidations()`

### Struct Validation Example
```go
type Model struct {
    Field string `validation:"rule=value"`
}

func (m *Model) Valid() (bool, []ValidationError) {
    // Run struct tag validations
    valid, errs := validate.IsValid(*m)

    // Add custom validations
    if ok, err := validate.IsValidEmail(m.Email, true); !ok {
        errs = append(errs, ValidationError{
            Key:     "Email",
            Message: err.Error(),
        })
    }

    return len(errs) == 0, errs
}
```

## ⚡ Quick Fixes

**Common Issues:**
- Missing validation registration → Call `InitValidations()`
- Thread safety → Use `sync.Map` and `sync.Once`
- Reflection errors → Check field types match validator expectations
- Tag parsing → Format: `validation:"rule1=value1 rule2=value2"`

**File Locations:**
- Core validation: `validate.go`
- String validators: `string.go`
- Numeric validators: `numeric.go`
- Extra utilities: `extra_validations.go`
- Errors: `errors.go`, `validation_error.go`
- Initialization: `init.go`