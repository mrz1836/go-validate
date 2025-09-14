package validate

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// invalidNumericTypes is for the types not allowed (numeric tests)
var invalidNumericTypes = []reflect.Kind{ //nolint:gochecknoglobals // Test data - acceptable pattern
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
}

// TestValidationMap_Atomicity
func TestValidationMap_Atomicity(_ *testing.T) {
	vm := Map{}
	// typ := reflect.TypeOf(vm)
	typ := reflect.TypeOf(&vm) // todo: go vet: call of reflect.TypeOf copies lock value: govalidation.Map contains sync.Map contains sync.Mutex
	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	wg2 := sync.WaitGroup{}
	wg2.Add(2)
	count := 10000
	go func() {
		wg1.Wait()
		for i := 0; i < count; i++ {
			vm.set(typ, nil)
		}
		wg2.Done()
	}()
	go func() {
		wg1.Wait()
		for i := 0; i < count; i++ {
			vm.get(typ)
		}
		wg2.Done()
	}()
	wg1.Done() // start !
	wg2.Wait()
}

// TestValidationSetFieldName test setting and getting field name
func TestValidationSetFieldName(t *testing.T) {
	inter, err := minValueValidation("10", reflect.Int)
	require.NoError(t, err)

	// Set the name
	testField := "test_field"
	inter.SetFieldName(testField)

	// Get the name
	name := inter.FieldName()
	assert.Equal(t, testField, name)
}

// TestValidationSetFieldIndex test setting and getting field index
func TestValidationSetFieldIndex(t *testing.T) {
	inter, err := minValueValidation("10", reflect.Int)
	require.NoError(t, err)

	// Set the index
	indexNumber := 18
	inter.SetFieldIndex(indexNumber)

	// Get the index
	index := inter.FieldIndex()
	assert.Equal(t, indexNumber, index)
}

// TestValidationValidate test using the Validate method (valid and invalid)
func TestValidationValidate(t *testing.T) {
	// Test making an interface
	minInterface, err := minValueValidation("10", reflect.Int)
	require.NoError(t, err)

	// Set the field name
	minInterface.SetFieldName("field")

	// Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := minInterface.Validate(8, testVal)
	require.NotNil(t, errs, "Expected to fail, 8 < 10")

	// Test failure error response
	errs = minInterface.Validate(8, testVal)
	assert.Equal(t, "field must be greater than or equal to 10", errs.Error())
}

// TestValidationErrorError tests using the Error method
func TestValidationErrorError(t *testing.T) {
	newError := ValidationError{
		Key:     "the_key",
		Message: "the_message",
	}

	// test it if correct
	errorResponse := newError.Error()
	expected := newError.Key + " " + newError.Message
	assert.Equal(t, expected, errorResponse)
}

// TestValidationErrorsError tests using the Error method
func TestValidationErrorsError(t *testing.T) {
	newError := ValidationError{
		Key:     "the_key",
		Message: "the_message",
	}

	newError2 := ValidationError{
		Key:     "the_key2",
		Message: "the_message2",
	}

	sliceOfErrors := ValidationErrors{}
	sliceOfErrors = append(sliceOfErrors, newError, newError2)

	// test it if correct
	errorResponse := sliceOfErrors.Error()
	expected := newError.Key + " " + newError.Message + " and 1 other errors"
	assert.Equal(t, expected, errorResponse)
}

// ExampleValidationError_Error is showing how to use the errors
func ExampleValidationError_Error() {
	type Person struct {
		// Gender must not be longer than 10 characters
		Gender string `validation:"max_length=10"`
	}

	var p Person
	p.Gender = "This is invalid!" // Will fail since it's > 10 characters

	_, errs := IsValid(p)
	fmt.Println(errs[0].Error())
	// Output: Gender must be no more than 10 characters
}

// TestValidationValidateFunc tests the Validate method of the Validation struct
func TestValidationValidateFunc(t *testing.T) {
	// Create a Validation instance with a field name
	v := &Validation{fieldName: "testField"}

	// Call the Validate method
	result := v.Validate(nil, reflect.Value{})

	// Assert that the result is not nil
	require.NotNil(t, result, "expected a ValidationError, got nil")

	// Assert that the Key is correct
	assert.Equal(t, "testField", result.Key)

	// Assert that the Message is correct
	expectedMessage := "validation not implemented"
	assert.Equal(t, expectedMessage, result.Message)
}

// Tests that are still needed for full package coverage
// todo:  TestMap_AddValidation(t *testing.T)
// todo:  TestMap_IsValid(t *testing.T)
// todo:  TestAddValidation(t *testing.T)
// todo:  TestIsValid(t *testing.T)
