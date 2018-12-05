/*
Package govalidation provides validations for struct fields based on a validation tag and offers additional validation functions.

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.
*/
package govalidation

import (
	"reflect"
	"testing"
)

//invalidNumericTypes is for the types not allowed (numeric tests)
var invalidNumericTypes []reflect.Kind

//init load the default invalid types
func init() {
	invalidNumericTypes = append(
		invalidNumericTypes,
		reflect.Array,
		reflect.Bool,
		reflect.Chan,
		reflect.Complex128,
		reflect.Complex64,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.String,
		reflect.Struct,
		reflect.UnsafePointer,
	)
}

//
// Generic tests
//

//TestMinValueValidation
func TestMinValueValidation(t *testing.T) {

	//Test invalid types
	var err error
	for i := 0; i < len(invalidNumericTypes); i++ {
		_, err = minValueValidation("10", invalidNumericTypes[i])
		if err == nil {
			t.Fatal("Expected error - cannot use: ", invalidNumericTypes[i])
		}
	}

	//Fail if string submitted or Parse int fails
	_, err = minValueValidation("foo", reflect.Int)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}

	//Fail if string submitted or Parse uint fails
	_, err = minValueValidation("foo", reflect.Uint64)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}

	//Fail if string submitted or Parse float fails
	_, err = minValueValidation("foo", reflect.Float32)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}
}

//TestMaxValueValidation
func TestMaxValueValidation(t *testing.T) {

	//Test invalid types
	var err error
	for i := 0; i < len(invalidNumericTypes); i++ {
		_, err = maxValueValidation("10", invalidNumericTypes[i])
		if err == nil {
			t.Fatal("Expected error - cannot use: ", invalidNumericTypes[i])
		}
	}

	//Fail if string submitted or Parse int fails
	_, err = maxValueValidation("foo", reflect.Int)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}

	//Fail if string submitted or Parse uint fails
	_, err = maxValueValidation("foo", reflect.Uint64)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}

	//Fail if string submitted or Parse float fails
	_, err = maxValueValidation("foo", reflect.Float32)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}
}

//
// Integer tests
//

//TestMinValueInt8Positive tests min value on int8
func TestMinValueInt8Positive(t *testing.T) {
	type minValueTestType struct {
		Value int8 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueInt8Negative tests min value on int8
func TestMinValueInt8Negative(t *testing.T) {
	type minValueTestType struct {
		Value int8 `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMinValueInt16Positive tests min value on int16
func TestMinValueInt16Positive(t *testing.T) {
	type minValueTestType struct {
		Value int16 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueInt16Negative tests min value on int16
func TestMinValueInt16Negative(t *testing.T) {
	type minValueTestType struct {
		Value int16 `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMinValueInt32Positive tests min value on int32
func TestMinValueInt32Positive(t *testing.T) {
	type minValueTestType struct {
		Value int32 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueInt32Negative tests min value on int32
func TestMinValueInt32Negative(t *testing.T) {
	type minValueTestType struct {
		Value int32 `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMinValueInt64Positive tests min value on int64
func TestMinValueInt64Positive(t *testing.T) {
	type minValueTestType struct {
		Value int64 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueInt64Negative tests min value on int64
func TestMinValueInt64Negative(t *testing.T) {
	type minValueTestType struct {
		Value int64 `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMinValueIntPositive tests min value on int
func TestMinValueIntPositive(t *testing.T) {
	type minValueTestType struct {
		Value int `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueIntNegative tests min value on int
func TestMinValueIntNegative(t *testing.T) {
	type minValueTestType struct {
		Value int `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMaxValueInt8Positive tests max value on int8
func TestMaxValueInt8Positive(t *testing.T) {
	type maxValueTestType struct {
		Value int8 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueInt8Negative tests max value on int8
func TestMaxValueInt8Negative(t *testing.T) {
	type maxValueTestType struct {
		Value int8 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Empty Int8(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}

}

//TestMaxValueInt16Positive tests max value on int16
func TestMaxValueInt16Positive(t *testing.T) {
	type maxValueTestType struct {
		Value int16 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueInt16Negative tests max value on int16
func TestMaxValueInt16Negative(t *testing.T) {
	type maxValueTestType struct {
		Value int16 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Empty Int16(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}

}

//TestMaxValueInt32Positive tests max value on int32
func TestMaxValueInt32Positive(t *testing.T) {
	type maxValueTestType struct {
		Value int32 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueInt32Negative tests max value on int32
func TestMaxValueInt32Negative(t *testing.T) {
	type maxValueTestType struct {
		Value int32 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Empty Int32(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

//TestMaxValueInt64Positive tests max value on int64
func TestMaxValueInt64Positive(t *testing.T) {
	type maxValueTestType struct {
		Value int64 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueInt64Negative tests max value on int64
func TestMaxValueInt64Negative(t *testing.T) {
	type maxValueTestType struct {
		Value int64 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Empty Int64(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

//TestMaxValueIntPositive tests max value on int
func TestMaxValueIntPositive(t *testing.T) {
	type maxValueTestType struct {
		Value int `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueIntNegative tests max value on int
func TestMaxValueIntNegative(t *testing.T) {
	type maxValueTestType struct {
		Value int `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Empty Int(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

//
// Unsigned integer tests
//

//TestMinValueUint8Positive tests min value on uint8
func TestMinValueUint8Positive(t *testing.T) {
	type minValueTestType struct {
		Value uint8 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint8Negative tests min value on uint8
func TestMinValueUint8Negative(t *testing.T) {
	type minValueTestType struct {
		Value uint8 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint16Positive tests min value on uint16
func TestMinValueUint16Positive(t *testing.T) {
	type minValueTestType struct {
		Value uint16 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint16Negative tests min value on uint16
func TestMinValueUint16Negative(t *testing.T) {
	type minValueTestType struct {
		Value uint16 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint32Positive tests min value on uint32
func TestMinValueUint32Positive(t *testing.T) {
	type minValueTestType struct {
		Value uint32 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint32Negative tests min value on uint32
func TestMinValueUint32Negative(t *testing.T) {
	type minValueTestType struct {
		Value uint32 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint64Positive tests min value on uint64
func TestMinValueUint64Positive(t *testing.T) {
	type minValueTestType struct {
		Value uint64 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint64Negative tests min value on uint64
func TestMinValueUint64Negative(t *testing.T) {
	type minValueTestType struct {
		Value uint64 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUintPositive tests min value on uint
func TestMinValueUintPositive(t *testing.T) {
	type minValueTestType struct {
		Value uint `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as min was set to 20", obj.Value)
	}

	obj.Value = 19

	ok, _ = IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUintNegative tests min value on uint
func TestMinValueUintNegative(t *testing.T) {
	type minValueTestType struct {
		Value uint `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUint8Positive tests max value on uint8
func TestMaxValueUint8Positive(t *testing.T) {
	type maxValueTestType struct {
		Value uint8 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUint8Negative tests max value on uint8
func TestMaxValueUint8Negative(t *testing.T) {
	type maxValueTestType struct {
		Value uint8 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUint16Positive tests max value on uint16
func TestMaxValueUint16Positive(t *testing.T) {
	type maxValueTestType struct {
		Value uint16 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUint16Negative tests max value on uint16
func TestMaxValueUint16Negative(t *testing.T) {
	type maxValueTestType struct {
		Value uint16 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUint32Positive tests max value on uint32
func TestMaxValueUint32Positive(t *testing.T) {
	type maxValueTestType struct {
		Value uint32 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUint32Negative tests max value on uint32
func TestMaxValueUint32Negative(t *testing.T) {
	type maxValueTestType struct {
		Value uint32 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUint64Positive tests max value on uint64
func TestMaxValueUint64Positive(t *testing.T) {
	type maxValueTestType struct {
		Value uint64 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUint64Negative tests max value on uint64
func TestMaxValueUint64Negative(t *testing.T) {
	type maxValueTestType struct {
		Value uint64 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUintPositive tests max value on uint
func TestMaxValueUintPositive(t *testing.T) {
	type maxValueTestType struct {
		Value uint `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUintNegative tests max value on uint
func TestMaxValueUintNegative(t *testing.T) {
	type maxValueTestType struct {
		Value uint `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//
// Float tests
//

//TestMinValueFloat32 tests min value on float32
func TestMinValueFloat32(t *testing.T) {
	type minValueTestType struct {
		Value float32 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueFloat64 tests min value on float64
func TestMinValueFloat64(t *testing.T) {
	type minValueTestType struct {
		Value float64 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueFloat32 tests max value on float32
func TestMaxValueFloat32(t *testing.T) {
	type maxValueTestType struct {
		Value float32 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than max -20")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueFloat64 tests max value on float64
func TestMaxValueFloat64(t *testing.T) {
	type maxValueTestType struct {
		Value float64 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)
	if ok {
		t.Fatal("Expected failure as 0 is less than max -20")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}
