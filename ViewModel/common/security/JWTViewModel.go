package security

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	UserName string `json:"username"`
	UserId   string `json:"userId"`
	jwt.StandardClaims
}
