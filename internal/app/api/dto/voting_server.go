package dto

import "e-voting-mater/internal/pkg/service/votingserver"

type CreateVotingServerRequest struct {
	AdminId               string                    `json:"admin_id"`
	NumberOfCandidates    int64                     `json:"number_of_candidates"`
	MaximumNumberOfVoters int64                     `json:"maximum_number_of_voters"`
	ServerName            string                    `json:"server_name"`
	Candidates            []*votingserver.Candidate `json:"candidates"`
	ContractAddress       string                    `json:"contract_address"`
	ExpTime               int64                     `json:"exp_time"`
	SignatureHex          string                    `json:"signature_hex"`
}

type GetVotingServerByIdRequest struct {
	ServerId string `json:"server_id" form:"server_id" binding:"required"`
}

type OpenVoteRequest struct {
	AdminId      string `json:"admin_id"`
	ServerId     string `json:"server_id"`
	ServerName   string `json:"server_name"`
	SignatureHex string `json:"signature_hex"`
}

type OpenVoteResponse struct {
	Candidate *votingserver.Candidate `json:"candidate"`
	Results   map[string]int64        `json:"results"`
}

type PublishVoteRequest struct {
	ServerId string `json:"server_id" form:"server_id" binding:"required"`
}

type ActiveVotingServerRequest struct {
	ServerName   string `json:"server_name"`
	AdminId      string `json:"admin_id"`
	SignatureHex string `json:"signature_hex"`
}
