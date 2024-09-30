package service

import (
	"fmt"
	"gin-gorm-crud/internal/api/dto"
	"gin-gorm-crud/internal/models"
	"gin-gorm-crud/internal/repository"
	"gin-gorm-crud/internal/utils"
	"strconv"
)

type UserService interface {
	Login(email, password string) (dto.LoginUserRes, error)
	CreateUser(userRegisterInput *dto.UserRegisterInput) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) Login(email string, password string) (dto.LoginUserRes, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return dto.LoginUserRes{}, err
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return dto.LoginUserRes{}, err
	}

	return dto.LoginUserRes{
		AccessToken: token,
		ID:          strconv.Itoa(int(user.ID)),
		Username:    user.UserName,
	}, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(userRegisterInput *dto.UserRegisterInput) (*models.User, error) {

	// check if password and confirm password is the same
	if userRegisterInput.Password != userRegisterInput.ConfirmPassword {
		return nil, fmt.Errorf("password and confirm password do not match")
	}

	// check if email is already exists
	existingUser, err := s.userRepo.GetUserByEmail(userRegisterInput.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("email already exists")
	}

	// create user
	user := &models.User{
		UserName: userRegisterInput.UserName,
		Email:    userRegisterInput.Email,
		Password: userRegisterInput.Password,
	}

	user.HashPassword(userRegisterInput.Password)
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}
