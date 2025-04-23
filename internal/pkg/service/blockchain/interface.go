package blockchain

type IService interface {
	Vote(hexAddress string, index uint64) (string, error)
	AddCandidates(candidates []string) (string, error)
	EndVote() (string, error)
	StartVote() (string, error)
	GiveRightToVote(address string) (string, error)
	Admin() (string, error)
	GetResults() ([]*Result, error)
	WinningCandidate() (string, error)
}
