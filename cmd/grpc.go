package cmd

import (
	"wings-of-liberty/grpc"

	"github.com/spf13/cobra"
)

var grpcServiceCmd = &cobra.Command{
	Use:   "grpc",
	Short: "wings-of-liberty app grpc service for get encryption array",
	Run: func(cmd *cobra.Command, args []string) {
		grpc.Run()
	},
}

func init() {
	rootCmd.AddCommand(grpcServiceCmd)
}
