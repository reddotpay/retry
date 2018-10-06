# retry
--
    import "github.com/reddotpay/retry"


## Usage

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

#### type Retryer

```go
type Retryer struct {
	// Attempts defines the number of retries if an error occurs
	Attempts uint

	// Delay defines a delay function that returns time.Duration
	Delay func() time.Duration
}
```

Retryer defines a Retryer function

#### func (Retryer) Try

```go
func (retryer Retryer) Try(fn func() error) error
```
Try executes `fn` and retries on error
