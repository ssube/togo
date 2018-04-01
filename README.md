# togo

CLI for Todoist's v8 API written in Go.

## Build

```shell
dep ensure
go build
```

## Config

togo expects a `~/.togo.yml` file to exist in the current user's home directory.

The config file should contain:

```yaml
token: api-token
```

## Run

```shell
togo list
togo add "task"
togo done "task"
```