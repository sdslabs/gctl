package middlewares

import (
	"github.com/go-git/go-git/v5"
)

func GitInit(path string) error {
	_, err := git.PlainInit(path, false)
	return err
}
