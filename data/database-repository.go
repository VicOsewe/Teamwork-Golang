package data

import (
	"Teamwork-Golang/registering"

	"Teamwork-Golang/getting"

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

func (repo UserRepository) UserSignIn(user getting.UserSignInfo) error {

	userDetails := User{}
	if err := repo.db.Where("email = ?", user.Email).First(&userDetails).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		// w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	return nil

}
