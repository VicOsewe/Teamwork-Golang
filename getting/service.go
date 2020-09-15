package getting

import "strings"

type GettingService interface {
}

type GettingRepository interface {
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
