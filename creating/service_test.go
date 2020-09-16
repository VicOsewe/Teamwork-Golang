package creating

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockUserRepo struct {
	mock.Mock
}

type CreatingTestSuite struct {
	suite.Suite
	service CreatingService
	repo    *mockUserRepo
}

func (m *mockUserRepo) CreateUser(user Users) (uuid.UUID, error) {
	args := m.Called(user)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *mockUserRepo) UserSignIn(user UserSignInfo) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepo) CreateArticle(art Article) (articleID uuid.UUID, createdAt time.Time, ArticleTitle string, erro error) {
	args := m.Called(art)
	return args.Get(0).(uuid.UUID), args.Get(1).(time.Time), args.String(2), args.Error(3)
}

func (s *CreatingTestSuite) SetupTest() {
	s.repo = new(mockUserRepo)
	s.service = NewcreatingService(s.repo)
}

func (s *CreatingTestSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}

func TestCreatingService(t *testing.T) {
	suite.Run(t, new(CreatingTestSuite))
}

func (s *CreatingTestSuite) TestUserCreationSuccess() {
	user := fakeUserInfo()

	uuid := uuid.New()
	s.repo.On("CreateUser", mock.Anything).Return(uuid, nil)
	_, err := s.service.CreateUser(user)
	// assert the CDS returned a nil error
	s.Nil(err)
	s.repo.AssertCalled(s.T(), "CreateUser", user)
}

func (s *CreatingTestSuite) TestUserCreationValidateFailure() {

	userInfo := Users{}
	userInfo.Firstname = ""
	userInfo.Lastname = ""

	_, err := s.service.CreateUser(userInfo)
	s.NotNil(err)
	s.repo.AssertNotCalled(s.T(), "CreateUser", userInfo)

}

func (s *CreatingTestSuite) TestCreateUserDatabaseFailure() {
	user := fakeUserInfo()

	uuid := uuid.New()
	s.repo.On("CreateUser", mock.Anything).Return(uuid, errors.New("database error"))
	_, err := s.service.CreateUser(user)
	// assert the CDS returned a nil error
	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "CreateUser", user)
}

func (s *CreatingTestSuite) TestUserSignInSuccess() {
	signInfo := fakeSignInfo()
	s.repo.On("UserSignIn", mock.Anything).Return(nil)
	err := s.service.UserSignIn(signInfo)
	s.Nil(err)
	s.repo.AssertCalled(s.T(), "UserSignIn", signInfo)
}

func (s *CreatingTestSuite) TestUserSignInValidateFailure() {
	signInfo := UserSignInfo{}
	signInfo.Email = ""
	signInfo.Password = ""
	err := s.service.UserSignIn(signInfo)
	s.NotNil(err)
	s.repo.AssertNotCalled(s.T(), "UserSignIn", signInfo)
}

func (s *CreatingTestSuite) TestUserSignInDatabaseFailure() {
	signInfo := fakeSignInfo()
	s.repo.On("UserSignIn", mock.Anything).Return(errors.New("database error"))
	err := s.service.UserSignIn(signInfo)
	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "UserSignIn", signInfo)

}

func (s *CreatingTestSuite) TestCreateArticleSuccess() {
	articleInfo := fakeArticleInfo()
	uuid := uuid.New()
	time := time.Now()
	articleID := "Lipsum Lorem"
	s.repo.On("CreateArticle", mock.Anything).Return(uuid, time, articleID, nil)
	_, _, _, err := s.service.CreateArticle(articleInfo)
	s.Nil(err)

}

func (s *CreatingTestSuite) TestCreateArticleValidationFailure() {
	article := Article{}
	article.Article = ""
	article.Title = ""
	_, _, _, err := s.service.CreateArticle(article)
	s.NotNil(err)
	s.repo.AssertNotCalled(s.T(), "CreateArticle", article)
}

func (s *CreatingTestSuite) TestCreateArticleDatabaseFailure() {
	article := fakeArticleInfo()
	uuid := uuid.New()
	time := time.Now()
	articleID := ""

	s.repo.On("CreateArticle", mock.Anything).Return(uuid, time, articleID, errors.New("database error"))
	_, _, _, err := s.service.CreateArticle(article)
	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "CreateArticle", article)
}

func fakeUserInfo() Users {
	return Users{
		Firstname: "Jane",
		Lastname:  "Doe",
		Password:  "MyPassword",
	}
}

func fakeSignInfo() UserSignInfo {
	return UserSignInfo{
		Email:    "example@gmail.com",
		Password: "MyPassword",
	}
}

func fakeArticleInfo() Article {

	return Article{
		Title:   "Lipsum Lorem",
		Article: "fake article body",
	}
}
