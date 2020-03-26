package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestJwtRepository_Encode(t *testing.T) {
	r := NewJwtRepository("0123456789123456")
	data := NewData("测试数据", "test", 24 * 3600 * 365 * 30)
	token, err := r.Encode(data)
	if err != nil {
		assert.Error(t, err)
	}
	assert.GreaterOrEqual(t, len(token), 1, token, err)
}

func TestJwtRepository_Decode(t *testing.T) {
	// "测试数据", "test", 24 * 3600 * 365 * 30
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI1MzEzMTU2ODMsImlzcyI6ImV5SlNZWGRFWVhSaElqb2k1cldMNksrVjVwV3c1bzJ1SWl3aVZIbHdaU0k2SW5SbGMzUWlMQ0pVYVcxbElqb2lNakF5TUMwd015MHlObFF5TXpveE5EbzBNeTQ0T1RnMk1UZ3JNRGc2TURBaUxDSkZlSEJwY21WeklqbzVORFl3T0RBd01EQXNJbEpsWm5KbGMyZ2lPbVpoYkhObGZRPT0ifQ.5uwHm94K9Ob3b4C3Z0DqWgBmWTVx0xEa2KBmHGI2kTU"
	r := NewJwtRepository("0123456789123456")
	data, err := r.Decode(token)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, int64(24 * 3600 * 365 * 30), data.Expires)
	assert.Equal(t, "测试数据", data.RawData)
	assert.Equal(t, "test", data.Type)
}

func TestJwtRepository_Validate(t *testing.T) {
	// "测试数据", "test", 24 * 3600 * 365 * 30
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI1MzEzMTU2ODMsImlzcyI6ImV5SlNZWGRFWVhSaElqb2k1cldMNksrVjVwV3c1bzJ1SWl3aVZIbHdaU0k2SW5SbGMzUWlMQ0pVYVcxbElqb2lNakF5TUMwd015MHlObFF5TXpveE5EbzBNeTQ0T1RnMk1UZ3JNRGc2TURBaUxDSkZlSEJwY21WeklqbzVORFl3T0RBd01EQXNJbEpsWm5KbGMyZ2lPbVpoYkhObGZRPT0ifQ.5uwHm94K9Ob3b4C3Z0DqWgBmWTVx0xEa2KBmHGI2kTU"
	r := NewJwtRepository("0123456789123456")
	data, err := r.Validate(token)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, int64(24 * 3600 * 365 * 30), data.Expires)
	assert.Equal(t, "测试数据", data.RawData)
	assert.Equal(t, "test", data.Type)
}