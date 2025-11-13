package grpcerrs

import (
	"log/slog"
	"strconv"

	"google.golang.org/grpc/codes"
)

type Code uint32

const (
	OK                 Code = 0
	Canceled           Code = 1
	Unknown            Code = 2
	InvalidArgument    Code = 3
	DeadlineExceeded   Code = 4
	NotFound           Code = 5
	AlreadyExists      Code = 6
	PermissionDenied   Code = 7
	ResourceExhausted  Code = 8
	FailedPrecondition Code = 9
	Aborted            Code = 10
	OutOfRange         Code = 11
	Unimplemented      Code = 12
	Internal           Code = 13
	Unavailable        Code = 14
	DataLoss           Code = 15
	Unauthenticated    Code = 16
)

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}

func (c Code) GRPCStatusCode() codes.Code {
	switch c {
	case OK:
		return codes.OK
	case Canceled:
		return codes.Canceled
	case Unknown:
		return codes.Unknown
	case InvalidArgument:
		return codes.InvalidArgument
	case DeadlineExceeded:
		return codes.DeadlineExceeded
	case NotFound:
		return codes.NotFound
	case AlreadyExists:
		return codes.AlreadyExists
	case PermissionDenied:
		return codes.PermissionDenied
	case ResourceExhausted:
		return codes.ResourceExhausted
	case FailedPrecondition:
		return codes.FailedPrecondition
	case Aborted:
		return codes.Aborted
	case OutOfRange:
		return codes.OutOfRange
	case Unimplemented:
		return codes.Unimplemented
	case Internal:
		return codes.Internal
	case Unavailable:
		return codes.Unavailable
	case DataLoss:
		return codes.DataLoss
	case Unauthenticated:
		return codes.Unauthenticated
	}
	return codes.Unknown
}

func (c Code) SlogLevel() slog.Level {
	switch c {
	case
		Canceled,
		NotFound,
		AlreadyExists:
		return slog.LevelInfo
	case
		Unauthenticated,
		InvalidArgument,
		ResourceExhausted,
		FailedPrecondition,
		PermissionDenied,
		Unavailable,
		DeadlineExceeded,
		Aborted,
		OutOfRange:
		return slog.LevelWarn
	case
		Unimplemented,
		Unknown,
		Internal,
		DataLoss:
		return slog.LevelError
	default:
		return slog.LevelError
	}
}
