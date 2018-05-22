# togo

CLI for [Todoist's](https://todoist.com) [v8 API](https://developer.todoist.com/rest/v8/) written in Go.

[![docker size](https://img.shields.io/microbadger/image-size/ssube/togo.svg)](https://hub.docker.com/r/ssube/togo/)
[![pipeline status](https://git.apextoaster.com/ssube/togo/badges/master/pipeline.svg)](https://git.apextoaster.com/ssube/togo/commits/master)
[![license](https://img.shields.io/github/license/ssube/togo.svg)](https://github.com/ssube/togo/blob/master/LICENSE.md)

Binaries are available from [the Github releases page](https://github.com/ssube/togo/releases) and container images
from [the Docker hub](https://hub.docker.com/r/ssube/togo/).

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

You may also set the default columns and sort order for tables:

```yaml
default:
  labels:
    columns: [ID, Name]
    sort: Name
  projects:
    columns: [ID, Name]
    sort: Name
  tasks:
    columns: [ID, Content]
    sort: ID
```

These defaults are used unless the `--columns` and `--sort` flags are passed. This section is optional; if omitted, the
values shown above will be used.

## Usage

```none
togo [--columns col1,col2] [--sort col] cmd [flags...]
```

You can run togo as a binary or from a Docker container (mounting the config):

```shell
$ docker run -it -v ${HOME}:/root:ro ssube/togo list

    ID   Priority  Content
...
```

### Count

```none
togo
```

Count incomplete tasks:

```shell
$ togo

3 tasks to go
```

Count uses the `--filter` parameter, defaulting to `today | overdue`.

### List

```none
togo [options...] list [--labels label1,label2,label3] [--project project_id] [filter...]
```

List incomplete tasks:

```shell
$ togo list

    ID   Priority  Content
 01231          1  clean desk
 01232          1  update gitlab
 01233          1  clean computer monitor

$ togo --sort Content list

    ID   Priority  Content
 01233          1  clean computer monitor
 01231          1  clean desk
 01232          1  update gitlab

$ togo list --project 03211 --labels computer,desk "search: monitor"

    ID   Priority  Content
 01233          1  clean computer monitor
```

The `columns` parameter selects fields from the tasks, in order, and displays them in a table with headers:

```shell
$ togo --columns ID,Order list

    ID   Order
 01231       1
 01232       2
 01233       3
```

The `project` parameter only lists tasks from a single project, `labels` are applied with an `|` operator, and
any trailing arguments are passed as an `&` filter.

[Filters are documented here](https://support.todoist.com/hc/en-us/articles/205248842) and limited to Todoist Premium.
If no filter is provided, the parameter is omitted, which is equivalent to `"all"`. Labels are combined with the `|`
operator. Additional filters are wrapped in parentheses and joined with `&`. In examples:

|                       options |                         filter |
| ----------------------------- | ------------------------------ |
|            `--labels foo,bar` |                 `@foo \| @bar` |
|               `"search: foo"` |                `"search: foo"` |
|  `--labels foo "search: bar"` |     `(@foo) & ("search: bar")` |
|  `"overdue \| today" "#Work"` | `(overdue \| today) & (#Work)` |

### Add

```none
togo add [--done] [content...]
```

Add a new task:

```shell
$ togo add "task"

    ID  Content
 01234  hello world
```

Trailing arguments are merged with `" "` (a space), so loose words will be combined but special characters should be
quoted.

The `--done` parameter completes the task immediately after adding it.

### Done

```none
togo done [id...]
```

Complete a task:

```shell
$ togo done 01231 01232 01233 01234

closing 01231
closing 01232
closing 01233
closing 01234
```

### Projects

```none
togo projects
```

List projects:

```shell
$ todo projects

     ID  Name
  01231  Inbox
  01232  Personal
```

Project IDs may be used with `list --project`.

## Todo

Features:

- [x] labels
- [x] filters
- [x] custom columns
- [x] list projects
- [x] sort order
- [x] add task & complete
- [x] custom root filter
- [x] add task project
- [x] list labels
- [x] add task labels
- [x] config defaults
- [x] columns & sort on root
- [x] add due date column
- [ ] list project names
- [ ] list label names
- [ ] test coverage
- [ ] edit task
- [ ] postpone task
- [ ] create config

## Legal

togo was written by [ssube](https://github.com/ssube) and is not created by, affiliated with, or supported by Doist.

Source, documentation, and everything else in this repository is distributed under the included MIT license.
