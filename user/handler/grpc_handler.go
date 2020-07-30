package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sangianpatrick/grpc-service-demo/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userGRPCHandler struct{}

// NewUserGRPCHandler is a constructor
func NewUserGRPCHandler() pb.UserServiceServer {
	return userGRPCHandler{}
}

func (userGRPCHandler) Register(ctx context.Context, pbUser *pb.User) (*empty.Empty, error) {
	userBuff, _ := json.Marshal(pbUser)

	fmt.Println("NAME", pbUser.Name)

	fmt.Println(string(userBuff))

	return nil, status.Error(codes.InvalidArgument, "bla bla")
}
