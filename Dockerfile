FROM golang:latest

WORKDIR /go

COPY * /go/src/

RUN go get github.com/dtauraso/CS385lab3

CMD ["./bin/my_app"]
