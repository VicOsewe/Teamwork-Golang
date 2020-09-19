package data

import (
	"Teamwork-Golang/creating"
	"Teamwork-Golang/deleting"
	"Teamwork-Golang/updating"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return UserRepository{database}
}

func (repo UserRepository) CreateUser(user creating.Users) (userID uuid.UUID, erro error) {
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return uuid.Nil, err
	}

	User := User{
		ID:         newUUID(),
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   string(hashedPassword),
		Gender:     user.Gender,
		Jobrole:    user.Jobrole,
		Department: user.Department,
		Address:    user.Address,
	}

	if err := repo.db.Create(&User).Error; err != nil {
		return uuid.Nil, err
	}
	return User.ID, nil

}

func newUUID() uuid.UUID {
	uuid, _ := uuid.NewUUID()
	return uuid
}

func (repo UserRepository) UserSignIn(user creating.UserSignInfo) error {

	userDetails := User{}
	if err := repo.db.Where("email = ?", user.Email).First(&userDetails).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		return err
	}
	return nil

}

func (repo UserRepository) CreateArticle(art creating.Article) (ArticleID uuid.UUID, CreatedAt time.Time, articleTitle string, erro error) {
	article := Article{
		ID:          newUUID(),
		Title:       art.Title,
		Article:     art.Article,
		DateCreated: time.Now(),
	}
	timecreated := time.Now()
	if err := repo.db.Create(&article).Error; err != nil {
		return uuid.Nil, timecreated, article.Title, err
	}
	return article.ID, article.DateCreated, article.Title, nil

}

func (repo UserRepository) UpdateArticle(art updating.UpdateAtricle) (articleTitle string, message string, err error) {
	article := Article{}

	timecreated := time.Now()
	if err := repo.db.Debug().Model(&article).Where(Article{ID: art.ArticleID}).Find(&article).Update(Article{Title: art.Title, Article: art.Message, DateLastUpdated: timecreated}).Error; err != nil {
		return art.Title, art.Message, err
	}
	return art.Title, art.Message, nil

}

func (repo UserRepository) DeleteArticle(ID deleting.DeleteArt) error {
	article := Article{}

	if err := repo.db.Debug().Model(&article).Where(Article{ID: ID.ArticleID}).Find(&article).Delete(&article).Error; err != nil {
		return err

	}
	return nil
}
