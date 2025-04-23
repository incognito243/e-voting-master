package api

import (
	"errors"
	"fmt"

	"e-voting-mater/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	"e-voting-mater/configs"
	"e-voting-mater/internal/app/api/component"
	"e-voting-mater/internal/app/api/route"
)

var Command = &cobra.Command{
	Use:   "api",
	Short: "Run api server",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := runMigrations(); err != nil {
			return err
		}
		return RunE(cmd, args)
	},
}

func RunE(_ *cobra.Command, _ []string) error {
	err := component.InitComponents()
	if err != nil {
		return err
	}
	s := api.Init(configs.G.Server.Mode, func(engine *gin.Engine) {
		route.Register(engine)
	})
	return s.Run(configs.G.Server.APIBindAddress)
}

func runMigrations() error {
	m, err := migrate.New(
		"file://migrations",
		configs.DBUrl(),
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	fmt.Println("Migration successful")
	return nil
}
