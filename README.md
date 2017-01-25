# rebasebot [![Circle CI](https://circleci.com/gh/chrisledet/rebasebot.svg?style=svg)](https://circleci.com/gh/chrisledet/rebasebot)

A GitHub bot that rebases your pull request branches when you ask

## How it works

1. Make a dedicated GitHub account for the bot
2. Grant the GitHub account read and write access to your repositories
3. [Setup](#setup) the bot on your own server
4. Type a comment "**@{github bot username} rebase**" in a pull request
5. The bot will then kick off a rebase and push (if rebase successful) to your repository
6. You can then delete the comment (in step 4) if you want to, including the rebase comment from the bot.

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

### Configuration

Here are the environment variables rebasebot uses:

* `GITHUB_USERNAME`: GitHub username for bot. Required.
* `GITHUB_PASSWORD`: GitHub password for bot. Required.
* `PORT`: HTTP server port for the bot. Required.
* `TMPDIR`: A path to a writable directory. All local copies will live here. Defaults to OS tmp. **Strongly recommended**.
* `SECRET`: A token used to verify web hook requests from GitHub. **Strongly recommended**.

The `GITHUB_*` are needed so the bot can post activity updates to GitHub as well as push to accessible Git repositories. Using your personal credentials is **not recommended**.

### Run

```shell
$ $GOPATH/bin/rebasebot
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
