package service

import (
	"CRUD/pkg/model"
	"CRUD/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"math/rand"
	"os"
	time "time"
)

type AuthService struct {
	repo repository.Authorisation
}

func (s *AuthService) RefreshToken(refreshToken string) (string, string, error) {
	userId, err := s.repo.CheckRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return "", "", err
	}

	return s.GenerateTokens(refreshToken, user)
}

func (s *AuthService) ParseToken(accessToken string) (int, int64, error) {
	signingKey := os.Getenv("SIGNINGKEY")
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, 0, err
	}
	claimes, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claimes.Userid, claimes.ExpiresAt, nil
}

type tokenClaims struct {
	jwt.StandardClaims
	Userid int `json:"user_id"`
}

func (s *AuthService) FindUserByUsernameAndPswd(username, password string) (model.User, error) {
	user, err := s.repo.FindUserByUserNameAndPswd(username, s.generatePasswordHash(password))
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *AuthService) GenerateTokens(oldToken string, user model.User) (string, string, error) {

	accessTokenTTL := time.Duration(viper.GetInt("tokens.accessTokenTTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	signingKey := os.Getenv("SIGNINGKEY")

	accessToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return "", "", err
	}

	err = s.repo.SaveRefreshToken(oldToken, refreshToken, user.Id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateRefreshToken() (string, error) {
	sl := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if _, err := r.Read(sl); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sl), nil
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	salt := os.Getenv("SALT")
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
