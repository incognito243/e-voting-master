package blockchain

type IService interface {
	Vote(hexAddress string, index uint64, contract string) (string, error)
	AddCandidates(candidates []string, contract string) (string, error)
	EndVote(contract string) (string, error)
	StartVote(contract string) (string, error)
	GiveRightToVote(address string, contract string) (string, error)
	Admin(contract string) (string, error)
	GetResults(contract string) ([]*Result, error)
	WinningCandidate(contract string) (string, error)
}
