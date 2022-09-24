package user

import (
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	phone, err := strconv.Atoi(input.Phone)
	if err != nil {
		return User{}, err
	}
	postCode, err := strconv.Atoi(input.PostCode)
	if err != nil {
		return User{}, err
	}
	user := User{}
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Phone = phone
	user.Email = input.Email
	user.PasswordHash = input.Password
	user.Role = input.Role
	user.PostCode = postCode
	user.Address = input.Address
	user.City = input.City

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Email tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Tidak ditemukan user berdasakan ID ini")
	}

	return user, nil
}
