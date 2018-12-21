package validate

import (
	"fmt"
	"reflect"
	"testing"
)

//
// Generic numeric struct and function tests
//

//TestMinValueValidation - series of different tests
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

	//Test making an interface
	minInterface, err := minValueValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := minInterface.Validate(8, testVal)
	if errs == nil {
		t.Fatal("Expected to fail, 8 < 10")
	}

	//Test converting a string
	errs = minInterface.Validate("ddd", testVal)
	if errs == nil {
		t.Fatal("Expected to fail, value is not convertible to type int64")
	}

	//Test making an interface
	minInterface, err = minValueValidation("10", reflect.Float32)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Test in converting to float
	errs = minInterface.Validate("ddd", testVal)
	if errs == nil {
		t.Fatal("Expected to fail, value is not convertible to type float")
	}
}

//TestMaxValueValidation - series of different tests
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

	//Test making an interface
	maxInterface, err := maxValueValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := maxInterface.Validate(14, testVal)
	if errs == nil {
		t.Fatal("Expected to fail, 14 > 10")
	}

	//Test converting a string
	errs = maxInterface.Validate("ddd", testVal)
	if errs == nil {
		t.Fatal("Expected to fail, value is not convertible to type int64")
	}

	//Test making an interface
	maxInterface, err = maxValueValidation("10", reflect.Float32)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Test converting a string
	errs = maxInterface.Validate("ddd", testVal)
	if errs == nil {
		t.Fatal("Expected to fail, value is not convertible to type float")
	}
}

//BenchmarkTestMinIntValue benchmarks the Min Value (valid value)
func BenchmarkTestMinIntValue(b *testing.B) {
	type testModel struct {
		Value int8 `validation:"min=20"`
	}
	model := testModel{
		Value: 21,
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

//BenchmarkTestMinUintValue benchmarks the Min Value (valid value)
func BenchmarkTestMinUintValue(b *testing.B) {
	type testModel struct {
		Value uint8 `validation:"min=20"`
	}
	model := testModel{
		Value: 21,
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

//BenchmarkTestMinFloatValue benchmarks the Min Value (valid value)
func BenchmarkTestMinFloatValue(b *testing.B) {
	type testModel struct {
		Value float32 `validation:"min=20"`
	}
	model := testModel{
		Value: 21.22,
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

//BenchmarkTestMaxIntValue benchmarks the Max Value (valid value)
func BenchmarkTestMaxIntValue(b *testing.B) {
	type testModel struct {
		Value int8 `validation:"max=20"`
	}
	model := testModel{
		Value: 19,
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

//BenchmarkTestMaxUintValue benchmarks the Max Value (valid value)
func BenchmarkTestMaxUintValue(b *testing.B) {
	type testModel struct {
		Value uint8 `validation:"max=20"`
	}
	model := testModel{
		Value: 19,
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

//BenchmarkTestMaxFloatValue benchmarks the Max Value (valid value)
func BenchmarkTestMaxFloatValue(b *testing.B) {
	type testModel struct {
		Value float32 `validation:"max=20"`
	}
	model := testModel{
		Value: 19.22,
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

//ExampleIsValid_MinInt is an example for Int Value validation (min)
func ExampleIsValid_minInt() {

	type Product struct {
		// Quantity must more than 1
		Quantity int8 `validation:"min=1"`
	}

	var p Product
	// Will fail since its Quantity = 0

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
	// Output: false [{Quantity must be greater than or equal to 1}]
}

//ExampleIsValid_MinFloat is an example for Float Value validation (min)
func ExampleIsValid_minFloat() {

	type Product struct {
		// Price must more than 0.01
		Price float32 `validation:"min=0.01"`
	}

	var p Product
	// Will fail since its Price = 0

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
	// Output: false [{Price must be greater than or equal to 1E-02}]
}

//ExampleIsValid_MaxInt is an example for Int Value validation (max)
func ExampleIsValid_maxInt() {

	type Product struct {
		// Quantity must more than 1 but less than 99
		Quantity int8 `validation:"min=1 max=99"`
	}

	var p Product
	p.Quantity = 101 // Will fail since it's greater than 99

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
	// Output: false [{Quantity must be less than or equal to 99}]
}

//ExampleIsValid_MaxFloat is an example for Float Value validation (max)
func ExampleIsValid_maxFloat() {

	type Product struct {
		// Price must more than 0.01 but less than 999.99
		Price float32 `validation:"min=0.01 max=999.99"`
	}

	var p Product
	p.Price = 10000.00 // Will fail since it's greater than 999.99

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
	// Output: false [{Price must be less than or equal to 9.9999E+02}]
}

//
// Integer tests (positive and negative)
//

//TestMinValueInt8Positive tests min value on int8
func TestMinValueInt8Positive(t *testing.T) {
	type testModel struct {
		Value int8 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueInt8Negative tests min value on int8
func TestMinValueInt8Negative(t *testing.T) {
	type testModel struct {
		Value int8 `validation:"min=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	model.Value = -40

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMinValueInt16Positive tests min value on int16
func TestMinValueInt16Positive(t *testing.T) {
	type testModel struct {
		Value int16 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueInt16Negative tests min value on int16
func TestMinValueInt16Negative(t *testing.T) {
	type testModel struct {
		Value int16 `validation:"min=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	model.Value = -40

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMinValueInt32Positive tests min value on int32
func TestMinValueInt32Positive(t *testing.T) {
	type testModel struct {
		Value int32 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueInt32Negative tests min value on int32
func TestMinValueInt32Negative(t *testing.T) {
	type testModel struct {
		Value int32 `validation:"min=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	model.Value = -40

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMinValueInt64Positive tests min value on int64
func TestMinValueInt64Positive(t *testing.T) {
	type testModel struct {
		Value int64 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueInt64Negative tests min value on int64
func TestMinValueInt64Negative(t *testing.T) {
	type testModel struct {
		Value int64 `validation:"min=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	model.Value = -40

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMinValueIntPositive tests min value on int
func TestMinValueIntPositive(t *testing.T) {
	type testModel struct {
		Value int `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueIntNegative tests min value on int
func TestMinValueIntNegative(t *testing.T) {
	type testModel struct {
		Value int `validation:"min=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	model.Value = -40

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

//TestMaxValueInt8Positive tests max value on int8
func TestMaxValueInt8Positive(t *testing.T) {
	type testModel struct {
		Value int8 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueInt8Negative tests max value on int8
func TestMaxValueInt8Negative(t *testing.T) {
	type testModel struct {
		Value int8 `validation:"max=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty Int8(0) should be invalid (>= -20)")
	}

	model.Value = -40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}

}

//TestMaxValueInt16Positive tests max value on int16
func TestMaxValueInt16Positive(t *testing.T) {
	type testModel struct {
		Value int16 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueInt16Negative tests max value on int16
func TestMaxValueInt16Negative(t *testing.T) {
	type testModel struct {
		Value int16 `validation:"max=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty Int16(0) should be invalid (>= -20)")
	}

	model.Value = -40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}

}

//TestMaxValueInt32Positive tests max value on int32
func TestMaxValueInt32Positive(t *testing.T) {
	type testModel struct {
		Value int32 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueInt32Negative tests max value on int32
func TestMaxValueInt32Negative(t *testing.T) {
	type testModel struct {
		Value int32 `validation:"max=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty Int32(0) should be invalid (>= -20)")
	}

	model.Value = -40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

//TestMaxValueInt64Positive tests max value on int64
func TestMaxValueInt64Positive(t *testing.T) {
	type testModel struct {
		Value int64 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueInt64Negative tests max value on int64
func TestMaxValueInt64Negative(t *testing.T) {
	type testModel struct {
		Value int64 `validation:"max=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty Int64(0) should be invalid (>= -20)")
	}

	model.Value = -40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

//TestMaxValueIntPositive tests max value on int
func TestMaxValueIntPositive(t *testing.T) {
	type testModel struct {
		Value int `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueIntNegative tests max value on int
func TestMaxValueIntNegative(t *testing.T) {
	type testModel struct {
		Value int `validation:"max=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty Int(0) should be invalid (>= -20)")
	}

	model.Value = -40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

//
// Unsigned integer tests (positive and negative)
//

//TestMinValueUint8Positive tests min value on uint8
func TestMinValueUint8Positive(t *testing.T) {
	type testModel struct {
		Value uint8 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint8Negative tests min value on uint8
func TestMinValueUint8Negative(t *testing.T) {
	type testModel struct {
		Value uint8 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint16Positive tests min value on uint16
func TestMinValueUint16Positive(t *testing.T) {
	type testModel struct {
		Value uint16 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint16Negative tests min value on uint16
func TestMinValueUint16Negative(t *testing.T) {
	type testModel struct {
		Value uint16 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint32Positive tests min value on uint32
func TestMinValueUint32Positive(t *testing.T) {
	type testModel struct {
		Value uint32 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint32Negative tests min value on uint32
func TestMinValueUint32Negative(t *testing.T) {
	type testModel struct {
		Value uint32 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint64Positive tests min value on uint64
func TestMinValueUint64Positive(t *testing.T) {
	type testModel struct {
		Value uint64 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUint64Negative tests min value on uint64
func TestMinValueUint64Negative(t *testing.T) {
	type testModel struct {
		Value uint64 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUintPositive tests min value on uint
func TestMinValueUintPositive(t *testing.T) {
	type testModel struct {
		Value uint `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as min was set to 20", model.Value)
	}

	model.Value = 19

	ok, _ = IsValid(model)
	if ok {
		t.Fatal("Expected failure as 19 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueUintNegative tests min value on uint
func TestMinValueUintNegative(t *testing.T) {
	type testModel struct {
		Value uint `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUint8Positive tests max value on uint8
func TestMaxValueUint8Positive(t *testing.T) {
	type testModel struct {
		Value uint8 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUint8Negative tests max value on uint8
func TestMaxValueUint8Negative(t *testing.T) {
	type testModel struct {
		Value uint8 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUint16Positive tests max value on uint16
func TestMaxValueUint16Positive(t *testing.T) {
	type testModel struct {
		Value uint16 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUint16Negative tests max value on uint16
func TestMaxValueUint16Negative(t *testing.T) {
	type testModel struct {
		Value uint16 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUint32Positive tests max value on uint32
func TestMaxValueUint32Positive(t *testing.T) {
	type testModel struct {
		Value uint32 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUint32Negative tests max value on uint32
func TestMaxValueUint32Negative(t *testing.T) {
	type testModel struct {
		Value uint32 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUint64Positive tests max value on uint64
func TestMaxValueUint64Positive(t *testing.T) {
	type testModel struct {
		Value uint64 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUint64Negative tests max value on uint64
func TestMaxValueUint64Negative(t *testing.T) {
	type testModel struct {
		Value uint64 `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueUintPositive tests max value on uint
func TestMaxValueUintPositive(t *testing.T) {
	type testModel struct {
		Value uint `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as value is less than 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Expected failure as value is > max value", errs)
	}

}

//TestMaxValueUintNegative tests max value on uint
func TestMaxValueUintNegative(t *testing.T) {
	type testModel struct {
		Value uint `validation:"max=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//
// Float tests (min / max)
//

//TestMinValueFloat32 tests min value on float32
func TestMinValueFloat32(t *testing.T) {
	type testModel struct {
		Value float32 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMinValueFloat64 tests min value on float64
func TestMinValueFloat64(t *testing.T) {
	type testModel struct {
		Value float64 `validation:"min=20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	model.Value = 40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueFloat32 tests max value on float32
func TestMaxValueFloat32(t *testing.T) {
	type testModel struct {
		Value float32 `validation:"max=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than max -20")
	}

	model.Value = -40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

//TestMaxValueFloat64 tests max value on float64
func TestMaxValueFloat64(t *testing.T) {
	type testModel struct {
		Value float64 `validation:"max=-20"`
	}
	model := testModel{}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Expected failure as 0 is less than max -20")
	}

	model.Value = -40

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}
