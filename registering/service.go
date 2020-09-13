package registering

import (
	"errors"
	"strings"
)

type RegisterService interface {
	CreateUser(user Users) error
}

type RegisterRepository interface {
	CreateUser(user Users) error
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

func (s *service) CreateUser(user Users) error {
	regError := RegisteringError{}
	err := s.validateUserInfo(user)
	if err != nil {
		regError.add("Invalid user info" + err.Error())
		return &regError
	}
	errUser := s.repo.CreateUser(user)
	if errUser != nil {
		regError.add("Failed to add user to the database" + err.Error())
	}
	return nil

}

func (s *service) validateUserInfo(user Users) error {
	if len(user.Firstname+user.Lastname) == 0 {
		return errors.New("Missing user name")
	}
	return nil
}
