package cmd

import (
	"github.com/davido912/cshex-utils/internal/cli"
	"github.com/davido912/cshex-utils/internal/log"
	"github.com/spf13/cobra"
)

func Run() {
	log.InitLogging()
	if err := cli.NewRootCmd(func(cmd *cobra.Command, args []string) error {
		return nil
	}).Execute(); err != nil {
		panic(err)
	}
}
