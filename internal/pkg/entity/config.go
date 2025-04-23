package entity

type Configs struct {
	Key   string `json:"key" mapstructure:"key"`
	Value string `json:"value" mapstructure:"value"`
}

func (Configs) TableName() string {
	return "configs"
}
