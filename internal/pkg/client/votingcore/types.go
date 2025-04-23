package votingcore

import "math/big"

type CreateVotingServerRequest struct {
	NumberOfCandidates    int64  `json:"number_of_candidates"`
	MaximumNumberOfVoters int64  `json:"maximum_number_of_voters"`
	ServerName            string `json:"server_name"`
}

type CreateVotingServerResponse struct {
	ServerId string `json:"server_id"`
}

type CreateUserRequest struct {
	Username string `json:"user_name"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}

type VoteRequest struct {
	UserID      string `json:"user_id"`
	ServerID    string `json:"server_id"`
	CandidateID int64  `json:"candidate_id"`
}

type EndVoteRequest struct {
	ServerId string `json:"server_id"`
}

type EndVoteResponse struct {
	Results []int64 `json:"results"`
}

type ECCPoint struct {
	X      *big.Int `json:"_x"`
	Y      *big.Int `json:"_y"`
	Origin bool     `json:"_origin"`
}

type IntPair struct {
	X *big.Int `json:"x"`
	Y *big.Int `json:"y"`
}

type EccPointPair struct {
	First  ECCPoint `json:"first"`
	Second ECCPoint `json:"second"`
}

type ProofOfWork struct {
	A []ECCPoint `json:"A"`
	B []ECCPoint `json:"B"`
	U []*big.Int `json:"u"`
	W []*big.Int `json:"w"`
}

type PublishResultResponse struct {
	VoterPublicKey     []IntPair      `json:"voter_public_key"`
	VoterVote          []EccPointPair `json:"voter_vote"`
	VoterSignedMessage []EccPointPair `json:"voter_signed_message"`
	VoterProveOfWork   []ProofOfWork  `json:"voter_prove_of_work"`
	EncryptedPackage   []EccPointPair `json:"encrypted_package"`
	DecryptedPackage   []ECCPoint     `json:"decrypted_package"`
	ResultPackage      [][]*big.Int   `json:"result_package"`
	Results            []*big.Int     `json:"results"`
}
