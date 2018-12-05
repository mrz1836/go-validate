/*
Package govalidation provides validations for struct fields based on a validation tag and offers additional validation functions.

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.
*/
package govalidation

import (
	"reflect"
	"sync"
	"testing"
)

//TestValidationMap_Atomicity
func TestValidationMap_Atomicity(t *testing.T) {
	vm := Map{}
	typ := reflect.TypeOf(vm)
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
	inter.SetFieldName("test_field")

	//Get the name
	name := inter.FieldName()
	if name != "test_field" {
		t.Fatal("Field name was not the same as when set")
	}
}

//TestValidation_SetFieldIndex
func TestValidation_SetFieldIndex(t *testing.T) {
	inter, err := minValueValidation("10", reflect.Int)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Set the index
	inter.SetFieldIndex(18)

	//Get the index
	index := inter.FieldIndex()
	if index != 18 {
		t.Fatal("Field index was not the same as when set")
	}
}
