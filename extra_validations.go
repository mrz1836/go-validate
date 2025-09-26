package validate

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strings"
)

const (
	// socialBasicRawRegex social Security number regex for validation
	socialBasicRawRegex = `^\d{3}-\d{2}-\d{4}$`
)

// Package-level variables for validation data.
// These are global by design to provide shared validation data across the package.
var (

	// blacklistedSocials known blacklisted socials (exclude automatically)
	blacklistedSocials = []string{ //nolint:gochecknoglobals // Shared validation data
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

	// numericRegExp numeric regex
	numericRegExp = regexp.MustCompile(`[^0-9]`)

	// blacklistedDomains known blacklisted domains for email addresses
	// Global for shared access across validation functions
	blacklistedDomains = []string{ //nolint:gochecknoglobals // Shared validation data
		"aol.con",     // Does not exist, but valid TLD in regex
		"example.com", // Invalid domain - used for testing but should not work in production
		"gmail.con",   // Does not exist, but valid TLD in regex
		"gnail.com",   // Does not exist, but valid TLD in regex
		"hotmail.con", // Does not exist, but valid TLD in regex
		"yahoo.con",   // Does not exist, but valid TLD in regex
	}

	// acceptedCountryCodes is the countries this phone number validation can currently accept
	// Global for shared access across validation functions
	acceptedCountryCodes = []string{ //nolint:gochecknoglobals // Shared validation data
		"1",  // USA and CAN
		"52", // Mexico
		// todo: support more countries in phone number validation (@mrz)
	}

	// dnsRegEx is the regex for a DNS name
	dnsRegEx = regexp.MustCompile(`^([a-zA-Z0-9_][a-zA-Z0-9_-]{0,62})(\.[a-zA-Z0-9_][a-zA-Z0-9_-]{0,62})*[._]?$`)
)

// IsValidEnum validates an enum given the required parameters and tests if the supplied value is valid from accepted values
func IsValidEnum(enum string, allowedValues *[]string, emptyValueAllowed bool) (success bool, err error) {
	// Empty is true and no value given?
	if emptyValueAllowed && len(enum) == 0 {
		success = true
		return success, err
	}

	// Check that the value is an allowed value (case-insensitive)
	for _, value := range *allowedValues {
		// Compare both in the lowercase
		if strings.EqualFold(enum, value) {
			success = true
			return success, err
		}
	}

	// We must have an error
	err = fmt.Errorf("%w: %s", ErrEnumValueNotAllowed, enum)
	return success, err
}

// IsValidEmail validate an email address using regex, checking name and host, and even MX record check
func IsValidEmail(email string, mxCheck bool) (success bool, err error) {
	// Minimum / Maximum sizes
	if len(email) < 5 || len(email) > 254 {
		err = ErrEmailLengthInvalid
		return success, err
	}

	// Validate first using regex
	if !emailRegex.MatchString(email) {
		err = ErrEmailFormatInvalid
		return success, err
	}

	// Find the @ sign (redundant with regex being first)
	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		err = ErrEmailMissingAtSign
		return success, err
	}

	// More than one at sign?
	if strings.Count(email, "@") > 1 {
		err = ErrEmailMultipleAtSigns
		return success, err
	}

	// Split the user and host
	user := email[:at]
	host := email[at+1:]

	// User cannot be more than 64 characters
	if len(user) > 64 {
		err = ErrEmailLengthInvalid
		return success, err
	}

	// Invalid domains
	// Check banned/blacklisted numbers
	if ok, _ := IsValidEnum(host, &blacklistedDomains, false); ok {
		err = ErrEmailDomainNotAccepted
		return success, err
	}

	// Validate the host
	if ok := IsValidHost(host); !ok {
		err = ErrEmailDomainInvalidHost
		return success, err
	}

	// Check for mx record or A record
	if mxCheck {
		ctx := context.Background()
		resolver := &net.Resolver{}
		if _, err = resolver.LookupMX(ctx, host); err != nil {
			if _, err = resolver.LookupIPAddr(ctx, host); err != nil {
				// Only fail if both MX and A records are missing - any of the
				// two is enough for an email to be deliverable
				err = fmt.Errorf("%w: %s", ErrEmailDomainCannotReceive, err.Error())
				return success, err
			}
		}
	}

	// All good
	success = true
	return success, err
}

// IsValidSocial validates the USA social security number using ATS rules
func IsValidSocial(social string) (success bool, err error) {
	// Sanitize
	social = strings.TrimSpace(social)

	// No value?
	if len(social) == 0 {
		err = ErrSocialEmpty
		return success, err
	}

	// Determine if it is missing hyphens
	if count := strings.Count(social, "-"); count != 2 {

		// Reduce to only numbers
		social = string(numericRegExp.ReplaceAll([]byte(social), []byte("")))

		// We do NOT have 9 digits
		if len(social) != 9 {
			err = ErrSocialLengthInvalid
			return success, err
		}

		// Break it up
		firstPart := social[0:3]
		secondPart := social[3:5]
		thirdPart := social[5:9]

		// Build it back up
		social = firstPart + "-" + secondPart + "-" + thirdPart
	}

	// Check the basics
	if match, _ := regexp.MatchString(socialBasicRawRegex, social); !match {
		err = ErrSocialRegexMismatch
		return success, err
	}

	// Break into three parts
	firstPart := social[0:3]
	secondPart := social[4:6]
	thirdPart := social[7:11]

	// Split the first section (not 000 or 666)
	if firstPart == "000" || firstPart == "666" || secondPart == "00" || thirdPart == "0000" {
		err = ErrSocialSectionInvalid
		return success, err
	}

	// Check banned/blacklisted numbers
	if ok, _ := IsValidEnum(social, &blacklistedSocials, false); ok {
		err = ErrSocialBlacklisted
		return success, err
	}

	// All good!
	success = true
	return success, err
}

// validateCountryCode validates and sanitizes the country code
func validateCountryCode(countryCode string) (string, error) {
	if len(countryCode) == 0 || len(countryCode) > 3 {
		return "", ErrCountryCodeLengthInvalid
	}

	// Sanitize the code
	countryCode = string(numericRegExp.ReplaceAll([]byte(countryCode), []byte("")))

	// Country code is not accepted
	if ok, _ := IsValidEnum(countryCode, &acceptedCountryCodes, false); !ok {
		return "", fmt.Errorf("%w: %s", ErrCountryCodeNotAccepted, countryCode)
	}

	return countryCode, nil
}

// validateUSACanadaPhone validates USA/Canada phone numbers (country code 1)
func validateUSACanadaPhone(phone string) error {
	if len(phone) != 10 {
		return ErrPhoneMustBeTenDigits
	}

	// Break up the phone number into NPA-NXX-XXXX
	npa := phone[0:3]
	nxx := phone[3:6]
	firstDigitOfNpa := npa[0:1]
	firstDigitOfNxx := nxx[0:1]
	secondThirdDigitOfNxx := nxx[1:3]

	// NPA Cannot start with 1 or 0
	if firstDigitOfNpa == "1" || firstDigitOfNpa == "0" {
		return fmt.Errorf("%w: %s", ErrPhoneNPAInvalidStart, firstDigitOfNpa)
	}

	// NPA Cannot contain 555 as leading value
	if npa == "555" {
		return fmt.Errorf("%w: %s", ErrPhoneNPAInvalidStart, npa)
	}

	// NXX Cannot start with 1 or 0
	if firstDigitOfNxx == "1" || firstDigitOfNxx == "0" {
		return fmt.Errorf("%w: %s", ErrPhoneNXXInvalidDigits, "cannot start with "+firstDigitOfNxx)
	}

	// NXX cannot be N11
	if secondThirdDigitOfNxx == "11" {
		return fmt.Errorf("%w: %s", ErrPhoneNXXInvalidDigits, "cannot be X"+secondThirdDigitOfNxx)
	}

	return nil
}

// validateMexicoPhone validates Mexico phone numbers (country code 52)
func validateMexicoPhone(phone string) error {
	// Validate the proper length
	if len(phone) != 8 && len(phone) != 10 { // 2002 mexico had 8-digit numbers and went to 10 digits
		return ErrPhoneMustBeEightOrTen
	}

	// Break up the phone number into NPA-NXX-XXXX
	npa := phone[0:3]
	firstDigitOfNpa := npa[0:1]

	// NPA Cannot start with 1 or 0
	if firstDigitOfNpa == "1" || firstDigitOfNpa == "0" {
		return fmt.Errorf("%w: %s", ErrPhoneNPAInvalidStart, firstDigitOfNpa)
	}

	return nil
}

// IsValidPhoneNumber validates a given phone number and country code
func IsValidPhoneNumber(phone, countryCode string) (success bool, err error) {
	// Validate and sanitize country code
	countryCode, err = validateCountryCode(countryCode)
	if err != nil {
		return false, err
	}

	// No phone number
	if len(phone) == 0 {
		return false, ErrPhoneLengthInvalid
	}

	// Sanitize the phone
	phone = string(numericRegExp.ReplaceAll([]byte(phone), []byte("")))

	// Phone number format validation by country code
	switch countryCode {
	case "1": // USA and CAN
		if err := validateUSACanadaPhone(phone); err != nil {
			return false, err
		}
	case "52": // Mexico
		if err := validateMexicoPhone(phone); err != nil {
			return false, err
		}
	default:
		return false, fmt.Errorf("%w: %s", ErrCountryCodeNotAccepted, countryCode)
	}

	return true, nil
}

// IsValidHost checks if the string is a valid IP (both v4 and v6) or a valid DNS name
func IsValidHost(host string) bool {
	return IsValidIP(host) || IsValidDNSName(host)
}

// IsValidIP checks if a string is either IP version 4 or 6. Alias for `net.ParseIP`
func IsValidIP(ipAddress string) bool {
	return net.ParseIP(ipAddress) != nil
}

// IsValidIPv4 check if the string is IP version 4.
func IsValidIPv4(ipAddress string) bool {
	ip := net.ParseIP(ipAddress)
	return ip != nil && ip.To4() != nil && !strings.Contains(ipAddress, ":")
}

// IsValidIPv6 check if the string is IP version 6.
func IsValidIPv6(ipAddress string) bool {
	return net.ParseIP(ipAddress) != nil && strings.Contains(ipAddress, ":")
}

// IsValidDNSName will validate the given string as a DNS name
func IsValidDNSName(dnsName string) bool {
	if dnsName == "" || len(strings.ReplaceAll(dnsName, ".", "")) > 255 {
		return false
	}
	return !IsValidIP(dnsName) && dnsRegEx.MatchString(dnsName)
}
