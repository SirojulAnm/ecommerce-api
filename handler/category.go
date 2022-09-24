package handler

import (
	"net/http"
	"superindo/category"
	"superindo/helper"
	"superindo/user"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(categoryService category.Service) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) AddCategory(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	if currentUser.Role != "admin" {
		response := helper.APIResponse("Add category must be role admin", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var input category.CategoryInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add category product failed saat input json", http.StatusUnprocessableEntity, "error", errorsMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCategory, err := h.categoryService.AddCategory(input)
	if err != nil {
		response := helper.APIResponse("Add category product failed saat insert db", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Add category product", http.StatusOK, "success", category.FormatCategory(newCategory))

	ctx.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategories(ctx *gin.Context) {
	categories, err := h.categoryService.GetAll()
	if err != nil {
		response := helper.APIResponse("Get all category product failed saat get db", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get all category product", http.StatusOK, "success", category.FormatCategories(categories))

	ctx.JSON(http.StatusOK, response)
}
