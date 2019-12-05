package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// root cmd
var rootCmd = &cobra.Command{
	Use:   "wings-of-liberty",
	Short: "wings-of-liberty is a tool for scientific internet",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

// Execute running app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
