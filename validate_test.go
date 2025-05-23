package validate

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

// invalidNumericTypes is for the types not allowed (numeric tests)
var (
	invalidNumericTypes []reflect.Kind
)

// init loads the default test data
func init() {

	// Build the invalid numeric types
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

// TestValidation_SetFieldName test setting and getting field name
func TestValidation_SetFieldName(t *testing.T) {
	inter, err := minValueValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Set the name
	testField := "test_field"
	inter.SetFieldName(testField)

	// Get the name
	name := inter.FieldName()
	if name != testField {
		t.Fatal("Field name was not the same as when set", testField, name)
	}
}

// TestValidation_SetFieldIndex test setting and getting field index
func TestValidation_SetFieldIndex(t *testing.T) {
	inter, err := minValueValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Set the index
	indexNumber := 18
	inter.SetFieldIndex(indexNumber)

	// Get the index
	index := inter.FieldIndex()
	if index != indexNumber {
		t.Fatal("Field index was not the same as when set", index, indexNumber)
	}
}

// TestValidation_Validate test using the Validate method (valid and invalid)
func TestValidation_Validate(t *testing.T) {

	// Test making an interface
	minInterface, err := minValueValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Set the field name
	minInterface.SetFieldName("field")

	// Test running the validate method
	var testInt int32 = 1
	testVal := reflect.ValueOf(testInt)
	errs := minInterface.Validate(8, testVal)
	if errs == nil {
		t.Fatal("Expected to fail, 8 < 10")
	}

	// Test failure error response
	errs = minInterface.Validate(8, testVal)
	if errs.Error() != "field must be greater than or equal to 10" {
		t.Fatal("Expected to fail, 8 < 10", errs)
	}
}

// TestValidationError_Error tests using the Error method
func TestValidationError_Error(t *testing.T) {
	newError := ValidationError{
		Key:     "the_key",
		Message: "the_message",
	}

	// test it if correct
	errorResponse := newError.Error()
	if errorResponse != newError.Key+" "+newError.Message {
		t.Fatal("Error response was not `key message` as expected", errorResponse)
	}
}

// TestValidationErrors_Error tests using the Error method
func TestValidationErrors_Error(t *testing.T) {
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
	if errorResponse != newError.Key+" "+newError.Message+" and 1 other errors" {
		t.Fatal("Error response was not `key message` as expected", errorResponse)
	}
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

// TestValidation_ValidateFunc tests the Validate method of the Validation struct
func TestValidation_ValidateFunc(t *testing.T) {
	// Create a Validation instance with a field name
	v := &Validation{fieldName: "testField"}

	// Call the Validate method
	result := v.Validate(nil, reflect.Value{})

	// Assert that the result is not nil
	if result == nil {
		t.Fatalf("expected a ValidationError, got nil")
	}

	// Assert that the Key is correct
	if result.Key != "testField" {
		t.Errorf("expected Key to be 'testField', got %q", result.Key)
	}

	// Assert that the Message is correct
	expectedMessage := "validation not implemented"
	if result.Message != expectedMessage {
		t.Errorf("expected Message to be %q, got %q", expectedMessage, result.Message)
	}
}

// Tests that are still needed for full package coverage
// todo:  TestMap_AddValidation(t *testing.T)
// todo:  TestMap_IsValid(t *testing.T)
// todo:  TestAddValidation(t *testing.T)
// todo:  TestIsValid(t *testing.T)
