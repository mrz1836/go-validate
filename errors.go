package validate

import "errors"

// Static error definitions to satisfy err113 linter
var (
	// Enum validation errors
	ErrEnumValueNotAllowed = errors.New("value is not allowed")

	// Email validation errors
	ErrEmailLengthInvalid       = errors.New("email length is invalid")
	ErrEmailFormatInvalid       = errors.New("email is not a valid address format")
	ErrEmailMissingAtSign       = errors.New("email is missing the @ sign")
	ErrEmailMultipleAtSigns     = errors.New("email contains more than one @ sign")
	ErrEmailDomainNotAccepted   = errors.New("email domain is not accepted")
	ErrEmailDomainInvalidHost   = errors.New("email domain is not a valid host")
	ErrEmailDomainCannotReceive = errors.New("email domain invalid/cannot receive mail")

	// Social Security validation errors
	ErrSocialEmpty          = errors.New("social is empty")
	ErrSocialLengthInvalid  = errors.New("social is not nine digits in length")
	ErrSocialRegexMismatch  = errors.New("social does not match the regex pattern")
	ErrSocialSectionInvalid = errors.New("social section was found invalid (cannot be 000 or 666)")
	ErrSocialBlacklisted    = errors.New("social was found to be blacklisted")

	// Phone number validation errors
	ErrCountryCodeLengthInvalid = errors.New("country code length is invalid")
	ErrCountryCodeNotAccepted   = errors.New("country code is not accepted")
	ErrPhoneLengthInvalid       = errors.New("phone number length is invalid")
	ErrPhoneMustBeTenDigits     = errors.New("phone number must be ten digits")
	ErrPhoneNPAInvalidStart     = errors.New("phone number NPA cannot start with specified digit")
	ErrPhoneNXXInvalidDigits    = errors.New("phone number NXX cannot be specified digits")
	ErrPhoneMustBeEightOrTen    = errors.New("phone number must be either eight or ten digits")
)
