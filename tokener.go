package token

type Tokener interface {
	GetDecoder() Decoder
	GetType() string
	GetExpires() int64
	Get(data interface{}) (token string, error error)
	GetRefresh(data interface{}) (newRefreshToken string, error error)
	Refresh(refreshToken string) (newRefreshToken string, err error)
	Validate(token string) (data *Data, err error)
}
