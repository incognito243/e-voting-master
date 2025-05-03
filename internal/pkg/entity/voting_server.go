package entity

type VotingServer struct {
	BaseID
	BaseCreatedUpdated
	AdminId               string `json:"admin_id" gorm:"column:admin_id"`
	NumberOfCandidates    int64  `json:"number_of_candidates" gorm:"column:number_of_candidates"`
	MaximumNumberOfVoters int64  `json:"maximum_number_of_voters" gorm:"column:maximum_number_of_voters"`
	ServerName            string `json:"server_name" gorm:"column:server_name"`
	ExpTime               int64  `json:"exp_time" gorm:"column:exp_time"`

	ContractAddress string `json:"contract_address" gorm:"column:contract_address"`
	ServerId        string `json:"server_id" gorm:"column:server_id"`
	OpenedVote      bool   `json:"opened_vote" gorm:"column:opened_vote"`
	Results         string `json:"results" gorm:"column:results"`
	Active          bool   `json:"active" gorm:"column:active"`
}

func (VotingServer) TableName() string {
	return "voting_servers"
}
