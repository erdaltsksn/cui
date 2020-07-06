# CUI (Commandline User Interface)

[![GoDoc](https://godoc.org/github.com/erdaltsksn/cui?status.svg)](https://godoc.org/github.com/erdaltsksn/cui)
![Go](https://github.com/erdaltsksn/cui/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/cui)](https://goreportcard.com/report/github.com/erdaltsksn/cui)

CUI gives us a set of useful tools while developing cli apps.

## Features

- Alerts
- Pre-made Cobra commands

## Getting Started

```sh
go get github.com/erdaltsksn/cui
touch main.go
```

**main.go:**

```go
package main

import (
	"fmt"

	"github.com/erdaltsksn/cui"
)

func main() {
	cui.Info("Information message")
}
```

```sh
go run main.go
```

**Output:**

![output](/assets/output.png)

Check out [examples](examples/simple) directory for more.

## Installation

```sh
go get github.com/erdaltsksn/cui
```

## Updating / Upgrading

```sh
go get -u github.com/erdaltsksn/cui
```

## Usage

```go
// Success
cui.Success("This is a success message")

// Info
cui.Info("Information message")

// Warning
details := "This is the details"
cui.Warning("Warning message", details)

// Error
err := errors.New("Wrapped error message")
if err != nil {
    cui.Error("Error message", err)
}
```

### Adding Version Command

**.goreleaser.yml:**

```yaml
builds:
- main:
  ldflags: -s -w -X github.com/erdaltsksn/cui.appVersion={{.Tag}}
```

**cmd/root.go:**

```go
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/erdaltsksn/cui"
)

func init() {
	rootCmd.AddCommand(cui.VersionCmd)
}
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](docs/CONTRIBUTING.md) for more information.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.
