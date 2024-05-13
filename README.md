# l33tspeak

[![Go pkg.go.dev](https://pkg.go.dev/badge/github.com/deadpyxel/l33tspeak.svg)](https://pkg.go.dev/github.com/deadpyxel/l33tspeak)


A simple CLI written in go to translate to and from leetspeak so you can work with strings like an "1337 HaX0r".

## Features

- Simple command structure (inspired by `base64`)
- Read input from `STDIN` using pipes, or pass it directly as arguments
- Very small footprint (~2.2MiB binary)
- Cross platform


## Installation

Install `l33tspeak` with go

```bash
go install github.com/deadpyxel/workday@latest
```

Or for an specific version:
```bash
go install github.com/deadpyxel/workday@v0.4.0
```

## Documentation

WIP


## Running Tests

To run tests, run the following command

```bash
go test -cover -v ./...
```
If you want to run the benchmarks:

```bash
go test -bench=. -v ./...
```

## Developing Locally

Clone the project

```bash
git clone git@github.com:deadpyxel/l33tspeak.git
```
> If you cannot use SSH based cloning, use the url https://github.com/deadpyxel/l33tspeak.git instead

Go to the project directory

```bash
cd l33tspeak
```

Build the project locally

```bash
go build -o bin/ -v ./...
```

Run the app

```bash
./bin/l33tspeak
```


## Usage/Examples

Working with input from STDIN:
```bash
echo "Lorem ipsum dolor sit amet" | l33tspeak
# Results: l0R3m 1p5Um d0l0r 517 4M37
```

Working with input as arguments:
```bash
l33tspeak "Lorem ipsum dolor sit amet"
# Results: l0R3m 1p5Um d0l0r 517 4M37
```

Controlling probability of change (E.g 50% of chance the character is replaced):
```bash
l33tspeak -p 0.5 "Lorem ipsum dolor sit amet"
# Results: l0r3M IpSuM D0lOR si7 am3T
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
