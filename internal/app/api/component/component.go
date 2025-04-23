package component

import (
	"e-voting-mater/configs"
	"e-voting-mater/internal/pkg/client/blockchain"
	"e-voting-mater/internal/pkg/client/votingcore"
	"e-voting-mater/internal/pkg/repository"
	"e-voting-mater/internal/pkg/service/auth"
	blockchainSrv "e-voting-mater/internal/pkg/service/blockchain"
	"e-voting-mater/internal/pkg/service/pw"
	"e-voting-mater/internal/pkg/service/user"
	"e-voting-mater/internal/pkg/service/votingserver"
	"e-voting-mater/pkg/db"
	"e-voting-mater/pkg/http"
	"e-voting-mater/pkg/logger"

	"github.com/aptos-labs/aptos-go-sdk"
)

// InitComponents init all components (singleton) using for dependency injection
func InitComponents() error {
	if err := logger.InitLogger(configs.G.Log.Level); err != nil {
		return err
	}
	// Infra
	postgresDb, err := db.InitPostgres()
	if err != nil {
		return err
	}

	//redisDb, err := db.InitRedis()
	//if err != nil {
	//	return err
	//}

	// Repo
	userRepo := repository.NewUserRepo(postgresDb)
	candidateRepo := repository.NewCandidateRepo(postgresDb)
	votingServer := repository.NewVotingServerRepo(postgresDb)

	// client
	httpClient := http.NewClient(
		configs.G.HttpClient.RetryCount,
		configs.G.HttpClient.RetryWaitTimeSeconds,
		configs.G.HttpClient.RetryMaxWaitTimeSeconds,
	)
	aptosClient, err := aptos.NewClient(aptos.NetworkConfig{
		Name:       configs.G.Aptos.Name,
		ChainId:    configs.G.Aptos.ChainId,
		NodeUrl:    configs.G.Aptos.NodeUrl,
		IndexerUrl: configs.G.Aptos.IndexerUrl,
	})
	if err != nil {
		return err
	}
	coreClient := votingcore.NewClient(httpClient, configs.G.VotingCore.BaseUrl, configs.G.VotingCore.ApiKey)
	blockChainClient, err := blockchain.NewClient(configs.G.Voting, aptosClient)
	if err != nil {
		return err
	}

	// Service
	auth.NewJWTService(configs.G.Jwt.SecretKey, configs.G.Jwt.Expire)
	_, err = pw.NewService(configs.G.PasswordKey)
	if err != nil {
		return err
	}
	blockchainSrv.NewService(blockChainClient, configs.G.Voting.ContractAddress)
	user.NewService(coreClient, userRepo, candidateRepo, votingServer)
	votingserver.NewService(votingServer, candidateRepo, userRepo, coreClient)
	return nil
}
