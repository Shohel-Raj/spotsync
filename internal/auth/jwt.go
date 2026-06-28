package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtSecretKey         = "your_secret"
	defaultTokenDuration = 7 * 24 * time.Hour
	jwtIssuer            = "spotsync"
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID uint, email, name, role string) (string, error)
	ValidateToken(tokenStr string) (*JWTClaims, error)
}

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTService(secretKey string) JWTService {
	if secretKey == "" {
		secretKey = jwtSecretKey
	}

	return &jwtService{
		secretKey:     secretKey,
		tokenDuration: defaultTokenDuration,
	}
}

func (js *jwtService) GenerateToken(userID uint, email, name, role string) (string, error) {
	now := time.Now()

	claims := JWTClaims{
		UserID: userID,
		Name:   name,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(js.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    jwtIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.secretKey))
}

func (js *jwtService) ValidateToken(tokenStr string) (*JWTClaims, error) {
	if tokenStr == "" {
		return nil, errors.New("token is required")
	}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&JWTClaims{},
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(js.secretKey), nil
		},
		jwt.WithIssuer(jwtIssuer),
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
