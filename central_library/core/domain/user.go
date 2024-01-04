package domain

type User struct {
	ID          uint `gorm:"primaryKey" json:",omitempty"`
	Name        string
	Surname     string
	Address     string
	Jmbg        string `gorm:"unique;not null"`
	IssuedBooks uint   `json:",omitempty"`
}
