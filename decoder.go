package token

type Decoder interface {
	GetTokenType() string
	Encode(data *Data) (token string, err error)
	Decode(token string) (data *Data, err error)
}
