package user

import "time"

type User struct {
	ID           int
	FirstName    string
	LastName     string
	Phone        int
	Email        string
	PasswordHash string
	// AvatarFileName string
	Role      string
	PostCode  int
	Address   string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
