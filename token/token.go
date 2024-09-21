package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*payload, error)
}

type payload struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTMaker struct {
	secretKey string
}

func NewMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey: secretKey}
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	expiredAt := time.Now().Add(duration).Unix()

	p := &payload{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	return token.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*payload, error) {
	tkn, err := jwt.ParseWithClaims(token, &payload{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(maker.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := tkn.Claims.(*payload); ok && tkn.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
