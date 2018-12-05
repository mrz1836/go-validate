/*
Package validation provides validations for struct fields based on a validation tag and offers additional validation functions.

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.
*/
package validation

import (
	"reflect"
	"strconv"
)

//intValueValidation
type intValueValidation struct {
	//Validation is the validation interface
	Validation

	//Testing value to compare
	value int64

	//Boolean for determining if less (min) or not (max)
	less bool
}

//Validate validate function for the intValueValidation type
func (m *intValueValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {

	//Compare the value to see if it is convertible to type int64
	var compareValue int64
	switch value := value.(type) {
	case int:
		compareValue = int64(value)
	case int8:
		compareValue = int64(value)
	case int16:
		compareValue = int64(value)
	case int32:
		compareValue = int64(value)
	case int64:
		compareValue = int64(value)
	default:
		return &ValidationError{
			Key:     m.FieldName(),
			Message: "is not convertible to type int64",
		}
	}

	//Check min
	if m.less {
		if compareValue < m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be greater than or equal to " + strconv.FormatInt(m.value, 10),
			}
		}
	} else { //Check max
		if compareValue > m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be less than or equal to " + strconv.FormatInt(m.value, 10),
			}
		}
	}

	return nil
}

//uintValueValidation
type uintValueValidation struct {

	//Validation is the validation interface
	Validation

	//Testing value to compare
	value uint64

	//Boolean for determining if less (min) or not (max)
	less bool
}

//Validate validate function for the uintValueValidation type
func (m *uintValueValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {

	//Compare the value to see if it is convertible to type int64
	var compareValue uint64
	switch value := value.(type) {
	case uint:
		compareValue = uint64(value)
	case uint8:
		compareValue = uint64(value)
	case uint16:
		compareValue = uint64(value)
	case uint32:
		compareValue = uint64(value)
	case uint64:
		compareValue = uint64(value)
	default:
		return &ValidationError{
			Key:     m.FieldName(),
			Message: "is not convertible to type uint64",
		}
	}

	//Check min
	if m.less {
		if compareValue < m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be greater than or equal to " + strconv.FormatUint(m.value, 10),
			}
		}
	} else { //Check max
		if compareValue > m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be less than or equal to " + strconv.FormatUint(m.value, 10),
			}
		}
	}

	return nil
}

//floatValueValidation
type floatValueValidation struct {

	//Validation is the validation interface
	Validation

	//Testing value to compare
	value float64

	//Boolean for determining if less (min) or not (max)
	less bool
}

//Validate validate function for the floatValueValidation type
func (m *floatValueValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {

	//Compare the value to see if it is convertible to type int64
	var compareValue float64
	switch value := value.(type) {
	case float32:
		compareValue = float64(value)
	case float64:
		compareValue = float64(value)
	default:
		return &ValidationError{
			Key:     m.FieldName(),
			Message: "is not convertible to type float64",
		}
	}

	//Check min
	if m.less {
		if compareValue < m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be greater than or equal to " + strconv.FormatFloat(m.value, 'E', -1, 64),
			}
		}
	} else { //Check max
		if compareValue > m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be less than or equal to " + strconv.FormatFloat(m.value, 'E', -1, 64),
			}
		}
	}

	return nil
}

//newMinValueValidation
func newMinValueValidation(options string, kind reflect.Kind) (Interface, error) {
	switch kind {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value, err := strconv.ParseInt(options, 10, 0)
		if err != nil {
			return nil, err
		}
		return &intValueValidation{
			value: value,
			less:  true,
		}, nil
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value, err := strconv.ParseUint(options, 10, 0)
		if err != nil {
			return nil, err
		}
		return &uintValueValidation{
			value: value,
			less:  true,
		}, nil
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value, err := strconv.ParseFloat(options, 64)
		if err != nil {
			return nil, err
		}
		return &floatValueValidation{
			value: value,
			less:  true,
		}, nil
	default:
		return nil, &ValidationError{
			Key:     "invalid_validation",
			Message: "field is not of numeric type and min validation only accepts numeric types",
		}
	}
}

//newMaxValueValidation
func newMaxValueValidation(options string, kind reflect.Kind) (Interface, error) {
	switch kind {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value, err := strconv.ParseInt(options, 10, 0)
		if err != nil {
			return nil, err
		}
		return &intValueValidation{
			value: value,
			less:  false,
		}, nil
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value, err := strconv.ParseUint(options, 10, 0)
		if err != nil {
			return nil, err
		}
		return &uintValueValidation{
			value: value,
			less:  false,
		}, nil
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value, err := strconv.ParseFloat(options, 64)
		if err != nil {
			return nil, err
		}
		return &floatValueValidation{
			value: value,
			less:  false,
		}, nil
	default:
		return nil, &ValidationError{
			Key:     "invalid_validation",
			Message: "field is not of numeric type and max validation only accepts numeric types",
		}
	}
}

//init add the numeric validations
func init() {

	//Min validation is where X cannot be less then Y
	AddValidation("min", newMinValueValidation)

	//Max validation is where X cannot be greater than Y
	AddValidation("max", newMaxValueValidation)
}
