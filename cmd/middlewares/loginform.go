package middlewares

import (
	"bufio"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
)

func LoginForm() openapi.Login {
	scanner := bufio.NewScanner(os.Stdin)
	var loginCreds openapi.Login
	for !ValidateEmail(loginCreds.Email) {
		fmt.Printf("Email: ")
		scanner.Scan()
		loginCreds.Email = scanner.Text()
		if !ValidateEmail(loginCreds.Email) {
			fmt.Println("Please enter a valid email id")
		}
	}
	for loginCreds.Password == "" {
		fmt.Printf("Password: ")
		maskedPasswd, _ := gopass.GetPasswdMasked()
		loginCreds.Password = string(maskedPasswd)
		if loginCreds.Password == "" {
			fmt.Println("Password cannot be empty")
		}
	}
	return loginCreds
}
