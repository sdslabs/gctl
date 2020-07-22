package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gctl",
	Short: "Your Cloud in a Binary",
	Long:  "Gasper is an intelligent Platform as a Service (PaaS) used for deploying and managing applications and databases in any cloud topology.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gasper is an intelligent Platform as a Service (PaaS) used for deploying and managing applications and databases in any cloud topology.")
	},
}

//Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
