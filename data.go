package token

import "time"

// Data
type Data struct {
	RawData   interface{}
	TokenType string
	Timestamp int64
	Expires   int64
	Refresh   bool
	Signature string
}

// NewData
func NewData(data interface{}, tokenType string, expires int64) *Data {
	return &Data{
		RawData:   data,
		TokenType: tokenType,
		Timestamp: time.Now().Unix(),
		Expires:   expires,
	}
}

// NewRefreshData
func NewRefreshData(data interface{}, tokenType string, expires int64) *Data {
	d := NewData(data, tokenType, expires)
	d.Refresh = true
	return d
}
