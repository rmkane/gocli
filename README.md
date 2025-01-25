<!-- omit in toc -->
# Go CLI Application

This is a simple CLI application built with Go and the `cobra` package. It includes three commands: `dump`, `encode`, and `decode`.

<!-- omit in toc -->
## Table of contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Run locally](#run-locally)
- [Building the project](#building-the-project)
- [Running the CLI](#running-the-cli)
- [Commands](#commands)
  - [Dump command](#dump-command)
  - [Encode command](#encode-command)
  - [Decode command](#decode-command)
- [License](#license)

## Prerequisites

- Go 1.23.5 or later

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/rmkane/go-cli.git
   cd go-cli
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

## Run locally

```sh
go run ./cmd/gocli
```

## Building the project

To build the project as a CLI binary, run the following command:

```sh
go build -o ./bin/gocli ./cmd/gocli
```

This will create an executable binary named gocli in the current directory.

## Running the CLI

You can run the CLI application using the following command:

```sh
./gocli
```

On Windows, use:

```sh
gocli.exe
```

## Commands

### Dump command

The dump command takes positional arguments and an optional --verbose flag:

### Encode command

The encode command takes positional arguments:

### Decode command

The decode command takes positional arguments:

## License

This project is licensed under the MIT License.
