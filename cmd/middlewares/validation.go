package middlewares

import (
	"net/url"
	"regexp"
	"strings"
)

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
	if port > 0 && port <= 65535 {
		return true
	}
	return false
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

func ValidateEnvVars(vars []string) bool {
	for v := 0; v < len(vars); v++ {
		a := strings.Split(vars[v], ":")
		if len(a) != 2 {
			return false
		}
	}
	return true
}
