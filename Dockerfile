FROM golang

ADD . /go/src/github.com/dbd/dbd.sh

RUN go install github.com/dbd/dbd.sh

ENTRYPOINT /go/bin/dbd.sh

EXPOSE 20300
