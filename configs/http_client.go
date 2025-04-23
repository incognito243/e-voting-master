package configs

type HttpClientConfig struct {
	RetryCount              int `default:"3" json:"retry_count" mapstructure:"retry_count"`
	RetryWaitTimeSeconds    int `default:"5" json:"retry_wait_time_seconds" mapstructure:"retry_wait_time_seconds"`
	RetryMaxWaitTimeSeconds int `default:"30" json:"retry_max_wait_time_seconds" mapstructure:"retry_max_wait_time_seconds"`
}

type ClientConfig struct {
	BaseURL string `json:"base_url" mapstructure:"base_url"`
	APIKey  string `json:"api_key" mapstructure:"api_key"`
}
