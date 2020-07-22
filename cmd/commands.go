package cmd

import (
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

var cfg = openapi.NewConfiguration()
var client = openapi.NewAPIClient(cfg)
var gctltoken string

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(updateCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create database or app",
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete database or app",
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch user details, databases and apps",
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an app or user password",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gasper",
	//TODO RUN
}
