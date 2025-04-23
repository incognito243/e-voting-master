package votingcore

import (
	"context"
	"errors"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	httpClient *resty.Client
	baseUrl    string
	apiKey     string
}

func NewClient(httpClient *resty.Client, baseUrl string, apiKey string) IClient {
	return &Client{
		httpClient: httpClient,
		baseUrl:    baseUrl,
		apiKey:     apiKey,
	}
}

func (c *Client) CreateVotingServer(ctx context.Context, serverConfig *CreateVotingServerRequest) (*CreateVotingServerResponse, error) {
	var response *CreateVotingServerResponse

	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("x-api-key", c.apiKey).
		SetBody(serverConfig).
		SetResult(&response).
		Post(c.baseUrl + CreateVotingServerPath)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}

	return response, nil
}

func (c *Client) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	var response *CreateUserResponse

	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("x-api-key", c.apiKey).
		SetBody(req).
		SetResult(&response).
		Post(c.baseUrl + CreateUserPath)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}

	return response, nil
}

func (c *Client) Vote(ctx context.Context, req *VoteRequest) error {
	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("x-api-key", c.apiKey).
		SetBody(req).
		Post(c.baseUrl + VotePath)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}

	return nil
}

func (c *Client) EndVote(ctx context.Context, req *EndVoteRequest) (*EndVoteResponse, error) {
	var response *EndVoteResponse

	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("x-api-key", c.apiKey).
		SetBody(req).
		SetResult(&response).
		Post(c.baseUrl + OpenVotePath)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}

	return response, nil
}

func (c *Client) PublishResult(ctx context.Context, serverId string) (*PublishResultResponse, error) {
	var response *PublishResultResponse

	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("x-api-key", c.apiKey).
		SetPathParam("server_id", serverId).
		SetResult(&response).
		Get(c.baseUrl + PublishResultPath)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}

	return response, nil
}
