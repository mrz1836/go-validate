<div align="center">

# ⚡&nbsp;&nbsp;go-validate

**Powerful struct field validation via tags plus extra validation utilities.**

<br/>

<a href="https://github.com/mrz1836/go-validate/releases"><img src="https://img.shields.io/github/release-pre/mrz1836/go-validate?include_prereleases&style=flat-square&logo=github&color=black" alt="Release"></a>
<a href="https://golang.org/"><img src="https://img.shields.io/github/go-mod/go-version/mrz1836/go-validate?style=flat-square&logo=go&color=00ADD8" alt="Go Version"></a>
<a href="https://github.com/mrz1836/go-validate/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mrz1836/go-validate?style=flat-square&color=blue" alt="License"></a>

<br/>

<table align="center" border="0">
  <tr>
    <td align="right">
       <code>CI / CD</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://github.com/mrz1836/go-validate/actions"><img src="https://img.shields.io/github/actions/workflow/status/mrz1836/go-validate/fortress.yml?branch=master&label=build&logo=github&style=flat-square" alt="Build"></a>
       <a href="https://github.com/mrz1836/go-validate/actions"><img src="https://img.shields.io/github/last-commit/mrz1836/go-validate?style=flat-square&logo=git&logoColor=white&label=last%20update" alt="Last Commit"></a>
    </td>
    <td align="right">
       &nbsp;&nbsp;&nbsp;&nbsp; <code>Quality</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://goreportcard.com/report/github.com/mrz1836/go-validate"><img src="https://goreportcard.com/badge/github.com/mrz1836/go-validate?style=flat-square" alt="Go Report"></a>
       <a href="https://codecov.io/gh/mrz1836/go-validate"><img src="https://codecov.io/gh/mrz1836/go-validate/branch/master/graph/badge.svg?style=flat-square" alt="Coverage"></a>
    </td>
  </tr>

  <tr>
    <td align="right">
       <code>Security</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://scorecard.dev/viewer/?uri=github.com/mrz1836/go-validate"><img src="https://api.scorecard.dev/projects/github.com/mrz1836/go-validate/badge?style=flat-square" alt="Scorecard"></a>
       <a href=".github/SECURITY.md"><img src="https://img.shields.io/badge/policy-active-success?style=flat-square&logo=security&logoColor=white" alt="Security"></a>
    </td>
    <td align="right">
       &nbsp;&nbsp;&nbsp;&nbsp; <code>Community</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://github.com/mrz1836/go-validate/graphs/contributors"><img src="https://img.shields.io/github/contributors/mrz1836/go-validate?style=flat-square&color=orange" alt="Contributors"></a>
       <a href="https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-validate&utm_term=go-validate&utm_content=go-validate"><img src="https://img.shields.io/badge/donate-bitcoin-ff9900?style=flat-square&logo=bitcoin" alt="Bitcoin"></a>
    </td>
  </tr>
</table>

</div>

<br/>
<br/>

<div align="center">

### <code>Project Navigation</code>

</div>

<table align="center">
  <tr>
    <td align="center" width="33%">
       🚀&nbsp;<a href="#installation"><code>Installation</code></a>
    </td>
    <td align="center" width="33%">
       🧪&nbsp;<a href="#examples--tests"><code>Examples&nbsp;&&nbsp;Tests</code></a>
    </td>
    <td align="center" width="33%">
       📚&nbsp;<a href="#documentation"><code>Documentation</code></a>
    </td>
  </tr>
  <tr>
    <td align="center">
       🤝&nbsp;<a href="#contributing"><code>Contributing</code></a>
    </td>
    <td align="center">
      🛠️&nbsp;<a href="#code-standards"><code>Code&nbsp;Standards</code></a>
    </td>
    <td align="center">
      ⚡&nbsp;<a href="#benchmarks"><code>Benchmarks</code></a>
    </td>
  </tr>
  <tr>
    <td align="center">
      🤖&nbsp;<a href="#-ai-usage--assistant-guidelines"><code>AI&nbsp;Usage</code></a>
    </td>
    <td align="center">
       ⚖️&nbsp;<a href="#license"><code>License</code></a>
    </td>
    <td align="center">
       👥&nbsp;<a href="#maintainers"><code>Maintainers</code></a>
    </td>
  </tr>
</table>

<br/>

## Installation

**go-validate** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get github.com/mrz1836/go-validate
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-validate)

> **Heads up!** `go-validate` is intentionally light on dependencies. The only
external package it uses is the excellent `testify` suite—and that's just for
our tests. You can drop this library into your projects without dragging along
extra baggage.

<details>
<summary><strong><code>Development Setup (Getting Started)</code></strong></summary>
<br/>

Install [MAGE-X](https://github.com/mrz1836/mage-x) build tool for development:

```bash
# Install MAGE-X for development and building
go install github.com/mrz1836/mage-x/cmd/magex@latest
magex update:install
```
</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

This project uses [goreleaser](https://github.com/goreleaser/goreleaser) for streamlined binary and library deployment to GitHub. To get started, install it via:

```bash
brew install goreleaser
```

The release process is defined in the [.goreleaser.yml](.goreleaser.yml) configuration file.

Then create and push a new Git tag using:

```bash
magex version:bump bump=patch push=true branch=master
```

This process ensures consistent, repeatable releases with properly versioned artifacts and citation metadata.

</details>

<details>
<summary><strong><code>Build Commands</code></strong></summary>
<br/>

View all build commands

```bash script
magex help
```

</details>

<details>
<summary><strong>GitHub Workflows</strong></summary>
<br/>

All workflows are driven by modular configuration in [`.github/env/`](.github/env/README.md) — no YAML editing required.

**[View all workflows and the control center →](.github/docs/workflows.md)**

</details>

<details>
<summary><strong><code>Updating Dependencies</code></strong></summary>
<br/>

To update all dependencies (Go modules, linters, and related tools), run:

```bash
magex deps:update
```

This command ensures all dependencies are brought up to date in a single step, including Go modules and any managed tools. It is the recommended way to keep your development environment and CI in sync with the latest versions.

</details>

<br/>

## Examples & Tests
All unit tests and fuzz tests run via [GitHub Actions](https://github.com/mrz1836/go-pre-commit/actions) and use [Go version 1.18.x](https://go.dev/doc/go1.18). View the [configuration file](.github/workflows/fortress.yml).

Run all tests (fast):

```bash script
magex test
```

Run all tests with race detector (slower):
```bash script
magex test:race
```

### Examples demonstrating various validation scenarios.

<details>
<summary><strong><code>Basic Struct Validation</code></strong></summary>
<br/>

```go
package main

import (
    "fmt"
    "log"

    "github.com/mrz1836/go-validate"
)

type User struct {
    Name     string `validation:"min_length=2 max_length=50"`
    Email    string `validation:"format=email"`
    Age      uint   `validation:"min=18 max=120"`
    Username string `validation:"min_length=3 max_length=20"`
}

func main() {
    // Initialize validations (required)
    validate.InitValidations()

    user := User{
        Name:     "John Doe",
        Email:    "john@example.com",
        Age:      25,
        Username: "johndoe",
    }

    // Validate the struct
    isValid, errors := validate.IsValid(user)
    if !isValid {
        for _, err := range errors {
            fmt.Printf("Validation error: %s\n", err.Error())
        }
    } else {
        fmt.Println("User is valid!")
    }
}
```
</details>

<details>
<summary><strong><code>Password Confirmation Example</code></strong></summary>
<br/>

```go
package main

import (
    "fmt"

    "github.com/mrz1836/go-validate"
)

type RegistrationForm struct {
    Email                string `validation:"format=email"`
    Password             string `validation:"min_length=8"`
    PasswordConfirmation string `validation:"compare=Password"`
    TermsAccepted        bool   // No validation needed
}

func main() {
    validate.InitValidations()

    form := RegistrationForm{
        Email:                "user@domain.com",
        Password:             "SecurePass123",
        PasswordConfirmation: "SecurePass123", // Must match Password field
        TermsAccepted:        true,
    }

    isValid, errors := validate.IsValid(form)
    if !isValid {
        fmt.Println("Registration form has errors:")
        for _, err := range errors {
            fmt.Printf("- %s\n", err.Error())
        }
    } else {
        fmt.Println("Registration form is valid!")
    }
}
```
</details>

<details>
<summary><strong><code>Custom Regex Validation</code></strong></summary>
<br/>

```go
package main

import (
    "fmt"

    "github.com/mrz1836/go-validate"
)

type Product struct {
    SKU         string  `validation:"format=regexp:^[A-Z]{2,3}[0-9]{4,6}$"`
    Name        string  `validation:"min_length=1 max_length=100"`
    Price       float64 `validation:"min=0.01"`
    Description string  `validation:"max_length=500"`
}

func main() {
    validate.InitValidations()

    product := Product{
        SKU:         "AB12345",  // Must match pattern: 2-3 uppercase letters + 4-6 digits
        Name:        "Wireless Headphones",
        Price:       99.99,
        Description: "High-quality wireless headphones with noise cancellation",
    }

    isValid, errors := validate.IsValid(product)
    if !isValid {
        fmt.Println("Product validation failed:")
        for _, err := range errors {
            fmt.Printf("- %s\n", err.Error())
        }
    } else {
        fmt.Println("Product is valid!")
    }
}
```
</details>

<details>
<summary><strong><code>Using Extra Validation Functions</code></strong></summary>
<br/>

```go
package main

import (
    "fmt"

    "github.com/mrz1836/go-validate"
)

type Contact struct {
    Name        string
    Email       string
    Phone       string
    Website     string
    CountryCode string
}

func (c *Contact) Validate() (bool, []validate.ValidationError) {
    var errors []validate.ValidationError

    // Email validation with MX record check
    if valid, err := validate.IsValidEmail(c.Email, true); !valid {
        errors = append(errors, validate.ValidationError{
            Key:     "Email",
            Message: err.Error(),
        })
    }

    // Phone number validation (US/Canada/Mexico)
    if valid, err := validate.IsValidPhoneNumber(c.Phone, c.CountryCode); !valid {
        errors = append(errors, validate.ValidationError{
            Key:     "Phone",
            Message: err.Error(),
        })
    }

    // Website host validation
    if c.Website != "" && !validate.IsValidHost(c.Website) {
        errors = append(errors, validate.ValidationError{
            Key:     "Website",
            Message: "is not a valid host or IP address",
        })
    }

    return len(errors) == 0, errors
}

func main() {
    contact := Contact{
        Name:        "Jane Smith",
        Email:       "jane@protonmail.com",
        Phone:       "555-123-4567",
        Website:     "janesmith.dev",
        CountryCode: "1", // USA/Canada
    }

    isValid, errors := contact.Validate()
    if !isValid {
        fmt.Println("Contact validation failed:")
        for _, err := range errors {
            fmt.Printf("- %s\n", err.Error())
        }
    } else {
        fmt.Println("Contact is valid!")
    }
}
```
</details>

<details>
<summary><strong><code>Custom Validation Implementation</code></strong></summary>
<br/>

```go
package main

import (
    "fmt"
    "reflect"
    "strings"

    "github.com/mrz1836/go-validate"
)

// Custom validation for allowed colors
type colorValidation struct {
    validate.Validation
    allowedColors []string
}

func (c *colorValidation) Validate(value interface{}, _ reflect.Value) *validate.ValidationError {
    strValue, ok := value.(string)
    if !ok {
        return &validate.ValidationError{
            Key:     c.FieldName(),
            Message: "must be a string",
        }
    }

    // Check if color is in allowed list
    for _, color := range c.allowedColors {
        if strings.EqualFold(strValue, color) {
            return nil // Valid
        }
    }

    return &validate.ValidationError{
        Key:     c.FieldName(),
        Message: fmt.Sprintf("must be one of: %s", strings.Join(c.allowedColors, ", ")),
    }
}

// Builder function for the custom validation
func colorValidationBuilder(colors string, _ reflect.Kind) (validate.Interface, error) {
    allowedColors := strings.Split(colors, ",")
    for i, color := range allowedColors {
        allowedColors[i] = strings.TrimSpace(color)
    }

    return &colorValidation{
        allowedColors: allowedColors,
    }, nil
}

type Car struct {
    Make  string `validation:"min_length=1"`
    Model string `validation:"min_length=1"`
    Year  int    `validation:"min=1900 max=2024"`
    Color string `validation:"color=red,blue,green,black,white,silver"`
}

func main() {
    // Register our custom validation
    validate.AddValidation("color", colorValidationBuilder)
    validate.InitValidations()

    car := Car{
        Make:  "Toyota",
        Model: "Camry",
        Year:  2023,
        Color: "blue", // Must be one of the allowed colors
    }

    isValid, errors := validate.IsValid(car)
    if !isValid {
        fmt.Println("Car validation failed:")
        for _, err := range errors {
            fmt.Printf("- %s\n", err.Error())
        }
    } else {
        fmt.Println("Car is valid!")
    }
}
```
</details>

<details>
<summary><strong><code>Enum Validation Example</code></strong></summary>
<br/>

```go
package main

import (
    "fmt"

    "github.com/mrz1836/go-validate"
)

type OrderStatus string

const (
    StatusPending   OrderStatus = "pending"
    StatusShipped   OrderStatus = "shipped"
    StatusDelivered OrderStatus = "delivered"
    StatusCancelled OrderStatus = "cancelled"
)

type Order struct {
    ID       string
    Status   string
    Priority string
}

func (o *Order) Validate() (bool, []validate.ValidationError) {
    var errors []validate.ValidationError

    // Validate order status using enum validation
    allowedStatuses := []string{"pending", "shipped", "delivered", "cancelled"}
    if valid, err := validate.IsValidEnum(o.Status, &allowedStatuses, false); !valid {
        errors = append(errors, validate.ValidationError{
            Key:     "Status",
            Message: err.Error(),
        })
    }

    // Validate priority with empty allowed
    allowedPriorities := []string{"low", "medium", "high", "urgent"}
    if valid, err := validate.IsValidEnum(o.Priority, &allowedPriorities, true); !valid {
        errors = append(errors, validate.ValidationError{
            Key:     "Priority",
            Message: err.Error(),
        })
    }

    return len(errors) == 0, errors
}

func main() {
    order := Order{
        ID:       "ORD-12345",
        Status:   "shipped",    // Valid status
        Priority: "",           // Empty allowed for priority
    }

    isValid, errors := order.Validate()
    if !isValid {
        fmt.Println("Order validation failed:")
        for _, err := range errors {
            fmt.Printf("- %s\n", err.Error())
        }
    } else {
        fmt.Println("Order is valid!")
    }
}
```
</details>

<details>
<summary><strong><code>Complete Model Validation</code></strong></summary>
<br/>

```go
package main

import (
    "fmt"

    "github.com/mrz1836/go-validate"
)

// Customer model with comprehensive validation
type Customer struct {
    // Basic info with struct tags
    FirstName string `validation:"min_length=2 max_length=50"`
    LastName  string `validation:"min_length=2 max_length=50"`
    Email     string `validation:"format=email"`
    Age       uint8  `validation:"min=18 max=120"`

    // Address info
    Address string `validation:"min_length=10 max_length=200"`
    City    string `validation:"min_length=2 max_length=50"`
    State   string `validation:"min_length=2 max_length=2"` // US state codes
    ZipCode string `validation:"format=regexp:^[0-9]{5}(-[0-9]{4})?$"`

    // Contact info (validated separately)
    Phone               string `json:"phone"`
    SocialSecurityNumber string `json:"-"`

    // Account info
    AccountType    string  `validation:"min_length=1"`
    InitialBalance float64 `validation:"min=0"`
}

// Custom validation method that combines struct tags with utility functions
func (c *Customer) Validate() (bool, []validate.ValidationError) {
    // First run struct tag validations
    _, errors := validate.IsValid(*c)

    // Add phone number validation
    if valid, err := validate.IsValidPhoneNumber(c.Phone, "1"); !valid && c.Phone != "" {
        errors = append(errors, validate.ValidationError{
            Key:     "Phone",
            Message: err.Error(),
        })
    }

    // Add SSN validation
    if valid, err := validate.IsValidSocial(c.SocialSecurityNumber); !valid {
        errors = append(errors, validate.ValidationError{
            Key:     "SocialSecurityNumber",
            Message: err.Error(),
        })
    }

    // Add account type enum validation
    allowedTypes := []string{"checking", "savings", "business", "premium"}
    if valid, err := validate.IsValidEnum(c.AccountType, &allowedTypes, false); !valid {
        errors = append(errors, validate.ValidationError{
            Key:     "AccountType",
            Message: err.Error(),
        })
    }

    return len(errors) == 0, errors
}

func main() {
    validate.InitValidations()

    customer := Customer{
        FirstName:            "Alice",
        LastName:             "Johnson",
        Email:                "alice.johnson@email.com",
        Age:                  28,
        Address:              "123 Main Street, Apt 4B",
        City:                 "New York",
        State:                "NY",
        ZipCode:              "10001",
        Phone:                "555-123-4567",
        SocialSecurityNumber: "123-45-6789", // This will fail validation (blacklisted)
        AccountType:          "checking",
        InitialBalance:       1000.00,
    }

    isValid, errors := customer.Validate()
    if !isValid {
        fmt.Printf("Customer validation failed with %d errors:\n", len(errors))
        for i, err := range errors {
            fmt.Printf("%d. %s\n", i+1, err.Error())
        }
    } else {
        fmt.Println("Customer validation passed! Ready to save to database.")
    }
}
```
</details>

<br/>

## Benchmarks
Run the Go benchmarks:
```shell script
magex bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## 🤖 AI Usage & Assistant Guidelines
Read the [AI Usage & Assistant Guidelines](.github/tech-conventions/ai-compliance.md) for details on how AI is used in this project and how to interact with AI assistants.

<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |

<br/>

## Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap:
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-validate&utm_term=go-validate&utm_content=go-validate) to ensure this journey continues indefinitely! :rocket:


[![Stars](https://img.shields.io/github/stars/mrz1836/go-validate?label=Please%20like%20us&style=social)](https://github.com/mrz1836/go-validate/stargazers)

<br/>

## License

[![License](https://img.shields.io/github/license/mrz1836/go-validate.svg?style=flat)](LICENSE)
