package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const Signature = "my-secret"

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    username,
		ID:        username,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
	})
	return token.SignedString([]byte(Signature))
}

func ValidateToken(token string) (jwt.Claims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(Signature), nil
	})

	return t.Claims, err
}
