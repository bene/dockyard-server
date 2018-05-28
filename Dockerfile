FROM golang

ADD . /go/src/github.com/bene/dockyard-server
WORKDIR /go/src/github.com/bene/dockyard-server

RUN go get -u github.com/golang/dep/...
RUN dep ensure
RUN go install

ENV ADDRESS=:6551

EXPOSE 6551

ENTRYPOINT /go/bin/dockyard-server
