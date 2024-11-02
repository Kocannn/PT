package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kocannn/api-programming-tadulako/domain"
	"github.com/kocannn/api-programming-tadulako/helper"
)

type UserServices struct {
	UserRepo domain.UsersRepository
}

func NewUserServices(userRepo domain.UsersRepository) domain.UserServices {
	return &UserServices{
		UserRepo: userRepo,
	}
}

func (s *UserServices) Register(user *domain.User) error {
	if err := validateUserInput(user); err != nil {
		return err
	}
	hassPass, err := helper.HassPass(user.Password)
	if err != nil {
		return err
	}
	user.Password = hassPass

	return s.UserRepo.Create(user)
}

func (s *UserServices) Login(email, password string) (string, error) {
	user, err := s.UserRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	isPasswordValid, err := helper.ComparePass([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	if !isPasswordValid {
		return "", errors.New("invalid password")
	}

	token, err := generateJWTToken(user, nil)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWTToken(user *domain.User, admin *domain.Admin) (string, error) {
	type CustomUserAdmin struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	var newUserAdmin CustomUserAdmin
	if user != nil {
		newUserAdmin.Email = user.Email
		newUserAdmin.Username = user.Username
	}
	if admin != nil {
		newUserAdmin.Email = admin.Email
		newUserAdmin.Username = admin.Username
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": newUserAdmin.Username,
		"email":    newUserAdmin.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	secret := os.Getenv("SECRET")
	if secret == "" {
		return "", errors.New("secret key is not set in environment variables")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateUserInput(user *domain.User) error {
	if user.Username == "" || user.Email == "" || user.Phone == "" || user.Full_name == "" || user.Address == "" || user.Password == "" {
		return errors.New("missing fields input")
	}
	return nil
}
