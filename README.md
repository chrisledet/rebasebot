# rebasebot

A GitHub integration service that rebases your Pull request branches when you ask

## How it works

1. Make a dedicated GitHub account for the bot
2. Grant the GitHub account read and write access to your repositories
3. [Setup](#setup) the bot on your own server
4. **@mention** the bot in a pull request comment
5. The bot will then kick off a rebase and push (if rebase successful) to your repository

## Dependencies

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

This is a required step to complete the setup.

1. Go into your GitHub repository's Webhooks and services page
2. Add webhook
  1. Enter your rebasebot's host public URL
  2. Only send "Issue comment" events. All other ones will be ignored.
3. GitHub should succesfully ping the service and receive a HTTP 200 OK
