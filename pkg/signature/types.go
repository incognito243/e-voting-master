package signature

type Voting struct {
	Username string
	ServerId string
}

type Admin struct {
	AdminId  string
	ServerId string
}

type VotingCandidate struct {
	CandidateIndex int64
}

type ApproveUsers struct {
	AdminId string
}
