package cmd

import (
	"github.com/spf13/cobra"
)

// some global var
const (
	ws      = "/Users/hnh/src"
	blogDir = "github.com/huynguyenh/huyng.dev"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hnh",
		Short: "hnh's mighty warhammer",
	}
)

func init() {
	rootCmd.AddCommand(configureTil())
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
