package retry

import (
	"time"
)

// Retryer defines a Retryer function
type Retryer struct {
	// Attempts defines the number of retries if an error occurs
	Attempts uint

	// Delay defines a delay function that returns time.Duration
	Delay func() time.Duration
}

// NoDelay defines a delay function with no delays
func NoDelay() time.Duration { return 0 }

// Try executes `fn` and retries on error
func (retryer Retryer) Try(fn func() error) error {
	var (
		n    uint
		errs Errors
	)

	for n < retryer.Attempts {
		err := fn()
		if err == nil {
			return nil
		}

		// Add err to err list
		errs = append(errs, err)

		n++

		time.Sleep(retryer.Delay())
	}

	return errs
}
