### Golang Project

#### USE DOCKER

```
docker run -it -d -v $(pwd):/go -p 8080:8080 --name go-template golang bash

docker exec -it go-template bash
```

#### COMPILE

```
go install kaemiin.com/user/hello
```

or

```
cd $GOPATH/src/kaemiin.com/user/hello
go install
```

#### RUN

```
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