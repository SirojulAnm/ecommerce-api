package user

type UserFormatter struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     int    `json:"phone"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	PostCode  int    `json:"post_code"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Token     string `json:"access_token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Email:     user.Email,
		Role:      user.Role,
		PostCode:  user.PostCode,
		Address:   user.Address,
		City:      user.City,
		Token:     token,
	}

	return formatter
}
