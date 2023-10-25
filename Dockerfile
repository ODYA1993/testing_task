FROM golang:alpine3.18 AS builder

COPY . /go/src/app

WORKDIR /go/src/app/cmd

RUN go build -o app app/cmd/main.go

EXPOSE 8080

CMD ["./app"]