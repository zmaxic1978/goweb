package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo "github.com/zmaxic1978/goweb"
	"github.com/zmaxic1978/goweb/pkg/repository"
	"time"
)

const (
	salt      = "jsl$567jDF6%7Gas!d2#a21SD^?fgdTU&"
	tokenSign = "GHTR576GHGGDHFSFST45"
	tokenTTL  = 1 * time.Hour
	loginErr  = "login or password is invalid. check and repeat again"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) CreateToken(login todo.Login) (string, error) {
	login.Password = generatePasswordHash(login.Password)
	user, err := s.repo.GetUser(login)
	if err != nil {

		return "", todo.AuthorizationError{Message: loginErr}
	}

	tokenClms := &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(tokenTTL).Unix(), IssuedAt: time.Now().Unix()},
		user.Id}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClms)
	signedToken, err := token.SignedString([]byte(tokenSign))
	if err != nil {
		return "", todo.InternalError{Message: err.Error()}
	}

	return signedToken, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, todo.AuthorizationError{Message: "invalid signing method"}
		}

		return []byte(tokenSign), nil
	})

	if err != nil {
		return 0, todo.AuthorizationError{Message: err.Error()}
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, todo.AuthorizationError{Message: "invalid token claims type"}
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
