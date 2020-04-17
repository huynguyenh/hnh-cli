package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hnh",
		Short: "hnh's mighty warhammer",
	}
)

func init() {
	rootCmd.AddCommand(tilCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
