package cmd

import (
	"github.com/spf13/cobra"
)

var createClusterCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"cc"},
	Run: func(cmd *cobra.Command, args []string) {

	},
}
