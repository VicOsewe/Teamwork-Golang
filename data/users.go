package data

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"column:id"`
	Firstname  string    `gorm:"column:firstname"`
	Lastname   string    `gorm:"column:lastname"`
	Email      string    `gorm:"column:email"`
	Password   string    `gorm:"column:password"`
	Gender     string    `gorm:"column:gender"`
	Jobrole    string    `gorm:"column:jobrole"`
	Department string    `gorm:"column:department"`
	Address    string    `gorm:"column:address"`
}

func (User) TableName() string {
	return "public.users"
}

type Article struct {
	ID              uuid.UUID `gorm:"column:id"`
	Title           string    `gorm:"column:title"`
	Article         string    `gorm:"column:article"`
	DateCreated     time.Time `gorm:"column:date_created"`
	DateLastUpdated time.Time `gorm:"column:date_last_updated"`
}

func (Article) TableName() string {
	return "public.article"
}
