package configs

type JwtConfig struct {
	SecretKey string `json:"secret_key" mapstructure:"secret_key"`
	Expire    int64  `json:"expire_time" mapstructure:"expire_time"`
}
