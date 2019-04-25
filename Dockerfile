FROM golang:1.12.4

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

ADD src $GOPATH/src
ADD bin $GOPATH/bin
ADD pkg $GOPATH/pkg

EXPOSE 9090

WORKDIR $GOPATH/src/kaemiin.com/user/beeserver

RUN echo "Start Service"

CMD ["/go/bin/beeserver"]
