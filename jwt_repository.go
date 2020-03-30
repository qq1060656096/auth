package token

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/wumansgy/goEncrypt"
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
	// 1. data json
	// 2. aes cbc
	// 3. 签名
	// 4. base64
	// 5. jwt
	skBytes := []byte(r.SecurityKey)
	// 传入明文和自己定义的密钥，密钥为8字节
	aesCbcData, err := goEncrypt.AesCbcEncrypt(data.Bytes(), skBytes) //得到密文,可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	if err != nil {
		return
	}
	sd := SignatureData{
		Raw: base64.StdEncoding.EncodeToString(aesCbcData),
		Signature: r.Signature(data),
	}
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(data.Expires) * time.Second).Unix(),
		Issuer:    base64.StdEncoding.EncodeToString(sd.Bytes()),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString(skBytes)
	return
}

func (r *JwtRepository) Decode(token string) (data *Data, err error) {
	// 1. jwt
	// 2. base64
	// 3. 验证签名
	skBytes := []byte(r.SecurityKey)
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
	signatureData := &SignatureData{}
	err = json.Unmarshal(issuer, &signatureData)
	if err != nil {
		return
	}
	b, err := base64.StdEncoding.DecodeString(signatureData.Raw)
	if err != nil {
		return
	}
	// 传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错 可以自己传入初始化向量,如果不传就使用默认的初始化向量,16字节
	dataStr, err := goEncrypt.AesCbcDecrypt(b, skBytes)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataStr, &data)
	if err != nil {
		return
	}
	return
}


func (r *JwtRepository) Validate(token string) (data *Data, err error) {
	data, err = r.Decode(token)
	return
}

func (r *JwtRepository) Signature(data *Data) string {
	// 1. 数据json
	// 2. sha256签名
	// 3. sha1签名
	bytes := data.Bytes()
	h := sha256.New()
	h.Write(bytes)
	h.Write([]byte(r.SecurityKey))
	signatureSha256 := fmt.Sprintf("%x", h.Sum(nil))

	h1 := sha1.New()
	h1.Write([]byte(signatureSha256))
	signature := fmt.Sprintf("%x", h1.Sum(nil))
	return signature
}
