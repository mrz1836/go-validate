package validate

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

//Extra validation variables
var (

	//blacklistedSocials known blacklisted socials (exclude automatically)
	blacklistedSocials = []string{
		"123-45-6789",
		"219-09-9999",
		"078-05-1120",
		"987-65-4320",
		"987-65-4321",
		"987-65-4322",
		"987-65-4323",
		"987-65-4324",
		"987-65-4325",
		"987-65-4326",
		"987-65-4327",
		"987-65-4328",
		"987-65-4329",
		"111-11-1111",
		"222-22-2222",
		"333-33-3333",
		"444-44-4444",
		"555-55-5555",
		"777-77-7777",
		"888-88-8888",
		"999-99-9999",
		"012-34-5678",
	}

	//numericRegExp numeric regex
	numericRegExp = regexp.MustCompile(`[^0-9]`)

	//blacklistedDomains known blacklisted domains for email addresses
	blacklistedDomains = []string{
		"aol.con",     //Does not exist, but valid TLD in regex
		"example.com", //Invalid domain - used for testing but should not work in production
		"gmail.con",   //Does not exist, but valid TLD in regex
		"hotmail.con", //Does not exist, but valid TLD in regex
		"yahoo.con",   //Does not exist, but valid TLD in regex
	}

	//acceptedCountryCodes is the countries this phone number validation can currently accept
	acceptedCountryCodes = []string{
		"1",  //USA and CAN
		"52", //Mexico
		//todo: support more countries in phone number validation (@mrz)
	}
)

//Extra validation constants
const (
	//socialBasicRawRegex social Security number regex for validation
	socialBasicRawRegex = `^\d{3}-\d{2}-\d{4}$`
)

//IsValidEnum validates an enum given the required parameters and tests if the supplied value is valid from accepted values
func IsValidEnum(enum string, allowedValues *[]string, emptyValueAllowed bool) (success bool, err error) {

	//Empty is true and no value given?
	if emptyValueAllowed == true && len(enum) == 0 {
		success = true
		return
	}

	//Check that the value is an allowed value (case insensitive)
	for _, value := range *allowedValues {

		//Compare both in lowercase
		if strings.ToLower(enum) == strings.ToLower(value) {
			success = true
			return
		}
	}

	//We must have an error
	err = fmt.Errorf("value %s is not allowed", enum)
	return
}

//IsValidEmail validate an email address using regex, checking name and host, and even MX record check
func IsValidEmail(email string, mxCheck bool) (success bool, err error) {

	//Minimum / Maximum sizes
	if len(email) < 5 || len(email) > 254 {
		err = fmt.Errorf("email length is invalid")
		return
	}

	//Validate first using regex
	if !emailRegex.Match([]byte(email)) {
		err = fmt.Errorf("email is not a valid address format")
		return
	}

	//Find the @ sign (redundant with regex being first)
	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		err = fmt.Errorf("email is missing the @ sign")
		return
	}

	//More than one at sign?
	if strings.Count(email, "@") > 1 {
		err = fmt.Errorf("email contains more than one @ sign")
		return
	}

	//Split the user and host
	user := email[:at]
	host := email[at+1:]

	//User cannot be more than 64 characters
	if len(user) > 64 {
		err = fmt.Errorf("email length is invalid")
		return
	}

	//Invalid domains
	//Check banned/blacklisted numbers
	ok, _ := IsValidEnum(host, &blacklistedDomains, false)
	if ok {
		err = fmt.Errorf("email domain is not accepted")
		return
	}

	//Check for mx record or A record
	if mxCheck {
		if _, err = net.LookupMX(host); err != nil {
			if _, err = net.LookupIP(host); err != nil {
				// Only fail if both MX and A records are missing - any of the
				// two is enough for an email to be deliverable
				err = fmt.Errorf("email domain invalid/cannot receive mail: " + err.Error())
				return
			}
		}
	}

	//All good
	success = true
	return
}

//IsValidSocial validates the USA social security number using ATS rules
func IsValidSocial(social string) (success bool, err error) {

	//Sanitize
	social = strings.TrimSpace(social)

	//No value?
	if len(social) == 0 {
		err = fmt.Errorf("social is empty")
		return
	}

	//Determine if it is missing hyphens
	if count := strings.Count(social, "-"); count != 2 {

		//Reduce to only numbers
		social = string(numericRegExp.ReplaceAll([]byte(social), []byte("")))

		//We do NOT have 9 digits
		if len(social) != 9 {
			err = fmt.Errorf("social is not nine digits in length")
			return
		}

		//Break it up
		firstPart := social[0:3]
		secondPart := social[3:5]
		thirdPart := social[5:9]

		//Build it back up
		social = firstPart + "-" + secondPart + "-" + thirdPart
	}

	//Check the basics
	if match, _ := regexp.MatchString(socialBasicRawRegex, social); !match {
		err = fmt.Errorf("social does not match the regex pattern")
		return
	}

	//Break into three parts
	firstPart := social[0:3]
	secondPart := social[4:6]
	thirdPart := social[7:11]

	//Split the first section (not 000 or 666)
	if firstPart == "000" || firstPart == "666" || secondPart == "00" || thirdPart == "0000" {
		err = fmt.Errorf("social section was found invalid (cannot be 000 or 666)")
		return
	}

	//Check banned/blacklisted numbers
	ok, _ := IsValidEnum(social, &blacklistedSocials, false)
	if ok {
		err = fmt.Errorf("social was found to be blacklisted")
		return
	}

	//All good!
	success = true
	return
}

//IsValidPhoneNumber validates a given phone number and country code
func IsValidPhoneNumber(phone string, countryCode string) (success bool, err error) {

	//No country code or country code is greater than expected
	if len(countryCode) == 0 || len(countryCode) > 3 {
		err = fmt.Errorf("country code length is invalid")
		return
	}

	//Sanitize the code
	countryCode = string(numericRegExp.ReplaceAll([]byte(countryCode), []byte("")))

	//Country code not accepted
	ok, _ := IsValidEnum(countryCode, &acceptedCountryCodes, false)
	if !ok {
		err = fmt.Errorf("country code %s is not accepted", countryCode)
		return
	}

	//No phone number
	if len(phone) == 0 {
		err = fmt.Errorf("phone number length is invalid")
		return
	}

	//Sanitize the phone
	phone = string(numericRegExp.ReplaceAll([]byte(phone), []byte("")))

	//Phone number format does not match the country code
	switch countryCode {
	case "1": //USA and CAN

		//Validate the proper length
		if len(phone) != 10 {
			err = fmt.Errorf("phone number must be ten digits")
			return
		}

		//todo: finish validation on NPA / NXX

	case "52": //Mexico

		//Validate the proper length
		if len(phone) != 10 {
			err = fmt.Errorf("phone number must be ten digits")
			return
		}

		//todo: validate MX number

	default:
		err = fmt.Errorf("country code %s is not accepted", countryCode)
		return
	}

	//All good
	success = true
	return
}
