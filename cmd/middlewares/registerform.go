package middlewares

import (
	"bufio"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
)

func RegisterForm() openapi.User {
	scanner := bufio.NewScanner(os.Stdin)
	var registerCreds openapi.User
	for !ValidateName(registerCreds.Username) {
		fmt.Printf("Username: ")
		scanner.Scan()
		registerCreds.Username = scanner.Text()
		if !ValidateName(registerCreds.Username) {
			fmt.Println("Username should have only alphanumeric characters, lowercase alphabets and should be of length 3-40")
		}
	}
	for !ValidateEmail(registerCreds.Email) {
		fmt.Printf("Email: ")
		scanner.Scan()
		registerCreds.Email = scanner.Text()
		if !ValidateEmail(registerCreds.Email) {
			fmt.Println("Please enter a valid email id")
		}
	}
	for registerCreds.Password == "" {
		fmt.Printf("Password: ")
		maskedPasswd, _ := gopass.GetPasswdMasked()
		registerCreds.Password = string(maskedPasswd)
		if registerCreds.Password == "" {
			fmt.Println("Password cannot be empty")
		}
	}
	return registerCreds
}
