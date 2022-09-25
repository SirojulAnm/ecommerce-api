package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"superindo/auth"
	"superindo/category"
	"superindo/handler"
	"superindo/helper"
	"superindo/order"
	"superindo/product"
	"superindo/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	categoryRepository := category.NewRepository(db)
	productRepository := product.NewRepository(db)
	orderRepository := order.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	categoryService := category.NewService(categoryRepository)
	productService := product.NewService(productRepository)
	orderService := order.NewService(orderRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService, categoryService)
	orderHandler := handler.NewOrderHandler(orderService, productService)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/user", authMiddleware(authService, userService), userHandler.FetchUser)

	api.POST("/category", authMiddleware(authService, userService), categoryHandler.AddCategory)
	api.GET("/category", authMiddleware(authService, userService), categoryHandler.GetCategories)

	api.POST("/product", authMiddleware(authService, userService), productHandler.AddProduct)
	api.GET("/product/:id", authMiddleware(authService, userService), productHandler.GetProductDetail)
	api.GET("/category/:id/product", authMiddleware(authService, userService), productHandler.GetProductsByCategoryID)

	api.POST("/order", authMiddleware(authService, userService), orderHandler.AddOrder)

	router.Run(":8083")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Beraer tokentoken
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("currentUser", user)
	}
}
