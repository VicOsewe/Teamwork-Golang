package creating

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CreatingService interface {
	CreateUser(user Users) (userID uuid.UUID, erro error)
	UserSignIn(user UserSignInfo) error
	CreateArticle(art Article) (articleID uuid.UUID, createdAt time.Time, articleTitle string, erro error)
}

type CreatingRepository interface {
	CreateUser(user Users) (userId uuid.UUID, erro error)
	UserSignIn(user UserSignInfo) error
	CreateArticle(art Article) (articleID uuid.UUID, createdAt time.Time, articleTitle string, erro error)
}

type CreatingError struct {
	errorList []string
}

func (e *CreatingError) add(message string) {
	e.errorList = append(e.errorList, message)
}

func (e *CreatingError) Error() string {
	return strings.Join(e.errorList, ", ")
}

type service struct {
	repo CreatingRepository
}

func NewcreatingService(r CreatingRepository) CreatingService {
	return &service{
		repo: r,
	}
}

func (s *service) CreateUser(user Users) (userID uuid.UUID, erro error) {
	var regError CreatingError
	err := s.validateUserInfo(user)
	if err != nil {
		regError.add("Invalid user info " + err.Error())
		return uuid.Nil, &regError
	}
	UserID, errUser := s.repo.CreateUser(user)
	if errUser != nil {
		regError.add("user not registered in database")
		return UserID, &regError
	}

	return UserID, nil

}

func (s *service) UserSignIn(user UserSignInfo) error {
	var getError CreatingError
	err := s.validateSignInInfo(user)
	if err != nil {
		getError.add("user log info not provided " + err.Error())
		return err
	}

	erro := s.repo.UserSignIn(user)
	if erro != nil {
		getError.add("user cannot login")
		return &getError
	}
	return nil

}

func (s *service) CreateArticle(art Article) (ArticleID uuid.UUID, CreatedAt time.Time, articleTitle string, erro error) {

	var regError CreatingError
	create := time.Now()
	err := s.validateArticleInfo(art)
	if err != nil {
		regError.add("Info not provided for:" + err.Error())
		return uuid.Nil, create, art.Title, &regError
	}
	articleID, createdAt, articleTitle, errro := s.repo.CreateArticle(art)
	if errro != nil {
		regError.add("Article not created")
		return articleID, create, art.Title, &regError
	}
	return articleID, createdAt, articleTitle, nil

}

func (s *service) validateUserInfo(user Users) error {
	if len(user.Firstname+user.Lastname) == 0 {
		return errors.New("Missing user name")
	}
	if len(user.Password) == 0 {
		return errors.New("Please provide a password")
	}
	return nil
}

func (s *service) validateArticleInfo(art Article) error {
	if len(art.Title) == 0 {
		return errors.New("Article must have a title")
	}
	if len(art.Article) == 0 {
		return errors.New("Please provide an article body")
	}
	return nil
}

func (s *service) validateSignInInfo(user UserSignInfo) error {
	if len(user.Email) == 0 {
		return errors.New("Email not provided")
	}

	if len(user.Password) == 0 {
		return errors.New("Password not provided")
	}
	return nil
}
