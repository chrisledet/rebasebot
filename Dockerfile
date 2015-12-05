FROM golang:1.5
MAINTAINER Chris Ledet <me@chrisledet.com>

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin

# Install rebasebot
RUN go get -u github.com/chrisledet/rebasebot
RUN go install github.com/chrisledet/rebasebot

RUN ln -sf /dev/stdout /var/log/rebasebot/access.log
RUN ln -sf /dev/stderr /var/log/rebasebot/error.log

# Set default container command
ENTRYPOINT $GOPATH/bin/rebasebot
