package token

import (
	"encoding/json"
	"time"
)

type Data struct {
	RawData interface{}
	Type string
	Time time.Time
	Expires int64
	Refresh bool
}

func NewData(data interface{}, tokenType string, expires int64) *Data {
	return &Data{
		RawData: data,
		Type: tokenType,
		Time: time.Now(),
		Expires: expires,
	}
}

func NewRefreshData(data interface{}, tokenType string, expires int64) *Data {
	d := NewData(data, tokenType, expires)
	d.Refresh = true
	return d
}

func (d *Data) Bytes() []byte {
	bytes, _:= json.Marshal(d)
	return bytes
}