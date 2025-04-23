package configs

type AptosConfig struct {
	Name       string `mapstructure:"name" json:"name"`
	ChainId    uint8  `mapstructure:"chain_id" json:"chain_id"`
	NodeUrl    string `mapstructure:"node_url" json:"node_url"`
	IndexerUrl string `mapstructure:"indexer_url" json:"indexer_url"`
}
