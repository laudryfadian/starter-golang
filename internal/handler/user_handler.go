package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laudryfadian/starter-golang/internal/domain"
	"github.com/laudryfadian/starter-golang/internal/entity"
	"github.com/laudryfadian/starter-golang/internal/helpers"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req entity.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.userUsecase.Register(c.Request.Context(), user)
	if err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "User registered successfully", http.StatusCreated, nil)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req entity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.userUsecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}

	helpers.SuccessResponse(c, "Login successful", http.StatusOK, res)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("userId")

	user, err := h.userUsecase.GetProfile(c.Request.Context(), userID)
	if err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusNotFound)
		return
	}

	helpers.SuccessResponse(c, "User profile retrieved successfully", http.StatusOK, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("userId")

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.userUsecase.UpdateProfile(c.Request.Context(), userID, &user)
	if err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "Profile updated successfully", http.StatusOK, nil)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.GetString("userId")

	err := h.userUsecase.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "User deleted successfully", http.StatusOK, nil)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers(c.Request.Context())
	if err != nil {
		helpers.ErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "Users retrieved successfully", http.StatusOK, users)
}
