package persistence

import (
	"central_library/core/domain"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

var (
	ErrJmbgExists = errors.New("user with this jmbg already exists")
	ErrIssuedMax  = errors.New("user has already 3 issued books")
	ErrNoIssued   = errors.New("user hasn't issued books")
	ErrNotFound   = errors.New("user not found")
)

type UserRepoPg struct {
	dbClient *gorm.DB
}

func NewUserRepoPg(dbClient *gorm.DB) *UserRepoPg {
	return &UserRepoPg{dbClient: dbClient}
}

func (rep *UserRepoPg) GetAll() ([]domain.User, error) {
	var users []domain.User

	err := rep.dbClient.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (rep *UserRepoPg) Register(user *domain.User) (*domain.User, error) {
	err := rep.dbClient.Create(user).Error
	if err != nil {
		if isDuplicateJmbgError(err) {
			return nil, ErrJmbgExists
		}
		return nil, err
	}

	//ID is populated after successful create
	return user, nil
}

func (rep *UserRepoPg) RecordBookIssue(userID uint) error {
	var user domain.User

	err := rep.dbClient.First(&user, userID).Error
	if err != nil {
		return ErrNotFound
	}

	if user.IssuedBooks == 3 {
		return ErrIssuedMax
	}

	user.IssuedBooks += 1
	rep.dbClient.Save(&user)

	return nil
}

func (rep *UserRepoPg) RecordBookReturn(userID uint) error {
	var user domain.User

	err := rep.dbClient.First(&user, userID).Error
	if err != nil {
		return ErrNotFound
	}

	if user.IssuedBooks == 0 {
		return ErrNoIssued
	}

	user.IssuedBooks -= 1
	rep.dbClient.Save(&user)

	return nil
}

func isDuplicateJmbgError(err error) bool {
	if pgError, ok := err.(*pgconn.PgError); ok {
		return pgError.Code == "23505" && pgError.ConstraintName == "users_jmbg_key"
	}
	return false
}
