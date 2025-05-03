package entity

type Tracking struct {
	BaseID
	BaseCreatedUpdated
	Username string `json:"username" gorm:"column:username"`
	ServerId string `json:"server_id" gorm:"column:server_id"`
}
