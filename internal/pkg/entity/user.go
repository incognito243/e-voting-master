package entity

type User struct {
	BaseID
	BaseCreatedUpdated
	UserIdCore    string `json:"user_id_core" gorm:"column:user_id_core"`
	Username      string `json:"username" gorm:"column:username"`
	CitizenID     string `json:"citizen_id" gorm:"column:citizen_id"`
	CitizenName   string `json:"citizen_name" gorm:"column:citizen_name"`
	Verified      bool   `json:"verified" gorm:"column:verified"`
	Email         string `json:"email" gorm:"column:email"`
	PublicKey     string `json:"public_key" gorm:"column:public_key"`
	AptosAddress  string `json:"aptos_address" gorm:"column:aptos_address"`
	EncryptedHash string `json:"encrypted_hash" gorm:"column:encrypted_hash"`
	Nonce         string `json:"nonce" gorm:"column:nonce"`
	IsAdmin       bool   `json:"is_admin" gorm:"column:is_admin"`
}

func (User) TableName() string {
	return "users"
}
