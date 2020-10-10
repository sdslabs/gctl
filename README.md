# Overview

```gctl``` is [Gasper](https://gasper-docs.netlify.app/) on the command line. It allows user to deploy and manage applications and databases from terminal.

# Installation

> Using this package requires a working Go environment. [See the install instructions for Go](https://golang.org/doc/install).

**gctl** can be installed using following command.

``` $ go get github.com/sdslabs/gctl ```

If you get following output, that means gctl is successfully installed.
```
$ gctl
Gasper is an intelligent Platform as a Service (PaaS) used for deploying and managing applications and databases in any cloud topology.
```

To execute commands of gctl, run [Gasper](https://gasper-docs.netlify.app/) on your local environment.

# Login

After Gasper is up and successfully running, we need to generate PA Token first to login in command-line.

Login to gasper using following command to generate the PA token

```
$ curl -X POST \
  http://localhost:3000/auth/login \
  -H 'Content-Type: application/json' \
  -H 'Authorization-Type: gctlToken' \
  -d '{
    "email": "anish.mukherjee1996@gmail.com",
    "password": "alphadose"
  }'

{
    "code":200,
    "expire":"2020-10-10T21:27:30+05:30","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImVtYWlsIjoiZ21haGFrMUBnbWFpbC5jb20iLCJleHAiOjE2MDIzNTE4MTAsImdjdGxfdXVpZCI6IiIsIm9yaWdfaWF0IjoxNjAyMzQ4MjEwLCJ1c2VybmFtZSI6Im1haGFrIn0.bImaUw9p8K_2QMpMqCAyHQHzX2aukDaRpXTDXmAkAoc"
}
```

After getting the PA token, login to gctl using command ```gctl login``` with flag e for email and t for token. Both of the flags are required.

```
$ gctl login -e anish.mukherjee1996@gmail.com -t eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImVtYWlsIjoiZ21haGFrMUBnbWFpbC5jb20iLCJleHAiOjE2MDIzNTE4MTAsImdjdGxfdXVpZCI6IiIsIm9yaWdfaWF0IjoxNjAyMzQ4MjEwLCJ1c2VybmFtZSI6Im1haGFrIn0.bImaUw9p8K_2QMpMqCAyHQHzX2aukDaRpXTDXmAkAoc
Logged in successfully
```
