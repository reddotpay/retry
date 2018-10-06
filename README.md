# retry
--
    import "github.com/reddotpay/retry"


## Usage

#### func  Execute

```go
func Execute(retryable Retryable) error
```
Execute executes a Retryable

#### type Retryable

```go
type Retryable struct {
	// Func defines a function to be retried
	Func func() error
	// Attempts defines the number of retries if an error occurs
	Attempts uint
	// Delay defines the delay between attempts
	Delay time.Duration
}
```

Retryable defines a retryable function
