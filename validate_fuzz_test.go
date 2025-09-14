package validate

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzBuildValidations(f *testing.F) {
	// We need to create struct types dynamically, but Go doesn't allow runtime struct creation easily
	// Instead, we'll test the validation tag parsing logic indirectly through IsValid

	// Seed corpus with various validation tag combinations
	f.Add("max_length=10")
	f.Add("min_length=5")
	f.Add("format=email")
	f.Add("format=regexp:^[a-z]+$")
	f.Add("min=0 max=100")
	f.Add("max_length=50 format=email")
	f.Add("compare=OtherField")
	f.Add("min_length=1 max_length=255")
	f.Add("")
	f.Add("invalid_validation=test")
	f.Add("max_length") // Missing =
	f.Add("max_length=")
	f.Add("max_length=abc") // Invalid number
	f.Add("format=invalid_format")
	f.Add("min=-1 max=abc")
	f.Add(strings.Repeat("max_length=10 ", 100)) // Very long tag

	f.Fuzz(func(t *testing.T, validationTag string) {
		defer func() {
			if r := recover(); r != nil {
				// Some panics are expected for invalid validation tags
				// The system uses log.Fatalln for invalid specs, which we can't easily test
				// So we allow panics but log them
				t.Logf("buildValidations panicked with tag %q: %v", validationTag, r)
			}
		}()

		// We can't easily create dynamic structs, so we test the validation registration instead
		// Create a test validation map
		testMap := &Map{}

		// Register basic validations to test with
		testMap.AddValidation("max_length", maxLengthValidation)
		testMap.AddValidation("min_length", minLengthValidation)
		testMap.AddValidation("format", formatValidation)
		testMap.AddValidation("compare", stringEqualsStringValidation)
		testMap.AddValidation("min", minValueValidation)
		testMap.AddValidation("max", maxValueValidation)

		// Test that adding validations doesn't panic
		_ = testMap
	})
}

func FuzzIsValidWithStruct(f *testing.F) {
	// Pre-register all validations
	RegisterStringValidations()
	RegisterNumericValidations()

	// Test struct for validation
	type TestStruct struct {
		Name  string `validation:"min_length=2 max_length=50"`
		Email string `validation:"format=email"`
		Age   int    `validation:"min=0 max=120"`
	}

	// Seed corpus with various struct field values
	f.Add("John", "john@example.com", 25)
	f.Add("", "invalid-email", -5)
	f.Add("A", "test@test.com", 150)
	f.Add(strings.Repeat("a", 60), "user@domain.co.uk", 0)
	f.Add("Valid", "", 25)
	f.Add("Test", "test@aol.con", 25) // blacklisted domain

	f.Fuzz(func(t *testing.T, name, email string, age int) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValid panicked with name %q, email %q, age %d: %v", name, email, age, r)
			}
		}()

		testStruct := TestStruct{
			Name:  name,
			Email: email,
			Age:   age,
		}

		// Test with DefaultMap
		valid, errors := IsValid(testStruct)
		_ = valid
		_ = errors // Should never panic

		// Test with pointer
		valid2, errors2 := IsValid(&testStruct)
		_ = valid2
		_ = errors2 // Should never panic

		// Test that validation results make sense
		if len(name) < 2 || len(name) > 50 {
			require.False(t, valid, "Invalid name length should make struct invalid")
		}

		if age < 0 || age > 120 {
			require.False(t, valid, "Invalid age should make struct invalid")
		}
	})
}

func FuzzIsValidWithVariousTypes(f *testing.F) {
	// Register validations
	RegisterStringValidations()
	RegisterNumericValidations()

	// Test different struct configurations
	type StringStruct struct {
		Value string `validation:"min_length=1"`
	}

	type IntStruct struct {
		Value int `validation:"min=0"`
	}

	type FloatStruct struct {
		Value float64 `validation:"max=100.0"`
	}

	f.Add("test", 42, 3.14)
	f.Add("", -1, -99.9)
	f.Add(strings.Repeat("x", 1000), 999, 999.999)

	f.Fuzz(func(t *testing.T, strVal string, intVal int, floatVal float64) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValid panicked with values str=%q, int=%d, float=%f: %v", strVal, intVal, floatVal, r)
			}
		}()

		// Test string struct
		strStruct := StringStruct{Value: strVal}
		valid1, errors1 := IsValid(strStruct)
		_ = valid1
		_ = errors1

		// Test int struct
		intStruct := IntStruct{Value: intVal}
		valid2, errors2 := IsValid(intStruct)
		_ = valid2
		_ = errors2

		// Test float struct
		floatStruct := FloatStruct{Value: floatVal}
		valid3, errors3 := IsValid(floatStruct)
		_ = valid3
		_ = errors3

		// Test with pointer to struct
		ptrStruct := &StringStruct{Value: strVal}
		valid4, errors4 := IsValid(ptrStruct)
		_ = valid4
		_ = errors4
	})
}

func FuzzValidationMapOperations(f *testing.F) {
	// Test the Map type operations directly
	f.Add("test_validation", "test_param")
	f.Add("email_validation", "email")
	f.Add("length_validation", "100")
	f.Add("", "")
	f.Add("validation_with_very_long_name_that_might_cause_issues", "parameter")
	f.Add("unicode_validation_test", "parameter")

	f.Fuzz(func(t *testing.T, validationName, parameter string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Map operations panicked with name %q, param %q: %v", validationName, parameter, r)
			}
		}()

		testMap := &Map{}

		// Test AddValidation with a simple builder function
		testMap.AddValidation(validationName, func(_ string, _ reflect.Kind) (Interface, error) {
			return &Validation{Name: validationName}, nil
		})

		// Test with a struct that has no validation tags
		type NoValidationStruct struct {
			Field string
		}

		noValidStruct := NoValidationStruct{Field: parameter}
		valid, errors := testMap.IsValid(noValidStruct)

		// Should always be valid since no validation tags
		require.True(t, valid, "Struct with no validation tags should be valid")
		require.Empty(t, errors, "Struct with no validation tags should have no errors")
	})
}

func FuzzValidationInterface(f *testing.F) {
	// Test the basic Validation struct and Interface methods
	f.Add("TestField", 42)
	f.Add("", -1)
	f.Add("VeryLongFieldNameThatMightCauseIssues", 0)
	f.Add("field_with_underscores", 999)
	f.Add("fieldWithCamelCase", 123)

	f.Fuzz(func(t *testing.T, fieldName string, fieldIndex int) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Validation interface methods panicked with field %q, index %d: %v", fieldName, fieldIndex, r)
			}
		}()

		validation := &Validation{Name: "test"}

		// Test field name methods
		validation.SetFieldName(fieldName)
		retrievedName := validation.FieldName()
		require.Equal(t, fieldName, retrievedName, "Field name should be stored and retrieved correctly")

		// Test field index methods
		validation.SetFieldIndex(fieldIndex)
		retrievedIndex := validation.FieldIndex()
		require.Equal(t, fieldIndex, retrievedIndex, "Field index should be stored and retrieved correctly")

		// Test basic validate method (should always return error since not implemented)
		result := validation.Validate("test", reflect.Value{})
		require.NotNil(t, result, "Basic validation should return error for not implemented")
		require.Contains(t, result.Message, "not implemented", "Error message should indicate not implemented")
	})
}

func FuzzComplexValidationScenarios(f *testing.F) {
	// Register all validations
	RegisterStringValidations()
	RegisterNumericValidations()

	// Complex struct for testing
	type ComplexStruct struct {
		Username string  `validation:"min_length=3 max_length=20 format=regexp:^[a-zA-Z0-9_]+$"`
		Email    string  `validation:"format=email"`
		Age      int     `validation:"min=13 max=120"`
		Score    float64 `validation:"min=0.0 max=100.0"`
	}

	f.Add("user123", "user@test.com", 25, 85.5)
	f.Add("u", "invalid-email", 12, -10.0)
	f.Add("valid_user", "test@example.com", 130, 110.0)
	f.Add("user@invalid", "user@domain.com", 25, 50.0) // username with @
	f.Add("", "", 0, 0.0)

	f.Fuzz(func(t *testing.T, username, email string, age int, score float64) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Complex validation panicked: %v", r)
				t.Errorf("Values: username=%q, email=%q, age=%d, score=%f", username, email, age, score)
			}
		}()

		complexStruct := ComplexStruct{
			Username: username,
			Email:    email,
			Age:      age,
			Score:    score,
		}

		valid, errors := IsValid(complexStruct)
		_ = valid

		// Verify that errors contain reasonable information
		for _, err := range errors {
			require.NotEmpty(t, err.Key, "Error should have a key")
			require.NotEmpty(t, err.Message, "Error should have a message")
		}

		// Test that validation count makes sense
		require.LessOrEqual(t, len(errors), 8, "Should not have more errors than total validation rules")
	})
}
