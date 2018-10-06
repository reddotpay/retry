package retry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/reddotpay/retry"
	"github.com/stretchr/testify/assert"
)

func TestRetry_Execute(t *testing.T) {
	retryable := retry.Retryable{
		Func:     func() error { return nil },
		Attempts: 1,
		Delay:    func() time.Duration { return 0 },
	}
	err := retry.Execute(retryable)

	assert := assert.New(t)
	assert.NoError(err)
}

func TestRetry_Execute_WithErrors(t *testing.T) {
	var count uint

	retryable := retry.Retryable{
		Func: func() error {
			count++
			return errors.New("an error occurred")
		},
		Attempts: 3,
		Delay:    retry.NoDelay,
	}
	err := retry.Execute(retryable)

	assert := assert.New(t)
	assert.Equal(uint(3), count)
	assert.Equal(err.Error(), "an error occurred 3 time(s)\n")
}

func BenchmarkRetry_Execute(b *testing.B) {
	for n := 0; n < b.N; n++ {
		retryable := retry.Retryable{
			Func:     func() error { return nil },
			Attempts: 1,
			Delay:    func() time.Duration { return 0 },
		}
		retry.Execute(retryable)
	}
}

func BenchmarkRetry_Execute_WithErrors(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var count uint

		retryable := retry.Retryable{
			Func: func() error {
				count++
				return errors.New("an error occurred")
			},
			Attempts: 3,
			Delay:    retry.NoDelay,
		}
		retry.Execute(retryable)
	}
}
