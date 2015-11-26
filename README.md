# rebasebot

A GitHub integration service that rebases your Pull request branches when you ask.

## Setup

Build

```shell
$ go build
```

Config

```shell
$ cp rebaser.json{.sample,}
```

Edit `rebaser.json` and enter a GitHub account and a list of repositories to checkout locally:

```json
{
  "username": "rebaser",
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
