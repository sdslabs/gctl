package cmd

import (
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

var cfg = openapi.NewConfiguration()
var client = openapi.NewAPIClient(cfg)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(userCmd)
	createCmd.AddCommand(appmakerCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create database or app",
}
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gasper",
	//TODO RUN
}
