package user

type RegisterUserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
	PostCode  string `json:"post_code" binding:"required"`
	Address   string `json:"address" binding:"required"`
	City      string `json:"city" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
