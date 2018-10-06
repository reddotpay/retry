package retry

import "fmt"

// Errors contains a list of error that occurred
type Errors []error

func (errs Errors) Error() string {
	m := make(map[string]uint)
	for _, err := range errs {
		if err != nil {
			m[err.Error()]++
		}
	}

	var s string
	for k, v := range m {
		s = s + fmt.Sprintf("%s %d time(s)\n", k, v)
	}
	return s
}
