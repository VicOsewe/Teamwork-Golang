package registering

import (
	"testing"

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

func (m *mockUserRepo) CreateUser(user Users) error {
	args := m.Called(user)
	return args.Error(0)
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

	s.repo.On("CreateUser", mock.Anything).Return(nil)
	err := s.service.CreateUser(user)
	// assert the CDS returned a nil error
	s.Nil(err)
	s.repo.AssertCalled(s.T(), "CreateUser", user)
}

func (s *RegisteringTestSuite) TestUserCreationFailure() {

	userInfo := Users{}
	userInfo.Firstname = ""
	userInfo.Lastname = ""

	err := s.service.CreateUser(userInfo)
	s.NotNil(err)
	s.repo.AssertNotCalled(s.T(), "CreateUser", userInfo)

}

func fakeUserInfo() Users {
	return Users{
		Firstname: "Jane",
		Lastname:  "Doe",
	}
}
