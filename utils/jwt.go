package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJwt(user *jwt.Token) string {
	claims := user.Claims.(jwt.MapClaims)
	fmt.Printf("claims: %v\n", claims)
	name := claims["username"].(string)
	return name
}
