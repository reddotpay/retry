package retry_test

import (
	"errors"
	"testing"

	"github.com/reddotpay/retry"
	"github.com/stretchr/testify/assert"
)

func TestRetry_Error(t *testing.T) {
	errs := retry.Errors{
		errors.New("an error occurred"),
		errors.New("another error occurred"),
		errors.New("an error occurred"),
	}

	assert := assert.New(t)
	assert.Equal("an error occurred 2 time(s)\nanother error occurred 1 time(s)\n", errs.Error())
}
