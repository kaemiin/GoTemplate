### Golang Project

#### DEVELOP ON DOCKER

```
docker run -it -d -v $(pwd):/go --name go-template --link mysql:mysql --link redis:redis golang bash

docker exec -it go-template bash
```

#### DOWNLOAD & INSTALL

```
go get -u -v github.com/tools/godep
go get -u golang.org/x/lint/golint
go get github.com/joho/godotenv
go get github.com/astaxie/beego
go get github.com/jasonlvhit/gocron
go get github.com/go-sql-driver/mysql
go get gopkg.in/natefinch/lumberjack.v2
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

or

```
go run src/kaemiin.com/user/app/*.go
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
cd ./src/kaemiin.com/user/beeserver

docker build --rm -t beeserver-test .

docker run -it -d -p 8080:9090 --name bee-server beeserver-test bash
```

```
docker build --rm -t go-test .

docker run -it -d --name go-test go-test bash

docker exec -it go-test bash

rm /etc/localtime

echo "Asia/Taipei" > /etc/timezone

exit

docker cp /usr/share/zoneinfo/Asia/Taipei go-test:/etc/localtime
```
