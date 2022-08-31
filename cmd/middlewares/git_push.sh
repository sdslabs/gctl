#! /bin/sh

gitPush(){
    directory="$1" 
    repo_url="$2"
    first_init="$3"

    shopt -s cdable_vars
    export directory="$directory"

    cd directory
    git init

    # Display unstaged files
    git status

    git remote add origin $repo_url
    git add .
    git commit -m "latest commit"
    git push --set-upstream origin master
}

gitPush