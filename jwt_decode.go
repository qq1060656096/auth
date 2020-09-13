package token

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type JwtDecode struct {
	tokenType        string
	signingKey       string
	jwtSigningMethod jwt.SigningMethod
}

func (o *JwtDecode) GetTokenType() string {
	return o.tokenType
}

func (o *JwtDecode) Encode(data *Data) (tokenStr string, err error) {
	dataStr, err := o.dataEncode(data)
	if err != nil {
		return
	}
	mySigningKey := []byte(o.signingKey)
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: data.Expires,
		Issuer:    string(dataStr),
	}
	jwtToken := jwt.NewWithClaims(o.jwtSigningMethod, claims)
	tokenStr, err = jwtToken.SignedString(mySigningKey)
	return tokenStr, err
}

func (o *JwtDecode) Decode(tokenStr string) (data *Data, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(o.signingKey), nil
	})
	if !token.Valid {
		return
	}
	standardClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	var d Data
	dataStr, ok := standardClaims["iss"].(string)
	if !ok {
		return
	}
	return o.dataDecode(dataStr, d)
}

func (o *JwtDecode) dataEncode(data *Data) (dataStr string, err error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	dataStr = string(dataBytes)
	method := jwt.GetSigningMethod(jwt.SigningMethodHS256.Alg())
	signature, err := method.Sign(dataStr, []byte(o.signingKey))
	if err != nil {
		return "", err
	}
	dataStrBase64 := base64.StdEncoding.EncodeToString(dataBytes)
	signatureBase64 := base64.StdEncoding.EncodeToString([]byte(signature))
	if err != nil {
		return
	}
	dataStr = fmt.Sprintf("%s.%s", dataStrBase64, signatureBase64)
	return
}

func (o *JwtDecode) dataDecode(dataStr string, d Data) (data *Data, err error) {
	strSlice := strings.Split(dataStr, ".")
	var dataStrBase64, signatureBase64 string
	if len(strSlice) > 1 {
		dataStrBase64, signatureBase64 = strSlice[0], strSlice[1]
	} else {
		dataStrBase64 = strSlice[0]
	}

	dataStrBytes, err := base64.StdEncoding.DecodeString(dataStrBase64)
	if err != nil {
		return
	}
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return
	}
	method := jwt.GetSigningMethod(jwt.SigningMethodHS256.Alg())
	err = method.Verify(string(dataStrBytes), string(signature), []byte(o.signingKey))
	if err != nil {
		return
	}
	err = json.Unmarshal(dataStrBytes, &d)
	data = &d
	if err != nil {
		return nil, err
	}
	return
}

func NewJwtDecode(tokenType, signingKey string, jwtSigningMethod jwt.SigningMethod) *JwtDecode {
	return &JwtDecode{
		tokenType:        tokenType,
		signingKey:       signingKey,
		jwtSigningMethod: jwtSigningMethod,
	}
}
