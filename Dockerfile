FROM golang:latest

WORKDIR /go

COPY my_app /go/

RUN go get github.com/dtauraso/CS385lab3

CMD ["./my_app"]
