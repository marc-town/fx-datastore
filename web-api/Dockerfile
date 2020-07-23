FROM golang:1.13 AS build

WORKDIR /go/src/github.com/marc-town/fx-datastore/web-api
ENV GO111MODULE=on
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .

WORKDIR /go/src/github.com/marc-town/fx-datastore/web-api/src
RUN go build -o /go/bin \
    && go get gopkg.in/urfave/cli.v2@master \
    && go get github.com/oxequa/realize