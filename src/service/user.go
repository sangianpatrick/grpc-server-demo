package service

import (
	"github.com/sangianpatrick/grpc-service-demo/src/module/user/handler"
	"github.com/sangianpatrick/grpc-service-demo/src/pb"
	"google.golang.org/grpc"
)

// InitializeUser initialize user domain
func InitializeUser(grpcServer *grpc.Server) {
	grpcUserServiceServer := handler.NewUserGRPCHandler()
	pb.RegisterUserServiceServer(grpcServer, grpcUserServiceServer)
}
