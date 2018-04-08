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

You can run togo as a binary (build and copy to `/usr/local/bin`) or from a Docker container (mounting the config):

```shell
$ docker run -it -v ${HOME}:/root:ro ssube/togo:master list

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

### List

```none
togo list [--columns col1,col2] [--labels label1,label2,label3] [--project project_id] [--sort col] [filter...]
```

List incomplete tasks:

```shell
$ togo list

    ID   Priority  Content
 01231          1  clean desk
 01232          1  update gitlab
 01233          1  clean computer monitor

$ togo list --sort Content

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
$ togo list --columns ID,Order

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
- [x] add complete
- [ ] add project
- [ ] list labels
- [ ] add labels
- [ ] test coverage
- [ ] edit task
- [ ] postpone task
