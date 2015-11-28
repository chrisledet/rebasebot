FROM ubuntu:14.04
MAINTAINER Chris Ledet <me@chrisledet.com>

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin

RUN mkdir -p /data/tmp
WORKDIR /data

# Dependencies
RUN apt-get update && apt-get install -y \
    git \
    wget

# Install Go
RUN wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz -O tmp/go1.5.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf tmp/go1.5.1.linux-amd64.tar.gz

# Install rebasebot
RUN go get -u github.com/chrisledet/rebasebot
RUN go install github.com/chrisledet/rebasebot

# Download config
RUN echo '{ "username": "", "password": "", "port": "", "secret": "", "tmpdir": "tmp"}' > rebasebot.json

# Expose port
EXPOSE 8080

# Set default container command
ENTRYPOINT $GOPATH/bin/rebasebot
