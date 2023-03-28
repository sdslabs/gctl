package middlewares

import (
	"errors"
	_ "io/ioutil"
	"time"

	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

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

func GitPush(pathToApplication string, repoURL string, pat string, email string, username string, rebuild bool) error {
	var firstInit bool
	repo, err := git.PlainOpen(pathToApplication)
	if err != nil {
		firstInit = true
		repo, err = GitInit(pathToApplication)
		if err != nil {
			return err
		}
	} else {
		firstInit = false
	}

	if !rebuild {
		remote, err := repo.Remote("origin")
		if remote != nil {
			return errors.New("Remote of the local directory already exists, use the git remote URL")
		} else if err != nil {
			_, err = repo.CreateRemote(&config.RemoteConfig{
				Name: "origin",
				URLs: []string{repoURL},
			})
			if err != nil {
				return err
			}
		}
	}
	
	w, _ := repo.Worktree()
	if firstInit {
		err = w.AddGlob(".")
		if err != nil {
			return err
		}
	} else {
		_, err = w.Add(".")
		if err != nil {
			return err
		}
	}

	_, _ = w.Commit("latest commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: email,
			When:  time.Now(),
		},
	})

	auth := &http.BasicAuth{
		Username: username,
		Password: pat,
	}

	if err != nil {
		return err
	}
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	if err != nil {
		return err
	}

	return err
}

func GitRemote(pathToApplication string) (string, error){
	repo, err := git.PlainOpen(pathToApplication)
	if err != nil {
		return "", err
	}
	remote, err := repo.Remote("origin")
	if err != nil {
		return "", err
	} else {
		return remote.Config().URLs[0], nil
	}
}