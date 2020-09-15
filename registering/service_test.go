package registering

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

type RegisteringTestSuite struct {
	suite.Suite
	service RegisterService
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

func (m *mockUserRepo) CreateArticle(art Article) (articleID uuid.UUID, createdAt time.Time, erro error) {
	args := m.Called(art)
	return args.Get(0).(uuid.UUID), args.Get(1).(time.Time), args.Error(2)
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

func (s *RegisteringTestSuite) TestUserCreationValidateFailure() {

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

func (s *RegisteringTestSuite) TestUserSignInSuccess() {
	signInfo := fakeSignInfo()
	s.repo.On("UserSignIn", mock.Anything).Return(nil)
	err := s.service.UserSignIn(signInfo)
	s.Nil(err)
	s.repo.AssertCalled(s.T(), "UserSignIn", signInfo)
}

func (s *RegisteringTestSuite) TestUserSignInValidateFailure() {
	signInfo := UserSignInfo{}
	signInfo.Email = ""
	signInfo.Password = ""
	err := s.service.UserSignIn(signInfo)
	s.NotNil(err)
	s.repo.AssertNotCalled(s.T(), "UserSignIn", signInfo)
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
