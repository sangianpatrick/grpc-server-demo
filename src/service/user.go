package service

import (
	"database/sql"

	"github.com/sangianpatrick/grpc-service-demo/src/module/user/repository"
	"github.com/sangianpatrick/grpc-service-demo/src/module/user/usecase"
	"github.com/sirupsen/logrus"

	"github.com/sangianpatrick/grpc-service-demo/src/module/user/handler"
	"github.com/sangianpatrick/grpc-service-demo/src/pb"
	"google.golang.org/grpc"
)

// InitializeUser initialize user domain
func InitializeUser(logger *logrus.Logger, grpcServer *grpc.Server, db *sql.DB) {

	umr := repository.NewUserMariadbRepository(db)
	uu := usecase.NewUserUsecase(logger, umr)

	grpcUserServiceServer := handler.NewUserGRPCHandler(uu)
	pb.RegisterUserServiceServer(grpcServer, grpcUserServiceServer)
}
