# CUI (Commandline User Interface)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/erdaltsksn/cui)](https://pkg.go.dev/github.com/erdaltsksn/cui)
![Go](https://github.com/erdaltsksn/cui/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/cui)](https://goreportcard.com/report/github.com/erdaltsksn/cui)

CUI gives us a set of useful tools while developing cli apps.

## Features

- Alerts
- Pre-made Cobra commands

## Requirements

- [Golang](https://golang.org)

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

Check out [examples](examples) directory for more.

**Output:**

![output](/assets/output.png)

## Installation

```sh
go get github.com/erdaltsksn/cui
```

## Updating / Upgrading

```sh
go get -u github.com/erdaltsksn/cui
```

## Usage

### Alerts

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

### Version Command

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
welcome. See [CONTRIBUTING](.github/CONTRIBUTING.md) for more information.

## Security Policy

If you discover a security vulnerability within this project, please follow our
[Security Policy Guide](.github/SECURITY.md).

## Code of Conduct

This project adheres to the Contributor Covenant [Code of Conduct](.github/CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.
