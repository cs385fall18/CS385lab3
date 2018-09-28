FROM golang:latest

WORKDIR /go

COPY * /go/src/

RUN go get My_app

CMD ["./bin/my_app"]
