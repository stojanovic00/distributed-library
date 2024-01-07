package persistence

import (
	"errors"
	"gorm.io/gorm"
	"local_library/core/domain"
)

var (
	ErrAlreadyTaken = errors.New("this book is already taken")
)

type IssuingRepoPg struct {
	dbClient *gorm.DB
}

func NewIssuingRepoPg(dbClient *gorm.DB) *IssuingRepoPg {
	return &IssuingRepoPg{dbClient: dbClient}
}
func (rep *IssuingRepoPg) GetAll() ([]domain.IssuingRecord, error) {
	var records []domain.IssuingRecord

	err := rep.dbClient.Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}
func (rep *IssuingRepoPg) CheckIfTaken(isbn string) bool {
	foundRecord := domain.IssuingRecord{}
	err := rep.dbClient.Where("isbn = ? and returned = false", isbn).First(&foundRecord).Error
	if err == nil {
		return true
	}
	return false
}

func (rep *IssuingRepoPg) RecordBookIssue(record *domain.IssuingRecord) error {
	err := rep.dbClient.Create(record).Error
	if err != nil {
		return err
	}

	return nil
}

func (rep *IssuingRepoPg) RecordBookReturn(isbn string) (uint, error) {
	record := domain.IssuingRecord{}
	err := rep.dbClient.Where("isbn = ? and returned = false", isbn).First(&record).Error
	if err != nil {
		return 0, err
	}

	record.Returned = true

	rep.dbClient.Save(&record)
	return record.UserId, nil
}
