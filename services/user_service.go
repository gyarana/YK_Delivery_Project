package services

import (
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"nix_education/model"
	"nix_education/model/repositories"
)

func NewUserService(UserRepository *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: *UserRepository,
	}
}

type UserServiceI interface {
	CreateNewUser(user *model.User) error
	GetUserByID(userID int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
}

type UserService struct {
	UserRepository repositories.UserRepository
}

func (u UserService) CreateNewUser(user *model.User) error {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	user.PasswordHash = string(hashPassword)
	err := u.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) GetUserByID(userID int) (*model.User, error) {
	user, err := u.UserRepository.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetUserByEmail(email string) (*model.User, error) {
	user, err := u.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) UpdateUser(user *model.User) error {
	err := u.UserRepository.EditUser(user)
	if err != nil {
		return err
	}
	return nil
}
