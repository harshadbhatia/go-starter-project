FROM golang:alpine as development

WORKDIR /opt/src

COPY src/ /opt/src/

RUN go get github.com/githubnemo/CompileDaemon

RUN apk add bash

ENTRYPOINT go mod tidy && CompileDaemon -command="./dcot"


FROM golang:alpine as release

WORKDIR /opt/src
COPY src/ /opt/src/

RUN go mod download
RUN go mod tidy
RUN go build

# RUN apk add bash

ENTRYPOINT ["./dcot"]


