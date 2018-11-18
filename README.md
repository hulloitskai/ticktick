# ticktick

_An unofficial Go client library for [TickTick](https://ticktick.com)._

I reverse-engineered the TickTick API with the help of
[`RayBB/node-ticktick-api`](https://github.com/RayBB/node-ticktick-api), and a
lot of digging around with [the Insomnia REST client](https://insomnia.rest).

This client library is the result of those efforts.

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
  if err = client.Login("email@example.com", "supersecure"); err != nil {
    panic(err)
  }

  // List all remaining (incomplete) tasks.
  tasks, err := client.GetTasks()
  if err != nil {
    panic(err)
  }

  for _, task := range tasks {
    fmt.Println(task.Title)
  }
}
```

### Examples

See `example/` for a more complete example of using client to interact with
the TickTick API.

You can run this example with:

```bash
$ go run example/main.go
```

<br />

## Project Setup

```bash
## Clone the repository.
$ git clone git@github.com:stevenxie/ticktick

## Configure githooks, dependencies, etc.
$ make setup
```
