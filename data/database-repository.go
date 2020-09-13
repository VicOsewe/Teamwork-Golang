package data

import (
	"Teamwork-Golang/registering"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return UserRepository{database}
}

func (repo UserRepository) CreateUser(user registering.Users) (userID uuid.UUID, erro error) {
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return uuid.Nil, err
	}

	User := User{
		ID:         newUUID(),
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   string(hashedPassword),
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
