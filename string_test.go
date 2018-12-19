/*
Package validate (go-validate) provides validations for struct fields based on a validation tag and offers additional validation functions.
*/
package validate

import (
	"reflect"
	"testing"
)

//FormatTestStruct is for testing the format validation
type FormatTestStruct struct {

	//Email is an email field
	Email string `validation:"format=email"`

	//RegExp is a field using regular expression
	RegExp string `validation:"format=regexp:Test[0-9]+"`
}

//
// Test max length
//

//TestMaxLengthValid tests max length (valid length of characters)
func TestMaxLengthValid(t *testing.T) {

	type testModel struct {
		Value string `validation:"max_length=5"`
	}

	model := testModel{
		Value: "1234",
	}

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal(errs)
	} else if errs != nil {
		t.Fatal(errs)
	}
}

//TestMaxLengthInvalid tests max length (invalid length of characters)
func TestMaxLengthInvalid(t *testing.T) {
	type testModel struct {
		Value string `validation:"max_length=5"`
	}

	model := testModel{
		Value: "12345678",
	}

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Max length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Max length errs should have 1 item not: %d", len(errs))
	}
}

//TestMaxLengthWrongType tests to make sure only valid types
func TestMaxLengthWrongType(t *testing.T) {
	type testModel struct {
		Value int64 `validation:"max_length=5"`
	}

	model := testModel{
		Value: 1234,
	}

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("This should have failed since max_length is for strings only")
	}
	if errs[0].Error() != "Value is not of type string and MaxLengthValidation only accepts strings" {
		t.Fatal("Error was not recognized", errs)
	}
}

//
// Test min length
//

//TestMinLengthValid tests min length (valid length of characters)
func TestMinLengthValid(t *testing.T) {

	type testModel struct {
		Value string `validation:"min_length=5"`
	}

	model := testModel{
		Value: "123456",
	}

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal(errs)
	} else if errs != nil {
		t.Fatal(errs)
	}
}

//TestMinLengthInvalid tests min length (invalid length of characters)
func TestMinLengthInvalid(t *testing.T) {
	type testModel struct {
		Value string `validation:"min_length=5"`
	}

	model := testModel{
		Value: "123",
	}

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Min length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Min length errs should have 1 item not: %d", len(errs))
	}
}

//TestMinLengthWrongType tests to make sure only valid types
func TestMinLengthWrongType(t *testing.T) {
	type testModel struct {
		Value int64 `validation:"min_length=5"`
	}

	model := testModel{
		Value: 1234,
	}

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("This should have failed since min_length is for strings only")
	}
	if errs[0].Error() != "Value is not of type string and MinLengthValidation only accepts strings" {
		t.Fatal("Error was not recognized", errs)
	}
}

//
// Test format, regex
//

//TestEmail tests email format
func TestEmail(t *testing.T) {
	model := FormatTestStruct{
		Email:  "",
		RegExp: "Test123",
	}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty email should be invalid")
	}

	model.Email = "123"

	ok, _ = IsValid(model)
	if ok {
		t.Fatalf("Invalid email (%s) should be invalid", model.Email)
	}

	model.Email = "test@example.com"

	ok, errs := IsValid(model)
	if !ok {
		t.Fatalf("Valid email (%s) should be valid - errs: %x", model.Email, errs)
	}

	model.Email = "BaseMail0@Base.com"

	ok, errs = IsValid(model)
	if !ok {
		t.Fatalf("Valid email with a number(%s) should be valid- errs: %x", model.Email, errs)
	}

	model.Email = "BaseMail0@Base.consulting"

	ok, errs = IsValid(model)
	if !ok {
		t.Fatalf("Valid email with a new TLD(%s) should be valid - errs: %x", model.Email, errs)
	}

	//todo: more regex checks for adhering to TLD specs

}

//TestRegExp tests regex format
func TestRegExp(t *testing.T) {
	model := FormatTestStruct{
		RegExp: "",
		Email:  "valid@example.com",
	}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty regexp should be invalid")
	}

	model.RegExp = "invalid"
	ok, _ = IsValid(model)

	if ok {
		t.Fatalf("Invalid regexp (%s) should be invalid", model.RegExp)
	}

	model.RegExp = "Test123"
	ok, errs := IsValid(model)

	if !ok {
		t.Fatalf("Valid regexp (%s) should be valid - errs: %x", model.RegExp, errs)
	}
}

//
// Generic string struct and function tests
//

//TestMaxLengthValidation - series of different tests
func TestMaxLengthValidation(t *testing.T) {

	//Test invalid types
	var err error

	//Fail if string submitted or Parse int fails
	_, err = maxLengthValidation("foo", reflect.Int)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}

	//Test making an interface
	maxInterface, err := maxLengthValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := maxInterface.Validate("this should fail", testVal)
	if errs == nil {
		t.Fatal("Expected to fail, 8 < 10")
	}

	//Test converting a string
	var invalidSlice []int
	errs = maxInterface.Validate(invalidSlice, testVal)
	if errs == nil {
		t.Fatal("Expected to fail, value is not of type string and MaxLengthValidation only accepts strings")
	}

	//Test converting an unsigned int
	var invalid uint64
	errs = maxInterface.Validate(invalid, testVal)
	if errs == nil {
		t.Fatal("Expected to fail, value is not of type string and MaxLengthValidation only accepts strings")
	}
}

//TestMinLengthValidation - series of different tests
func TestMinLengthValidation(t *testing.T) {

	//Test invalid types
	var err error

	//Fail if string submitted or Parse int fails
	_, err = minLengthValidation("foo", reflect.Int)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}

	//Test making an interface
	minInterface, err := minLengthValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := minInterface.Validate("this", testVal)
	if errs == nil {
		t.Fatal("Expected to fail", "10", "this")
	}

	//Test converting a string
	var invalidSlice []int
	errs = minInterface.Validate(invalidSlice, testVal)
	if errs == nil {
		t.Fatal("Expected to fail, value is not of type string and MinLengthValidation only accepts strings")
	}

	//Test converting an unsigned int
	var invalid uint64
	errs = minInterface.Validate(invalid, testVal)
	if errs == nil {
		t.Fatal("Expected to fail, value is not of type string and MinLengthValidation only accepts strings")
	}
}

//TestFormatValidation - series of different tests
func TestFormatValidation(t *testing.T) {

	//Test invalid types
	var err error

	//Fail if string submitted or Parse int fails
	_, err = formatValidation("doesNotExist", reflect.String)
	if err == nil {
		t.Fatal("Expected to fail - foo is a string and not a number")
	}

	//Test making an interface
	formatInterface, err := formatValidation("email", reflect.String)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := formatInterface.Validate("this works?", testVal)
	if errs == nil {
		t.Fatal("Expected error, value is not email")
	}

	errs = formatInterface.Validate("this@", testVal)
	if errs == nil {
		t.Fatal("Expected error, value is not email")
	}

	errs = formatInterface.Validate("this@domain.com", testVal)
	if errs != nil {
		t.Fatal("Email is valid but failed")
	}

	var invalidValue []int
	errs = formatInterface.Validate(invalidValue, testVal)
	if errs == nil {
		t.Fatal("Expected error, value is not string")
	}

}

/*func ExampleIsValid_StringLength() {
	type Person struct {
		// Name must be between 1 and 5 characters inclusive
		Name string `validation:"min_length=1 max_length=5"`
	}

	var p Person

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
}*/

/*func ExampleIsValid_format() {
	type Person struct {
		// Email must be valid email
		Email string `validation:"format=email"`
	}

	var p Person

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
}
*/
