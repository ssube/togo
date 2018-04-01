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

```none
togo list [--project project_id] [--labels label1,label2,label3] [filter...]
```

List incomplete tasks:

```shell
$ togo list

    id   priority  content
 01231          1  clean desk
 01232          1  update gitlab
 01233          1  clean computer monitor

$ togo list --project 03211 --labels computer,desk "search: monitor"

    id   priority  content
 01233          1  clean computer monitor
```

The `project` parameter only lists tasks from a single project, `labels` are applied with an `|` operator, and
any trailing arguments are passed as an `&` filter.

[Filters are documented here](https://support.todoist.com/hc/en-us/articles/205248842) and limited to Todoist Premium.
If no filter is provided, the parameter is omitted, which is equivalent to `"all"`. Labels are combined with the `|`
operator. Additional filters are wrapped in parentheses and joined with `&`. In examples:

|                       options |                        filter |
| ----------------------------- | ----------------------------- |
|            `--labels foo,bar` |                 `@foo | @bar` |
|               `"search: foo"` |               `"search: foo"` |
|  `--labels foo "search: bar"` |    `(@foo) & ("search: bar")` |
|   `"overdue | today" "#Work"` | `(overdue | today) & (#Work)` |

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

## Todo

Future features:

- list projects
