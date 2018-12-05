/*
Package govalidation provides validations for struct fields based on a validation tag and offers additional validation functions.

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.
*/
package govalidation

import (
	"log"
	"reflect"
	"strings"
	"sync"
)

//Interface specifies the necessary methods a validation must implement to be compatible with this package
type Interface interface {

	//SetFieldIndex stores the index of the field the validation was applied to
	SetFieldIndex(index int)

	//FieldIndex retrieves the index of the field the validation was applied to
	FieldIndex() int

	//SetFieldName stores the name of the field the validation was applied to
	SetFieldName(name string)

	//FieldName retrieves the name of the field the validation was applied to
	FieldName() string

	//Validate determines if the value is valid. Nil is returned if it is valid
	Validate(value interface{}, obj reflect.Value) *ValidationError
}

//Validation is an implementation of a Interface and can be used to provide basic functionality to a new validation type through an anonymous field
type Validation struct {
	Name       string
	fieldIndex int
	fieldName  string
	options    string
}

//SetFieldIndex stores the index of the field the validation was applied to
func (v *Validation) SetFieldIndex(index int) {
	v.fieldIndex = index
}

//FieldIndex retrieves the index of the field the validation was applied to
func (v *Validation) FieldIndex() int {
	return v.fieldIndex
}

//SetFieldName stores the name of the field the validation was applied to
func (v *Validation) SetFieldName(name string) {
	v.fieldName = name
}

//FieldName retrieves the name of the field the validation was applied to
func (v *Validation) FieldName() string {
	return v.fieldName
}

//Validate determines if the value is valid. Nil is returned if it is valid
func (v *Validation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	return &ValidationError{
		Key:     v.fieldName,
		Message: "Validation not implemented",
	}
}

//DefaultMap is the default validation map used to tell if a struct is valid.
var DefaultMap = Map{}

//Map is an atomic validation map when two Set happen at the same time, latest that started wins.
type Map struct {
	validator               sync.Map // map[reflect.Type][]Interface
	validationNameToBuilder sync.Map // map[string]func(string, reflect.Kind) (Interface, error)
}

//get
func (vm *Map) get(k reflect.Type) []Interface {
	v, ok := vm.validator.Load(k)
	if !ok {
		return []Interface{}
	}
	return v.([]Interface)
}

//set
func (vm *Map) set(k reflect.Type, v []Interface) {
	vm.validator.Store(k, v)
}

//AddValidation registers the validation specified by key to the known
// validations. If more than one validation registers with the same key, the
// last one will become the validation for that key
// using DefaultValidationMap.
func AddValidation(key string, fn func(string, reflect.Kind) (Interface, error)) {
	DefaultMap.AddValidation(key, fn)
}

//AddValidation registers the validation specified by key to the known
// validations. If more than one validation registers with the same key, the
// last one will become the validation for that key.
func (vm *Map) AddValidation(key string, fn func(string, reflect.Kind) (Interface, error)) {
	vm.validationNameToBuilder.Store(key, fn)
}

//IsValid determines if an object is valid based on its validation tags using DefaultValidationMap.
func IsValid(object interface{}) (bool, []ValidationError) {
	return DefaultMap.IsValid(object)
}

//IsValid determines if an object is valid based on its validation tags.
func (vm *Map) IsValid(object interface{}) (bool, []ValidationError) {

	//Get the object's value and type
	objectValue := reflect.ValueOf(object)
	objectType := reflect.TypeOf(object)

	//Get the validations
	validations := vm.get(objectType)

	//Run IsValid our value is the pointer and not nil
	if objectValue.Kind() == reflect.Ptr && !objectValue.IsNil() {
		return IsValid(objectValue.Elem().Interface())
	}

	//Do we have some validations?
	if len(validations) == 0 {
		var err error

		//Loop the tags
		for i := objectType.NumField() - 1; i >= 0; i-- {
			field := objectType.Field(i)
			validationTag := field.Tag.Get("validation")

			//Do we have a validation tag?
			if len(validationTag) > 0 {
				validationComponent := strings.Split(validationTag, " ")

				//Loop each validation component
				for _, v := range validationComponent {
					component := strings.Split(v, "=")
					if len(component) != 2 {
						log.Fatalln("invalid validation specification:", objectType.Name(), field.Name, v)
					}

					//Create the validation
					var validation Interface
					if builder, ok := vm.validationNameToBuilder.Load(component[0]); ok && builder != nil {
						fn := builder.(func(string, reflect.Kind) (Interface, error))
						validation, err = fn(component[1], field.Type.Kind())

						if err != nil {
							log.Fatalln("error creating validation:", objectType.Name(), field.Name, v, err)
						}
					} else {
						log.Fatalln("unknown validation named:", component[0])
					}

					//Store the other properties
					validation.SetFieldName(field.Name)
					validation.SetFieldIndex(i)
					validations = append(validations, validation)
				}
			}
		}

		//Set the validations
		vm.set(objectType, validations)
	}

	//Loop and build errors
	var errors []ValidationError
	for _, validation := range validations {
		field := objectValue.Field(validation.FieldIndex())
		value := field.Interface()
		if err := validation.Validate(value, objectValue); err != nil {
			errors = append(errors, *err)
		}
	}

	//Return flag and errors
	return len(errors) == 0, errors
}
