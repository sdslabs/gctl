package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to get a bearer token ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Email: ")
		var email string
		fmt.Scanln(&email)
		fmt.Printf("Password: ")
		maskedPasswd, _ := gopass.GetPasswdMasked()
		l := openapi.Login{Email: email, Password: string(maskedPasswd)}
		res, _, _ := client.AuthApi.Login(context.Background(), l)
		fmt.Println("Token: ", res.Token, "\n", "Expires at: ", res.Expire)

	},
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a user",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Username: ")
		scanner.Scan()
		username := scanner.Text()
		fmt.Printf("Email: ")
		scanner.Scan()
		email := scanner.Text()
		fmt.Printf("Password: ")
		maskedPasswd, err := gopass.GetPasswdMasked()
		if err != nil {
			fmt.Println(err)
		}
		r := openapi.User{Username: username, Email: email, Password: string(maskedPasswd)}
		res, _, _ := client.AuthApi.Register(context.Background(), r)
		fmt.Println(res.Message)
	},
}

var fetchUserCmd = &cobra.Command{
	Use:   "user",
	Short: "View user details",
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		res, _, err := client.UserApi.FetchUser(auth)
		if err != nil {
			cmd.Print(err)
			os.Exit(1)
		}
		cmd.Print("Username: " + res.Username + "\n" + "Email: " + res.Email + "\n")
	},
}

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh JWT token using existing token",
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		fmt.Print(client.AuthApi.Refresh(auth, token))
	},
}
