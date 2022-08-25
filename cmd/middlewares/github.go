package middlewares

import (
	// "fmt"
	_ "io/ioutil"
	"os"
	_ "path/filepath"
	"time"

	"github.com/apex/log"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const defaultRemoteName = "origin"

func GitInit(directoryPath string) (*git.Repository, error) {
	var (
		err error
	)
	_, err = os.Stat(directoryPath)
	if err != nil {
		return nil, err
	}
	repository, err := git.PlainInit(directoryPath, false)
	return repository, err
}

// Commit creates a commit in the current repository
func Commit(directoryPath string) (success bool, err error) {
	firstInit := false
	repo, err := git.PlainOpen(directoryPath)

	if err != nil {
		//creating repository
		log.Info("creating repository...")
		repo, err = GitInit(directoryPath)
		firstInit = true
	}

	if err != nil {
		return false, err
	}

	w, _ := repo.Worktree()

	log.Info("Committing new changes...")
	if firstInit {
		err = w.AddGlob("*")
		if err != nil {
			return false, err
		}

		//TODO if its new git init, need to add `git pull` command  with remote branch
	} else {
		_, _ = w.Add(".")
	}

	_, _ = w.Commit("test", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "yashre-bh",
			Email: "shreyabhagat2002@gmail.com",
			When:  time.Now(),
		},
	})

	_, err = repo.Remote(defaultRemoteName)
	if err != nil {
		log.Info("Creating new Git remote named " + defaultRemoteName)
		_, err = repo.CreateRemote(&config.RemoteConfig{
			Name: defaultRemoteName,
			URLs: []string{""}, //add repo
		})

		if err != nil {
			return false, err
		}
	}

	//define auth
	auth := &http.BasicAuth{
		Username: "",
		Password: "",
	}

	log.Info("Pushing changes to remote")
	err = repo.Push(&git.PushOptions{
		RemoteName: defaultRemoteName,
		Auth:       auth,
	})

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
