FROM golang:1.19-alpine

WORKDIR /go/src/app/

RUN apk update && apk add git

COPY go.mod go.sum /go/src/app/
# COPY ./go.sum /go/src/app/

RUN go mod tidy
