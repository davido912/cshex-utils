package cli

import (
	"fmt"

	"github.com/davido912/cshex-utils/internal/teleport"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewDbCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "Database utilities",
		Long:  "Database utilities",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		SilenceUsage: false,
	}
	cmd.AddCommand(NewRefreshConnectionsCommand())
	return cmd
}

func NewRefreshConnectionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "login",
		Short:        "Refresh database logins",
		Long:         "Refresh database logins (can take multiple database names or even substrings)",
		RunE:         dbLogin,
		SilenceUsage: false,
		Example:      "cshex db login <db-name> <db-name>",
	}

	return cmd
}

func dbLogin(cmd *cobra.Command, args []string) error {
	tsh := teleport.New()

	if len(args) == 0 {
		return fmt.Errorf("must pass database names/substrings")
	}
	for _, tshEnv := range Envs {
		tsh.SetEnv(tshEnv)
		log.Info().Msgf("Refreshing connections for %s environment", tshEnv)
		dbListResult := tsh.ListDb(args...)
		for _, db := range dbListResult {
			tsh.DbConnect(db.Metadata.Name)
		}
	}

	return nil
}
