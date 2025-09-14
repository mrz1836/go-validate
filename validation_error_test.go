package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationErrorsErrorNoErrors(t *testing.T) {
	// Create an empty ValidationErrors slice
	var v ValidationErrors

	// Call the Error method
	result := v.Error()

	// Assert that the result is an empty string
	assert.Empty(t, result)
}
