package token

import "encoding/json"

type Token struct {
	Repository   Repository `json:"-"`
	AccessToken  string `json:"access_token"`
	Type    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expires    int64  `json:"expires_in"`
}


func New(tokenType string, expires int64) *Token {
	return &Token{
		Type: tokenType,
		Expires: expires,
	}
}

func (t *Token) GetAccessToken(data interface{}) (token *Token, err error){
	tokenData := NewData(data, t.Type, t.Expires)
	accessToken, err := t.Repository.Encode(tokenData)
	if err != nil {
		return
	}
	refreshTokenData := NewRefreshData(data, t.Type, t.Expires)
	refreshToken, err := t.Repository.Encode(refreshTokenData)
	if err != nil {
		return
	}
	token = New(t.Type, t.Expires)
	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	return
}

func (t *Token) Refresh(refreshToken string) (newToken *Token, err error) {
	refreshTokenData, err := t.Repository.Validate(refreshToken)
	if err != nil {
		return
	}
	refreshTokenNew, err := t.Repository.Encode(refreshTokenData)
	if err != nil {
		return
	}
	tokenData := NewData(refreshTokenData.RawData, refreshTokenData.Type, refreshTokenData.Expires)
	accessToken, err := t.Repository.Encode(tokenData)
	if err != nil {
		return
	}
	token := New(t.Type, t.Expires)
	token.AccessToken = accessToken
	token.RefreshToken = refreshTokenNew
	return
}

func (t *Token) Validate(token string) (data *Data, err error) {
	return t.Repository.Validate(token)
}

func (t *Token) String() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}