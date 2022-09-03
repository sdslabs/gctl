#! /bin/sh
directory="$1"
repo_url="$2"

gitPush(){
    directory="$1"
    repo_url="$2"
    
    shopt -s cdable_vars
    export dir="$directory"

    cd dir
    git init

    git status

    git remote add origin $repo_url
    git add .
    git commit -m "latest commit"
    git push --set-upstream origin master
}

gitPush $directory $repo_url