package validate

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzMinValueValidation(f *testing.F) {
	// Seed corpus with various numeric values and edge cases
	f.Add("0")
	f.Add("1")
	f.Add("-1")
	f.Add("100")
	f.Add("-100")
	f.Add("9223372036854775807")  // max int64
	f.Add("-9223372036854775808") // min int64
	f.Add("18446744073709551615") // max uint64
	f.Add("3.14159")
	f.Add("-3.14159")
	f.Add("1.7976931348623157e+308") // max float64
	f.Add("2.2250738585072014e-308") // min positive float64
	f.Add("inf")
	f.Add("-inf")
	f.Add("nan")
	f.Add("abc")
	f.Add("")
	f.Add("1e10")
	f.Add("1e-10")

	f.Fuzz(func(t *testing.T, minValue string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("minValueValidation panicked with input %q: %v", minValue, r)
			}
		}()

		// Test with different numeric kinds
		kinds := []reflect.Kind{
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.String, // Should return error
		}

		for _, kind := range kinds {
			validation, err := minValueValidation(minValue, kind)

			// If validation creation succeeds, test it
			if err == nil && validation != nil {
				switch v := validation.(type) {
				case *intValueValidation:
					testIntValidation(t, v, true)
				case *uintValueValidation:
					testUintValidation(t, v, true)
				case *floatValueValidation:
					testFloatValidation(t, v, true)
				}
			}
		}
	})
}

func FuzzMaxValueValidation(f *testing.F) {
	// Seed corpus with various numeric values and edge cases
	f.Add("0")
	f.Add("1")
	f.Add("-1")
	f.Add("1000")
	f.Add("-1000")
	f.Add("9223372036854775807")  // max int64
	f.Add("-9223372036854775808") // min int64
	f.Add("18446744073709551615") // max uint64
	f.Add("2.71828")
	f.Add("-2.71828")
	f.Add("1.7976931348623157e+308") // max float64
	f.Add("-1.7976931348623157e+308")
	f.Add("inf")
	f.Add("-inf")
	f.Add("nan")
	f.Add("invalid")
	f.Add("")
	f.Add("1.23e45")

	f.Fuzz(func(t *testing.T, maxValue string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("maxValueValidation panicked with input %q: %v", maxValue, r)
			}
		}()

		// Test with different numeric kinds
		kinds := []reflect.Kind{
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Bool, // Should return error
		}

		for _, kind := range kinds {
			validation, err := maxValueValidation(maxValue, kind)

			// If validation creation succeeds, test it
			if err == nil && validation != nil {
				switch v := validation.(type) {
				case *intValueValidation:
					testIntValidation(t, v, false)
				case *uintValueValidation:
					testUintValidation(t, v, false)
				case *floatValueValidation:
					testFloatValidation(t, v, false)
				}
			}
		}
	})
}

func testIntValidation(t *testing.T, validation *intValueValidation, _ bool) {
	// Test with various integer types and values
	testValues := []interface{}{
		int(0), int(1), int(-1), int(100), int(-100),
		int8(0), int8(1), int8(-1), int8(127), int8(-128),
		int16(0), int16(1), int16(-1), int16(32767), int16(-32768),
		int32(0), int32(1), int32(-1), int32(2147483647), int32(-2147483648),
		int64(0), int64(1), int64(-1), int64(9223372036854775807), int64(-9223372036854775808),
		"not an int", 3.14, true, nil,
		[]int{1, 2, 3},
	}

	for _, val := range testValues {
		result := validation.Validate(val, reflect.Value{})
		// Should never panic
		_ = result

		// Verify type checking works correctly for non-integer types
		switch val.(type) {
		case int, int8, int16, int32, int64:
			// Should not error for type conversion
		default:
			// Should error for non-integer types
			require.NotNil(t, result, "Expected error for non-integer type %T", val)
		}
	}
}

func testUintValidation(t *testing.T, validation *uintValueValidation, _ bool) {
	// Test with various unsigned integer types and values
	testValues := []interface{}{
		uint(0), uint(1), uint(100), uint(4294967295),
		uint8(0), uint8(1), uint8(255),
		uint16(0), uint16(1), uint16(65535),
		uint32(0), uint32(1), uint32(4294967295),
		uint64(0), uint64(1), uint64(18446744073709551615),
		-1, "not a uint", 2.718, false,
		map[string]int{},
	}

	for _, val := range testValues {
		result := validation.Validate(val, reflect.Value{})
		// Should never panic
		_ = result

		// Verify type checking works correctly for non-uint types
		switch val.(type) {
		case uint, uint8, uint16, uint32, uint64:
			// Should not error for type conversion
		default:
			// Should error for non-uint types
			require.NotNil(t, result, "Expected error for non-uint type %T", val)
		}
	}
}

func testFloatValidation(t *testing.T, validation *floatValueValidation, _ bool) {
	// Test with various float types and values
	testValues := []interface{}{
		float32(0.0), float32(1.0), float32(-1.0), float32(3.14), float32(-3.14),
		float32(math.MaxFloat32), float32(-math.MaxFloat32), float32(math.SmallestNonzeroFloat32),
		float64(0.0), float64(1.0), float64(-1.0), float64(2.718), float64(-2.718),
		float64(math.MaxFloat64), float64(-math.MaxFloat64), float64(math.SmallestNonzeroFloat64),
		float64(math.Inf(1)), float64(math.Inf(-1)), float64(math.NaN()),
		42, "not a float", true,
		[]float64{1.0, 2.0},
	}

	for _, val := range testValues {
		result := validation.Validate(val, reflect.Value{})
		// Should never panic
		_ = result

		// Verify type checking works correctly for non-float types
		switch val.(type) {
		case float32, float64:
			// Should not error for type conversion
		default:
			// Should error for non-float types
			require.NotNil(t, result, "Expected error for non-float type %T", val)
		}
	}
}

func verifyIntValidationResult(t *testing.T, result error, testInt, value int64, isMin bool) {
	if isMin {
		if testInt >= value {
			require.NoError(t, result, "Expected nil for %d >= %d (min)", testInt, value)
		} else {
			require.Error(t, result, "Expected error for %d < %d (min)", testInt, value)
		}
		return
	}
	if testInt <= value {
		require.NoError(t, result, "Expected nil for %d <= %d (max)", testInt, value)
	} else {
		require.Error(t, result, "Expected error for %d > %d (max)", testInt, value)
	}
}

func FuzzIntValueValidationDirectly(f *testing.F) {
	// Seed corpus for direct integer validation testing
	f.Add(int64(0), true)
	f.Add(int64(1), false)
	f.Add(int64(-1), true)
	f.Add(int64(9223372036854775807), false)
	f.Add(int64(-9223372036854775808), true)

	f.Fuzz(func(t *testing.T, value int64, isMin bool) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("intValueValidation panicked with value %d, isMin %t: %v", value, isMin, r)
			}
		}()

		validation := &intValueValidation{
			value: value,
			less:  isMin,
		}
		validation.SetFieldName("testField")

		// Test with various integer values
		testInts := []int64{
			math.MinInt64, math.MaxInt64, 0, 1, -1, 100, -100, value - 1, value, value + 1,
		}

		for _, testInt := range testInts {
			result := validation.Validate(testInt, reflect.Value{})
			// Should never panic
			_ = result

			// Verify logic - convert *ValidationError to error interface
			var err error
			if result != nil {
				err = result
			}
			verifyIntValidationResult(t, err, testInt, value, isMin)
		}
	})
}

func verifyUintValidationResult(t *testing.T, result error, testUint, value uint64, isMin bool) {
	if isMin {
		if testUint >= value {
			require.NoError(t, result, "Expected nil for %d >= %d (min)", testUint, value)
		} else {
			require.Error(t, result, "Expected error for %d < %d (min)", testUint, value)
		}
		return
	}
	if testUint <= value {
		require.NoError(t, result, "Expected nil for %d <= %d (max)", testUint, value)
	} else {
		require.Error(t, result, "Expected error for %d > %d (max)", testUint, value)
	}
}

func generateUintTestValues(value uint64) []uint64 {
	testUints := []uint64{0, math.MaxUint64, 1, 100, value}

	// Add boundary values safely
	if value > 0 {
		testUints = append(testUints, value-1)
	}
	if value < math.MaxUint64 {
		testUints = append(testUints, value+1)
	}

	return testUints
}

func FuzzUintValueValidationDirectly(f *testing.F) {
	// Seed corpus for direct unsigned integer validation testing
	f.Add(uint64(0), true)
	f.Add(uint64(1), false)
	f.Add(uint64(18446744073709551615), false) // max uint64

	f.Fuzz(func(t *testing.T, value uint64, isMin bool) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("uintValueValidation panicked with value %d, isMin %t: %v", value, isMin, r)
			}
		}()

		validation := &uintValueValidation{
			value: value,
			less:  isMin,
		}
		validation.SetFieldName("testField")

		// Test with various unsigned integer values
		testUints := generateUintTestValues(value)

		for _, testUint := range testUints {
			result := validation.Validate(testUint, reflect.Value{})
			// Should never panic
			_ = result

			// Verify logic - convert *ValidationError to error interface
			var err error
			if result != nil {
				err = result
			}
			verifyUintValidationResult(t, err, testUint, value, isMin)
		}
	})
}

func verifyFloatValidationResult(t *testing.T, result error, testFloat, value float64, isMin bool) {
	if isMin {
		if testFloat >= value {
			require.NoError(t, result, "Expected nil for %f >= %f (min)", testFloat, value)
		} else {
			require.Error(t, result, "Expected error for %f < %f (min)", testFloat, value)
		}
		return
	}
	if testFloat <= value {
		require.NoError(t, result, "Expected nil for %f <= %f (max)", testFloat, value)
	} else {
		require.Error(t, result, "Expected error for %f > %f (max)", testFloat, value)
	}
}

func generateFloatTestValues(value float64) []float64 {
	return []float64{
		-math.MaxFloat64, math.MaxFloat64, 0.0, 1.0, -1.0,
		math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64,
		value, value - 1.0, value + 1.0,
	}
}

func validateFloatTestValues(t *testing.T, validation *floatValueValidation, testFloats []float64, value float64, isMin bool) {
	for _, testFloat := range testFloats {
		if math.IsNaN(testFloat) || math.IsInf(testFloat, 0) {
			continue
		}

		result := validation.Validate(testFloat, reflect.Value{})
		// Should never panic
		_ = result

		// Verify logic - convert *ValidationError to error interface
		var err error
		if result != nil {
			err = result
		}
		verifyFloatValidationResult(t, err, testFloat, value, isMin)
	}
}

func FuzzFloatValueValidationDirectly(f *testing.F) {
	// Seed corpus for direct float validation testing
	f.Add(0.0, true)
	f.Add(1.0, false)
	f.Add(-1.0, true)
	f.Add(3.14159, false)
	f.Add(-2.71828, true)

	f.Fuzz(func(t *testing.T, value float64, isMin bool) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("floatValueValidation panicked with value %f, isMin %t: %v", value, isMin, r)
			}
		}()

		// Skip invalid float values that could cause issues
		if math.IsNaN(value) || math.IsInf(value, 0) {
			return
		}

		validation := &floatValueValidation{
			value: value,
			less:  isMin,
		}
		validation.SetFieldName("testField")

		// Test with various float values
		testFloats := generateFloatTestValues(value)
		validateFloatTestValues(t, validation, testFloats, value, isMin)
	})
}
