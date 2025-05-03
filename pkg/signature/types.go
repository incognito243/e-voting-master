package signature

type Voting struct {
	Username string
	ServerId string
}

type Admin struct {
	AdminId    string
	ServerName string
}

type VotingCandidate struct {
	CandidateIndex int64
}

type ApproveUsers struct {
	AdminId string
}

type CreateServer struct {
	AdminId    string
	ServerName string
}
