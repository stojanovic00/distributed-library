package domain

import "time"

type IssuingRecord struct {
	ID          uint `gorm:"primaryKey" json:",omitempty"`
	Title       string
	Author      string
	ISBN        string
	IssuingDate time.Time
	UserId      uint
	Returned    bool
}
