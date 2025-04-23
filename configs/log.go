package configs

type LogConfig struct {
	Level int `mapstructure:"level" json:"level" default:"1"`
}
