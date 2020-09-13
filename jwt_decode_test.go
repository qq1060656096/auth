package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type newJwtDecodeTest struct {
	tokenType        string
	signingKey       string
	jwtSigningMethod jwt.SigningMethod
	data             Data
	tokenStr         string
}

var newJwtDecodesOk = []newJwtDecodeTest{
	{
		tokenType:        "Bearer",
		signingKey:       "12345678901234567890123456789012",
		jwtSigningMethod: jwt.SigningMethodHS256,
		data: Data{
			RawData:   "d.1.ok",
			TokenType: "tokenType.1",
			Timestamp: 0,
			Expires:   2862226989,
			Refresh:   false,
			Signature: "",
		},
		tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI4NjIyMjY5ODksImlzcyI6ImV5SlNZWGRFWVhSaElqb2laQzR4TG05cklpd2lWRzlyWlc1VWVYQmxJam9pZEc5clpXNVVlWEJsTGpFaUxDSlVhVzFsYzNSaGJYQWlPakFzSWtWNGNHbHlaWE1pT2pJNE5qSXlNalk1T0Rrc0lsSmxabkpsYzJnaU9tWmhiSE5sTENKVGFXZHVZWFIxY21VaU9pSWlmUT09LlFtTnRWRVIyUjIxcVpXeEhaelJoUnpWa2VWSXpPSHA0TTNWRFJuUjJiRXRzU1dKWlFUSlZOVkZqZHc9PSJ9.CMfu88i7o141_ak4YjtjSMfPq7xLQadfy48OnSU4mn0",
	},
	{
		tokenType:        "default",
		signingKey:       "12345678",
		jwtSigningMethod: jwt.SigningMethodHS256,
		data: Data{
			RawData:   "d.2.ok",
			TokenType: "tokenType.2",
			Timestamp: 0,
			Expires:   2862226989,
			Refresh:   true,
			Signature: "",
		},
		tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI4NjIyMjY5ODksImlzcyI6ImV5SlNZWGRFWVhSaElqb2laQzR5TG05cklpd2lWRzlyWlc1VWVYQmxJam9pZEc5clpXNVVlWEJsTGpJaUxDSlVhVzFsYzNSaGJYQWlPakFzSWtWNGNHbHlaWE1pT2pJNE5qSXlNalk1T0Rrc0lsSmxabkpsYzJnaU9uUnlkV1VzSWxOcFoyNWhkSFZ5WlNJNklpSjkuU1ZOdlpGWjFlVU0yTWpGS2VsUlpTa3hGYVZoS2RYRnJiM0pJYlVsSVJtcEtOMUZ3VTBkR1ZsSXdXUT09In0.nW3_ntA0vhKOmfIqixYhcMl5DwdmP2obDyPS-Lw7AXg",
	},
}

func TestNewJwtDecode(t *testing.T) {
	for _, v := range newJwtDecodesOk {
		jwtDecode := NewJwtDecode(v.tokenType, v.signingKey, v.jwtSigningMethod)
		assert.Equal(t, v.tokenType, jwtDecode.tokenType)
		assert.Equal(t, v.signingKey, jwtDecode.signingKey)
		assert.Equal(t, v.jwtSigningMethod, jwtDecode.jwtSigningMethod)
	}
}

func TestJwtDecode_Encode(t *testing.T) {
	for _, v := range newJwtDecodesOk {
		jwtDecode := NewJwtDecode(v.tokenType, v.signingKey, v.jwtSigningMethod)
		tokenStr, err := jwtDecode.Encode(&v.data)
		assert.NoError(t, err, "TestJwtDecode_Encode")
		assert.Equal(t, v.tokenType, jwtDecode.tokenType)
		assert.Equal(t, 3, len(strings.Split(tokenStr, ".")))
	}
}

func TestJwtDecode_Decode(t *testing.T) {
	for _, v := range newJwtDecodesOk {
		jwtDecode := NewJwtDecode(v.tokenType, v.signingKey, v.jwtSigningMethod)
		data, err := jwtDecode.Decode(v.tokenStr)
		assert.NoError(t, err)
		assert.Equal(t, v.data.RawData, data.RawData)
		assert.Equal(t, v.data.TokenType, data.TokenType)
		assert.Equal(t, v.data.Timestamp, data.Timestamp)
		assert.Equal(t, v.data.Expires, data.Expires)
		assert.Equal(t, v.data.Refresh, data.Refresh)
		assert.Equal(t, v.data.Signature, data.Signature)
	}
}
