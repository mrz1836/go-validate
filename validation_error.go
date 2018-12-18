/*
Package validation (go-validation) provides validations for struct fields based on a validation tag and offers additional validation functions.
*/
package validation

import "fmt"

//Error is key and message of the corresponding error
type Error struct {

	//Field name, key name
	Key string

	//Error message
	Message string
}

//Error returns a string of key + message
func (e *Error) Error() string {
	return e.Key + " " + e.Message
}

//Errors is a slice of validation errors
type Errors []Error

//Error returns the list of errors from the slice of errors
func (e Errors) Error() (errors string) {

	//No errors?
	if len(e) == 0 {
		return
	}

	//Get the first error
	errors = e[0].Error()

	//Add x other errors
	if len(e) > 1 {
		errors += fmt.Sprintf(" and %d other errors", len(e))
	}

	return
}
