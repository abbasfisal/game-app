package authservice

import (
	"fmt"
	"github.com/abbasfisal/game-app/entity"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Config struct {
	SignKey               string
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}
type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

func (s Service) CreateAccessToken(u entity.User) (string, error) {
	return s.createToken(u.ID, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(u entity.User) (string, error) {
	return s.createToken(u.ID, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) createToken(userID uint, subject string, expiresDuration time.Duration) (string, error) {

	// set our claims
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresDuration)),
		},
		UserID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}
	// Creat token string
	return tokenString, nil
}

func (s Service) VerifyToken(bearerToken string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(bearerToken[len("Bearer "):], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("userID : %v , exp:  %v", claims.UserID, claims.ExpiresAt)
		return claims, nil
	} else {
		return nil, err
	}
}
