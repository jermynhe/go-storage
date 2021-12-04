package endpoint

import (
	"errors"
	"fmt"
)

var (
	// ErrUnsupportedProtocol will return if protocol is unsupported.
	//
	// Deprecated: Moved to github.com/minhjh/go-endpoint
	ErrUnsupportedProtocol = errors.New("unsupported protocol")
	// ErrInvalidValue means value is invalid.
	//
	// Deprecated: Moved to github.com/minhjh/go-endpoint
	ErrInvalidValue = errors.New("invalid value")
)

// Error represents error related to endpoint.
//
// Deprecated: Moved to github.com/minhjh/go-endpoint
type Error struct {
	Op  string
	Err error

	Protocol string
	Values   []string
}

func (e *Error) Error() string {
	if e.Values == nil {
		return fmt.Sprintf("%s: %s: %s", e.Op, e.Protocol, e.Err.Error())
	}
	return fmt.Sprintf("%s: %s, %s: %s", e.Op, e.Protocol, e.Values, e.Err.Error())
}

// Unwrap implements xerrors.Wrapper
func (e *Error) Unwrap() error {
	return e.Err
}
