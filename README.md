# Todo

A simple command-line todo application written in Go.

## Features

- Add tasks
- List tasks
- Delete tasks
- Store tasks in CSV format
- Show task creation times

## Installation

Clone the repository:

```bash
git clone https://github.com/artuacioli/todo.git
cd todo
```

Build the application:

```bash
go build
```

Or run directly:

```bash
go run .
```

## Usage

Add a task:

```bash
todo add "Buy groceries"
```

List active tasks:

```bash
todo list
```

List all tasks, including completed ones:

```bash
todo list -a
```

or

```bash
todo list --all
```

Delete a task:

```bash
todo delete 1
```

## Example

```bash
$ todo add "Learn Go"

$ todo list
ID  TASK      ADDED_AT
1   Learn Go  2 minutes ago
```



