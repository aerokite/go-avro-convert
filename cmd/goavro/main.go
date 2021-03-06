package main

import (
	"os"

	"github.com/spf13/cobra"
)

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goavro",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(newCmdGovro())
	return cmd
}

func main() {
	cmd := newCmdRoot()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
