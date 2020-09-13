package registering

import (
	"errors"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockUserRepo struct {
	mock.Mock
}

type RegisteringTestSuite struct {
	suite.Suite
	service RegisterService
	repo    *mockUserRepo
}

func (m *mockUserRepo) CreateUser(user Users) (uuid.UUID, error) {
	args := m.Called(user)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (s *RegisteringTestSuite) SetupTest() {
	s.repo = new(mockUserRepo)
	s.service = NewRegisteringService(s.repo)
}

func (s *RegisteringTestSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}

func TestRegisteringService(t *testing.T) {
	suite.Run(t, new(RegisteringTestSuite))
}

func (s *RegisteringTestSuite) TestUserCreationSuccess() {
	user := fakeUserInfo()

	uuid := uuid.New()
	s.repo.On("CreateUser", mock.Anything).Return(uuid, nil)
	_, err := s.service.CreateUser(user)
	// assert the CDS returned a nil error
	s.Nil(err)
	s.repo.AssertCalled(s.T(), "CreateUser", user)
}

func (s *RegisteringTestSuite) TestUserCreationFailure() {

	userInfo := Users{}
	userInfo.Firstname = ""
	userInfo.Lastname = ""

	_, err := s.service.CreateUser(userInfo)
	s.NotNil(err)
	s.repo.AssertNotCalled(s.T(), "CreateUser", userInfo)

}

func (s *RegisteringTestSuite) TestCreateUserDatabaseFailure() {
	user := fakeUserInfo()

	uuid := uuid.New()
	s.repo.On("CreateUser", mock.Anything).Return(uuid, errors.New("database error"))
	_, err := s.service.CreateUser(user)
	// assert the CDS returned a nil error
	s.NotNil(err)
	s.repo.AssertCalled(s.T(), "CreateUser", user)
}

func fakeUserInfo() Users {
	return Users{
		Firstname: "Jane",
		Lastname:  "Doe",
	}
}
