# go-validate [![Build Status](https://travis-ci.org/BakedSoftware/go-validation.svg?branch=master)](https://travis-ci.org/BakedSoftware/go-validation)

Provides validations for struct fields based on a validation tag and offers additional validation functions.

See godoc for more info: http://godoc.org/github.com/BakedSoftware/go-validation

# Usage Examples

```
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
```

