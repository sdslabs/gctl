package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(registerCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gasper",
	//TODO RUN
}
