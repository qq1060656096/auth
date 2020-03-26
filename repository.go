package token

type Repository interface {
	Encode(data *Data) (token string, err error)
	Decode(token string) (data *Data, err error)
	Validate(token string) (data *Data, err error)
}