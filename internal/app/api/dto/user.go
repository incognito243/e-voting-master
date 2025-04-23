package dto

import "e-voting-mater/internal/pkg/service/user"

type CreateUserRequest struct {
	Username    string `json:"username"`
	CitizenID   string `json:"citizen_id"`
	CitizenName string `json:"citizen_name"`
	Email       string `json:"email"`
	PublicKey   string `json:"public_key"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"is_admin"`
}

type LoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PersonalCode string `json:"personal_code"`
}

type LoginResponse struct {
	Token string         `json:"token"`
	User  *user.InfoUser `json:"user"`
}

type GetUserByUsernameRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
}

type GetUserByCitizenIdRequest struct {
	CitizenId string `json:"citizen_id" form:"citizen_id" binding:"required"`
}

type VotingRequest struct {
	Username     string `json:"username"`
	ServerId     string `json:"server_id"`
	VotingHex    string `json:"voting_hex"`
	SignatureHex string `json:"signature_hex"`
}

type VerifyUsers struct {
	Usernames    []string `json:"usernames"`
	AdminId      string   `json:"admin_id"`
	SignatureHex string   `json:"signature_hex"`
}
