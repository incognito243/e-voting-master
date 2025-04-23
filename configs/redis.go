package configs

type RedisConfig struct {
	InitAddress  []string `mapstructure:"init_address" json:"initAddress"`
	SelectDB     int      `mapstructure:"select_db" json:"selectDb"`
	MasterName   string   `mapstructure:"master_name" json:"masterName"`
	Password     string   `mapstructure:"password" json:"password"`
	DisableCache bool     `mapstructure:"disable_cache" json:"disableCache"`
}
