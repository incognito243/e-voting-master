package signature

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

const (
	verifyMessageTemplate      = `Vote From: {{.Username}} is voting in server {{.ServerId}}`
	votingMessageTemplate      = `Vote For: user is voting for {{.CandidateIndex}}`
	adminVerifyMessageTemplate = `Admin: {{.AdminId}} is acting in server {{.ServerId}}`
	approveUserTemplate        = `Approve User: {{.AdminId}} is approving users`
)

func BuildVerifyMessage(username, serverId string) (string, error) {
	tmpl, err := template.New("Vote").Parse(verifyMessageTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	data := Voting{
		Username: username,
		ServerId: serverId,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

func BuildAdminVerifyMessage(adminId, serverId string) (string, error) {
	tmpl, err := template.New("Vote").Parse(adminVerifyMessageTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	data := Admin{
		AdminId:  adminId,
		ServerId: serverId,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

func BuildApproveUser(adminId string) (string, error) {
	tmpl, err := template.New("Vote").Parse(approveUserTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	data := ApproveUsers{
		AdminId: adminId,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

func FindCandidateIndex(fromAddress, msgHex string, numberOfCandidates int64) (int64, error) {
	if numberOfCandidates <= 0 {
		return -1, fmt.Errorf("invalid number of candidates: %d", numberOfCandidates)
	}

	for i := int64(0); i < numberOfCandidates; i++ {
		tmpl, err := template.New("Vote").Parse(votingMessageTemplate)
		if err != nil {
			return -1, fmt.Errorf("failed to parse template: %v", err)
		}
		data := VotingCandidate{
			CandidateIndex: i,
		}
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return -1, fmt.Errorf("failed to execute template: %v", err)
		}
		err = VerifySignature(fromAddress, msgHex, buf.String())
		if err != nil {
			if errors.Is(err, ErrorFailedToVerifyMsg) {
				continue
			}
			return -1, err
		}
		return i, nil
	}
	return -1, fmt.Errorf("failed to find candidate index")
}
