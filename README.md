# Token
> 快速生成和刷新 Auth Token 

### 使用示例
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/qq1060656096/token"
)

func main() {
	var expires int64 = 7200000000
	signingKey := "123456"
	data := "api token data"
	// 创建token
	apiTokenData, err := token.CreateApiToken(expires, signingKey, data)
	if err != nil {
		// error todo
	}
	jsonBytes, err := json.Marshal(apiTokenData)
	if err != nil {
		// error todo
	}
	fmt.Println("CreateApiToken", string(jsonBytes))
	/**
	CreateApiToken
	{
	    "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdPVEV4Tnpnc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPbVpoYkhObExDSlRhV2R1WVhSMWNtVWlPaUlpZlE9PS5TR0ptVlV0d1pXTjFZM0U1TURCcExYb3pVbXBDWWtodmNVeDRNWE5IUkhCaFNXdHhRMmgwWW1kT1ZRPT0ifQ.JfIgN-I55vnullHTZ9EJVwVlGJX4bB4mdGYo9VJOOL8",
	    "token_type":"Bearer",
	    "expires_in":7200000000,
	    "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdPVEV4Tnpnc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPblJ5ZFdVc0lsTnBaMjVoZEhWeVpTSTZJaUo5LlJucEJXR05KVFcxeloweEtkSGcwU0V4S04wZGljMDV4YkRsYVlsUlFlWEZyVm1OamFFZzVNVkppZHc9PSJ9.xnp-jINwlu42Ty1UDUAJjJVHQcCebMJENAj_juNn5DI"
	}
	*/

	// 刷新token
	refreshTokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdPVEV4Tnpnc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPblJ5ZFdVc0lsTnBaMjVoZEhWeVpTSTZJaUo5LlJucEJXR05KVFcxeloweEtkSGcwU0V4S04wZGljMDV4YkRsYVlsUlFlWEZyVm1OamFFZzVNVkppZHc9PSJ9.xnp-jINwlu42Ty1UDUAJjJVHQcCebMJENAj_juNn5DI"
	apiTokenData, err = token.RefreshApiToken(expires, signingKey, refreshTokenStr)
	if err != nil {
		// error todo
	}
	jsonBytes, err = json.Marshal(apiTokenData)
	if err != nil {
		// error todo
	}
	fmt.Println("RefreshApiToken", string(jsonBytes))
	/**
	RefreshApiToken
	{
	    "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqcDdJbEpoZDBSaGRHRWlPaUpoY0drZ2RHOXJaVzRnWkdGMFlTSXNJbFJ2YTJWdVZIbHdaU0k2SWtKbFlYSmxjaUlzSWxScGJXVnpkR0Z0Y0NJNk1UWXdNREE1TVRFM09Dd2lSWGh3YVhKbGN5STZOekl3TURBd01EQXdNQ3dpVW1WbWNtVnphQ0k2ZEhKMVpTd2lVMmxuYm1GMGRYSmxJam9pSW4wc0lsUnZhMlZ1Vkhsd1pTSTZJa0psWVhKbGNpSXNJbFJwYldWemRHRnRjQ0k2TVRZd01EQTVNVEkzTXl3aVJYaHdhWEpsY3lJNk56SXdNREF3TURBd01Dd2lVbVZtY21WemFDSTZabUZzYzJVc0lsTnBaMjVoZEhWeVpTSTZJaUo5LmNYVkRlbGQ2V25JMlJHMU5aRkkxWVcxbU1tWldiak5hTVdSUE9IRlhSVmRIV1dGVVNXSTFkR2RpU1E9PSJ9.ffmMCUMsyZB36xweTpN-vrt4Tq7358RW4eKkPhigQoc",
	    "token_type":"Bearer",
	    "expires_in":7200000000,
	    "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdPVEV5TnpNc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPblJ5ZFdVc0lsTnBaMjVoZEhWeVpTSTZJaUo5LloyWmhhRFZGZDJReVJIbGtRMnB1TFZOT2EwdDNSelpqUzFrMGNqbHZialpTWm1JNWNGTXhaRWx0WXc9PSJ9.EeOnpLZXZ47cAYVxdView7ifJuIGEdSLyjpC8pTmH4c"
	}
	*/
}

```