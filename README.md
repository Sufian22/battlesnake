# Battlesnake

## Prerequisites

- Go: https://golang.org/dl/
- cli: https://github.com/BattlesnakeOfficial/rules
- [optional] make:
  - MacOS: https://formulae.brew.sh/formula/make
  - Ubuntu: https://zoomadmin.com/HowToInstall/UbuntuPackage/make
  - Windows: http://gnuwin32.sourceforge.net/packages/make.htm
- [optional] Docker: https://docs.docker.com/engine/install/

## Usage

### Running with Docker

In order to execute the server using Docker it will be as simple as executing:

```cmd
make run-docker
```

### Running with Go

In order to execute the server using Go compiler, you can just execute the following command:

```bash
go run cmd/battlesnake/main.go
```

I've enabled the specification of the configuration file using a flag:

```bash
Usage:
  -config string
        Battlesnake server configuration (default "config.json")
```

Alternatively, you can use the makefile to execute the program:

```bash
CONFIG_FILE="config.json" make run
```

If you want to see all the commands available in the Makefile execute:

```bash
make
```

Once executed, you can use the CLI to start playing:

```bash
battlesnake play -W 8 -H 8 --name my-snake --url http://localhost:3000 -g solo -v
```

## Tests

Tests have been implemented, in order to execute them and see the coverage just run:

```cmd
make test
```
