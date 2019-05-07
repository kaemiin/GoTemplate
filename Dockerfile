FROM golang:1.12.4

MAINTAINER kaemiin chang <kaemiinwork@gmail.com>

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

ADD ./src/kaemiin.com /go/src/kaemiin.com

RUN cd /go \
    && go get github.com/joho/godotenv \
    && go get github.com/jasonlvhit/gocron \
    && go get gopkg.in/natefinch/lumberjack.v2 \
	&& go install kaemiin.com/user/app

WORKDIR /go/src/kaemiin.com/user/app

RUN mkdir log

RUN echo "Start Service"

ENTRYPOINT ["app"]
