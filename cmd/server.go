package cmd

import (
	"wings-of-liberty/fly/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run wings-of-liberty app server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
