package validate

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testPhone               = "234-234-2345"
	testMxCountryCode       = "+52"
	emailInvalidFormatError = "email is not a valid address format"
)

// TestIsValidSocial testing the social security number (invalid and valid)
func TestIsValidSocial(t *testing.T) {
	tests := []struct {
		name     string
		social   string
		expected bool
		desc     string
	}{
		{"all zeros", "000-00-0000", false, "All zeros"},
		{"not enough digits", "434-43-433", false, "Not enough digits"},
		{"too many digits", "434-43-4334444", false, "Too many digits"},
		{"starts with 666", "666-00-0000", false, "Starts with 666"},
		{"first section zeros", "000-12-1235", false, "First section zeros"},
		{"middle section zeros", "888-00-1235", false, "Middle section zeros"},
		{"last section zeros", "888-14-0000", false, "Last section zeros"},
		{"valid with dashes", "434-43-4334", true, "Valid"},
		{"valid without dashes 1", "323126767", true, "Valid"},
		{"valid without dashes 2", "212126768", true, "Valid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := IsValidSocial(tt.social)
			if tt.expected {
				assert.True(t, ok, "Social test failed when it should have passed: %s", tt.social)
				assert.NoError(t, err)
			} else {
				assert.False(t, ok, "Social test passed when it should have failed: %s", tt.social)
			}
		})
	}

	// Test all blacklisted socials
	t.Run("blacklisted socials", func(t *testing.T) {
		for _, ssn := range blacklistedSocials {
			ok, _ := IsValidSocial(ssn)
			assert.False(t, ok, "Blacklisted social should be invalid: %s", ssn)
		}
	})
}

// TestIsValidSocialErrorResponses test the error responses
func TestIsValidSocialErrorResponses(t *testing.T) {
	tests := []struct {
		name          string
		social        string
		expectedError string
	}{
		{"empty", "", "social is empty"},
		{"not long enough", "2123", "social is not nine digits in length"},
		{"all zeros", "000-00-0000", "social section was found invalid (cannot be 000 or 666)"},
		{"invalid pattern", "1234-1-2342", "social does not match the regex pattern"},
		{"invalid section", "000-00-2345", "social section was found invalid (cannot be 000 or 666)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := IsValidSocial(tt.social)
			require.Error(t, err)
			assert.Equal(t, tt.expectedError, err.Error())
		})
	}

	// Test blacklisted socials
	t.Run("blacklisted socials", func(t *testing.T) {
		for _, ssn := range blacklistedSocials {
			_, err := IsValidSocial(ssn)
			require.Error(t, err)
			assert.Equal(t, "social was found to be blacklisted", err.Error())
		}
	})
}

// ExampleIsValidSocial_invalid example of an invalid social response
func ExampleIsValidSocial_invalid() {
	ok, err := IsValidSocial("666-00-0000") // Invalid
	fmt.Println(ok, err)
	// Output: false social section was found invalid (cannot be 000 or 666)
}

// ExampleIsValidSocial_invalid example of a valid social response
func ExampleIsValidSocial_valid() {
	ok, err := IsValidSocial("212126768") // Valid
	fmt.Println(ok, err)
	// Output: true <nil>
}

// BenchmarkIsValidSocial benchmarks the IsValidSocial (valid value)
func BenchmarkIsValidSocial(b *testing.B) {
	testSocialNumber := "212126768"

	for i := 0; i < b.N; i++ {
		_, _ = IsValidSocial(testSocialNumber)
	}
}

// TestIsValidEmail testing the email validation
func TestIsValidEmail(t *testing.T) { //nolint:gocognit // Test function complexity acceptable
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

	// ========= VALID EMAILS ================

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

	// ========= INVALID MX EMAILS ================

	email = "d-d-d.d.dt@testthissite.com.uk"
	if success, _ = IsValidEmail(email, true); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "tester@example.com"
	if success, _ = IsValidEmail(email, true); success {
		t.Fatal("Email value should be invalid! Value: " + email)
	}

	email = "yolanda@615.yt@gmail.com"
	var err error
	if success, err = IsValidEmail(email, true); success {
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

	// ========= TEST ALL TLDS ================

	// Test all blacklisted hosts and errors
	for _, host := range blacklistedDomains {
		if ok, err := IsValidEmail("someone@"+host, false); ok {
			t.Fatal("This should have failed, host is blacklisted", host, err)
		}
	}
}

// TestIsValidEmailErrorResponses test the error responses
func TestIsValidEmailErrorResponses(t *testing.T) {
	var err error

	email := "test"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != "email length is invalid" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "test@"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != emailInvalidFormatError {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "test.some"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != emailInvalidFormatError {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "@this.com"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != emailInvalidFormatError {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	email = "1234567890123456789012345678901234567890123456789012345678901234567890@this.com"
	if _, err = IsValidEmail(email, false); err != nil && err.Error() != "email length is invalid" {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}

	// Test all blacklisted hosts and errors
	for _, host := range blacklistedDomains {
		if _, err = IsValidEmail("someone@"+host, false); err != nil && err.Error() != "email domain is not accepted" {
			t.Fatal("Error response was not as expected - Value: " + err.Error())
		}
	}

	email = "someone@gmail.conn"
	_, err = IsValidEmail(email, true)

	// email domain invalid/cannot receive mail: lookup gmail.conn on 169.254.169.254:53: no such host
	// if err.Error() != "email domain invalid/cannot receive mail: lookup gmail.conn: no such host" {
	if err != nil && !strings.Contains(err.Error(), "email domain invalid/cannot receive mail:") {
		t.Fatal("Error response was not as expected - Value: " + err.Error())
	}
}

// ExampleIsValidEmail_invalid example of an invalid email address response
func ExampleIsValidEmail_invalid() {
	ok, err := IsValidEmail("notvalid@domain", false) // Invalid
	fmt.Println(ok, err)
	// Output: false email is not a valid address format
}

// ExampleIsValidEmail_valid example of a valid email address response
func ExampleIsValidEmail_valid() {
	ok, err := IsValidEmail("person@gmail.com", false) // Valid
	fmt.Println(ok, err)
	// Output: true <nil>
}

// BenchmarkIsValidEmail benchmarks the IsValidEmail (valid value)
func BenchmarkIsValidEmail(b *testing.B) {
	testEmail := "person@gmail.com"

	for i := 0; i < b.N; i++ {
		_, _ = IsValidEmail(testEmail, false)
	}
}

// TestIsValidEnum testing the enum value in an accepted list of values
func TestIsValidEnum(t *testing.T) {
	var ok bool
	var err error

	// Invalid value
	testEnumValue := "1"
	testAcceptedValues := []string{"123"}
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, false); ok {
		t.Fatal("This should have failed - value is not found", testEnumValue, testAcceptedValues, err)
	} else if err != nil && err.Error() != "value is not allowed: "+testEnumValue {
		t.Fatal("error message was not as expected", err.Error())
	}

	// Valid value
	testEnumValue = "123"
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, false); !ok {
		t.Fatal("This should have passed - value is valid", testEnumValue, testAcceptedValues, err)
	}

	// Empty valid is not allowed
	testEnumValue = ""
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, false); ok {
		t.Fatal("This should have failed - can be empty flag", testEnumValue, testAcceptedValues, err)
	}

	// Empty value allowed
	testEnumValue = ""
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, true); !ok {
		t.Fatal("This should have passed - can be empty flag", testEnumValue, testAcceptedValues, err)
	}

	// Test case-insensitive
	testEnumValue = "mystring"
	testAcceptedValues = []string{"myString"}
	if ok, err = IsValidEnum(testEnumValue, &testAcceptedValues, false); !ok {
		t.Fatal("This should have passed - can be empty flag", testEnumValue, testAcceptedValues, err)
	}
}

// ExampleIsValidEnum_invalid example of an invalid enum
func ExampleIsValidEnum_invalid() {
	testAcceptedValues := []string{"123"}
	ok, err := IsValidEnum("1", &testAcceptedValues, false) // Invalid
	fmt.Println(ok, err)
	// Output: false value is not allowed: 1
}

// ExampleIsValidEnum_valid example of an valid enum
func ExampleIsValidEnum_valid() {
	testAcceptedValues := []string{"123"}
	ok, err := IsValidEnum("123", &testAcceptedValues, false) // Valid
	fmt.Println(ok, err)
	// Output: true <nil>
}

// BenchmarkIsValidEnum benchmarks the IsValidEnum (valid value)
func BenchmarkIsValidEnum(b *testing.B) {
	testValue := "1"
	testAcceptedValues := []string{"123"}
	for i := 0; i < b.N; i++ {
		_, _ = IsValidEnum(testValue, &testAcceptedValues, false)
	}
}

// TestIsValidPhoneNumber testing the phone value
func TestIsValidPhoneNumber(t *testing.T) { //nolint:gocognit // Test function complexity acceptable
	var ok bool
	var err error

	// Missing country code (too short)
	phone := ""
	countryCode := ""

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "country code length is invalid" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Invalid country code (too long)
	countryCode = "6666"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "country code length is invalid" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Country code is not accepted
	countryCode = "+32"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "country code is not accepted: 32" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number missing
	countryCode = "+1" // USA / CAN

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number length is invalid" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number not right length (USA)
	countryCode = "+1" // USA / CAN
	phone = "555-444-3"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number must be ten digits" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number not right length (MX)
	phone = "555-444-3"

	if ok, err = IsValidPhoneNumber(phone, testMxCountryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, testMxCountryCode, err)
	} else if err != nil && err.Error() != "phone number must be either eight or ten digits" {
		t.Fatal("error message was not as expected", phone, testMxCountryCode, err.Error())
	}

	// Phone number not right length (USA)
	countryCode = "+1" // USA / CAN
	phone = "234-444-3"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number must be ten digits" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number cannot start with 1
	countryCode = "+1" // USA / CAN
	phone = "123-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with specified digit: 1" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number cannot start with 0
	countryCode = "+1" // USA / CAN
	phone = "023-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with specified digit: 0" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number cannot start with 1
	phone = "123-123-1234"

	if ok, err = IsValidPhoneNumber(phone, testMxCountryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, testMxCountryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with specified digit: 1" {
		t.Fatal("error message was not as expected", phone, testMxCountryCode, err.Error())
	}

	// Phone number cannot start with 0
	phone = "023-123-1234"

	if ok, err = IsValidPhoneNumber(phone, testMxCountryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, testMxCountryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with specified digit: 0" {
		t.Fatal("error message was not as expected", phone, testMxCountryCode, err.Error())
	}

	// Phone number cannot start with 555
	countryCode = "+1" // USA / CAN
	phone = "555-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NPA cannot start with specified digit: 555" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number NXX cannot start with 1
	countryCode = "+1" // USA / CAN
	phone = "234-123-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NXX cannot be specified digits: cannot start with 1" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number NXX cannot start with 0
	countryCode = "+1" // USA / CAN
	phone = "234-023-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NXX cannot be specified digits: cannot start with 0" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number NXX cannot be X11
	countryCode = "+1" // USA / CAN
	phone = "234-911-1234"

	if ok, err = IsValidPhoneNumber(phone, countryCode); ok {
		t.Fatal("This should have failed - phone is invalid for USA", phone, countryCode, err)
	} else if err != nil && err.Error() != "phone number NXX cannot be specified digits: cannot be X11" {
		t.Fatal("error message was not as expected", phone, countryCode, err.Error())
	}

	// Phone number valid
	countryCode = "+1" // USA / CAN

	if ok, err = IsValidPhoneNumber(testPhone, countryCode); !ok {
		t.Fatal("This should have passed - phone is valid for USA", testPhone, countryCode, err)
	}
}

// ExampleIsValidPhoneNumber_invalid example of an invalid phone number
func ExampleIsValidPhoneNumber_invalid() {
	countryCode := "+1"
	phone := "555-444-44"
	ok, err := IsValidPhoneNumber(phone, countryCode)
	fmt.Println(ok, err)
	// Output: false phone number must be ten digits
}

// ExampleIsValidPhoneNumber_valid example of a valid phone number
func ExampleIsValidPhoneNumber_valid() {
	countryCode := "+1"
	ok, err := IsValidPhoneNumber(testPhone, countryCode)
	fmt.Println(ok, err)
	// Output: true <nil>
}

// BenchmarkIsValidPhoneNumber benchmarks the IsValidPhoneNumber (valid value)
func BenchmarkIsValidPhoneNumber(b *testing.B) {
	countryCode := "+1"
	for i := 0; i < b.N; i++ {
		_, _ = IsValidPhoneNumber(testPhone, countryCode)
	}
}

// TestIsValidDNSName testing the dns name
func TestIsValidDNSName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		param    string
		expected bool
	}{
		{"localhost", true},
		{"a.bc", true},
		{"a.b.", true},
		{"a.b..", false},
		{"localhost.local", true},
		{"localhost.localdomain.intern", true},
		{"l.local.intern", true},
		{"ru.link.n.svpncloud.com", true},
		{"-localhost", false},
		{"localhost.-localdomain", false},
		{"localhost.localdomain.-int", false},
		{"_localhost", true},
		{"localhost._localdomain", true},
		{"localhost.localdomain._int", true},
		{"lÖcalhost", false},
		{"localhost.lÖcaldomain", false},
		{"localhost.localdomain.üntern", false},
		{"__", true},
		{"localhost/", false},
		{"127.0.0.1", false},
		{"[::1]", false},
		{"50.50.50.50", false},
		{"localhost.localdomain.intern:65535", false},
		{"漢字汉字", false}, //nolint:gosmopolitan // ignore for now
		{"www.jubfvq1v3p38i51622y0dvmdk1mymowjyeu26gbtw9andgynj1gg8z3msb1kl5z6906k846pj3sulm4kiyk82ln5teqj9nsht59opr0cs5ssltx78lfyvml19lfq1wp4usbl0o36cmiykch1vywbttcus1p9yu0669h8fj4ll7a6bmop505908s1m83q2ec2qr9nbvql2589adma3xsq2o38os2z3dmfh2tth4is4ixyfasasasefqwe4t2ub2fz1rme.de", false},
	}

	for _, test := range tests {
		actual := IsValidDNSName(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsValidDNSName(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

// ExampleIsValidDNSName_invalid example of an invalid dns name
func ExampleIsValidDNSName_invalid() {
	ok := IsValidDNSName("localhost.-localdomain")
	fmt.Println(ok)
	// Output: false
}

// ExampleIsValidDNSName_valid example of a valid dns name
func ExampleIsValidDNSName_valid() {
	ok := IsValidDNSName("localhost")
	fmt.Println(ok)
	// Output: true
}

// BenchmarkIsValidDNSName benchmarks the IsValidDNSName (valid value)
func BenchmarkIsValidDNSName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsValidDNSName("localhost")
	}
}

// TestIsValidIP will test IPv4 and IPv6
func TestIsValidIP(t *testing.T) {
	t.Parallel()

	// Without version
	tests := []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsValidIP(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsValidIP(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// IPv4
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", false},
		{"2001:db8:0000:1:1:1:1:1", false},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsValidIPv4(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsValidIPv4(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// IPv6
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", false},
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"1.2.3.4", false},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsValidIPv6(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsValidIPv6(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

// ExampleIsValidIP_invalid example of an invalid ip address
func ExampleIsValidIP_invalid() {
	ok := IsValidIP("300.0.0.0")
	fmt.Println(ok)
	// Output: false
}

// ExampleIsValidIP_valid example of a valid ip address
func ExampleIsValidIP_valid() {
	ok := IsValidIP("127.0.0.1")
	fmt.Println(ok)
	// Output: true
}

// BenchmarkIsValidIP benchmarks the IsValidIP (valid value)
func BenchmarkIsValidIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsValidIP("127.0.0.1")
	}
}

// TestIsValidHost will test a hostname
func TestIsValidHost(t *testing.T) {
	t.Parallel()
	tests := []struct {
		param    string
		expected bool
	}{
		{"localhost", true},
		{"localhost.localdomain", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"::1", true},
		{"play.golang.org", true},
		{"localhost.localdomain.intern:65535", false},
		{"-[::1]", false},
		{"-localhost", false},
		{".localhost", false},
	}
	for _, test := range tests {
		actual := IsValidHost(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsValidHost(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

// ExampleIsValidHost_invalid example of an invalid host
func ExampleIsValidHost_invalid() {
	ok := IsValidHost("-localhost")
	fmt.Println(ok)
	// Output: false
}

// ExampleIsValidHost_valid example of a valid host
func ExampleIsValidHost_valid() {
	ok := IsValidHost("localhost")
	fmt.Println(ok)
	// Output: true
}

// BenchmarkIsValidHost benchmarks the IsValidHost (valid value)
func BenchmarkIsValidHost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsValidHost("localhost")
	}
}
