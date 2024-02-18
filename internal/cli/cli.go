package cli

import (
	"github.com/spf13/cobra"
)

const (
	appName = "cshex"
	version = "0.0.1"
)

type CobraRunFunc func(cmd *cobra.Command, args []string) error

func NewRootCmd(entrypointFunc CobraRunFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   appName,
		Short: "CLI implementation for CashEx utilities",
		Long:  "CLI implementation for CashEx utilities",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Usage()
			}
		},
		RunE:         entrypointFunc,
		SilenceUsage: true,
	}
	cmd.AddCommand(NewIntjCommand())
	cmd.AddCommand(NewDbCommand())
	return cmd
}
