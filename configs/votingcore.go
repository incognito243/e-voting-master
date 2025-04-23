package configs

type VotingCoreConfig struct {
	BaseUrl string `mapstructure:"base_url" json:"base_url"`
	ApiKey  string `mapstructure:"api_key" json:"api_key"`
}
