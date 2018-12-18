/*
Package validation (go-validation) provides validations for struct fields based on a validation tag and offers additional validation functions.
*/
package validation

import (
	"reflect"
	"testing"
)

//LengthTestStruct is for testing string length
type LengthTestStruct struct {

	//LongTitle is to test a max value
	LongTitle string `validation:"max_length=5"`

	//ShortTitle is to test a min value
	ShortTitle string `validation:"min_length=3"`

	//Title is to test both a min and max length
	Title string `validation:"min_length=3 max_length=5"`
}

//FormatTestStruct is for testing the format validation
type FormatTestStruct struct {

	//Email is an email field
	Email string `validation:"format=email"`

	//RegExp is a field using regular expression
	RegExp string `validation:"format=regexp:Test[0-9]+"`
}

//TestMaxLengthValid tests max length
func TestMaxLengthValid(t *testing.T) {
	object := LengthTestStruct{
		LongTitle:  "123",
		ShortTitle: "123",
		Title:      "1234",
	}

	ok, errs := IsValid(object)
	if !ok {
		t.Fatal(errs)
	}
}

//TestMaxLengthInvalid tests max length
func TestMaxLengthInvalid(t *testing.T) {
	object := LengthTestStruct{
		Title: "123456",
	}

	ok, errs := IsValid(object)
	if ok {
		t.Fatal("Max length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Max length errs should have 1 item not: %d", len(errs))
	}
}

//TestMinLengthValid tests min length
func TestMinLengthValid(t *testing.T) {
	object := LengthTestStruct{
		Title:      "12345",
		ShortTitle: "123",
	}

	ok, errs := IsValid(object)
	if !ok {
		t.Fatal(errs)
	}
}

//TestMinLengthInvalid tests min length
func TestMinLengthInvalid(t *testing.T) {
	object := LengthTestStruct{
		Title: "1234",
	}

	ok, errs := IsValid(object)
	if ok {
		t.Fatal("Min length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Min length errs should have 1 item not: %d", len(errs))
	}
}

//TestLengthValid tests length
func TestLengthValid(t *testing.T) {
	object := LengthTestStruct{
		Title:      "12345",
		ShortTitle: "123",
	}

	ok, errs := IsValid(object)
	if !ok {
		t.Fatal(errs)
	}
}

//TestLengthInvalid tests length
func TestLengthInvalid(t *testing.T) {
	// Check min_length=3
	object := LengthTestStruct{
		Title:      "12",
		ShortTitle: "123",
	}

	ok, errs := IsValid(object)
	if ok {
		t.Fatal("Length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Length errs should have 1 item not: %d", len(errs))
	}

	// Check max_length=5
	object = LengthTestStruct{
		Title:      "123456",
		ShortTitle: "123",
	}

	ok, errs = IsValid(object)
	if ok {
		t.Fatal("Length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Length errs should have 1 item not: %d", len(errs))
	}
}

//TestEmail tests email format
func TestEmail(t *testing.T) {
	object := FormatTestStruct{
		Email:  "",
		RegExp: "Test123",
	}

	ok, _ := IsValid(object)
	if ok {
		t.Fatal("Empty email should be invalid")
	}

	object.Email = "123"

	ok, _ = IsValid(object)
	if ok {
		t.Fatalf("Invalid email (%s) should be invalid", object.Email)
	}

	object.Email = "test@example.com"

	ok, errs := IsValid(object)
	if !ok {
		t.Fatalf("Valid email (%s) should be valid - errs: %x", object.Email, errs)
	}

	object.Email = "BaseMail0@Base.com"

	ok, errs = IsValid(object)
	if !ok {
		t.Fatalf("Valid email with a number(%s) should be valid- errs: %x", object.Email, errs)
	}

	object.Email = "BaseMail0@Base.consulting"

	ok, errs = IsValid(object)
	if !ok {
		t.Fatalf("Valid email with a new TLD(%s) should be valid - errs: %x", object.Email, errs)
	}

	//todo: more regex checks for adhering to TLD specs

}

//TestRegExp tests regex format
func TestRegExp(t *testing.T) {
	object := FormatTestStruct{
		RegExp: "",
		Email:  "valid@example.com",
	}

	ok, _ := IsValid(object)
	if ok {
		t.Fatal("Empty regexp should be invalid")
	}

	object.RegExp = "invalid"
	ok, _ = IsValid(object)

	if ok {
		t.Fatalf("Invalid regexp (%s) should be invalid", object.RegExp)
	}

	object.RegExp = "Test123"
	ok, errs := IsValid(object)

	if !ok {
		t.Fatalf("Valid regexp (%s) should be valid - errs: %x", object.RegExp, errs)
	}
}

//TestCompare tests comparing fields
//todo: finish the compare string tests
/*func TestCompare(){

}*/

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
