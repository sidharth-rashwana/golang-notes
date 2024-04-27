package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secret = "supersecret"

func GenerateToken(email string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"nbf":   time.Date(2023, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenStr, nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad signed method received")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func TokenCheck(jwtToken string) error {
	token, err := parseToken(jwtToken)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("unable to map claims")
	}
	return nil // Token is successfully parsed and claims are mapped
}
