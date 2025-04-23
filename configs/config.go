package configs

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	// G global configuration variable
	G *Config
)

type Config struct {
	Server      ServerConfig     `mapstructure:"server" json:"server"`
	DB          DBConfig         `mapstructure:"db" json:"db"`
	Redis       RedisConfig      `mapstructure:"redis" json:"redis"`
	Log         LogConfig        `mapstructure:"log" json:"log"`
	HttpClient  HttpClientConfig `mapstructure:"http_client" json:"http_client"`
	VotingCore  VotingCoreConfig `mapstructure:"voting_core" json:"voting_core"`
	Jwt         JwtConfig        `mapstructure:"jwt" json:"jwt"`
	PasswordKey string           `mapstructure:"password_key" json:"password_key"`
	Voting      VotingConfig     `mapstructure:"voting" json:"voting"`
	Aptos       AptosConfig      `mapstructure:"aptos" json:"aptos"`
}

func Init(path string) {
	if path != "" {
		viper.SetConfigFile(path)
	}

	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&G); err != nil {
		panic(err)
	}
}
