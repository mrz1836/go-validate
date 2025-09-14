package validate

import "sync"

var initOnce sync.Once //nolint:gochecknoglobals // Initialization synchronization

// InitValidations registers all built-in validations.
// This function is safe to call multiple times and will only execute once.
// It must be called before using any validation features.
func InitValidations() {
	initOnce.Do(func() {
		RegisterStringValidations()
		RegisterNumericValidations()
	})
}
