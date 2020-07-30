package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/sangianpatrick/grpc-service-demo/pkg/util"
	"github.com/sirupsen/logrus"

	"github.com/sangianpatrick/grpc-service-demo/pkg/grpcerrors"

	"github.com/sangianpatrick/grpc-service-demo/src/module/user"
	"github.com/sangianpatrick/grpc-service-demo/src/pb"
)

type userUsecase struct {
	logger *logrus.Logger
	umr    user.MariadbRepository
}

// NewUserUsecase is a constructor
func NewUserUsecase(logger *logrus.Logger, umr user.MariadbRepository) user.Usecase {
	return userUsecase{
		logger: logger,
		umr:    umr,
	}
}

func (uu userUsecase) Register(ctx context.Context, user *pb.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	_, err = uu.umr.FindByUsernameMobileNumberEmail(
		ctx, user.GetUsername(), user.GetMobileNumber(), user.GetEmail())

	if err == nil {
		return grpcerrors.ErrAlreadyExist
	}

	if err != sql.ErrNoRows {
		uu.logger.Error(err)
		return grpcerrors.ErrInternal
	}

	now, _ := util.ToMySQLDate(time.Now(), util.TimezoneAsiaJakarta)

	user.AccountStatus = pb.AccountStatus_Active
	user.CreatedAt = now
	user.UpdatedAt = now

	err = uu.umr.InsertOne(ctx, user)

	if err != nil {
		uu.logger.Error(err)
		return grpcerrors.ErrInternal
	}

	return
}

func (uu userUsecase) GetByUsername(ctx context.Context, username string) (user *pb.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	user, err = uu.umr.FindByUsername(ctx, username)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, grpcerrors.ErrDeadlineExceeded
		}
		if err == sql.ErrNoRows {
			return nil, grpcerrors.ErrNotFound
		}

		return nil, grpcerrors.ErrInternal
	}

	return
}
