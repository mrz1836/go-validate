package validate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzIsValidEmail(f *testing.F) {
	// Seed corpus with various email formats and edge cases
	f.Add("test@example.com", true)
	f.Add("test@example.com", false)
	f.Add("user@domain.co.uk", true)
	f.Add("invalid@", false)
	f.Add("@invalid.com", false)
	f.Add("", false)
	f.Add("a@b.c", true)
	f.Add("user+tag@example.com", false)
	f.Add("user.name@example-domain.com", true)
	f.Add("test@test", false)
	f.Add("user@domain@domain.com", false)
	f.Add(strings.Repeat("a", 250)+"@example.com", false)
	f.Add("user@"+strings.Repeat("a", 250)+".com", false)
	f.Add("test@aol.con", false)     // blacklisted domain
	f.Add("test@example.com", false) // blacklisted domain
	f.Add("unicode-test@example.com", false)
	f.Add("test@[192.168.1.1]", false)

	f.Fuzz(func(t *testing.T, email string, mxCheck bool) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidEmail panicked with input %q, mxCheck %t: %v", email, mxCheck, r)
			}
		}()

		success, err := IsValidEmail(email, mxCheck)

		// Function should never panic
		_ = success
		_ = err

		// Basic sanity checks
		if len(email) == 0 {
			require.False(t, success, "Empty email should not be valid")
			require.Error(t, err, "Empty email should return error")
		}

		if len(email) > 254 {
			require.False(t, success, "Email longer than 254 chars should not be valid")
			require.Error(t, err, "Long email should return error")
		}

		if len(email) < 5 && len(email) > 0 {
			require.False(t, success, "Email shorter than 5 chars should not be valid")
			require.Error(t, err, "Short email should return error")
		}
	})
}

func FuzzIsValidSocial(f *testing.F) {
	// Seed corpus with various SSN formats and edge cases
	f.Add("123-45-6789")
	f.Add("987-65-4321") // blacklisted
	f.Add("000-12-3456") // invalid first part
	f.Add("123-00-4567") // invalid second part
	f.Add("123-45-0000") // invalid third part
	f.Add("666-12-3456") // invalid first part
	f.Add("123456789")   // no hyphens
	f.Add("12-345-6789") // wrong format
	f.Add("")
	f.Add("abc-de-fghi")
	f.Add("123-45-67890")      // too long
	f.Add("12-34-5678")        // too short
	f.Add("   123-45-6789   ") // with whitespace
	f.Add("111-11-1111")       // blacklisted
	f.Add("555-55-5555")       // blacklisted

	f.Fuzz(func(t *testing.T, social string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidSocial panicked with input %q: %v", social, r)
			}
		}()

		success, err := IsValidSocial(social)

		// Function should never panic
		_ = success
		_ = err

		// Basic sanity checks
		trimmed := strings.TrimSpace(social)
		if len(trimmed) == 0 {
			require.False(t, success, "Empty social should not be valid")
			require.Error(t, err, "Empty social should return error")
		}

		// Check blacklisted socials
		for _, blacklisted := range blacklistedSocials {
			if strings.Contains(social, blacklisted) || social == blacklisted {
				if success {
					t.Logf("Blacklisted social %q was marked as valid", social)
				}
			}
		}
	})
}

func FuzzIsValidPhoneNumber(f *testing.F) {
	// Seed corpus with various phone number formats and country codes
	f.Add("5551234567", "1")
	f.Add("15551234567", "1")
	f.Add("(555) 123-4567", "1")
	f.Add("555-123-4567", "1")
	f.Add("555.123.4567", "1")
	f.Add("5551234567", "52") // Mexico
	f.Add("12345678", "52")   // Mexico 8-digit
	f.Add("1234567890", "52") // Mexico 10-digit
	f.Add("", "1")
	f.Add("123", "1")
	f.Add("0551234567", "1")   // Invalid NPA start
	f.Add("1551234567", "1")   // Invalid NPA start
	f.Add("5550234567", "1")   // Invalid NXX start
	f.Add("5551114567", "1")   // Invalid NXX N11
	f.Add("5551234567", "")    // Empty country code
	f.Add("5551234567", "999") // Invalid country code
	f.Add("abc1234567", "1")
	f.Add("5551234567", "52")
	f.Add("5551234567", "44") // UK (not supported)

	f.Fuzz(func(t *testing.T, phone, countryCode string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidPhoneNumber panicked with phone %q, countryCode %q: %v", phone, countryCode, r)
			}
		}()

		success, err := IsValidPhoneNumber(phone, countryCode)

		// Function should never panic
		_ = success
		_ = err

		// Basic sanity checks
		if len(phone) == 0 {
			require.False(t, success, "Empty phone should not be valid")
			require.Error(t, err, "Empty phone should return error")
		}

		if len(countryCode) == 0 || len(countryCode) > 3 {
			require.False(t, success, "Invalid country code length should not be valid")
			require.Error(t, err, "Invalid country code length should return error")
		}
	})
}

func FuzzIsValidEnum(f *testing.F) {
	// Seed corpus with various enum values and allowed values
	f.Add("red", true)
	f.Add("RED", true)
	f.Add("blue", false)
	f.Add("", true)
	f.Add("", false)
	f.Add("green", true)
	f.Add("invalid", false)
	f.Add("BLUE", true)
	f.Add("Green", false)
	f.Add("white", true)
	f.Add(strings.Repeat("a", 1000), false)

	f.Fuzz(func(t *testing.T, enum string, emptyValueAllowed bool) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidEnum panicked with enum %q, emptyValueAllowed %t: %v", enum, emptyValueAllowed, r)
			}
		}()

		// Create test allowed values
		allowedValues := []string{"red", "green", "blue", "Red", "GREEN"}

		success, err := IsValidEnum(enum, &allowedValues, emptyValueAllowed)

		// Function should never panic
		_ = success
		_ = err

		// Basic sanity checks
		if len(enum) == 0 && emptyValueAllowed {
			require.True(t, success, "Empty value should be valid when allowed")
			require.NoError(t, err, "Empty value should not return error when allowed")
		} else if len(enum) == 0 && !emptyValueAllowed {
			require.False(t, success, "Empty value should not be valid when not allowed")
			require.Error(t, err, "Empty value should return error when not allowed")
		}

		// Test case insensitive matching
		for _, allowed := range allowedValues {
			if strings.EqualFold(enum, allowed) {
				require.True(t, success, "Case-insensitive match should be valid: %q vs %q", enum, allowed)
				require.NoError(t, err, "Case-insensitive match should not return error")
				return
			}
		}
	})
}

func FuzzIsValidHost(f *testing.F) {
	// Seed corpus with various host formats
	f.Add("example.com")
	f.Add("192.168.1.1")
	f.Add("2001:db8::1")
	f.Add("localhost")
	f.Add("sub.domain.example.com")
	f.Add("192.168.1.256") // Invalid IP
	f.Add("::1")
	f.Add("127.0.0.1")
	f.Add("")
	f.Add("invalid..domain")
	f.Add(".invalid")
	f.Add("invalid.")
	f.Add("host-name")
	f.Add("host_name")
	f.Add("123.abc.def")
	f.Add(strings.Repeat("a", 256))

	f.Fuzz(func(t *testing.T, host string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidHost panicked with input %q: %v", host, r)
			}
		}()

		result := IsValidHost(host)
		_ = result // Should never panic

		// Also test individual functions
		ipResult := IsValidIP(host)
		dnsResult := IsValidDNSName(host)
		_ = ipResult
		_ = dnsResult

		// Test specific IP versions
		ipv4Result := IsValidIPv4(host)
		ipv6Result := IsValidIPv6(host)
		_ = ipv4Result
		_ = ipv6Result

		// Basic sanity check
		if result {
			require.True(t, ipResult || dnsResult, "Valid host should be either valid IP or DNS name")
		}
	})
}

func FuzzIsValidDNSName(f *testing.F) {
	// Seed corpus with various DNS name formats
	f.Add("example.com")
	f.Add("sub.domain.example.com")
	f.Add("localhost")
	f.Add("host-name")
	f.Add("host123")
	f.Add("123host")
	f.Add("")
	f.Add(".")
	f.Add("..")
	f.Add("host.")
	f.Add(".host")
	f.Add("host..name")
	f.Add("very-long-subdomain-name-that-might-exceed-limits.example.com")
	f.Add(strings.Repeat("a", 255))
	f.Add(strings.Repeat("a", 256))
	f.Add("host_with_underscore")
	f.Add("192.168.1.1") // Should be invalid as DNS (it's an IP)

	f.Fuzz(func(t *testing.T, dnsName string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidDNSName panicked with input %q: %v", dnsName, r)
			}
		}()

		result := IsValidDNSName(dnsName)
		_ = result // Should never panic

		// Basic sanity checks
		if len(dnsName) == 0 {
			require.False(t, result, "Empty DNS name should not be valid")
		}

		if len(strings.ReplaceAll(dnsName, ".", "")) > 255 {
			require.False(t, result, "DNS name longer than 255 chars (without dots) should not be valid")
		}

		// If it's a valid IP, it should not be a valid DNS name
		if IsValidIP(dnsName) {
			require.False(t, result, "Valid IP should not be valid DNS name: %q", dnsName)
		}
	})
}

func FuzzIsValidIP(f *testing.F) {
	// Seed corpus with various IP address formats
	f.Add("192.168.1.1")
	f.Add("127.0.0.1")
	f.Add("0.0.0.0")
	f.Add("255.255.255.255")
	f.Add("192.168.1.256") // Invalid
	f.Add("2001:db8::1")
	f.Add("::1")
	f.Add("fe80::1")
	f.Add("2001:db8:85a3::8a2e:370:7334")
	f.Add("invalid")
	f.Add("")
	f.Add("192.168.1")
	f.Add("192.168.1.1.1")
	f.Add(":::")
	f.Add("127.0.0.1:80")     // With port
	f.Add("[2001:db8::1]:80") // IPv6 with port

	f.Fuzz(func(t *testing.T, ipAddress string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidIP panicked with input %q: %v", ipAddress, r)
			}
		}()

		result := IsValidIP(ipAddress)
		ipv4Result := IsValidIPv4(ipAddress)
		ipv6Result := IsValidIPv6(ipAddress)

		// Should never panic
		_ = result
		_ = ipv4Result
		_ = ipv6Result

		// Basic sanity checks
		if result {
			require.True(t, ipv4Result || ipv6Result, "Valid IP should be either IPv4 or IPv6")
		}

		if ipv4Result {
			require.Contains(t, ipAddress, ".", "IPv4 should contain dots")
			require.NotContains(t, ipAddress, ":", "IPv4 should not contain colons")
		}

		if ipv6Result {
			require.Contains(t, ipAddress, ":", "IPv6 should contain colons")
		}

		if len(ipAddress) == 0 {
			require.False(t, result, "Empty IP should not be valid")
			require.False(t, ipv4Result, "Empty IP should not be valid IPv4")
			require.False(t, ipv6Result, "Empty IP should not be valid IPv6")
		}
	})
}
