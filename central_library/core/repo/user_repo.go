package repo

import "central_library/core/domain"

type UserRepo interface {
	Register(user *domain.User) (*domain.User, error)
	RecordBookIssue(userID uint) error
	RecordBookReturn(userID uint) error
}
