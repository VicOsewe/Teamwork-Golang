package updating

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockUpdateRepo struct {
	mock.Mock
}

type UpdatingTestSuite struct {
	suite.Suite
	service UpdateService
	repo    *mockUpdateRepo
}

func (m *mockUpdateRepo) UpdateArticle(art UpdateAtricle) (articleTitle string, message string, err error) {
	args := m.Called(art)
	return args.String(0), args.String(1), args.Error(2)
}

func (s *UpdatingTestSuite) SetupTest() {
	s.repo = new(mockUpdateRepo)
	s.service = NewUpdatingService(s.repo)
}

func (s *UpdatingTestSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}

func TestUpdatingService(t *testing.T) {
	suite.Run(t, new(UpdatingTestSuite))
}

func (s *UpdatingTestSuite) TestUpdateArticleSuccess() {
	updateInfo := UpdateAtricle{}
	updateInfo.ArticleID = uuid.New()
	updateInfo.Message = "fake update message"
	updateInfo.Title = "fake title"
	s.repo.On("UpdateArticle", mock.Anything).Return(updateInfo.Title, updateInfo.Message, nil)
	_, _, err := s.service.UpdateArticle(updateInfo)

	s.Nil(err)
	s.repo.AssertCalled(s.T(), "UpdateArticle", updateInfo)

}

func (s *UpdatingTestSuite) TestUpdateArticleValidationFailure() {
	updateInfo := UpdateAtricle{}
	updateInfo.ArticleID = uuid.Nil
	updateInfo.Message = ""
	updateInfo.Title = "fake title"
	_, _, err := s.service.UpdateArticle(updateInfo)
	s.NotNil(err)
	s.repo.AssertNotCalled(s.T(), "UpdateArticle", updateInfo)

}

func (s *UpdatingTestSuite) TestUpdateArticledatabaseFailure() {
	updateInfo := UpdateAtricle{}
	updateInfo.ArticleID = uuid.New()
	updateInfo.Message = "fake update message"
	updateInfo.Title = "fake title"
	s.repo.On("UpdateArticle", mock.Anything).Return(updateInfo.Title, updateInfo.Message, errors.New("database error"))
	_, _, err := s.service.UpdateArticle(updateInfo)

	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "UpdateArticle", updateInfo)
}
