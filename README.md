### Golang Project

#### DEVELOP ON DOCKER

```
docker run -it -d -v $(pwd):/go --name go-template golang bash

docker exec -it go-template bash
```

#### DOWNLOAD & INSTALL

```
go get -u -v github.com/tools/godep
go get -u golang.org/x/lint/golint
go get github.com/joho/godotenv
go get github.com/astaxie/beego
```

#### BUILD & INSTALL

```
go install kaemiin.com/user/hello
```

or

```
cd $GOPATH/src/kaemiin.com/user/hello
go install
```

#### BUILD

```
go build kaemiin.com/user/stringutil
```

#### RUN

```
cd $GOPATH/src/kaemiin.com/user/hello

$GOPATH/bin/hello
```

or

```
hello
```

#### LINT

```
golint kaemiin.com/user/hello
```

### TEST

```
go test kaemiin.com/user/stringutil
```

### FORMATTING

```
gofmt $GOPATH/src/kaemiin.com/user/hello
```

### USE godep

save your dependencies:

```
cd /go/src/kaemiin.com/user/stringutil

godep save
```

add a dependency:

```
1. go get foo/bar
2. Edit your code to import foo/bar.
3. Run godep save
```

update a dependency:

```
1. go get -u foo/bar
2. godep update foo/bar
```

#### RUN ON DOCKER ?

```
docker build --rm -t kaemiin-golang .

docker run -it -d -p 8080:9090 --name go-server kaemiin-golang bash
```
