package service

import (
	"github.com/sangianpatrick/grpc-service-demo/pb"
	"github.com/sangianpatrick/grpc-service-demo/user/handler"
	"google.golang.org/grpc"
)

// InitializeUser initialize user domain
func InitializeUser(grpcServer *grpc.Server) {
	grpcUserServiceServer := handler.NewUserGRPCHandler()
	pb.RegisterUserServiceServer(grpcServer, grpcUserServiceServer)
}
