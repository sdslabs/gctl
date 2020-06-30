package cmd

import (
	"context"

	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(refreshCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to get a bearer token ",
	Run: func(cmd *cobra.Command, args []string) {
		loginCreds := middlewares.LoginForm()
		res, _, err := client.AuthApi.Login(context.Background(), loginCreds)
		if res.Token != "" {
			cmd.Println("Token: ", res.Token, "\n", "Expires at: ", res.Expire)
		} else {
			cmd.Println(err)
		}
	},
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a user",
	Run: func(cmd *cobra.Command, args []string) {
		registerCreds := middlewares.RegisterForm()
		res, _, err := client.AuthApi.Register(context.Background(), registerCreds)
		if res.Success {
			cmd.Println(res.Message)
		} else {
			cmd.Println(err)
		}
	},
}

var refreshCmd = &cobra.Command{
	Use:   "refresh [BEARER_TOKEN]",
	Short: "Refresh JWT token using existing token",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//TODO
	},
}
