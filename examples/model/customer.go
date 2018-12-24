/*
Package main is an example of using validate in a basic model (validating data before persisting into a database)
*/
package main

import (
	"fmt"

	"github.com/mrz1836/go-validate"
)

//Customer is a very basic model showing validation mixed with JSON
type Customer struct {
	Age                  uint    `validation:"min=18" json:"age"`
	Balance              float32 `validation:"min=0" json:"balance"`
	Class                string  `validation:"min_length=5 max_length=10" json:"class"`
	Email                string  `validation:"format=email" json:"email"`
	Name                 string  `validation:"format=regexp:[A-Z][a-z]{3,12}" json:"name"`
	Password             string  `validation:"compare=PasswordConfirmation" json:"-"`
	PasswordConfirmation string  `json:"-"`
	Phone                string  `json:"phone"`
	Region               uint    `validation:"min=1 max=5" json:"region"`
	SocialSecurityNumber string  `json:"-"`
}

//Valid is a custom method for the model that will run all built-in validations and also run any custom validations
func (c *Customer) Valid() (bool, []validate.ValidationError) {

	//Run all validations off the struct configuration
	_, errs := validate.IsValid(*c)

	//
	//Customize: errs (you can add/remove your own errs)
	//

	//Showing use of a public validation method (extra validations outside of the struct built-in validations)
	ok, err := validate.IsValidSocial(c.SocialSecurityNumber)
	if !ok {
		errs = append(errs, validate.ValidationError{
			Key:     "SocialSecurityNumber",
			Message: err.Error(),
		})
	}

	//Return error if found
	return len(errs) == 0, errs
}

//Add your custom validations
func init() {

	//Add your own custom validations
	//
	//
}

//main example (just an example of validating a model's data before persisting into a database)
func main() {

	//Start with some model & data
	customer := &Customer{
		Age:                  21,
		Balance:              1.00,
		Class:                "executive",
		Email:                "john@protonmail.com",
		Name:                 "John",
		Password:             "MyNewPassword123!",
		PasswordConfirmation: "MyNewPassword123",
		Region:               2,
		SocialSecurityNumber: "212126768",
	}

	//Validate the model - Run before saving into database
	_, errs := customer.Valid()
	if errs != nil {
		fmt.Printf("Customer validation failed! %+v", errs)
	} else {
		fmt.Println("Customer validation succeed!")
	}

	//customer.Save() //Now save to a database
}
