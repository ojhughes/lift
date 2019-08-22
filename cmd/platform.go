package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Platform commands",
	Long:  `Commands related to platform operations`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(platformCmd)
}
