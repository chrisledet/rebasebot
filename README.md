# rebasebot

A GitHub integration service that rebases your Pull request branches when you ask.

## Setup

Install

```shell
$ go get github.com/chrisledet/rebasebot
```

Build

```shell
$ go build
```

Config

```shell
$ cp rebasebot.json{.sample,}
```

Edit `rebasebot.json` and enter a GitHub account and a list of repositories to checkout locally:

```json
{
  "username": "rebasebot",
  "password": "1s0l33t",
  "tmpdir": "tmp",
  "repositories": [
    "chrisledet/dotfiles"
  ]
}
```

Run

```shell
$ ./rebasebot
```

Test

```shell
$ go test ./...
```
