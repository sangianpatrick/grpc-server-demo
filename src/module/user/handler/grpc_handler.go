package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sangianpatrick/grpc-service-demo/src/pb"
)

type userGRPCHandler struct{}

// NewUserGRPCHandler is a constructor
func NewUserGRPCHandler() pb.UserServiceServer {
	return userGRPCHandler{}
}

func (userGRPCHandler) Register(ctx context.Context, pbUser *pb.User) (*pb.Empty, error) {
	userBuff, _ := json.Marshal(pbUser)

	fmt.Println("NAME", pbUser.Name)

	fmt.Println(string(userBuff))

	return new(pb.Empty), nil
}
