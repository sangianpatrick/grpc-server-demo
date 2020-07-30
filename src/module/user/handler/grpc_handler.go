package handler

import (
	"context"

	"github.com/sangianpatrick/grpc-service-demo/src/module/user"

	"github.com/sangianpatrick/grpc-service-demo/src/pb"
)

type userGRPCHandler struct {
	uu user.Usecase
}

// NewUserGRPCHandler is a constructor
func NewUserGRPCHandler(uu user.Usecase) pb.UserServiceServer {
	return userGRPCHandler{
		uu: uu,
	}
}

func (ugh userGRPCHandler) Register(ctx context.Context, pbUser *pb.User) (empty *pb.Empty, err error) {
	err = ugh.uu.Register(ctx, pbUser)
	if err != nil {
		return
	}

	empty = &pb.Empty{}
	return
}

func (ugh userGRPCHandler) GetByUsername(ctx context.Context, request *pb.UserByUsernameRequest) (userResponse *pb.UserResponse, err error) {
	user, err := ugh.uu.GetByUsername(ctx, request.GetUsername())
	if err != nil {
		return
	}

	userResponse = new(pb.UserResponse)
	userResponse.Data = user
	return
}
