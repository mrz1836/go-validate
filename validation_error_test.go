package validate

import "testing"

func TestValidationErrors_Error_NoErrors(t *testing.T) {
	// Create an empty ValidationErrors slice
	var v ValidationErrors

	// Call the Error method
	result := v.Error()

	// Assert that the result is an empty string
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}
