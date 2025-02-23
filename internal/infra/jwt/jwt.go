package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTItf interface {
	GenerateToken(userId uuid.UUID, isAdmin bool) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, bool, error)
}

type JWT struct {
	SecretKey   string
	ExpiredTime time.Duration
}

func NewJWT(secretKey string, expiredTime time.Duration) *JWT {
	return &JWT{
		SecretKey:   secretKey,
		ExpiredTime: expiredTime,
	}
}

type Claims struct {
	UserId  uuid.UUID
	IsAdmin bool
	jwt.RegisteredClaims
}

func (j *JWT) GenerateToken(userId uuid.UUID, isAdmin bool) (string, error) {
	claim := Claims{
		UserId:  userId,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ValidateToken(tokenString string) (uuid.UUID, bool, error) {
	var claim Claims
	var id uuid.UUID

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return id, false, err
	}

	if !token.Valid {
		return id, false, err
	}

	id = claim.UserId

	return id, claim.IsAdmin, nil
}
