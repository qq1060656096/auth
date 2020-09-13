package token

import "github.com/dgrijalva/jwt-go"

func NewJwtToken(tokenType string, expires int64, signingKey string, jwtSigningMethod jwt.SigningMethod) *token {
	decoder := NewJwtDecode(tokenType, signingKey, jwtSigningMethod)
	return NewToken(decoder, expires)
}

func NewBearerJwtToken(expires int64, signingKey string) *token {
	tokenType := "Bearer"
	jwtSigningMethod := jwt.SigningMethodHS256
	decoder := NewJwtDecode(tokenType, signingKey, jwtSigningMethod)
	return NewToken(decoder, expires)
}
