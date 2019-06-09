# go-validate
**go-validate** provides validations for struct fields based on a validation tag and offers additional validation functions.

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-validate.svg?style=flat) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-validate?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-validate)  | [![Codacy Badge](https://api.codacy.com/project/badge/Grade/e25f7c37ecb246fba1cabf1000aa76a3)](https://www.codacy.com/app/mrz1818/go-validate?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-validate&amp;utm_campaign=Badge_Grade) |  [![Build Status](https://travis-ci.com/mrz1836/go-validate.svg?branch=master)](https://travis-ci.com/mrz1836/go-validate)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-validate.svg?style=flat)](https://github.com/mrz1836/go-validate/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-validate?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-validate) |

## Table of Contents
- [Installation](https://github.com/mrz1836/go-validate#installation)
- [Documentation](https://github.com/mrz1836/go-validate#documentation)
- [Examples & Tests](https://github.com/mrz1836/go-validate#examples--tests)
- [Benchmarks](https://github.com/mrz1836/go-validate#benchmarks)
- [Code Standards](https://github.com/mrz1836/go-validate#code-standards)
- [Usage](https://github.com/mrz1836/go-validate#usage)
- [Maintainers](https://github.com/mrz1836/go-validate#maintainers)
- [Contributing](https://github.com/mrz1836/go-validate#contributing)
- [License](https://github.com/mrz1836/go-validate#license)

## Installation

**go-validate** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/mrz1836/go-validate
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-validate).

## Examples & Tests
All unit tests and examples run via [Travis CI](https://travis-ci.com/mrz1836/go-validate) and uses [Go version 1.12.x](https://golang.org/doc/go1.12). View the [deployment configuration file](https://github.com/mrz1836/go-validate/blob/master/.travis.yml).
```bash
$ cd ../go-validate
$ go test ./... -v
```

## Benchmarks
Run the Go benchmarks:
```bash
$ cd ../go-validate
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](https://github.com/mrz1836/go-validate/blob/master/CODE_STANDARDS.md).

## Usage
- View the [full model example](https://github.com/mrz1836/go-validate/blob/master/examples/model/customer.go)
- View the [numeric examples](https://github.com/mrz1836/go-validate/blob/master/numeric_test.go) or [string examples](https://github.com/mrz1836/go-validate/blob/master/string_test.go)
- View the [numeric benchmarks](https://github.com/mrz1836/go-validate/blob/master/numeric_test.go) or [string benchmarks](https://github.com/mrz1836/go-validate/blob/master/string_test.go)
- View the [numeric tests](https://github.com/mrz1836/go-validate/blob/master/numeric_test.go) or [string tests](https://github.com/mrz1836/go-validate/blob/master/string_test.go)
- View the [generic tests](https://github.com/mrz1836/go-validate/blob/master/validate_test.go)

Basic model implementation:
```golang

//ExampleModel shows inline validations via the struct tag
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

//Example showing extra validation functions for additional use
ok, err := validate.IsValidEmail("someones@email.com")
if !ok {
    errs = append(errs, validate.ValidationError{
        Key:     "Email",
        Message: err.Error(),
    })
}
```

## Maintainers

[@kayleg](https://github.com/kayleg) - [@MrZ1836](https://github.com/mrz1836)

## Contributing

View the [contributing guidelines](https://github.com/mrz1836/go-validate/blob/master/CONTRIBUTING.md) and follow the [code of conduct](https://github.com/mrz1836/go-validate/blob/master/CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-validate)

## License

![License](https://img.shields.io/github/license/mrz1836/go-validate.svg?style=flat)