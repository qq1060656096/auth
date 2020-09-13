package token

import (
	"encoding/json"
	"log"
)


func ExampleNewBearerJwtToken() {
	expires := int64(2862226989)
	signingKey := "123456"
	token := NewBearerJwtToken(expires, signingKey)
	apiTokenMap := make(map[string]interface{})
	data := "user.id.1"
	tokenStr, err := token.Get(data)
	if err != nil {
		// error todo
	}
	refreshTokenStr, err := token.GetRefresh(data)
	if err != nil {
		// error todo
	}

	apiTokenMap["access_token"] = tokenStr
	apiTokenMap["token_type"] = token.GetType()
	apiTokenMap["expires_in"] = token.GetExpires()
	apiTokenMap["refresh_token"] = refreshTokenStr
	jsonStr, err := json.Marshal(apiTokenMap)
	if err != nil {
		// error todo
	}
	log.Println(string(jsonStr))
	/**
	{
	    "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI4NjIyMjY5ODksImlzcyI6ImV5SlNZWGRFWVhSaElqb2lkWE5sY2k1cFpDNHhJaXdpVkc5clpXNVVlWEJsSWpvaVFtVmhjbVZ5SWl3aVZHbHRaWE4wWVcxd0lqb3hOakF3TURBNU9USXpMQ0pGZUhCcGNtVnpJam95T0RZeU1qSTJPVGc1TENKU1pXWnlaWE5vSWpwbVlXeHpaU3dpVTJsbmJtRjBkWEpsSWpvaUluMD0uUjBaQlp6bDBXWEpoYm14S1J6UllTbWwyWTB4ZlUyVTFlbXRoT1VKbE9FNXVkREpEVW5GMGExVkhRUT09In0.ZwSwy7LFsRd7DpWZ9v_o09wXisZoqmAzo2YvdaLVG0w",
	    "expires_in":2862226989,
	    "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI4NjIyMjY5ODksImlzcyI6ImV5SlNZWGRFWVhSaElqb2lkWE5sY2k1cFpDNHhJaXdpVkc5clpXNVVlWEJsSWpvaVFtVmhjbVZ5SWl3aVZHbHRaWE4wWVcxd0lqb3hOakF3TURBNU9USXpMQ0pGZUhCcGNtVnpJam95T0RZeU1qSTJPVGc1TENKU1pXWnlaWE5vSWpwMGNuVmxMQ0pUYVdkdVlYUjFjbVVpT2lJaWZRPT0uTVZWWlMySjBhMlZPUm5ZM01tOTZObVZvWDNWTlRHTldiMlpwY2pWd05HTXRSRTQwVjI1QlJXbEhOQT09In0.l2WQwypvE-RhrG6KZK_eO5Yp0JRiAJMbSjbq1Qxlib8",
	    "token_type":"Bearer"
	}
	 */
}
