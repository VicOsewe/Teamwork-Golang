package updating

import "strings"

type UpdateService interface {
	UpdateArticle(art UpdateAtricle) (articleTitle string, message string, err error)
}

type UpdateRepository interface {
	UpdateArticle(art UpdateAtricle) (articleTitle string, message string, err error)
}

type UpdateError struct {
	errorList []string
}

func (e *UpdateError) add(message string) {
	e.errorList = append(e.errorList, message)
}

func (e *UpdateError) Error() string {
	return strings.Join(e.errorList, ", ")
}

type service struct {
	repo UpdateRepository
}

func NewUpdatingService(r UpdateRepository) UpdateService {
	return &service{
		repo: r,
	}
}

func (s *service) UpdateArticle(art UpdateAtricle) (articleTitle string, message string, err error) {
	var updateErr UpdateError
	if len(art.Message) == 0 {
		updateErr.add("Message not provided ")
		return art.Title, art.Message, &updateErr
	}

	Title, message, err := s.repo.UpdateArticle(art)
	if err != nil {
		updateErr.add("article not updated in the database")
		return Title, message, &updateErr
	}
	return Title, message, nil
}
