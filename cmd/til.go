package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	tilCmd = &cobra.Command{
		Use:   "til",
		Short: "#til generator for huyng.dev",
		Run:   tilcmd,
	}
)

func tilcmd(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
