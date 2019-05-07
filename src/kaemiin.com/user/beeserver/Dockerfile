FROM golang:1.12.4

MAINTAINER kaemiin chang <kaemiinwork@gmail.com>

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

ADD ./ /go/src/app

RUN cd /go/src/app \
	&& go get github.com/astaxie/beego \
	&& go build

EXPOSE 9090

RUN echo "Start Service"

ENTRYPOINT ["/go/src/app/app"]
