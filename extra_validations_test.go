package validate

import (
	"fmt"
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
	if _, err = IsValidSocial(testSocialNumber); err.Error() != "social is empty" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	testSocialNumber = "2123" //Not long enough
	if _, err = IsValidSocial(testSocialNumber); err.Error() != "social is not nine digits in length" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	testSocialNumber = "000-00-0000" //All zeros
	if _, err = IsValidSocial(testSocialNumber); err.Error() != "social section was found invalid (cannot be 000 or 666)" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	testSocialNumber = "1234-1-2342" //Invalid Pattern
	if _, err = IsValidSocial(testSocialNumber); err.Error() != "social does not match the regex pattern" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	testSocialNumber = "000-00-2345" //Invalid section
	if _, err = IsValidSocial(testSocialNumber); err.Error() != "social section was found invalid (cannot be 000 or 666)" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	for _, ssn := range blacklistedSocials { //Blacklisted
		if _, err = IsValidSocial(ssn); err.Error() != "social was found to be blacklisted" {
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
	if _, err = IsValidEmail(email, false); err.Error() != "email length is invalid" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "test@"
	if _, err = IsValidEmail(email, false); err.Error() != "email is not a valid address format" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "test.some"
	if _, err = IsValidEmail(email, false); err.Error() != "email is not a valid address format" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "@this.com"
	if _, err = IsValidEmail(email, false); err.Error() != "email is not a valid address format" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "1234567890123456789012345678901234567890123456789012345678901234567890@this.com"
	if _, err = IsValidEmail(email, false); err.Error() != "email length is invalid" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	//Test all blacklisted hosts and errors
	for _, host := range blacklistedDomains {
		if _, err = IsValidEmail("someone@"+host, false); err.Error() != "email domain is not accepted" {
			t.Fatal("Error response was not as expected - Value: " + err.Error())
		}
	}

	email = "someone@gmail.conn"
	if _, err = IsValidEmail(email, true); err.Error() != "email domain invalid/cannot receive mail: lookup gmail.conn: no such host" {
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
