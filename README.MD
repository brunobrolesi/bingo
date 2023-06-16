# go-meliplaces-commons

![technology Go](https://img.shields.io/badge/technology-go-blue.svg?style=flat-square)

This is a CLI for bingo draw.

- [Getting Started](#Getting-Started)
- [Build](#Build)

## Getting Started

Download the repository using `git clone`:

```shell
git clone git@github.com:brunobrolesi/bingo.git
```

To run all the tests of this repository use the included `Makefile`.

```shell
make
```

## Build

To build the project use the included `Makefile`.

### Windows 64 bits

```shell
make build-windows
```

### MacOS Intel

```shell
make build-macos-intel
```

### MacOS ARM

```shell
make build-macos-arm
```

### Linux 64 bits

```shell
make build-linux
```

The binaries will be available in the `bin` folder. Run the binary acording to your OS.

