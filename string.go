/*
Package govalidation provides validations for struct fields based on a validation tag and offers additional validation functions.

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.
*/
package govalidation

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//Common regular expressions
var emailRegex = regexp.MustCompile(`(?i)^[a-z0-9._%+\-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?`)

//maxLengthStringValidation type used for string length values
type maxLengthStringValidation struct {
	//Validation is the validation interface
	Validation

	//The max string length value
	length int
}

//Validate is for the maxLengthStringValidation type and will test the max string length
func (v *maxLengthStringValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	strValue, ok := value.(string)
	if !ok {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "is not of type string and MaxLengthValidation only accepts strings",
		}
	}

	if len(strValue) > v.length {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "must be no more than " + strconv.Itoa(v.length) + " characters",
		}
	}

	return nil
}

//minLengthStringValidation type used for string length values
type minLengthStringValidation struct {
	//Validation is the validation interface
	Validation

	//The max string length value
	length int
}

//Validate is for the minLengthStringValidation type and will test the min string length
func (v *minLengthStringValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	strValue, ok := value.(string)
	if !ok {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "is not of type string and MinLengthValidation only accepts strings",
		}
	}

	if len(strValue) < v.length {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "must be at least " + strconv.Itoa(v.length) + " characters",
		}
	}

	return nil
}

//formatStringValidation type used for string pattern testing using regular expressions
type formatStringValidation struct {
	Validation
	pattern     *regexp.Regexp
	patternName string
}

//Validate is for the formatStringValidation type and will test the given regular expression
func (v *formatStringValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	strValue, ok := value.(string)
	if !ok {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "is not of type string and FormatValidation only accepts strings",
		}
	}

	if !v.pattern.MatchString(strValue) {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "does not match " + v.patternName + " format",
		}
	}

	return nil
}

//maxLengthValidation creates an interface based on the "kind" type
func maxLengthValidation(options string, kind reflect.Kind) (Interface, error) {
	length, err := strconv.ParseInt(options, 10, 0)
	if err != nil {
		return nil, err
	}

	return &maxLengthStringValidation{
		length: int(length),
	}, nil
}

//minLengthValidation creates an interface based on the "kind" type
func minLengthValidation(options string, kind reflect.Kind) (Interface, error) {
	length, err := strconv.ParseInt(options, 10, 0)
	if err != nil {
		return nil, err
	}

	return &minLengthStringValidation{
		length: int(length),
	}, nil
}

//formatValidation creates an interface based on the "kind" type
func formatValidation(options string, kind reflect.Kind) (Interface, error) {
	if strings.ToLower(options) == "email" {
		return &formatStringValidation{
			pattern:     emailRegex,
			patternName: "email",
		}, nil
	} else if strings.Contains(options, "regexp:") {
		patternStr := options[strings.Index(options, ":")+1:]
		pattern, err := regexp.Compile(patternStr)

		if err != nil {
			return nil, &ValidationError{Key: "regexp:", Message: err.Error()}
		}

		return &formatStringValidation{
			pattern:     pattern,
			patternName: "regexp",
		}, nil
	}

	return nil, &ValidationError{Key: "format", Message: "Has no pattern " + options}
}

//init add the string validations when this package is loaded
func init() {

	//Max length validation is len(string) < X
	AddValidation("max_length", maxLengthValidation)

	//Min length validation is len(string) > X
	AddValidation("min_length", minLengthValidation)

	//Format validation uses a given regular expression to match
	AddValidation("format", formatValidation)
}
