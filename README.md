
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