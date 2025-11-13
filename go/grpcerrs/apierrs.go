package grpcerrs

import (
	"fmt"
	"runtime"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type APIError struct {
	code               Code
	message            string
	location           string
	stackTrace         string
	originalGRPCStatus *status.Status
}

func (a *APIError) Error() string {
	return fmt.Sprintf("[%s] %s", a.code.String(), a.message)
}

func (a *APIError) Code() Code {
	return a.code
}

func (a *APIError) Stack() string {
	return a.stackTrace
}

func (a *APIError) Location() string {
	return a.location
}

func (a *APIError) GRPCStatus() *status.Status {
	return status.New(a.code.GRPCStatusCode(), a.message)
}

func newAPIError(code Code, message string) *APIError {
	location := "unknown"

	// 引数は遡るstack frameの数
	// 例えばGetCost() -> InternalError() -> newAPIError()と呼ばれていればGetCost()のfile, lineを指す
	_, file, line, ok := runtime.Caller(2)
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)

	return &APIError{
		code:       code,
		message:    message,
		location:   location,
		stackTrace: string(buf[:n]),
	}
}

func CanceledError(msg string) *APIError {
	return newAPIError(Canceled, msg)
}

func UnknownError(msg string) *APIError {
	return newAPIError(Unknown, msg)
}

func InvalidArgumentError(msg string) *APIError {
	return newAPIError(InvalidArgument, msg)
}

func DeadlineExceededError(msg string) *APIError {
	return newAPIError(DeadlineExceeded, msg)
}

func NotFoundError(msg string) *APIError {
	return newAPIError(NotFound, msg)
}

func AlreadyExistsError(msg string) *APIError {
	return newAPIError(AlreadyExists, msg)
}

func PermissionDeniedError(msg string) *APIError {
	return newAPIError(PermissionDenied, msg)
}

func ResourceExhaustedError(msg string) *APIError {
	return newAPIError(ResourceExhausted, msg)
}

func FailedPreconditionError(msg string) *APIError {
	return newAPIError(FailedPrecondition, msg)
}

func AbortedError(msg string) *APIError {
	return newAPIError(Aborted, msg)
}

func OutOfRangeError(msg string) *APIError {
	return newAPIError(OutOfRange, msg)
}

func UnimplementedError(msg string) *APIError {
	return newAPIError(Unimplemented, msg)
}

func InternalError(msg string) *APIError {
	return newAPIError(Internal, msg)
}

func UnavailableError(msg string) *APIError {
	return newAPIError(Unavailable, msg)
}

func DataLossError(msg string) *APIError {
	return newAPIError(DataLoss, msg)
}

func UnauthenticatedError(msg string) *APIError {
	return newAPIError(Unauthenticated, msg)
}

func FromGRPCStatus(err error, msg string) *APIError {
	s, ok := status.FromError(err)
	if !ok {
		return newAPIError(Internal, fmt.Sprintf("it's not a grpc error. message:%s", msg))
	}

	var code Code
	switch s.Code() {
	case codes.NotFound:
		code = NotFound
	case codes.Canceled:
		code = Canceled
	case codes.DeadlineExceeded:
		code = DeadlineExceeded
	case codes.ResourceExhausted:
		code = ResourceExhausted
	case codes.OutOfRange:
		code = OutOfRange
	case codes.DataLoss:
		code = DataLoss
	case codes.InvalidArgument:
		code = InvalidArgument
	case codes.AlreadyExists:
		code = AlreadyExists
	case codes.PermissionDenied:
		code = PermissionDenied
	case codes.FailedPrecondition:
		code = FailedPrecondition
	case codes.Aborted:
		code = Aborted
	case codes.Unimplemented:
		code = Unimplemented
	case codes.Internal:
		code = Internal
	case codes.Unavailable:
		code = Internal
	case codes.Unauthenticated:
		code = Unauthenticated
	case codes.Unknown:
		code = Unknown
	default:
		code = Internal
		msg = fmt.Sprintf("unknown grpc code:%s message:%s", s.Code().String(), msg)
	}

	e := newAPIError(code, msg)
	e.originalGRPCStatus = s
	return e
}
