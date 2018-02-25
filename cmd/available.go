package cmd

import (
	"encoding/json"
	"os"

	"github.com/gobuffalo/buffalo/plugins"
	"github.com/spf13/cobra"
)

// availableCmd represents the available command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "a list of available buffalo plugins",
	Run: func(cmd *cobra.Command, args []string) {
		p := plugins.Commands{
			{Name: swaggerCmd.Use, BuffaloCommand: "generate", Description: swaggerCmd.Short, Aliases: []string{"s"}},
		}
		json.NewEncoder(os.Stdout).Encode(p)
	},
}

func init() {
	rootCmd.AddCommand(availableCmd)
}
