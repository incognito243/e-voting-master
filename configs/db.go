package configs

import "fmt"

type DBConfig struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	DBName       string `mapstructure:"dbname" json:"dbname"`
	SSLMode      string `mapstructure:"ssl_mode" json:"ssl_mode"`
	ConnLifeTime int    `mapstructure:"conn_life_time" json:"conn_life_time" default:"300"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" default:"10"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" default:"80"`
	LogLevel     int    `mapstructure:"log_level" json:"log_level" default:"1"`
}

func DBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		G.DB.User,
		G.DB.Password,
		G.DB.Host,
		G.DB.Port,
		G.DB.DBName,
		G.DB.SSLMode,
	)
}

func PConn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		G.DB.Host,
		G.DB.Port,
		G.DB.User,
		G.DB.Password,
		G.DB.DBName,
		G.DB.SSLMode,
	)
}
