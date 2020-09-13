package token

import "time"

type token struct {
	Tokener
	decoder   Decoder
	tokenType string
	expires   int64
}

func (t *token) GetDecoder() Decoder {
	return t.decoder
}

func (t *token) GetType() string {
	return t.tokenType
}

func (t *token) GetExpires() int64 {
	return t.expires
}

func (t *token) Get(data interface{}) (token string, error error) {
	time.Now().Unix()
	d := NewData(data, t.GetType(), t.GetExpires())
	return t.GetDecoder().Encode(d)
}

func (t *token) GetRefresh(data interface{}) (newRefreshToken string, error error) {
	time.Now().Unix()
	d := NewData(data, t.GetType(), t.GetExpires())
	d.Refresh = true
	return t.GetDecoder().Encode(d)
}

func (t *token) Refresh(refreshToken string) (newRefreshToken string, err error) {
	data, err := t.GetDecoder().Decode(refreshToken)
	if err != nil {
		return "", err
	}
	data.Timestamp = time.Now().Unix()
	data.Refresh = true
	return t.GetDecoder().Encode(data)
}

func (t *token) Validate(token string) (data *Data, err error) {
	return t.GetDecoder().Decode(token)
}

func NewToken(decoder Decoder, expires int64) *token {
	return &token{
		decoder:   decoder,
		tokenType: decoder.GetTokenType(),
		expires:   expires,
	}
}
