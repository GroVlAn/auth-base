package httpx

import (
	"errors"
	"net/http"

	"github.com/GroVlAn/auth-base/ew"
)

const (
	internalServerError = "internal server error"
)

var httpCodes = map[ew.ErrorType]int{
	ew.ErrorTypeNotFound:     http.StatusNotFound,
	ew.ErrorTypeConflict:     http.StatusConflict,
	ew.ErrorTypeUnauthorized: http.StatusUnauthorized,
	ew.ErrorTypeInternal:     http.StatusInternalServerError,
}

type ResponseError struct {
	Status  int
	Message string
	Fields  []ew.ValidationField `json:"fields"`
}

func HandleError(err error) ResponseError {
	var errValidation *ew.ErrValidation
	var errWrapper *ew.Error

	if errors.As(err, &errValidation) {
		return ResponseError{
			Status:  http.StatusBadRequest,
			Message: errValidation.Error(),
			Fields:  errValidation.Fields(),
		}
	}

	if errors.As(err, &errWrapper) {
		return ResponseError{
			Status:  httpCodes[errWrapper.ErrorType()],
			Message: errWrapper.Error(),
			Fields:  nil,
		}
	}

	return ResponseError{
		Status:  http.StatusInternalServerError,
		Message: internalServerError,
		Fields:  nil,
	}
}
