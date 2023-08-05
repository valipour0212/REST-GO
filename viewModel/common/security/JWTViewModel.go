package security

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	UserID   string `json:"UserID"`
	UserName string `json:"UserName"`
	jwt.StandardClaims
}
