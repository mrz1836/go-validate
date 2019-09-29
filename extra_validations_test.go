package validate

import (
	"fmt"
	"strings"
	"testing"
)

//TestIsValidSocial testing the social security number (invalid and valid)
func TestIsValidSocial(t *testing.T) {

	var ok bool
	var err error

	testSocialNumber := "000-00-0000" //All zeros
	if ok, err = IsValidSocial(testSocialNumber); ok {
		t.Fatal("Social test passed when it should have failed.", testSocialNumber, err)
	}

	testSocialNumber = "434-43-433" //Not enough digits
	if ok, err = IsValidSocial(testSocialNumber); ok {
		t.Fatal("Social test passed when it should have failed.", testSocialNumber, err)
	}

	testSocialNumber = "434-43-4334444" //Too many digits
	if ok, err = IsValidSocial(testSocialNumber); ok {
		t.Fatal("Social test passed when it should have failed.", testSocialNumber, err)
	}

	testSocialNumber = "666-00-0000" //Starts with 666
	if ok, err = IsValidSocial(testSocialNumber); ok {
		t.Fatal("Social test passed when it should have failed.", testSocialNumber, err)
	}

	testSocialNumber = "000-12-1235" //First section zeros
	if ok, err = IsValidSocial(testSocialNumber); ok {
		t.Fatal("Social test passed when it should have failed.", testSocialNumber, err)
	}

	testSocialNumber = "888-00-1235" //Middle section zeros
	if ok, err = IsValidSocial(testSocialNumber); ok {
		t.Fatal("Social test passed when it should have failed.", testSocialNumber, err)
	}

	testSocialNumber = "888-14-0000" //Last section zeros
	if ok, err = IsValidSocial(testSocialNumber); ok {
		t.Fatal("Social test passed when it should have failed.", testSocialNumber, err)
	}

	//Test all blacklisted socials
	for _, ssn := range blacklistedSocials {
		if ok, err = IsValidSocial(ssn); ok {
			t.Fatal("Social test passed when it should have failed.", testSocialNumber, err)
		}
	}

	testSocialNumber = "434-43-4334" //Valid
	if ok, err = IsValidSocial(testSocialNumber); !ok {
		t.Fatal("Social test failed when it should have passed.", testSocialNumber, err)
	}

	testSocialNumber = "323126767" //Valid
	if ok, err = IsValidSocial(testSocialNumber); !ok {
		t.Fatal("Social test failed when it should have passed.", testSocialNumber, err)
	}

	testSocialNumber = "212126768" //Valid
	if ok, err = IsValidSocial(testSocialNumber); !ok {
		t.Fatal("Social test failed when it should have passed.", testSocialNumber, err)
	}
}

//TestIsValidSocialErrorResponses test the error responses
func TestIsValidSocialErrorResponses(t *testing.T) {
	var err error

	testSocialNumber := "" //Empty
	if _, err = IsValidSocial(testSocialNumber); err != nil && err.Error() != "social is empty" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	testSocialNumber = "2123" //Not long enough
	if _, err = IsValidSocial(testSocialNumber); err != nil && err.Error() != "social is not nine digits in length" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	testSocialNumber = "000-00-0000" //All zeros
	if _, err = IsValidSocial(testSocialNumber); err != nil && err.Error() != "social section was found invalid (cannot be 000 or 666)" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	testSocialNumber = "1234-1-2342" //Invalid Pattern
	if _, err = IsValidSocial(testSocialNumber); err != nil && err.Error() != "social does not match the regex pattern" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	testSocialNumber = "000-00-2345" //Invalid section
	if _, err = IsValidSocial(testSocialNumber); err != nil && err.Error() != "social section was found invalid (cannot be 000 or 666)" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	for _, ssn := range blacklistedSocials { //Blacklisted
		if _, err = IsValidSocial(ssn); err != nil && err.Error() != "social was found to be blacklisted" {
			t.Fatal("Error response was not as expected - Value: " + err.Error())
		}
	}
}

//ExampleIsValidSocial_invalid example of an invalid social response
func ExampleIsValidSocial_invalid() {
	ok, err := IsValidSocial("666-00-0000") //Invalid
	fmt.Println(ok, err)
	// Output: false social section was found invalid (cannot be 000 or 666)
}

//ExampleIsValidSocial_invalid example of a valid social response
func ExampleIsValidSocial_valid() {
	ok, err := IsValidSocial("212126768") //Valid
	fmt.Println(ok, err)
	// Output: true <nil>
}

//BenchmarkIsValidSocial benchmarks the IsValidSocial (valid value)
func BenchmarkIsValidSocial(b *testing.B) {
	testSocialNumber := "212126768"

	for i := 0; i < b.N; i++ {
		_, _ = IsValidSocial(testSocialNumber)
	}
}

//TestIsValidEmail testing the email validation
func TestIsValidEmail(t *testing.T) {

	var success bool

	email := "test"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "test@"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "test@some"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "test.some"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "t@t"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "T@T."
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = ".com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "@.com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "a@.com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "a@..com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}
	email = "a@...com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "a@something..com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "a@something.-.com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "a@---.com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "a@a---.com"
	if success, _ = IsValidEmail(email, false); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	//========= VALID EMAILS ================

	email = "test@test.com"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@dd.com"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@t2.com"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@2t.com"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@t.co"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@dekora.fashion"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@sierra.finance"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@money.cash.co"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "jp.power.co@money.we.cash.co"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@t.co.uk"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "t@test.com.uk"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "d-d-d.d.dt@test.com.uk"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "john_doe@test.com.uk"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	email = "john+doe@test.com.uk"
	if success, _ = IsValidEmail(email, false); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	//========= INVALID MX EMAILS ================

	email = "d-d-d.d.dt@testthissite.com.uk"
	if success, _ = IsValidEmail(email, true); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "tester@example.com"
	if success, _ = IsValidEmail(email, true); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "yolanda@615.yt@gmail.com"
	if success, err := IsValidEmail(email, true); success {
		t.Fatal("Email value should be invalid! Value: "+email, success, err)
	}

	email = "d-d-d.d.dt@gmail.c"
	if success, _ = IsValidEmail(email, true); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "someone@gmail.com"
	if success, _ = IsValidEmail(email, true); !success {
		t.Fatal("Email value should be valid! Value: " + email)
	}

	//========= TEST ALL TLDS ================

	//Test all blacklisted hosts and errors
	for _, host := range blacklistedDomains {
		if ok, err := IsValidEmail("someone@"+host, false); ok {
			t.Fatal("This should have failed, host is blacklisted", host, err)
		}
	}
}

//TestIsValidEmailErrorResponses test the error responses
func TestIsValidEmailErrorResponses(t *testing.T) {
	var err error

	email := "test"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != "email length is invalid" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "test@"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != "email is not a valid address format" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "test.some"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != "email is not a valid address format" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "@this.com"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != "email is not a valid address format" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "1234567890123456789012345678901234567890123456789012345678901234567890@this.com"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != "email length is invalid" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	//Test all blacklisted hosts and errors
	for _, host := range blacklistedDomains {
		if _, err = IsValidEmail("someone@"+host, false); err != nil && err.Error() != "email domain is not accepted" {
			t.Fatal("Error response was not as expected - Value: " + err.Error())
		}
	}

	email = "someone@gmail.conn"
	_, err = IsValidEmail(email, true)

	//email domain invalid/cannot receive mail: lookup gmail.conn on 169.254.169.254:53: no such host
	//if err.Error() != "email domain invalid/cannot receive mail: lookup gmail.conn: no such host"  {
	if err != nil && !strings.Contains(err.Error(), "email domain invalid/cannot receive mail:") {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}
}

//ExampleIsValidEmail_invalid example of an invalid email address response
func ExampleIsValidEmail_invalid() {
	ok, err := IsValidEmail("notvalid@domain", false) //Invalid
	fmt.Println(ok, err)
	// Output: false email is not a valid address format
}

//ExampleIsValidEmail_valid example of a valid email address response
func ExampleIsValidEmail_valid() {
	ok, err := IsValidEmail("person@gmail.com", false) //Valid
	fmt.Println(ok, err)
	// Output: true <nil>
}

//BenchmarkIsValidEmail benchmarks the IsValidEmail (valid value)
func BenchmarkIsValidEmail(b *testing.B) {
	testEmail := "person@gmail.com"

	for i := 0; i < b.N; i++ {
		_, _ = IsValidEmail(testEmail, false)
	}
}

//TestIsValidEnum testing the enum value in an accepted list of values
func TestIsValidEnum(t *testing.T) {

	var ok bool
	var err error

	//Invalid value
	testEnumValue := "1"
	testAcceptedValues := []string{"123"}
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, false); ok {
		t.Fatal("This should have failed - value is not found", testEnumValue, testAcceptedValues, err)
	} else if err != nil && err.Error() != "value "+testEnumValue+" is not allowed" {
		t.Fatal("error message was not as expected", err.Error())
	}

	//Valid value
	testEnumValue = "123"
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, false); !ok {
		t.Fatal("This should have passed - value is valid", testEnumValue, testAcceptedValues, err)
	}

	//Empty valid not allowed
	testEnumValue = ""
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, false); ok {
		t.Fatal("This should have failed - can be empty flag", testEnumValue, testAcceptedValues, err)
	}

	//Empty value allowed
	testEnumValue = ""
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, true); !ok {
		t.Fatal("This should have passed - can be empty flag", testEnumValue, testAcceptedValues, err)
	}

}

//ExampleIsValidEnum_invalid example of an invalid enum
func ExampleIsValidEnum_invalid() {
	testAcceptedValues := []string{"123"}
	ok, err := IsValidEnum("1", &testAcceptedValues, false) //Invalid
	fmt.Println(ok, err)
	// Output: false value 1 is not allowed
}

//ExampleIsValidEnum_valid example of an valid enum
func ExampleIsValidEnum_valid() {
	testAcceptedValues := []string{"123"}
	ok, err := IsValidEnum("123", &testAcceptedValues, false) //Valid
	fmt.Println(ok, err)
	// Output: true <nil>
}

//BenchmarkIsValidEnum benchmarks the IsValidEnum (valid value)
func BenchmarkIsValidEnum(b *testing.B) {
	testValue := "1"
	testAcceptedValues := []string{"123"}
	for i := 0; i < b.N; i++ {
		_, _ = IsValidEnum(testValue, &testAcceptedValues, false)
	}
}

//TestIsValidPhoneNumber testing the phone value
func TestIsValidPhoneNumber(t *testing.T) {

	var ok bool
	var err error

	//Missing country code (too short)
	phone := ""
	countryCode := ""

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "country code length is invalid" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Invalid country code (too long)
	countryCode = "6666"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "country code length is invalid" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Country code not accepted
	countryCode = "+32"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "country code 32 is not accepted" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number missing
	countryCode = "+1" //USA / CAN

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number length is invalid" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number not right length (USA)
	countryCode = "+1" //USA / CAN
	phone = "555-444-3"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number must be ten digits" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number not right length (MX)
	countryCode = "+52" //Mexico
	phone = "555-444-3"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number must be either eight or ten digits" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number not right length (USA)
	countryCode = "+1" //USA / CAN
	phone = "234-444-3"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number must be ten digits" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number cannot start with 1
	countryCode = "+1" //USA / CAN
	phone = "123-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with 1" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number cannot start with 0
	countryCode = "+1" //USA / CAN
	phone = "023-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with 0" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number cannot start with 1
	countryCode = "+52" //MX
	phone = "123-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with 1" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number cannot start with 0
	countryCode = "+52" //MX
	phone = "023-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with 0" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number cannot start with 555
	countryCode = "+1" //USA / CAN
	phone = "555-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with 555" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number NXX cannot start with 1
	countryCode = "+1" //USA / CAN
	phone = "234-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NXX cannot start with 1" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number NXX cannot start with 0
	countryCode = "+1" //USA / CAN
	phone = "234-023-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NXX cannot start with 0" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number NXX cannot be X11
	countryCode = "+1" //USA / CAN
	phone = "234-911-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NXX cannot be X11" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	//Phone number valid
	countryCode = "+1" //USA / CAN
	phone = "234-234-2345"

	if ok, err = IsValidPhoneNumber(phone, countryCode); !ok {
		t.Fatal("This should have passed - phone is valid for USA", phone, countryCode, err)
	}
}

//ExampleIsValidPhoneNumber_invalid example of an invalid phone number
func ExampleIsValidPhoneNumber_invalid() {
	countryCode := "+1"
	phone := "555-444-44"
	ok, err := IsValidPhoneNumber(phone, countryCode)
	fmt.Println(ok, err)
	// Output: false phone number must be ten digits
}

//ExampleIsValidPhoneNumber_valid example of an valid phone number
func ExampleIsValidPhoneNumber_valid() {
	countryCode := "+1"
	phone := "234-234-2345"
	ok, err := IsValidPhoneNumber(phone, countryCode)
	fmt.Println(ok, err)
	// Output: true <nil>
}

//BenchmarkIsValidPhoneNumber benchmarks the IsValidPhoneNumber (valid value)
func BenchmarkIsValidPhoneNumber(b *testing.B) {
	countryCode := "+1"
	phone := "234-234-2345"
	for i := 0; i < b.N; i++ {
		_, _ = IsValidPhoneNumber(phone, countryCode)
	}
}
