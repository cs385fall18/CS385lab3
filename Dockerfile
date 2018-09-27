FROM golang:latest

WORKDIR /go

COPY * /go/src/

RUN go get my_app

CMD ["./bin/my_app"]
