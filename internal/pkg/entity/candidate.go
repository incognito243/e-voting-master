package entity

type Candidate struct {
	BaseID
	BaseCreatedUpdated
	CandidateName string `json:"candidate_name" gorm:"column:candidate_name"`
	CitizenID     string `json:"citizen_id" gorm:"column:citizen_id"`
	AvatarURL     string `json:"avatar_url" gorm:"column:avatar_url"`

	ServerId       string `json:"server_id" gorm:"column:server_id"`
	CandidateIndex int64  `json:"candidate_index" gorm:"column:candidate_index"`
}
