package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sangianpatrick/grpc-service-demo/pkg/util"

	"github.com/sangianpatrick/grpc-service-demo/src/module/user"

	"github.com/sangianpatrick/grpc-service-demo/src/pb"
)

type userMariadbRepository struct {
	db *sql.DB
}

// NewUserMariadbRepository is a constructor
func NewUserMariadbRepository(db *sql.DB) user.MariadbRepository {
	return userMariadbRepository{
		db: db,
	}
}

func (umr userMariadbRepository) InsertOne(ctx context.Context, user *pb.User) (err error) {
	queryUser := `INSERT INTO user SET username=?, email=?, mobile_number=?, password=?, name=?, account_status=?, created_at=?, updated_at=?`
	queryAddress := `INSERT INTO address SET province=?, district=?, sub_district=?, village=?, street=?, zip_code=?, user_id=?`

	conn, err := umr.db.Conn(ctx)
	if err != nil {
		return
	}

	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	userStmt, err := tx.PrepareContext(ctx, queryUser)
	if err != nil {
		return
	}
	userResult, err := userStmt.ExecContext(ctx, user.Username, user.Email, user.MobileNumber, user.Password, user.Name, user.AccountStatus, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return
	}

	userID, _ := userResult.LastInsertId()
	addressStmt, err := tx.PrepareContext(ctx, queryAddress)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = addressStmt.ExecContext(ctx, user.Address.Province, user.Address.District, user.Address.SubDistrict, user.Address.Village, user.Address.Street, user.Address.ZipCode, userID)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}

	return nil
}

func (umr userMariadbRepository) FindByUsernameMobileNumberEmail(ctx context.Context, username, mobileNumber, email string) (user *pb.User, err error) {
	query := `SELECT u.id, u.username, u.email, u.mobile_number, u.name, u.account_status, a.province, a.district, a.sub_district, a.village, a.street, a.zip_code, created_at, updated_at from user u 
	LEFT JOIN address a ON u.id = a.user_id WHERE u.username = ? OR u.mobile_number=? OR u.email=?`

	user, err = umr.findOne(ctx, query, username, mobileNumber, email)
	return
}

func (umr userMariadbRepository) FindByUsername(ctx context.Context, username string) (user *pb.User, err error) {
	query := `SELECT u.id, u.username, u.email, u.mobile_number, u.name, u.account_status, a.province, a.district, a.sub_district, a.village, a.street, a.zip_code, created_at, updated_at from user u 
	LEFT JOIN address a ON u.id = a.user_id WHERE u.username = ?`

	user, err = umr.findOne(ctx, query, username)
	return
}

func (umr userMariadbRepository) findOne(ctx context.Context, query string, value ...interface{}) (user *pb.User, err error) {
	conn, err := umr.db.Conn(ctx)
	if err != nil {
		return
	}

	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	row := stmt.QueryRowContext(ctx, value...)
	u := pb.User{}
	a := pb.Address{}

	var createdAt string
	var updatedAt string

	err = row.Scan(
		&u.Id,
		&u.Username,
		&u.Email,
		&u.MobileNumber,
		&u.Name,
		&u.AccountStatus,
		&a.Province,
		&a.District,
		&a.SubDistrict,
		&a.Village,
		&a.Street,
		&a.ZipCode,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return
	}

	fmt.Println(createdAt)

	l, _ := time.LoadLocation(util.TimezoneAsiaJakarta)

	cat, _ := time.ParseInLocation(util.ISOFormat, createdAt, l)
	uat, _ := time.ParseInLocation(util.ISOFormat, updatedAt, l)

	u.CreatedAt, _ = util.ToISODate(cat, "UTC")
	u.UpdatedAt, _ = util.ToISODate(uat, "UTC")
	u.Address = &a

	return &u, nil
}
