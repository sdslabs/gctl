package middlewares

import (
	"bufio"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
)

//DbForm takes input for openapi.Database
func DbForm() (string, openapi.Database) {
	var database openapi.Database
	var dbtype string
	scanner := bufio.NewScanner(os.Stdin)
	for !ValidateName(database.Name) {
		fmt.Printf("*Database Name: ")
		scanner.Scan()
		database.Name = scanner.Text()
		if !ValidateName(database.Name) {
			fmt.Println("Database Name should have only alphanumeric characters, lowercase alphabets and should be of length 3-40.")
		}
	}
	for database.Password == "" {
		fmt.Printf("*Application Password: ")
		maskedPasswd, _ := gopass.GetPasswdMasked()
		database.Password = string(maskedPasswd)
		if database.Password == "" {
			fmt.Println("This field is required. Please enter a valid password.")
		}
	}
	for !ValidateDbType(dbtype) {
		fmt.Printf("Database Type: ")
		scanner.Scan()
		dbtype = scanner.Text()
		if !ValidateDbType(dbtype) {
			fmt.Println("This is field is required. supported database types are mysql, mongodb, postgresql and redis")
		}
	}
	return dbtype, database
}
