# Token
> 快速生成和刷新 Auth Token 

```go
package main

import (
	"fmt"
	"github.com/qq1060656096/token"
)

func main()  {
	var expires int64 = 7200000000
	signingKey := "123456"
	data := "api token data"
	// 创建token
	jsonStr, err := token.CreateApiToken(expires, signingKey, data)
	if err != nil {
		// error todo
	}
	fmt.Println("CreateApiToken", jsonStr)
	/**
	CreateApiToken
	{
	    "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdOREEwTkRNc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPbVpoYkhObExDSlRhV2R1WVhSMWNtVWlPaUlpZlE9PS5SaTF6WXpabU9WQlBSR1JPY2tWVlVHeHlhM0JtU3pGa1EyZFRkMVZpV0RsdGFuUkNOMWswTkVoWFdRPT0ifQ.YKLDHu48xq8hguGexDKOtbxK2m5fYpBsFNcfuToSxwI",
	    "expires_in":7200000000,
	    "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdOREEwTkRNc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPblJ5ZFdVc0lsTnBaMjVoZEhWeVpTSTZJaUo5LlYwc3RRbWxvUlVWelNEZzNZVU0wY1VKd1oyUnhlbmR3TFhWbVFVaGhabWhOZUhWdGRHOTZTVXd3Wnc9PSJ9.YNWtDodFyjKes2FQBl2SATK6G3OEdLYQFly8ReQ1JHg",
	    "token_type":"Bearer"
	}
	 */

	// 刷新token
	refreshTokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdOREEwTkRNc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPblJ5ZFdVc0lsTnBaMjVoZEhWeVpTSTZJaUo5LlYwc3RRbWxvUlVWelNEZzNZVU0wY1VKd1oyUnhlbmR3TFhWbVFVaGhabWhOZUhWdGRHOTZTVXd3Wnc9PSJ9.YNWtDodFyjKes2FQBl2SATK6G3OEdLYQFly8ReQ1JHg"
	jsonStr, err = token.RefreshApiToken(expires, signingKey, refreshTokenStr)
	if err != nil {
		// error todo
	}
	fmt.Println("RefreshApiToken", jsonStr)
	/**
	RefreshApiToken
	{
	    "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdOREExTXpFc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPblJ5ZFdVc0lsTnBaMjVoZEhWeVpTSTZJaUo5LmRYaE5NR3BWUTIxelJqZDBTVlJJWHpoblprMTZSMmxUYkdOMmRFRm9SMHMzYVhkYVUzWjVVelpCWXc9PSJ9.-z33MwAwXFd-GXQnZgWj5JPBujya5zOlngrFJinV9UE",
	    "expires_in":7200000000,
	    "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjcyMDAwMDAwMDAsImlzcyI6ImV5SlNZWGRFWVhSaElqb2lZWEJwSUhSdmEyVnVJR1JoZEdFaUxDSlViMnRsYmxSNWNHVWlPaUpDWldGeVpYSWlMQ0pVYVcxbGMzUmhiWEFpT2pFMk1EQXdOREEwTkRNc0lrVjRjR2x5WlhNaU9qY3lNREF3TURBd01EQXNJbEpsWm5KbGMyZ2lPblJ5ZFdVc0lsTnBaMjVoZEhWeVpTSTZJaUo5LlYwc3RRbWxvUlVWelNEZzNZVU0wY1VKd1oyUnhlbmR3TFhWbVFVaGhabWhOZUhWdGRHOTZTVXd3Wnc9PSJ9.YNWtDodFyjKes2FQBl2SATK6G3OEdLYQFly8ReQ1JHg",
	    "token_type":"Bearer"
	}
	 */
}
```