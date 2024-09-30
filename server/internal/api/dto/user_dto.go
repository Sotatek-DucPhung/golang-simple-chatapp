package dto

import "gin-gorm-crud/internal/models"

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterInput struct {
	UserName        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserRes struct {
	AccessToken string
	ID          string `json:"id"`
	Username    string `json:"username"`
}

func ToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}
}

func ToUserResponses(users []models.User) []UserResponse {
	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = *ToUserResponse(&user)
	}
	return userResponses
}
