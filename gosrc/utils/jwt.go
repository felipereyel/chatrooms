package utils

import (
	"chatrooms/gosrc/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id, username string) (string, time.Time, error) {
	secret := config.Configs.JWTSecret

	expiration := time.Now().Add(7 * 24 * time.Hour)

	claims := &jwt.RegisteredClaims{
		ID:        id,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expiration),
		Subject:   username,
	}

	jwtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := jwtoken.SignedString([]byte(secret))

	return jwt, expiration, err
}

func ParseJWT(token string) (string, error) {
	secret := config.Configs.JWTSecret
	claims := jwt.RegisteredClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims.ID == "" {
		return "", jwt.ErrInvalidKey
	}

	return claims.ID, nil
}
