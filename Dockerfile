FROM golang:latest
MAINTAINER xingzheng "xzheng0624@gmail.com"
RUN apt-get update
RUN apt-get install net-tools

RUN go get -u github.com/xxx0624/tinyRPC
RUN echo "tinyRPC download successfully"

ENV APP_HOME /go/src/github.com/xxx0624/tinyRPC
WORKDIR $APP_HOME
RUN echo "current directory is set"

CMD [ "go", "run", "example/server.go"]
EXPOSE 8080