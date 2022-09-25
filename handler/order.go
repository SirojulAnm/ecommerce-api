package handler

import (
	"fmt"
	"net/http"
	"superindo/helper"
	"superindo/order"
	"superindo/product"
	"superindo/user"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService   order.Service
	productService product.Service
}

func NewOrderHandler(orderService order.Service, productService product.Service) *orderHandler {
	return &orderHandler{orderService, productService}
}

func (h *orderHandler) AddOrder(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	if currentUser.Role != "user" {
		response := helper.APIResponse("Failed add shopping cart, must be role user", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var input order.ShoppingCartInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errorsMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed add shopping cart saat input json", http.StatusUnprocessableEntity, "error", errorsMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var totalPrice float32
	var totalQty float32
	for _, item := range input.Cart {
		cekIdProduct, err := h.productService.CekProductByID(item.ProductId)
		if cekIdProduct != true {
			errorsMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Failed add shopping cart, id product not found", http.StatusBadRequest, "error", errorsMessage)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		getProduct, err := h.productService.GetProductDetailByID(item.ProductId)
		if err != nil {
			errorsMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Failed add shopping cart, id product not found", http.StatusBadRequest, "error", errorsMessage)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		totalPrice = totalPrice + (getProduct.Price * float32(item.ProductQty))
		totalQty = totalQty + float32(item.ProductQty)
	}

	var orderInput order.OrderInput
	orderInput.UserId = currentUser.ID
	orderInput.OrderTotalPrice = totalPrice
	orderInput.OrderStatus = "menunggu pembayaran"
	orderInput.IsDelivered = false
	paymentURL, err := h.orderService.GetPaymentURL(int64(totalPrice), currentUser.FirstName, currentUser.Email)
	orderInput.PaymentURL = paymentURL

	newOrder, err := h.orderService.AddOrder(orderInput)
	if err != nil {
		response := helper.APIResponse("Failed add shopping cart, saat insert db order", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	for _, item := range input.Cart {
		getProduct, err := h.productService.GetProductDetailByID(item.ProductId)
		fmt.Println(getProduct)
		if err != nil {
			errorsMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Failed add shopping cart, id product not found", http.StatusBadRequest, "error", errorsMessage)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		addItem := order.OrderDetailsInput{
			OrderId:           newOrder.ID,
			ProductId:         item.ProductId,
			ProductName:       getProduct.Name,
			ProductQty:        item.ProductQty,
			ProductPrice:      getProduct.Price,
			TotalProductPrice: getProduct.Price * float32(item.ProductQty),
		}

		_, err = h.orderService.AddOrderDetail(addItem)
		if err != nil {
			response := helper.APIResponse("Failed add shopping cart, saat insert db order details", http.StatusBadRequest, "error", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		currentStock := getProduct.Stock - item.ProductQty
		err = h.productService.UpdateProductByID(getProduct.ID, currentStock)
	}

	getOrder, _ := h.orderService.GetOrderByID(newOrder.ID)

	response := helper.APIResponse("Success add shopping cart", http.StatusOK, "success", order.FormatOrder(getOrder))

	ctx.JSON(http.StatusOK, response)
}
