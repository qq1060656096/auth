package token

import (
	"encoding/base64"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"time"
)


type JwtRepository struct {
	Repository
	SecurityKey string

}

func NewJwtRepository(securityKey string) *JwtRepository {
	return &JwtRepository{
		SecurityKey: securityKey,
	}
}

func (r *JwtRepository) Encode(data *Data) (token string, err error) {
	// 1. json
	// 2. base64
	// 3. jwt
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(data.Expires) * time.Second).Unix(),
		Issuer:    base64.StdEncoding.EncodeToString(data.Bytes()),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(r.SecurityKey))
	return
}

func (r *JwtRepository) Decode(token string) (data *Data, err error) {
	// 1. jwt
	// 2. base64
	// 3. json
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.SecurityKey), nil
	})
	if err != nil {
		return
	}
	if !jwtToken.Valid {
		return
	}
	tokenMsp, _ := jwtToken.Claims.(jwt.MapClaims)
	base64Issuer, _ := tokenMsp["iss"].(string)
	issuer, err := base64.StdEncoding.DecodeString(base64Issuer)
	if err != nil {
		return
	}
	err = json.Unmarshal(issuer, &data)
	if err != nil {
		return
	}
	return
}


func (r *JwtRepository) Validate(token string) (data *Data, err error) {
	data, err = r.Decode(token)
	return
}
