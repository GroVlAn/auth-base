package grpcx

import (
	"errors"

	"github.com/GroVlAn/auth-base/ew"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	internalServerError = "internal server error"
)

var grpcCodes = map[ew.ErrorType]codes.Code{
	ew.ErrorTypeNotFound:     codes.NotFound,
	ew.ErrorTypeConflict:     codes.AlreadyExists,
	ew.ErrorTypeUnauthorized: codes.Unauthenticated,
	ew.ErrorTypeInternal:     codes.Internal,
}

func HandleError(err error) error {
	var errValidation *ew.ErrValidation
	var errWrapper *ew.Error

	if errors.As(err, &errValidation) {
		st := status.New(
			codes.InvalidArgument,
			errValidation.Error(),
		)

		br := &errdetails.BadRequest{}

		for _, field := range errValidation.Fields() {
			br.FieldViolations = append(
				br.FieldViolations,
				&errdetails.BadRequest_FieldViolation{
					Field:       field.Field,
					Description: field.Reason,
				},
			)
		}

		st, err := st.WithDetails(br)
		if err != nil {
			return status.Error(
				codes.Internal,
				"failed tot attach validation details",
			)
		}

		return st.Err()
	}

	if errors.As(err, &errWrapper) {
		return status.Error(
			grpcCodes[errWrapper.ErrorType()],
			errWrapper.Error(),
		)
	}

	return status.Error(
		codes.Internal,
		internalServerError,
	)
}
