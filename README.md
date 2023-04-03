
# Golang REST API - Clean Architecture Principles

Source codes for blog series: https://www.favtuts.com/golang-go-crash-course-series/


# install golang

* Download: 
    - Access `golang.org` redirects `go.dev`
    - Download page: https://go.dev/dl/
    - Install Guide: https://go.dev/doc/install

```
wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz
```

**For linux:**

1. **Remove any previous Go installation** by deleting the `/usr/local/go` folder (if it exists), then extract the archive you just downloaded into `/usr/local`, creating a fresh Go tree in `/usr/local/go`:
```
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz
```

(You may need to run the command as root or through sudo).

**Do not** untar the archive into an existing `/usr/local/go` tree. This is known to produce broken Go installations.

2. Add `/usr/local/go/bin` to the PATH environment variable.
You can do this by adding the following line to your `$HOME/.profile` or `/etc/profile` (for a system-wide installation):
```
export PATH=$PATH:/usr/local/go/bin
```
**Note:** Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as `source $HOME/.profile`.

3. Verify that you've installed Go by opening a command prompt and typing the following command:
```
$ go version
```
Confirm that the command prints the installed version of Go.


# create project folder

```
cd $GOPATH/src
mkdir golang-mux-api
cd golang-mux-api
go mod init golang-mux-api
```

# Go Modules commands

Initialize the module for the app (from root dir)

```
$ go mod init github.com/favtuts/golang-mux-api
```

Edit the existing module file for the new name
```
$ go mod edit -module github.com/favtuts/golang-mux-api
```

To rename all imported module from OLD_MODULE to NEW_MODULE
```
find . -type f -name '*.go' \
  -exec sed -i -e 's,{OLD_MODULE},{NEW_MODULE},g' {} \;
```
In our cases
```
find . -type f -name '*.go' \
  -exec sed -i -e 's,golang-mux-api,github.com/favtuts/golang-mux-api,g' {} \;
```


Download external dependencies
```
$ go mod download  
```

The go command uses the `go.sum` file to ensure that future downloads of these modules retrieve the same bits as the first download (it keeps an initial hash)

Both `go.mod` and `go.sum` should be checked into version control.


# install Mux library

```
$ go get -u github.com/gorilla/mux

go: added github.com/gorilla/mux v1.8.0
```

# install Chi library
```
$ go get -u github.com/go-chi/chi

go: added github.com/go-chi/chi v1.5.4
```


# install Firestore library
```
$ go get -u cloud.google.com/go/firestore
```

# install MySQL library
```
$ go get -u github.com/go-sql-driver/mysql

go: added github.com/go-sql-driver/mysql v1.7.0
```

# install MongoDB library
```
$ go get -u go.mongodb.org/mongo-driver

go: added go.mongodb.org/mongo-driver v1.11.2
```

# install SQLite 3 Library
```
$ go get -u github.com/mattn/go-sqlite3

go: added github.com/mattn/go-sqlite3 v1.14.16
```

# install Testify library

```
$ go get -u github.com/stretchr/testify

go: added github.com/stretchr/testify v1.8.2
```

# export Environment variable
```
export GOOGLE_APPLICATION_CREDENTIALS='/path/to/project-private.key.json'

/home/tvt/FIREBASE/pragmatic-reviews-62551-firebase-adminsdk-l8p5o-65461f8c5b.json
```

# how to get the private key JSON file:

From the Firebase Console: **Project Overview** -> **Project Settings** -> **Service Accounts** -> **Generate new private key**


# build project

```
go build
```

# run project

```
go run .
```

or

```
go run *.go
```

# run Unit Test

Test (specific test)
```
$ go test -run NameOfTest
```

Test (all the tests within the service folder)
```
$ go test service/*.go
```

# set PORT environment

```
$ export PORT="8000"
$ echo $PORT
```

# use docker container

To build docker image
```
$ docker build -t golang-api .
```

To run docker container
```
$ docker run -p 8000:8000 golang-api
```


# install Redis on Linux (Ubuntu)

```
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

sudo apt-get update
sudo apt-get install redis

sudo service redis-server start
```


# testing REST API

```
$ cd controller/
$ go test -v -run TestAddPost
=== RUN   TestAddPost
--- PASS: TestAddPost (0.01s)
PASS
ok      github.com/favtuts/golang-mux-api/controller    0.019s


$ go test -v -run TestGetPosts
=== RUN   TestGetPosts
--- PASS: TestGetPosts (0.01s)
PASS
ok      github.com/favtuts/golang-mux-api/controller    0.034s


$ go test -v -run TestGetPostByID
=== RUN   TestGetPostByID
--- PASS: TestGetPostByID (0.01s)
PASS
ok      github.com/favtuts/golang-mux-api/controller    0.028s
```