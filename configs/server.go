package configs

type ServerConfig struct {
	APIBindAddress string `mapstructure:"api_bind_address" json:"api_bind_address"`
	Mode           string `mapstructure:"mode" json:"mode"`
	AdminAPIKey    string `mapstructure:"admin_api_key" json:"admin_api_key"`
}
