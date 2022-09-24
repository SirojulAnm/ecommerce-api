package handler

import (
	"net/http"
	"strconv"
	"superindo/category"
	"superindo/helper"
	"superindo/product"
	"superindo/user"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService  product.Service
	categoryService category.Service
}

func NewProductHandler(productService product.Service, categoryService category.Service) *productHandler {
	return &productHandler{productService, categoryService}
}

func (h *productHandler) AddProduct(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	if currentUser.Role != "admin" {
		response := helper.APIResponse("Add category must be role admin", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var input product.ProductInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add product failed saat input json", http.StatusUnprocessableEntity, "error", errorsMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	categoryID := input.CategoryId
	getCategoryByID, err := h.categoryService.GetCategoryByID(categoryID)
	if err != nil {
		response := helper.APIResponse("Add product failed saat cek category db, record id "+strconv.Itoa(categoryID)+" not found in table category", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	input.CategoryId = getCategoryByID.ID
	newProduct, err := h.productService.AddProduct(input)
	if err != nil {
		response := helper.APIResponse("Add product failed saat insert db", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Add product", http.StatusOK, "success", product.FormatProduct(newProduct))

	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProductsByCategoryID(ctx *gin.Context) {
	var input product.IdInput

	err := ctx.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed get product's by category saat BindUri", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	listProducts, err := h.productService.ListProductsByCategoryID(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed get product's by category saat get db", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get product's by category", http.StatusOK, "success", product.ListFormatProducts(listProducts))
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProductDetail(ctx *gin.Context) {
	var input product.IdInput

	err := ctx.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed get product detail saat BindUri", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.productService.GetProductDetailByID(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed get product detail saat get db", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get product detail", http.StatusOK, "success", product.FormatProduct(products))

	ctx.JSON(http.StatusOK, response)
}
