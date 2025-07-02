# go-kubernetes-controllers

Project for Kubernetes controllers course by fwdays
---

## step 4 **FastHTTP Server Command**

   - Added a new `server` command using [fasthttp](https://github.com/valyala/fasthttp).
   - The command starts a FastHTTP server with a configurable port (default: 8088).
   - Supports the `--log-level` flag for controlling log verbosity.
   - Uses zerolog for logging.

   **Usage:**
   ```sh
   git switch feature/step4-fasthttp-server

   go run main.go server --port 8080 --log-level debug
   ```

   **What it does:**
   - Starts a FastHTTP server on the specified port.
   - Responds with "Hello from FastHTTP!" to any request.
   - Respects the log level set by the `--log-level` flag.

## Project Structure

- `cmd/` — Contains your CLI commands.
- `main.go` — Entry point for your application.
- `server.go` - fasthttp server

## step 3 **Support log-level flags**

   Example usage in your `main.go`:
   ```go

       rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Set log level: trace, debug, info, warn, error")
       rootCmd.Execute()
   }
   ```

   Build your CLI:
   ```sh
   go build -o controller
   ```

   You can now run your CLI with different log levels:
   ```sh
   ./controller --log-level debug
   ./controller start --log-level trace
   ```
### Notes
Flags can be binded with viper [link](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#bind-flags-with-config)  

#### global logging setup:

`PersistentFlags` - package level flag  
```go
func init() {
rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "", "Logging level")
}
```

`PersistentPreRun` - 
```go
var rootCmd = &cobra.Command{
    ...
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogger()
    },
```

or use `OnInitialize()` in _init_:  
```go
func init() {
	cobra.OnInitialize(initConfig)
```

## step 2 zerolog

### Prerequisites
- [Go](https://golang.org/dl/) 1.24 or newer
- [cobra-cli](https://github.com/spf13/cobra-cli) installed:
  ```sh
  go install github.com/spf13/cobra-cli@latest
  ```
- [zerolog](https://github.com/rs/zerolog) installed:
    ```sh
    go get github.com/rs/zerolog/log
    ```

###

Build your CLI:
```sh
go build -o controller
```

Run your CLI (shows help by default):
```sh
./controller
```


### Zerolog usage quicknotes

**zerolog** allows for logging at the following levels (from highest to lowest):

* panic (`zerolog.PanicLevel`, 5)
* fatal (`zerolog.FatalLevel`, 4)
* error (`zerolog.ErrorLevel`, 3)
* warn (`zerolog.WarnLevel`, 2)
* info (`zerolog.InfoLevel`, 1)
* debug (`zerolog.DebugLevel`, 0)
* trace (`zerolog.TraceLevel`, -1)

**zerolog** allows data to be added to log messages in the form of key:value pairs. The data added to the message adds "context" about the log event that can be critical for debugging as well as myriad other purposes. An example of this is below:

```go
package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    log.Debug().
        Str("Scale", "833 cents").
        Float64("Interval", 833.09).
        Msg("Fibonacci is everywhere")
    
    log.Debug().
        Str("Name", "Tom").
        Send()
}

// Output: {"level":"debug","Scale":"833 cents","Interval":833.09,"time":1562212768,"message":"Fibonacci is everywhere"}
// Output: {"level":"debug","Name":"Tom","time":1562212768}
```

#### Add file and line number to log

Equivalent of `Llongfile`:

```go
log.Logger = log.With().Caller().Logger()
log.Info().Msg("hello world")

// Output: {"level": "info", "message": "hello world", "caller": "/go/src/your_project/some_file:21"}
```

Equivalent of `Lshortfile`:

```go
zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
    return filepath.Base(file) + ":" + strconv.Itoa(line)
}
log.Logger = log.With().Caller().Logger()
log.Info().Msg("hello world")

// Output: {"level": "info", "message": "hello world", "caller": "some_file:21"}
```


#### Pretty logging

To log a human-friendly, colorized output, use `zerolog.ConsoleWriter`:

```go
log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

log.Info().Str("foo", "bar").Msg("Hello world")
```
// Output: 3:04PM INF Hello World foo=bar

#### Error Logging with Stacktrace

Using `github.com/pkg/errors`, you can add a formatted stacktrace to your errors. 

```go
package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	err := outer()
	log.Error().Stack().Err(err).Msg("")
}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}

// Output: {"level":"error","stack":[{"func":"inner","line":"20","source":"errors.go"},{"func":"middle","line":"24","source":"errors.go"},{"func":"outer","line":"32","source":"errors.go"},{"func":"main","line":"15","source":"errors.go"},{"func":"main","line":"204","source":"proc.go"},{"func":"goexit","line":"1374","source":"asm_amd64.s"}],"error":"seems we have an error here","time":1609086683}
```

> zerolog.ErrorStackMarshaler must be set in order for the stack to output anything.



### step 1 cobra-cli 

#### Prerequisites

- [Go](https://golang.org/dl/) 1.24 or newer
- [cobra-cli](https://github.com/spf13/cobra-cli) installed:
  ```sh
  go install github.com/spf13/cobra-cli@latest
  ```

## Getting Started

1. **Clone this repository:**
   ```sh
   git clone https://github.com/saaverdo/k8s-controller-tutorial.git
   cd k8s-controller-tutorial
   ```

2. **Initialize Go module (if not already):**
   ```sh
   go mod init github.com/saaverdo/k8s-controller-tutorial
   ```

3. **Initialize Cobra:**
   ```sh
   cobra-cli init
   ```

4. **Build your CLI:**
   ```sh
   go build -o controller
   ```

5. **Run your CLI (shows help by default):**
   ```sh
   ./controller --help
   ```

## Project Structure

- `cmd/` — Contains your CLI commands.
- `main.go` — Entry point for your application.
- `cmd/go_basic.go`: Implements the command and struct logic
- `cmd/go_basic_test.go`: Unit tests for the struct methods 

This directory contains the `go_basic.go` file, which demonstrates basic usage of Go structs and methods within a Cobra CLI command.

## go_basic.go Overview
- Defines a `Kubernetes` struct with fields for name, version, users, and node number.
- Implements methods to print users and add a new user.
- Registers a `go-basic` Cobra command that
  - Initializes a sample `Kubernetes` struct
  - Prints the list of users
  - Adds a new user
  - Prints the updated list of users

## Usage

To run the `go-basic` command:

```sh
# From the project root
go run main.go go-basic
```

You should see output listing the initial users, then the updated list after adding a new user.