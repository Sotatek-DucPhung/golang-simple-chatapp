package handlers

import (
	"gin-gorm-crud/internal/api/dto"
	"gin-gorm-crud/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// @Summary Login user
// @Description Login user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.UserLoginInput true "User login data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {

	var input dto.UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, err := h.userService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.SetCookie("jwt", userInfo.AccessToken, 60*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, userInfo)
}

// @Summary Create a new user
// @Description Create a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.UserRegisterInput true "User register data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/signup [post]
func (h *UserHandler) CreateUser(c *gin.Context) {

	var input dto.UserRegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToUserResponse(user))
}

// @Summary Get user by email
// @Security ApiKeyAuth
// @Description Get user by email

// @Tags users
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/email/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserResponse(user))
}
