package ew

type ErrorType int

const (
	ErrorTypeNotFound ErrorType = iota
	ErrorTypeConflict
	ErrorTypeUnauthorized
	ErrorTypeInternal
)

type Error struct {
	errType ErrorType
	err     error
	msg     string
}

func New(errType ErrorType, err error, msg string) *Error {
	return &Error{
		errType: errType,
		err:     err,
		msg:     msg,
	}
}

func (ew *Error) Error() string {
	return ew.msg
}

func (ew *Error) Unwrap() error {
	return ew.err
}

func (ew *Error) ErrorType() ErrorType {
	return ew.errType
}

func (e ErrorType) String() string {
	switch e {
	case ErrorTypeNotFound:
		return "not_found"
	case ErrorTypeConflict:
		return "conflict"
	case ErrorTypeUnauthorized:
		return "unauthorized"
	case ErrorTypeInternal:
		return "internal"
	default:
		return "unknown"
	}
}
