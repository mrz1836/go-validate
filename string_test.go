package validate

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//
// Generic string struct and function tests
//

// TestMaxLengthValidation - a series of different tests
func TestMaxLengthValidation(t *testing.T) {
	// Test invalid types
	_, err := maxLengthValidation("foo", reflect.Int)
	require.Error(t, err, "Expected to fail - foo is a string and not a number")

	// Test making an interface
	maxInterface, err := maxLengthValidation("10", reflect.Int)
	require.NoError(t, err)

	// Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := maxInterface.Validate("this should fail", testVal)
	require.NotNil(t, errs, "Expected to fail")

	// Test converting a string
	var invalidSlice []int
	errs = maxInterface.Validate(invalidSlice, testVal)
	require.NotNil(t, errs, "Expected to fail, value is not of type string and MaxLengthValidation only accepts strings")

	// Test converting an unsigned int
	var invalid uint64
	errs = maxInterface.Validate(invalid, testVal)
	require.NotNil(t, errs, "Expected to fail, value is not of type string and MaxLengthValidation only accepts strings")
}

// TestMinLengthValidation - a series of different tests
func TestMinLengthValidation(t *testing.T) {
	// Test invalid types
	_, err := minLengthValidation("foo", reflect.Int)
	require.Error(t, err, "Expected to fail - foo is a string and not a number")

	// Test making an interface
	minInterface, err := minLengthValidation("10", reflect.Int)
	require.NoError(t, err)

	// Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := minInterface.Validate("this", testVal)
	require.NotNil(t, errs, "Expected to fail")

	// Test converting a string
	var invalidSlice []int
	errs = minInterface.Validate(invalidSlice, testVal)
	require.NotNil(t, errs, "Expected to fail, value is not of type string and MinLengthValidation only accepts strings")

	// Test converting an unsigned int
	var invalid uint64
	errs = minInterface.Validate(invalid, testVal)
	require.NotNil(t, errs, "Expected to fail, value is not of type string and MinLengthValidation only accepts strings")
}

// TestFormatValidation - a series of different tests
func TestFormatValidation(t *testing.T) {
	// Test invalid types
	_, err := formatValidation("doesNotExist", reflect.String)
	require.Error(t, err, "Expected to fail - format validation does not exist")

	// Test making an interface
	formatInterface, err := formatValidation("email", reflect.String)
	require.NoError(t, err)

	// Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := formatInterface.Validate("this works?", testVal)
	require.NotNil(t, errs, "Expected error, value is not email")

	errs = formatInterface.Validate("this@", testVal)
	require.NotNil(t, errs, "Expected error, value is not email")

	errs = formatInterface.Validate("this@domain.com", testVal)
	require.Nil(t, errs, "Email is valid but should not fail")

	var invalidValue []int
	errs = formatInterface.Validate(invalidValue, testVal)
	require.NotNil(t, errs, "Expected error, value is not string")
}

//
// Test max length
//

// TestMaxLengthValid tests max length (valid length of characters)
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

// TestMaxLengthInvalid tests max length (invalid length of characters)
func TestMaxLengthInvalid(t *testing.T) {
	type testModel struct {
		Value string `validation:"max_length=5"`
	}

	model := testModel{
		Value: "12345678",
	}

	ok, errs := IsValid(model)
	assert.False(t, ok, "Max length should have failed")
	require.NotEmpty(t, errs, "Max length errs should have at least 1 item")
}

// TestMaxLengthWrongType tests to make sure only valid types
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

// BenchmarkTestMaxLength benchmarks the Max Length Value (valid value)
func BenchmarkTestMaxLength(b *testing.B) {
	type testModel struct {
		Value string `validation:"max_length=20"`
	}
	model := testModel{
		Value: "12345",
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

// ExampleIsValid_MaxLength is an example for max length validation (max)
func ExampleIsValid_maxLength() {
	type Person struct {
		// Gender must not be longer than 10 characters
		Gender string `validation:"max_length=10"`
	}

	var p Person
	p.Gender = "This is invalid!" // Will fail since it's > 10 characters

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
	// Output: false [{Gender must be no more than 10 characters}]
}

//
// Test min length
//

// TestMinLengthValid tests min length (valid length of characters)
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

// TestMinLengthInvalid tests min length (invalid length of characters)
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

// TestMinLengthWrongType tests to make sure only valid types
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

// BenchmarkTestMinLength benchmarks the Min Length Value (valid value)
func BenchmarkTestMinLength(b *testing.B) {
	type testModel struct {
		Value string `validation:"min_length=3"`
	}
	model := testModel{
		Value: "12345",
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

// ExampleIsValid_MinLength is an example for min length validation (min)
func ExampleIsValid_minLength() {
	type Person struct {
		// Gender must be > 1 character
		Gender string `validation:"min_length=1"`
	}

	var p Person
	// Will fail since it's < 1 character

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
	// Output: false [{Gender must be at least 1 characters}]
}

//
// Test format, regex
//

// TestFormatEmail tests an email format (invalid and valid formats)
func TestFormatEmail(t *testing.T) {
	type testModel struct {
		Value string `validation:"format=email"`
	}

	model := testModel{
		Value: "",
	}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty email should be invalid")
	}

	model.Value = "123"

	ok, _ = IsValid(model)
	if ok {
		t.Fatalf("Invalid email (%s) should be invalid", model.Value)
	}

	model.Value = "test@example.com"

	ok, errs := IsValid(model)
	if !ok {
		t.Fatalf("Valid email (%s) should be valid - errs: %x", model.Value, errs)
	}

	model.Value = "BaseMail0@Base.com"

	ok, errs = IsValid(model)
	if !ok {
		t.Fatalf("Valid email with a number(%s) should be valid - errs: %x", model.Value, errs)
	}

	model.Value = "BaseMail0@Base.consulting"

	ok, errs = IsValid(model)
	if !ok {
		t.Fatalf("Valid email with a new TLD(%s) should be valid - errs: %x", model.Value, errs)
	}

	model.Value = "BaseMail0@Base.miami"

	ok, errs = IsValid(model)
	if !ok {
		t.Fatalf("Valid email with a new TLD(%s) should be valid - errs: %x", model.Value, errs)
	}

	model.Value = "BaseMail0@Base.co.uk"

	ok, errs = IsValid(model)
	if !ok {
		t.Fatalf("Valid email with a international TLD(%s) should be valid - errs: %x", model.Value, errs)
	}

	model.Value = "BaseMail0@email.Base.com"

	ok, errs = IsValid(model)
	if !ok {
		t.Fatalf("Valid email with a subdomain TLD(%s) should be valid - errs: %x", model.Value, errs)
	}

	// All TLD tests are in: TestFormatEmailAcceptedTLDs
}

// TestFormatEmailAcceptedTLDs tests an email format (all accepted TLDs)
func TestFormatEmailAcceptedTLDs(t *testing.T) {
	type testModel struct {
		Value string `validation:"format=email"`
	}

	model := testModel{
		Value: "",
	}

	// Loop all TLDs and try an email
	for _, tld := range TopLevelDomains {

		model.Value = "BaseMail@BaseDomain." + tld

		ok, _ := IsValid(model)
		if !ok {
			t.Fatal("TLD should be accepted but failed", tld, model.Value)
		}
	}
}

// BenchmarkTestFormatEmail benchmarks the format by email (valid value)
func BenchmarkTestFormatEmail(b *testing.B) {
	type testModel struct {
		Value string `validation:"format=email"`
	}
	model := testModel{
		Value: "BaseMail@Base.com",
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

// ExampleIsValid_FormatEmail is an example for format email validation
func ExampleIsValid_formatEmail() {
	type Person struct {
		// Email must be in valid email format
		Email string `validation:"format=email"`
	}

	var p Person
	// Will fail since the email is not valid

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
	// Output: false [{Email does not match email format}]
}

// TestFormatRegExp tests regex format (invalid and valid formats)
func TestFormatRegExp(t *testing.T) {
	type testModel struct {
		Value string `validation:"format=regexp:Test[0-9]+"`
	}

	model := testModel{
		Value: "",
	}

	ok, _ := IsValid(model)
	if ok {
		t.Fatal("Empty regexp should be invalid")
	}

	model.Value = "invalid"
	ok, _ = IsValid(model)

	if ok {
		t.Fatalf("Invalid regexp (%s) should be invalid", model.Value)
	}

	model.Value = "Test123"
	ok, errs := IsValid(model)

	if !ok {
		t.Fatalf("Valid regexp (%s) should be valid - errs: %x", model.Value, errs)
	}
}

// BenchmarkTestFormatRegEx benchmarks the format by regex (valid value)
func BenchmarkTestFormatRegEx(b *testing.B) {
	type testModel struct {
		Value string `validation:"format=regexp:Test[0-9]+"`
	}
	model := testModel{
		Value: "Test123",
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

// ExampleIsValid_FormatRegEx is an example for format regex validation
func ExampleIsValid_formatRegEx() {
	type Person struct {
		// Phone must be in valid phone regex format
		Phone string `validation:"format=regexp:[0-9]+"`
	}

	var p Person
	// Will fail since the email is not valid

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
	// Output: false [{Phone does not match regexp format}]
}

//
// Test compare string
//

// TestCompareStringValid tests compare string (valid comparison)
func TestCompareStringValid(t *testing.T) {
	type testModel struct {
		Value        string `validation:"compare=ValueCompare"`
		ValueCompare string
	}

	model := testModel{
		Value:        "123456",
		ValueCompare: "123456",
	}

	ok, errs := IsValid(model)
	if !ok {
		t.Fatal(errs)
	} else if errs != nil {
		t.Fatal(errs)
	}
}

// TestCompareStringInValid tests compare string (invalid comparison)
func TestCompareStringInValid(t *testing.T) {
	type testModel struct {
		Value        string `validation:"compare=ValueCompare"`
		ValueCompare string
	}

	model := testModel{
		Value:        "12345",
		ValueCompare: "123456789",
	}

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("Compare should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Compare errs should have 1 item not: %d", len(errs))
	}
}

// TestCompareStringWrongType tests to make sure only valid types
func TestCompareStringWrongType(t *testing.T) {
	type testModel struct {
		Value        int64 `validation:"compare=ValueCompare"`
		ValueCompare string
	}

	model := testModel{
		Value:        1234,
		ValueCompare: "123456789",
	}

	ok, errs := IsValid(model)
	if ok {
		t.Fatal("This should have failed since min_length is for strings only")
	}
	if errs[0].Error() != "Value is not of type string and StringEqualsStringValidation only accepts strings" {
		t.Fatal("Error was not recognized", errs)
	}
}

// BenchmarkTestCompareString benchmarks the comparing of string (valid value)
func BenchmarkTestCompareString(b *testing.B) {
	type testModel struct {
		Value        string `validation:"compare=ValueCompare"`
		ValueCompare string
	}
	model := testModel{
		Value:        "Test123",
		ValueCompare: "Test123",
	}

	for i := 0; i < b.N; i++ {
		_, _ = IsValid(model)
	}
}

// ExampleIsValid_CompareString is an example for compare string validation
func ExampleIsValid_compareString() {
	type User struct {
		// Password should match confirmation on submission
		Password             string `validation:"compare=PasswordConfirmation"`
		PasswordConfirmation string
	}

	var u User // User submits a new password and confirms wrong
	u.Password = "This"
	u.PasswordConfirmation = "That"

	ok, errs := IsValid(u)
	fmt.Println(ok, errs)
	// Output: false [{Password is not the same as the compare field PasswordConfirmation}]
}
