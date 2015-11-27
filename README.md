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
* Go 1.5
* Git
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

Copy the sample config and update with it GitHub credentials in the `username` and `password` option

```shell
$ cp rebasebot.json.sample rebasebot.json
```

There are additional config options for further customize:

* `secret`: A token used to verify web hook requests from GitHub. It's **strongly recommended** that you use this option.

* `port`: HTTP server port for the bot. Defaults to `8080`

### Run

Start rebasebot

```shell
$ rebasebot
```

By default, rebasebot will attempt to locate its config `rebasebot.json` in the current working directory

If you wish to store it in another location, you can specify that path via `CONFIG` environment variable

```shell
$ CONFIG=path/to/rebasebot.json rebasebot
```


### Add GitHub Webhook

This is a required step to complete the setup.

1. Go into your GitHub repository's Webhooks and services page
2. Add webhook
  1. Enter `http://<your host>/rebase` in the "Payload URL" field
  2. Content type should be set to "application/json"
  3. Generate a secret token and enter it in "Secret" field
  4. Only send "Issue comment" events. All other ones will be ignored.
3. GitHub should succesfully ping the service and receive a HTTP 200 OK

## Resources

* GitHub guide for [securing your webhooks](https://developer.github.com/webhooks/securing/)

* Generate secret token with Ruby

  ```shell
  $ ruby -rsecurerandom -e 'puts SecureRandom.hex(20)'
  ```
