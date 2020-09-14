package getting

import (
	"errors"
	"strings"
)

type GettingService interface {
	UserSignIn(user UserSignInfo) error
}

type GettingRepository interface {
	UserSignIn(user UserSignInfo) error
}

type GettingError struct {
	errorList []string
}

func (e *GettingError) add(message string) {
	e.errorList = append(e.errorList, message)
}

func (e *GettingError) Error() string {
	return strings.Join(e.errorList, ", ")
}

type service struct {
	repo GettingRepository
}

func NewGettingService(r GettingRepository) GettingService {
	return &service{
		repo: r,
	}
}

func (s *service) UserSignIn(user UserSignInfo) error {
	var getError GettingError
	err := s.validateUserInfo(user)
	if err != nil {
		getError.add("user log info not provided " + err.Error())
	}

	erro := s.repo.UserSignIn(user)
	if erro != nil {
		getError.add("user cannot login")
		return &getError
	}
	return nil

}

func (s *service) validateUserInfo(user UserSignInfo) error {
	if len(user.Email) == 0 {
		return errors.New("Email not provided")
	}

	if len(user.Password) == 0 {
		return errors.New("Password not provided")
	}
	return nil
}
