package middlewares

import (
	// "fmt"
	"context"
	"fmt"
	_ "io/ioutil"

	// "net/http"
	"os"
	_ "path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
	// "github.com/go-git/go-git/v5/plumbing/transport/ssh"
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

func CreateRepository(repoName string) (*github.Repository, *github.Response, error) {
	// create a new private repository named "foo"
	tc := oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: goDotEnvVariable("PAT")},
		),
	)
	client := github.NewClient(tc)
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(true),
	}
	repo, res, err := client.Repositories.Create(context.Background(), "", repo)
	fmt.Printf("hello %v", (&res))
	return repo, res, err
}

func AddDeployKeyToRepo(repoName string, deployKey string) error {
	tc := oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: goDotEnvVariable("PAT")},
		),
	)
	client := github.NewClient(tc)
	deployCred := &github.Key{
		Key: &deployKey,
	}

	_, _, err := client.Repositories.CreateKey(context.Background(), "username", repoName, deployCred)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// Commit creates a commit in the current repository
// func Commit(directoryPath string) (success bool, err error) {
// 	firstInit := false
// 	repo, err := git.PlainOpen(directoryPath)

// 	if err != nil {
// 		//creating repository
// 		log.Info("creating repository...")
// 		repo, err = GitInit(directoryPath)
// 		firstInit = true
// 	}

// 	if err != nil {
// 		return false, err
// 	}

// 	w, _ := repo.Worktree()

// 	log.Info("Committing new changes...")
// 	if firstInit {
// 		err = w.AddGlob("*")
// 		if err != nil {
// 			return false, err
// 		}

// 		//TODO if its new git init, need to add `git pull` command  with remote branch
// 	} else {
// 		_, _ = w.Add(".")
// 	}

// 	_, _ = w.Commit("test", &git.CommitOptions{
// 		Author: &object.Signature{
// 			Name:  "yashre-bh",
// 			Email: "shreyabhagat2002@gmail.com",
// 			When:  time.Now(),
// 		},
// 	})

// 	_, err = repo.Remote(defaultRemoteName)
// 	if err != nil {
// 		log.Info("Creating new Git remote named " + defaultRemoteName)
// 		_, err = repo.CreateRemote(&config.RemoteConfig{
// 			Name: defaultRemoteName,
// 			URLs: []string{"https://github.com/yashre-bh/test.git"}, //add repo
// 		})

// 		if err != nil {
// 			return false, err
// 		}
// 	}

// 	// currentUser, err := user.Current()
// 	// if err != nil {
// 	// 	log.Error(err.Error())
// 	// }
// 	// sshAuth, err := ssh.NewPublicKeysFromFile("git", currentUser.HomeDir+"/.ssh/id_rsa_deploy", "")
// 	auth := &http.BasicAuth{
// 		Username: "yashe-bh",
// 		Password: "",
// 	}

// 	if err != nil {
// 		log.Error(err.Error())
// 	}

// 	log.Info("Pushing changes to remote")
// 	err = repo.Push(&git.PushOptions{
// 		RemoteName: defaultRemoteName,
// 		Auth:       auth,
// 	})

// 	if err != nil {
// 		log.Info("error haha")
// 		return false, err
// 	} else {
// 		log.Info("success")
// 		return true, nil
// 	}
// }
