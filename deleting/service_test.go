package deleting

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockDeleteRepo struct {
	mock.Mock
}

type DeleteTestSuite struct {
	suite.Suite
	service DeleteService
	repo    *mockDeleteRepo
}

func (m *mockDeleteRepo) DeleteArticle(ID DeleteArt) error {
	args := m.Called(ID)
	return args.Error(0)
}

func (s *DeleteTestSuite) SetupTest() {
	s.repo = new(mockDeleteRepo)
	s.service = NewDeletingService(s.repo)
}

func (s *DeleteTestSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}

func TestDeleteService(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}

func (s *DeleteTestSuite) TestDeleteArticleSuccess() {

	id := DeleteArt{}
	id.ArticleID = uuid.New()
	s.repo.On("DeleteArticle", mock.Anything).Return(nil)
	err := s.service.DeleteArticle(id)

	s.Nil(err)
	s.repo.AssertCalled(s.T(), "DeleteArticle", id)

}

func (s *DeleteTestSuite) TestDeleteArticledatabaseFailure() {

	id := DeleteArt{}
	id.ArticleID = uuid.New()
	s.repo.On("DeleteArticle", mock.Anything).Return(errors.New("database error"))
	err := s.service.DeleteArticle(id)

	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "DeleteArticle", id)
}
