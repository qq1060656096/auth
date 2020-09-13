package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type tokenTest struct {
	decoder         Decoder
	tokenType       string
	expires         int64
	data            string
	tokenStr        string
	refreshTokenStr string
}

var tokens = []tokenTest{
	{
		decoder: &JwtDecode{
			tokenType:        "Bearer",
			signingKey:       "123456",
			jwtSigningMethod: jwt.SigningMethodHS256,
		},
		tokenType:       "Bearer",
		data:            "token.test.1",
		expires:         2862226989,
		tokenStr:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI4NjIyMjY5ODksImlzcyI6ImV5SlNZWGRFWVhSaElqb2lkRzlyWlc0dWRHVnpkQzR4SWl3aVZHOXJaVzVVZVhCbElqb2lRbVZoY21WeUlpd2lWR2x0WlhOMFlXMXdJam94TmpBd01EQTNNamswTENKRmVIQnBjbVZ6SWpveU9EWXlNakkyT1RnNUxDSlNaV1p5WlhOb0lqcG1ZV3h6WlN3aVUybG5ibUYwZFhKbElqb2lJbjA9LllYTTFUV013WDJKVGVGQlpZWEIxU0Y5dVVsOXlaakJ1UjBzelMyMTFiRmh5ZDNKUlpFTldMVko2VVE9PSJ9.1B5v7E4B-nu7GesH0sKv3DLyzrDU_EdWkrpTYsrGqiA",
		refreshTokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI4NjIyMjY5ODksImlzcyI6ImV5SlNZWGRFWVhSaElqb2lkRzlyWlc0dWRHVnpkQzR4SWl3aVZHOXJaVzVVZVhCbElqb2lRbVZoY21WeUlpd2lWR2x0WlhOMFlXMXdJam94TmpBd01EQTNOelV6TENKRmVIQnBjbVZ6SWpveU9EWXlNakkyT1RnNUxDSlNaV1p5WlhOb0lqcDBjblZsTENKVGFXZHVZWFIxY21VaU9pSWlmUT09LmFHcE1hR2N0U0hkM1prRjBOalV0YzNvMlJtbG9YMVZYYm5aRmRUZDFjMmhXUkhsTFRqSmthRXQxV1E9PSJ9.r5es993kJV-TKdglCcB6OBdHJrwpHw97Pc-v7DXiyTw",
	},
}

func TestToken_GetBaseInfo(t *testing.T) {
	for _, tt := range tokens {
		token := NewToken(tt.decoder, tt.expires)
		assert.Equal(t, tt.decoder, token.GetDecoder())
		assert.Equal(t, tt.tokenType, token.GetType())
		assert.Equal(t, tt.expires, token.GetExpires())
	}
}

func TestToken_GetExpires(t *testing.T) {

}

func TestToken_Get(t *testing.T) {
	for _, tt := range tokens {
		token := NewToken(tt.decoder, tt.expires)
		newTokenStr, err := token.Get(tt.data)
		assert.NoError(t, err, "TestToken_Get")
		assert.Equal(t, 3, len(strings.Split(newTokenStr, ".")))
	}
}

func TestToken_GetRefresh(t *testing.T) {
	for _, tt := range tokens {
		token := NewToken(tt.decoder, tt.expires)
		newRefreshTokenStr, err := token.GetRefresh(tt.data)
		assert.NoError(t, err, "TestToken_GetRefresh")
		assert.Equal(t, 3, len(strings.Split(newRefreshTokenStr, ".")))
	}
}

func TestToken_Refresh(t *testing.T) {
	for _, tt := range tokens {
		token := NewToken(tt.decoder, tt.expires)
		newRefreshTokenStr, err := token.Refresh(tt.refreshTokenStr)
		assert.NoError(t, err, "TestToken_Refresh")
		assert.Equal(t, 3, len(strings.Split(newRefreshTokenStr, ".")))
	}
}

func TestToken_Validate(t *testing.T) {
	for _, tt := range tokens {
		token := NewToken(tt.decoder, tt.expires)
		tokenData, err := token.Validate(tt.tokenStr)
		assert.NoError(t, err, "TestToken_Validate.1")
		assert.Equal(t, tokenData.RawData, tt.data)
		assert.Equal(t, tokenData.Refresh, false)

		refreshTokenData, err := token.Validate(tt.refreshTokenStr)
		assert.NoError(t, err, "TestToken_Validate.2")
		assert.Equal(t, refreshTokenData.RawData, tt.data)
		assert.Equal(t, refreshTokenData.Refresh, true)

	}
}
