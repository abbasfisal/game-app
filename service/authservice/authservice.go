package authservice

import (
	"fmt"
	"github.com/abbasfisal/game-app/entity"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service struct {
	signKey               string
	accessExpirationTime  time.Duration
	refreshExpirationTime time.Duration
	accessSubject         string
	refreshSubject        string
}

func New(signKey string, accessSubject, refreshSubject string, accessExpirationTime, refreshExpirationTime time.Duration) Service {
	return Service{
		signKey:               signKey,
		accessExpirationTime:  accessExpirationTime,
		refreshExpirationTime: refreshExpirationTime,
		accessSubject:         accessSubject,
		refreshSubject:        refreshSubject,
	}
}
func (s Service) CreateAccessToken(u entity.User) (string, error) {
	return s.createToken(u.ID, s.accessSubject, s.accessExpirationTime)
}

func (s Service) CreateRefreshToken(u entity.User) (string, error) {
	return s.createToken(u.ID, s.refreshSubject, s.refreshExpirationTime)
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
	tokenString, err := accessToken.SignedString([]byte(s.signKey))
	if err != nil {
		return "", err
	}
	// Creat token string
	return tokenString, nil
}

func (s Service) VerifyToken(bearerToken string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(bearerToken[len("Bearer "):], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
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
