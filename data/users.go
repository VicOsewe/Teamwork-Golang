package data

import "github.com/google/uuid"

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

