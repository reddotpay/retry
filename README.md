# retry
--
    import "github.com/reddotpay/retry"


## Usage

#### func  Execute

```go
func Execute(retryable Retryable) error
```
Execute executes a Retryable

#### func  NoDelay

```go
func NoDelay() time.Duration
```
NoDelay defines a delay function with no delays

#### type Errors

```go
type Errors []error
```

Errors contains a list of error that occurred

#### func (Errors) Error

```go
func (errs Errors) Error() string
```

#### type Retryable

```go
type Retryable struct {
	// Func defines a function to be retried
	Func func() error

	// Attempts defines the number of retries if an error occurs
	Attempts uint

	// Delay defines a delay function that returns time.Duration
	Delay func() time.Duration
}
```

Retryable defines a retryable function
