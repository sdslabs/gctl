package cmd

import (
	"bufio"
	"context"
	"os"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
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
		scanner := bufio.NewScanner(os.Stdin)
		cmd.Printf("Email: ")
		scanner.Scan()
		email := scanner.Text()
		cmd.Printf("Password: ")
		maskedPasswd, _ := gopass.GetPasswdMasked()
		l := openapi.Login{Email: email, Password: string(maskedPasswd)}
		res, _, err := client.AuthApi.Login(context.Background(), l)
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
		scanner := bufio.NewScanner(os.Stdin)
		cmd.Printf("Username: ")
		scanner.Scan()
		username := scanner.Text()
		cmd.Printf("Email: ")
		scanner.Scan()
		email := scanner.Text()
		cmd.Printf("Password: ")
		maskedPasswd, _ := gopass.GetPasswdMasked()
		r := openapi.User{Username: username, Email: email, Password: string(maskedPasswd)}
		res, _, err := client.AuthApi.Register(context.Background(), r)
		if res.Success {
			cmd.Println(res.Message)
		} else {
			cmd.Println(err)
		}
	},
}

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh JWT token using existing token",
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		cmd.Print(client.AuthApi.Refresh(auth, token))
	},
}
