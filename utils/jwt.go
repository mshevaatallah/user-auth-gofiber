package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("SECRET_COY")

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("asu")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOK := token.Claims.(jwt.MapClaims)
	if isOK && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("asu")

}
