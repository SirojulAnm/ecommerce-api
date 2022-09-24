package handler

import (
	"net/http"
	"superindo/auth"
	"superindo/helper"
	"superindo/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(ctx *gin.Context) {
	var input user.RegisterUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed saat input json", http.StatusUnprocessableEntity, "error", errorsMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed saat insert db", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed generate token", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var input user.LoginInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed saat input json", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed saat cek email atau password", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed saat generate token", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Success Log In", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}
