package middlewares

import (
	// "fmt"
	"context"
	_ "io/ioutil"

	// "net/http"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v41/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	// "github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func goDotEnvVariable(key string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	// load .env file
	err = godotenv.Load(environmentPath)

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

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
	log.Info(goDotEnvVariable("PAT"))
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

	_, _, err := client.Repositories.CreateKey(context.Background(), "yashre-bh", repoName, deployCred)
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
