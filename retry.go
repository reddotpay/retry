package retry

import (
	"time"
)

// Retryable defines a retryable function
type Retryable struct {
	// Func defines a function to be retried
	Func func() error

	// Attempts defines the number of retries if an error occurs
	Attempts uint

	// Delay defines a delay function that returns time.Duration
	Delay func() time.Duration
}

// NoDelay defines a delay function with no delays
func NoDelay() time.Duration { return 0 }

// Execute executes a Retryable
func Execute(retryable Retryable) error {
	var (
		n    uint
		errs Errors
	)

	for n < retryable.Attempts {
		err := retryable.Func()
		if err == nil {
			return nil
		}

		// Add err to err list
		errs = append(errs, err)

		n++

		time.Sleep(retryable.Delay())
	}

	return errs
}
