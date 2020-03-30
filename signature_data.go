package token

import 	"encoding/json"

type SignatureData struct {
	Raw string
	Signature string
}

func (d *SignatureData) Bytes() []byte {
	bytes, _:= json.Marshal(d)
	return bytes
}