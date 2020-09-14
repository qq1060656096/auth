package token

type ApiTokenData struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func CreateApiToken(expires int64, signingKey string, data interface{}) (td *ApiTokenData, err error) {

	token := NewBearerJwtToken(expires, signingKey)
	tokenStr, err := token.Get(data)
	if err != nil {
		// error todo
		return
	}
	refreshTokenStr, err := token.GetRefresh(data)
	if err != nil {
		// error todo
		return
	}
	td = &ApiTokenData{
		AccessToken:  tokenStr,
		TokenType:    token.GetType(),
		ExpiresIn:    token.GetExpires(),
		RefreshToken: refreshTokenStr,
	}
	return td, nil
}

func RefreshApiToken(expires int64, signingKey string, refreshTokenStr string) (td *ApiTokenData, err error) {
	token := NewBearerJwtToken(expires, signingKey)
	data, err := token.Validate(refreshTokenStr)
	if err != nil {
		return
	}
	tokenStr, err := token.Get(data)

	newRefreshTokenStr, err := token.Refresh(refreshTokenStr)
	if err != nil {
		// error todo
		return
	}
	td = &ApiTokenData{
		AccessToken:  tokenStr,
		TokenType:    token.GetType(),
		ExpiresIn:    token.GetExpires(),
		RefreshToken: newRefreshTokenStr,
	}
	return td, nil
}
