package grpcerrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Collection of errors
var (
	ErrAlreadyExist     = status.Error(codes.AlreadyExists, "Data is already exist")
	ErrInternal         = status.Error(codes.Internal, "An error occurred while processing request")
	ErrDeadlineExceeded = status.Error(codes.DeadlineExceeded, "Deadline Exceeded")
	ErrNotFound         = status.Error(codes.NotFound, "Data is not found")
)
