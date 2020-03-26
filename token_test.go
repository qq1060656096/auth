package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToken_GetAccessToken(t *testing.T) {
	token := New("test", 24 * 3600 * 365 * 30)
	token.Repository = NewJwtRepository("0123456789123456")
	tokenNew, err := token.GetAccessToken("test.token.data")
	if err != nil {
		assert.Error(t, err)
	}

	assert.True(t, false, tokenNew)
}

func TestToken_Refresh(t *testing.T) {

}

func TestToken_Validate(t *testing.T) {

}