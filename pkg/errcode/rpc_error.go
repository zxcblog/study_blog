package errcode

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zxcblog/study_blog/pb/base"
)

type Status struct {
	*status.Status
}

func ToGrpcStatus(err *Error) *Status {
	s, _ := status.New(ToRpcCode(err.Code()), err.Msg()).WithDetails(&base.Error{
		Code:    int32(err.Code()),
		Message: err.Msg(),
	})
	return &Status{s}
}

func ToGrpcError(err *Error) error {
	s, _ := status.New(ToRpcCode(err.Code()), err.Msg()).WithDetails(&base.Error{
		Code:    int32(err.Code()),
		Message: err.Msg(),
	})
	return s.Err()
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func ToRpcCode(code int) codes.Code {
	switch code {
	case Fail.Code():
		return codes.Internal
	case InvalidParams.Code():
		return codes.InvalidArgument
	case Unauthorized.Code():
		return codes.Unauthenticated
	case AccessDenied.Code():
		return codes.PermissionDenied
	case DeadlineExceeded.Code():
		return codes.DeadlineExceeded
	case NotFound.Code():
		return codes.NotFound
	case LimitExceed.Code():
		return codes.ResourceExhausted
	case MethodNotAllowed.Code():
		return codes.Unimplemented
	}
	return codes.Unknown
}
