FROM golang:latest

WORKDIR /go/src

# ENV GOPATH=/go/src
# ENV GOROOT=/go/root

COPY ./ ./ 