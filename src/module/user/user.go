package user

import (
	"context"

	"github.com/sangianpatrick/grpc-service-demo/src/pb"
)

// MariadbRepository is a collection of behaviors of mariadb
type MariadbRepository interface {
	InsertOne(ctx context.Context, user *pb.User) (err error)
	FindByUsername(ctx context.Context, username string) (user *pb.User, err error)
	FindByUsernameMobileNumberEmail(ctx context.Context, username, mobileNumber, email string) (user *pb.User, err error)
}

// Usecase is a collection of behaviors of user domain
type Usecase interface {
	Register(ctx context.Context, user *pb.User) (err error)
	GetByUsername(ctx context.Context, username string) (user *pb.User, err error)
}
