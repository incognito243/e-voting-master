package user

type InfoUser struct {
	Username      string `json:"username"`
	CitizenID     string `json:"citizen_id"`
	CitizenName   string `json:"citizen_name"`
	Verified      bool   `json:"verified"`
	Email         string `json:"email"`
	CompressedKey string `json:"compressed_key"`
	IsAdmin       bool   `json:"is_admin"`
}
