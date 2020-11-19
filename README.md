## Overview

```gctl``` is [Gasper](https://gasper-docs.netlify.app/) on the command line. It allows user to deploy and manage applications and databases from terminal.

## Installation

> Using this package requires a working Go environment. [See the install instructions for Go](https://golang.org/doc/install).

**gctl** can be installed using following command.

``` $ go get github.com/sdslabs/gctl ```

If you get following output, that means gctl is successfully installed.
```
$ gctl
Gasper is an intelligent Platform as a Service (PaaS) used for deploying and managing applications and databases in any cloud topology.
```

To execute commands of gctl, run [Gasper](https://gasper-docs.netlify.app/) on your local environment.

## Login

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

Now that we have logged in, we are ready deploy and maintain our applications and databases through command-line.

## Application Deployment

* To create a new application, we will use command ```gctl create app```. We can provide required details of the app either through a form in terminal or passing a json file in a flag.

  1. First we will create app by entering details in a form in the terminal.
      ```
      $ gctl create app
      *App Name: test
      *Language: php
      *Application Password: ****
      *Git URL: https://github.com/sdslabs/gasper-sample-php
      Is this repo private? [yes/no]: no
      Branch: 
      *Index: index.php
      Port: 8000
      Does this repo contain Gasperfile.txt? [yes/no]: no
      Build Commands: 
      Run Commands: 
      Environment Variables(key:value): 
      App created successfully 
      Container Id: 32e81f3d244d09da489aec03bea932ae7d96e8e2f5bd9484fc7a21a7e0e967dd Container Port:  44437 Docker Image: docker.io/sdsws/php:3.0 App Url: test.app.sdslabs.co Host Ip: 192.168.43.137 Name Servers:  [8.8.8.8 8.8.4.4] Instance Type: application Language: php Owner: gmahak1@gmail.com Ssh Cmd: ssh -p 2222 test@192.168.43.137 Id: 5f8c8d094374798e04edf3d6
      ```

      Fields with * are required.

  2. You can also provide app details in a config json file. Just create the json file with necessary data. Example for required json data to deploy an app can be found in example section of [Gasper Docs](https://gasper-docs.netlify.app/). Run the command ```gctl create app {filename} {language}``` in the same folder where your config file is and the app will be deployed.


- Fetch details of an app using the command ```gctl fetch app -n {name}``` where n is flag for the name of the app.

- Fetch details of all the apps using the command ```gctl fetch app```.

- Delete an app using the command ```gctl delete app {app name}```.

- Rebuild an app using the command ```gctl rebuild {app name}```.

- Fetch logs of apps using the command ```gctl fetch logs {app name} {number of logs}```. The second argument, which is for number of logs, is optional.

- Update an app using a json config file with command ```gctl update app {app name} {filename}```. Config file format should be like the one you provided while creating the app. You can also provide details by filling a form in terminal using the command ```gctl update app```.


## Database Deployment

- A new database can be created either by providing details in flags with the command or by filling a terminal form.
Following example shows how to create a mysql database via Gasper using gctl -

  1. Using flags -

      ```
      $ gctl create db -t mysql -n alphamysql -p alphamysql
      Database created
      ```
      Here flag t is for database type, n is for database name and p is for database password.

  2. Using terminal form -

      ```
      $ gctl create db
      *Database Name: alphatmysql       
      *Application Password: ***********
      Database Type: mysql
      Database created
      ```

- Fetch details of a database using the command ```gctl fetch db -n {name}``` where n is flag for the name of the db.

- Fetch details of all the databases using the command ```gctl fetch db```.

- Delete a database using the command ```gctl delete db {db name}```.

## Instances

Fetch details of all the instances using command ```gctl fetch instances```.

## Logout

Logout from a system using the command ```gctl logout```. You can also revoke the token from sws or using following curl command.

```
$ curl -X PUT   http://localhost:3000/auth/revoke   -H 'Content-Type: application/json'   -d '{
    "email": "anish.mukherjee1996@gmail.com",
    "password": "alphadose"
  }'
{"message":"token revoked","success":true}
```

## Contributing

We are always open for contributions. If you find any feature missing, or just want to report a bug, feel free to open an issue and/or submit a pull request regarding the same.

For more information on contribution, refer to the [contributing documentation](https://github.com/sdslabs/gctl/blob/development/CONTRIBUTING.md).

## Contact

If you have a query regarding the product or just want to say hello then feel free to visit
[chat.sdslabs.co](http://chat.sdslabs.co/) or drop a mail at [contact@sdslabs.co.in](mailto:contact@sdslabs.co.in).

Made with :heart: by [SDSLabs](https://github.com/sdslabs)
