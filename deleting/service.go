package deleting

import "strings"

type DeleteService interface {
	DeleteArticle(ID DeleteArt) error
}

type DeleteRepository interface {
	DeleteArticle(ID DeleteArt) error
}

type DeleteError struct {
	errorList []string
}

func (e *DeleteError) add(message string) {
	e.errorList = append(e.errorList, message)
}

func (e *DeleteError) Error() string {
	return strings.Join(e.errorList, ", ")
}

type service struct {
	repo DeleteRepository
}

func NewDeletingService(r DeleteRepository) DeleteService {
	return &service{
		repo: r,
	}
}

func (s *service) DeleteArticle(ID DeleteArt) error {
	var deleteErr DeleteError
	if len(ID.ArticleID) == 0 {
		deleteErr.add("Article ID not provided ")
		return &deleteErr
	}

	err := s.repo.DeleteArticle(ID)
	if err != nil {
		deleteErr.add("user not registered in database")
		return &deleteErr
	}
	return nil
}
