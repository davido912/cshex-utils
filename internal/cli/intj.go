package cli

import (
	"strings"

	"github.com/davido912/cshex-utils/internal/ide/intellij"
	"github.com/davido912/cshex-utils/internal/teleport"
	"github.com/spf13/cobra"
)

var (
	ProjecRootPath string
	Envs           []string
)

// Flag names
const (
	projectPathFlagName = "path"
	envsFlagName        = "envs"
)

func NewIntjCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "intj",
		Short: "Intellij utilities",
		Long:  "Intellij utilities",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		SilenceUsage: false,
	}

	cmd.AddCommand(NewIntjInitDbCommand())
	return cmd
}

func NewIntjInitDbCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "initdb",
		Short:        "Create datasource db connections",
		Long:         "Create datasource db connections (takes multiple database names or substrings)",
		RunE:         createDatasources,
		SilenceUsage: false,
		Example:      "cshex db intj -p <project-root-path> initdb <db-name>",
	}

	cmd.Flags().StringVarP(&ProjecRootPath, projectPathFlagName, "p", "", "Intellij project root path")
	cmd.Flags().StringSliceVarP(&Envs, envsFlagName, "e", []string{"production", "staging", "development"}, "Environments to create datasources for")
	cmd.MarkFlagRequired(projectPathFlagName)

	return cmd
}

func createDatasources(cmd *cobra.Command, args []string) error {
	tsh := teleport.New()
	intj := intellij.NewManager(tsh)
	projectRoot := strings.Join([]string{ProjecRootPath, ".idea"}, "/")
	intj.CreateDatasources(Envs, projectRoot, args...)
	return nil
}
