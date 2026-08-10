package usecase

import (
	"adoend/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	Repo      *repository.AuthRepository
	JWTSecret string
}

func NewAuthUsecase(r *repository.AuthRepository, secret string) *AuthUsecase {
	return &AuthUsecase{
		Repo:      r,
		JWTSecret: secret,
	}
}

func (u *AuthUsecase) Register(email, password, name, phone string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	return u.Repo.CreateUserWithProfile(
		email,
		string(hashed),
		name,
		"",
	)
}

func (u *AuthUsecase) Login(email, password string) (string, error) {
	user, err := u.Repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(u.JWTSecret))

	if err != nil {
		return "", err
	}

	return signed, nil
}
