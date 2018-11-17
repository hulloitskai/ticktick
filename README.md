# ticktick

_An unofficial Go client library for [TickTick](https://ticktick.com)._

## Usage

### Install

```bash
$ go get -u github.com/stevenxie/ticktick # install library
```

### Integrate

```go
import (
  "fmt"
  "github.com/stevenxie/ticktick"
)

// This program logs into TickTick, and lists all remaining tasks.
func main() {
  // Create a client instance.
  client, err := ticktick.NewClient()
  if err != nil {
    panic(err)
  }

  // Login as a user.
  if err = client.Login("example@email.com", "supersecure"); err != nil {
    panic(err)
  }

  // List all remaining (incomplete) tasks:
  todos, err := client.ListTasks()
  if err != nil {
    panic(err)
  }

  for _, todo := range todos {
    fmt.Println(todo.Title)
  }
}
```

<br />

## Project Setup

```bash
## Clone the repository.
$ git clone git@github.com:stevenxie/ticktick

## Configure githooks, dependencies, etc.
$ make setup
```
