package utils

import "github.com/golang-jwt/jwt/v5"

var secretKey = []byte("SECRET_COY")

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
