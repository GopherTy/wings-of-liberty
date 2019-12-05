package cmd

import (
	"wings-of-liberty/fly/client"

	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "run wings-of-liberty app client ",
	Run: func(cmd *cobra.Command, args []string) {
		client.Run()
	},
}

// regist client command
func init() {
	rootCmd.AddCommand(clientCmd)
}
