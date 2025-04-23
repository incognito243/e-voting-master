package configs

type VotingConfig struct {
	PrivateKeyName  string `mapstructure:"private_key_name" json:"private_key_name"`
	ContractAddress string `mapstructure:"contract_address" json:"contract_address"`
}
