# togo

CLI for [Todoist's](https://todoist.com) [v8 API](https://developer.todoist.com/rest/v8/) written in Go.

## Build

```shell
dep ensure
go build
```

## Config

togo expects a `~/.togo.yml` file to exist in the current user's home directory.

The config file should contain:

```yaml
root: "https://beta.todoist.com/API/v8"
token: "api-token"
```

You can find your API token [on the Integrations page](https://todoist.com/Users/viewPrefs?page=integrations).

## Usage

togo can list and complete pending tasks, as well as add new tasks.

### To Go

Count incomplete tasks:

```shell
$ togo

3 tasks to go
```

### List

List incomplete tasks:

```shell
$ togo list

    id   priority  content
 01231          1  clean desk
 01232          1  update gitlab
 01233          1  clean computer monitor
```

### Add

Add a new task:

```shell
$ togo add --content "task"

    id   priority  content
 01234          1  hello world
```

### Done

Complete an existing task, by id:

```shell
$ togo done 01231 01232 01233 01234

closing 01231
closing 01232
closing 01233
closing 01234
```
