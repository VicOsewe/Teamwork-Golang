package data

import (
	"Teamwork-Golang/registering"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return UserRepository{database}
}

func (repo UserRepository) CreateUser(user registering.Users) (userID uuid.UUID, erro error) {
	User := User{
		ID:         newUUID(),
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   user.Password,
		Gender:     user.Gender,
		Jobrole:    user.Jobrole,
		Department: user.Department,
		Address:    user.Address,
	}

	if err := repo.db.Create(&User).Error; err != nil {
		return uuid.Nil, err
	}
	return User.ID, nil

}

func newUUID() uuid.UUID {
	uuid, _ := uuid.NewUUID()
	return uuid
}
