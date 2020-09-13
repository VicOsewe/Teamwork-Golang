package registering

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type RegisterService interface {
	CreateUser(user Users) (userId uuid.UUID, erro error)
}

type RegisterRepository interface {
	CreateUser(user Users) (userId uuid.UUID, erro error)
}

type RegisteringError struct {
	errorList []string
}

func (e *RegisteringError) add(message string) {
	e.errorList = append(e.errorList, message)
}

func (e *RegisteringError) Error() string {
	return strings.Join(e.errorList, ", ")
}

type service struct {
	repo RegisterRepository
}

func NewRegisteringService(r RegisterRepository) RegisterService {
	return &service{
		repo: r,
	}
}

func (s *service) CreateUser(user Users) (userID uuid.UUID, erro error) {
	var regError RegisteringError
	err := s.validateUserInfo(user)
	if err != nil {
		regError.add("Invalid user info" + err.Error())
		return uuid.Nil, &regError
	}
	UserID, errUser := s.repo.CreateUser(user)
	fmt.Println(errUser)
	if errUser != nil {
		regError.add("user not registered in database")
		return UserID, &regError
	}

	return UserID, nil

}

func (s *service) validateUserInfo(user Users) error {
	if len(user.Firstname+user.Lastname) == 0 {
		return errors.New("Missing user name")
	}
	return nil
}
