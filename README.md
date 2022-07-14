# sutil

`sutil` is collection of useful commands for working with Slack.

## Install

```sh
go install github.com/s-beats/sutil@latest
```

## Usage

you configure the environment variable.

```sh
$ export SLACK_OAUTH_TOKEN=xxx...
```

### Aggregate messages 

```sh
$ sutil aggregate-messages --chanid xxxxxxxxxxx
```

```sh
USER NAME         MESSAGES COUNT
A                 10
B                 5
C                 21
```
