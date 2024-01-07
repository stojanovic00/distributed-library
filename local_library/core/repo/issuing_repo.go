package repo

import "local_library/core/domain"

type IssuingRepo interface {
	GetAll() ([]domain.IssuingRecord, error)
	RecordBookIssue(record *domain.IssuingRecord) error
	CheckIfTaken(isbn string) bool
	RecordBookReturn(isbn string) (uint, error)
}
