package votingserver

type ExtendedVotingServer struct {
	VotingServer *VotingServer `json:"voting_server"`
	Candidates   []*Candidate  `json:"candidates"`
}

type Candidate struct {
	Index         int64  `json:"index"`
	CandidateName string `json:"candidate_name"`
	CitizenID     string `json:"citizen_id"`
	AvatarURL     string `json:"avatar_url"`
}

type VotingServer struct {
	NumberOfCandidates    int64  `json:"number_of_candidates"`
	MaximumNumberOfVoters int64  `json:"maximum_number_of_voters"`
	ServerName            string `json:"server_name" `
	ServerId              string `json:"server_id" `
	OpenedVote            bool   `json:"opened_vote" `
	Results               string `json:"results" `
	Active                bool   `json:"active" `
	ExpTime               int64  `json:"exp_time" `
}
