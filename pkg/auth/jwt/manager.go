package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Manager struct {
	ttl  time.Duration
	sign string
}

func NewManager(ttl time.Duration, sign string) *Manager {
	return &Manager{ttl: ttl, sign: sign}
}

func (m *Manager) NewToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(m.ttl).Unix(),
		Subject:   id,
	})

	return token.SignedString([]byte(m.sign))
}

func (m *Manager) Parse(token string) (string, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.sign), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error getting user claims from token")
	}

	return claims["sub"].(string), nil
}
