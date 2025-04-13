package cmd

import (
	"forum/internal/server"

	"github.com/spf13/cobra"
)

func init() {
	serverCmd.Flags().StringP("port", "p", "4000", "port to run the server")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start the server",
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value.String()
		server.Run(port)
	},
}
