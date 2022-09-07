package middlewares

import (
	"context"
	"fmt"
	_ "io/ioutil"
	"time"

	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/v41/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func GoDotEnvVariable(key string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath)

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

func GitInit(directoryPath string) (*git.Repository, error) {
	fmt.Println("GITINIT DONE")
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

func CreateRepository(repoName string) (*github.Repository, *github.Response, error) {
	tc := oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: GoDotEnvVariable("PAT")},
		),
	)
	client := github.NewClient(tc)
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(true),
	}
	repo, res, err := client.Repositories.Create(context.Background(), "", repo)
	return repo, res, err
}

func AddDeployKeyToRepo(repoName string, deployKey string) error {
	tc := oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: GoDotEnvVariable("PAT")},
		),
	)
	client := github.NewClient(tc)
	deployCred := &github.Key{
		Key: &deployKey,
	}

	_, _, err := client.Repositories.CreateKey(context.Background(), "yashre-bh", repoName, deployCred)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GitPush(pathToApplication string, repoURL string) (*git.Repository, error) {
	var firstInit bool
	repo, err := git.PlainOpen(pathToApplication)
	if err != nil {
		firstInit = true
		repo, err = GitInit(pathToApplication)
		if err != nil {
			return nil, err
		}
	} else {
		firstInit = false
	}
	_, err = repo.Remote("origin")
	if err != nil {
		_, err = repo.CreateRemote(&config.RemoteConfig{
			Name: "origin",
			URLs: []string{repoURL},
		})
		if err != nil {
			return nil, err
		}
	}
	w, _ := repo.Worktree()
	if firstInit {
		err = w.AddGlob(".")
		if err != nil {
			return nil, err
		}
	} else {
		_, err = w.Add(".")
		if err != nil {
			return nil, err
		}
	}

	_, _ = w.Commit("latest commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  GoDotEnvVariable("USERNAME"),
			Email: GoDotEnvVariable("EMAIL"),
			When:  time.Now(),
		},
	})

	auth := &http.BasicAuth{
		Username: GoDotEnvVariable("USERNAME"),
		Password: GoDotEnvVariable("PAT"),
	}

	if err != nil {
		return nil, err
	}
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	if err != nil {
		return nil, err
	}

	return repo, err
}
