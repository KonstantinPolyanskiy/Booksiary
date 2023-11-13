package service

import (
	"Booksiary/auth-service/internal/domain"
	"Booksiary/auth-service/internal/repository"
	. "Booksiary/auth-service/pkg/password"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Token interface {
	Create(login, password string) (string, error)
	Parse(accessToken string) (domain.UserTokenData, error)
}
type TokenService struct {
	repository.Account
}

type TokenClaims struct {
	jwt.StandardClaims
	UserUUID uuid.UUID `json:"userUUID"`
	Role     int       `json:"role"`
}

func NewTokenService(repo repository.Account) *TokenService {
	return &TokenService{
		Account: repo,
	}
}

func (s *TokenService) Create(login, password string) (string, error) {
	account, err := s.Account.Get(login, Hash(password, passwordSalt))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&TokenClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
				IssuedAt:  time.Now().Unix(),
				Issuer:    "Auth service",
			},
			UserUUID: account.UUID,
			Role:     account.Role,
		})

	return token.SignedString([]byte(signingKey))

}

func (s *TokenService) Parse(accessToken string) (domain.UserTokenData, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return domain.UserTokenData{}, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return domain.UserTokenData{}, errors.New("claims is not a type")
	}

	return domain.UserTokenData{
		UUID:   claims.UserUUID,
		RoleId: claims.Role,
	}, nil
}
