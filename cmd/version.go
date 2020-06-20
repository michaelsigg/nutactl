package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "DEV"

func newVersionCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "version",
		Short:                 "Print version information",
		Args:                  cobra.NoArgs,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVersion),
	}
	return cmd
}

func runVersion(cli *CLI, cmd *cobra.Command, args []string) error {
	fmt.Printf("nutactl %s\n", Version)
	return nil
}
