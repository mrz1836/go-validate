# go-validate
> Validations for struct fields based on a validation tag and offers additional validation functions

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-validate.svg?logo=github&style=flat)](https://github.com/mrz1836/go-validate/releases)
[![Build Status](https://travis-ci.com/mrz1836/go-validate.svg?branch=master)](https://travis-ci.com/mrz1836/go-validate)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-validate?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-validate)
[![codecov](https://codecov.io/gh/mrz1836/go-validate/branch/master/graph/badge.svg)](https://codecov.io/gh/mrz1836/go-validate)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-validate)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat)](https://mrz1818.com/?tab=tips&af=go-validate)

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

<br/>

## Installation

**go-validate** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-validate
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-validate)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-validate?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-validate)

<details>
<summary><strong><code>Library Deployment</code></strong></summary>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                            Runs lint, test-short and vet
bench                          Run all benchmarks in the Go application
clean                          Remove previous builds and any test cache data
clean-mods                     Remove all the Go mod cache
coverage                       Shows the test coverage
godocs                         Sync the latest tag with GoDocs
help                           Show all make commands available
lint                           Run the Go lint application
release                        Full production release (creates release in Github)
release-test                   Full production test release (everything except deploy)
release-snap                   Test the full release (build binaries)
run-examples                   Runs all the examples
tag                            Generate a new tag and push (IE: tag version=0.0.0)
tag-remove                     Remove a tag if found (IE: tag-remove version=0.0.0)
tag-update                     Update an existing tag to current commit (IE: tag-update version=0.0.0)
test                           Runs vet, lint and ALL tests
test-short                     Runs vet, lint and tests (excludes integration tests)
test-travis                    Runs tests via Travis (also exports coverage)
update                         Update all project dependencies
update-releaser                Update the goreleaser application
vet                            Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and examples run via [Travis CI](https://travis-ci.org/mrz1836/go-validate) and uses [Go version 1.14.x](https://golang.org/doc/go1.14). View the [deployment configuration file](.travis.yml).
```shell script
make test
```

<br/>

## Benchmarks
Run the Go benchmarks:
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

<br/>

## Usage
- [Full model example](examples/model/customer.go)
- [Numeric examples](numeric_test.go) or [string examples](string_test.go)
- [Numeric benchmarks](numeric_test.go) or [string benchmarks](string_test.go)
- [Numeric tests](numeric_test.go) or [string tests](string_test.go)
- [Generic tests](validate_test.go)

Basic model implementation:
```go
// ExampleModel shows inline validations via the struct tag
type ExampleModel struct {
    Age             uint    `validation:"min=18"`
    Category        string  `validation:"min_length=5 max_length=10"`
    Email           string  `validation:"format=email"`
    Name            string  `validation:"format=regexp:[A-Z][a-z]{3,12}"`
    Password        string  `validation:"compare=PasswordConfirm"`
    PasswordConfirm string  `json:"-"`
    Quantity        uint    `validation:"min=1 max=5"`
    Total           float32 `validation:"min=0"`
}

// Example showing extra validation functions for additional use
ok, err := validate.IsValidEmail("someones@email.com")
if !ok {
    errs = append(errs, validate.ValidationError{
        Key:     "Email",
        Message: err.Error(),
    })
}
```

<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) | [<img src="https://github.com/kayleg.png" height="50" alt="kayleg" />](https://github.com/kayleg) |
|:---:|:---:|
| [MrZ](https://github.com/mrz1836) | [kayleg](https://github.com/kayleg) |

<br/>

## Contributing
View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&af=go-validate) to ensure this journey continues indefinitely! :rocket:

<br/>

## License

![License](https://img.shields.io/github/license/mrz1836/go-validate.svg?style=flat)