package token

import "encoding/json"

func CreateApiToken(expires int64, signingKey string) (jsonStr string, err error) {
	token := NewBearerJwtToken(expires, signingKey)
	apiTokenMap := make(map[string]interface{})
	data := "user.id.1"
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

	apiTokenMap["access_token"] = tokenStr
	apiTokenMap["token_type"] = token.GetType()
	apiTokenMap["expires_in"] = token.GetExpires()
	apiTokenMap["refresh_token"] = refreshTokenStr
	jsonBytes, err := json.Marshal(apiTokenMap)
	if err != nil {
		// error todo
		return
	}
	return string(jsonBytes), nil
}

func RefreshApiToken(refreshTokenStr string, expires int64, signingKey string) (newRefreshTokenStr string, err error) {
	token := NewBearerJwtToken(expires, signingKey)
	apiTokenMap := make(map[string]interface{})
	newRefreshTokenStr, err = token.Refresh(refreshTokenStr)
	if err != nil {
		// error todo
		return
	}
	apiTokenMap["access_token"] = newRefreshTokenStr
	apiTokenMap["token_type"] = token.GetType()
	apiTokenMap["expires_in"] = token.GetExpires()
	apiTokenMap["refresh_token"] = refreshTokenStr
	jsonBytes, err := json.Marshal(apiTokenMap)
	if err != nil {
		// error todo
		return
	}
	return string(jsonBytes), nil
}