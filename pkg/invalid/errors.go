package invalid

import "fmt"

type ErrDuplicate struct {
	Resource string
	Column   string
	Value    interface{}
}

func (err ErrDuplicate) Error() string {
	return fmt.Sprintf("Duplicate %s{%s: %v} ", err.Resource, err.Column, err.Value)
}

func IsErrDuplicate(err error) bool {
	_, ok := err.(ErrDuplicate)
	return ok
}

type ErrNotFound struct {
	Resource string
	Column   string
	Value    interface{}
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("%s{%s: %v} Not Found", err.Resource, err.Column, err.Value)
}

func IsErrorNotFound(err error) bool {
	_, ok := err.(ErrNotFound)
	return ok
}
