# rebasebot

A GitHub integration service that rebases your Pull request branches when you ask

## Depedencies

* Dedicated host (e.g. EC2, Digital Ocean, Rackspace)
* Go version 1.5
* Dedicated GitHub account

## Setup

### Download

```shell
$ go get github.com/chrisledet/rebasebot
```

### Build

```shell
$ cd $GOPATH/src/github.com/chrisledet/rebasebot
$ go build
```

### Install

Make sure `$GOPATH/bin` is located in your `$PATH`

```shell
$ go install
```

### Config

Copy from `rebasebot.json.sample` and replace with your GitHub login

```shell
$ cp rebasebot.json.sample rebasebot.json
```

### Run

Start rebasebot

```shell
$ rebasebot
```

if `rebasebot.json` isn't present in the current directory, you can specify its location

```shell
$ CONFIG=path/to/rebasebot.json rebasebot
```

By default, rebasebot binds to port `8080` however you can use a different port

```shell
$ PORT=80 rebasebot
```

### Add GitHub Webhook

1. Go into your GitHub repository's Webhooks and services page
2. Add webhook
  1. Enter your rebasebot's host public URL
  2. Only send "Issue comment" events. All other ones will be ignored.
3. GitHub should succesfully ping the service and receive a HTTP 200 OK
