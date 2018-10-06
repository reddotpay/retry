package retry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/reddotpay/retry"
	"github.com/stretchr/testify/assert"
)

func TestRetry_Execute(t *testing.T) {
	retryer := retry.Retryer{
		Attempts: 1,
		Delay:    func() time.Duration { return 0 },
	}
	err := retryer.Try(func() error { return nil })

	assert := assert.New(t)
	assert.NoError(err)
}

func TestRetry_Execute_WithErrors(t *testing.T) {
	var count uint

	retryer := retry.Retryer{
		Attempts: 3,
		Delay:    retry.NoDelay,
	}
	err := retryer.Try(func() error {
		count++
		return errors.New("an error occurred")
	})

	assert := assert.New(t)
	assert.Equal(uint(3), count)
	assert.Equal(err.Error(), "an error occurred 3 time(s)\n")
}

func BenchmarkRetry_Execute(b *testing.B) {
	for n := 0; n < b.N; n++ {
		retryer := retry.Retryer{
			Attempts: 1,
			Delay:    func() time.Duration { return 0 },
		}
		retryer.Try(func() error { return nil })
	}
}

func BenchmarkRetry_Execute_WithErrors(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var count uint

		retryer := retry.Retryer{
			Attempts: 3,
			Delay:    retry.NoDelay,
		}
		retryer.Try(func() error {
			count++
			return errors.New("an error occurred")
		})
	}
}
