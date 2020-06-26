package middlewares

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
)

func AppForm() (string, string, openapi.Application) {
	var token, language string
	var application openapi.Application
	scanner := bufio.NewScanner(os.Stdin)
	for token == "" {
		fmt.Printf("*Token:")
		scanner.Scan()
		token = scanner.Text()
		if token == "" {
			fmt.Println("Token is required!")
		}
	}
	for !ValidateName(application.Name) {
		fmt.Printf("*App Name: ")
		scanner.Scan()
		application.Name = scanner.Text()
		if !ValidateName(application.Name) {
			fmt.Println("App Name should have only alphanumeric characters, lowercase alphabets and should be of length 3-40.")
		}
	}
	for !ValidateLanguageApp(language) {
		fmt.Printf("*Language: ")
		scanner.Scan()
		language = scanner.Text()
		if !ValidateLanguageApp(language) {
			fmt.Println("This field is required. Supported languages are static, php, nodejs, python2, python3, golang, ruby and rust")
		}
	}
	for application.Password == "" {
		fmt.Printf("*Application Password: ")
		maskedPasswd, _ := gopass.GetPasswdMasked()
		application.Password = string(maskedPasswd)
		if application.Password == "" {
			fmt.Println("This field is required. Please enter a valid password.")
		}
	}
	for !ValidateURL(application.Git.RepoUrl) {
		fmt.Printf("*Git URL: ")
		scanner.Scan()
		application.Git.RepoUrl = scanner.Text()
		if !ValidateURL(application.Git.RepoUrl) {
			fmt.Println("Please enter a valid URL.")
		}
	}
	fmt.Printf("Is this repo private? [yes/no]: ")
	scanner.Scan()
	if scanner.Text() == "yes" {
		fmt.Printf("*Git Access Token: ")
		scanner.Scan()
		application.Git.AccessToken = scanner.Text()
	}
	fmt.Printf("Branch: ")
	scanner.Scan()
	application.Git.Branch = scanner.Text()
	for application.Context.Index == "" {
		fmt.Printf("*Index: ")
		scanner.Scan()
		application.Context.Index = scanner.Text()
		if application.Context.Index == "" {
			fmt.Println("Index cannot be empty.")
		}
	}
	for ValidatePort(application.Context.Port) {
		fmt.Printf("Port: ")
		fmt.Scanln(&application.Context.Port)
		if application.Context.Port == 0 {
			break
		} else if ValidatePort(application.Context.Port) {
			fmt.Println("Please enter valid port number.")
		}
	}
	fmt.Printf("Does this repo contain Gasperfile.txt? [yes/no]: ")
	scanner.Scan()
	if scanner.Text() == "no" {
		fmt.Printf("Build Commands: ")
		scanner.Scan()
		application.Context.Build = strings.Split(scanner.Text(), ",")
		fmt.Printf("Run Commands: ")
		scanner.Scan()
		application.Context.Run = strings.Split(scanner.Text(), ",")
	} else if scanner.Text() == "yes" {
		application.Context.RcFile = true
	}
	fmt.Printf("Environment Variables(key:value): ")
	scanner.Scan()
	if scanner.Text() != "" {
		m := make(map[string]string)
		vars := strings.Split(scanner.Text(), ",")
		for v := 0; v < len(vars); v++ {
			key := strings.Split(vars[v], ":")[0]
			value := strings.Split(vars[v], ":")[1]
			m[key] = value
		}
		application.Env = m
	}
	return token, language, application
}
