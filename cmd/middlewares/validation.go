package middlewares

import (
	"net/url"
	"regexp"
)

//ValidateEmail validates the email
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

//ValidateName validates the username
func ValidateName(username string) bool {
	if len(username) < 3 || len(username) > 40 {
		return false
	}
	re := regexp.MustCompile("^[a-z0-9_]*$")
	return re.MatchString(username)

}

//ValidateURL validates and sanitizes the url
func ValidateURL(u string) bool {
	if u == "" {
		return false
	}
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	return true
}

//ValidatePort validates port number
func ValidatePort(port int64) bool {
	re := regexp.MustCompile("^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$")
	return re.MatchString(string(port))
}

//ValidateLanguageApp validates language for a new app
func ValidateLanguageApp(language string) bool {
	validLanguages := [8]string{"static", "php", "nodejs", "python2", "python3", "golang", "ruby", "rust"}
	for _, a := range validLanguages {
		if a == language {
			return true
		}
	}
	return false
}

//ValidateDbType validates database type
func ValidateDbType(dbtype string) bool {
	validTypes := [4]string{"mysql", "mongodb", "postgresql", "redis"}
	for _, a := range validTypes {
		if a == dbtype {
			return true
		}
	}
	return false
}
