/*
Package validation (go-validation) provides validations for struct fields based on a validation tag and offers additional validation functions.
*/
package validation

import (
	"reflect"
	"sync"
	"testing"
)

//invalidNumericTypes is for the types not allowed (numeric tests)
var (
	invalidNumericTypes []reflect.Kind
)

//init load the default invalid types
func init() {

	//Build the invalid numeric types
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

//TestValidationMap_Atomicity
func TestValidationMap_Atomicity(t *testing.T) {
	vm := Map{}
	typ := reflect.TypeOf(vm) //todo: go vet: call of reflect.TypeOf copies lock value: govalidation.Map contains sync.Map contains sync.Mutex
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

//TestValidation_SetFieldName
func TestValidation_SetFieldName(t *testing.T) {
	inter, err := minValueValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Set the name
	testField := "test_field"
	inter.SetFieldName(testField)

	//Get the name
	name := inter.FieldName()
	if name != testField {
		t.Fatal("Field name was not the same as when set", testField, name)
	}
}

//TestValidation_SetFieldIndex
func TestValidation_SetFieldIndex(t *testing.T) {
	inter, err := minValueValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Set the index
	indexNumber := 18
	inter.SetFieldIndex(indexNumber)

	//Get the index
	index := inter.FieldIndex()
	if index != indexNumber {
		t.Fatal("Field index was not the same as when set", index, indexNumber)
	}
}

//todo: test the validation errors
