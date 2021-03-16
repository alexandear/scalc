# Sets Calculator

A set calculator that prints to stdout an evaluated expression.

Grammar of calculator is given:
```
expression := "[" operator N sets "]"
sets := set | set sets
set := file | expression
operator := "EQ" | "LE" | "GR"
```

Where:
- `file` is a file with sorted integers, one integer in a line.
- `N` is a positive integer

Meaning of operators:
- `EQ` - returns a set of integers which consists only from values which exists in exactly N sets - arguments of operator.
- `LE` - returns a set of integers which consists only from values which exists in less then N sets - arguments of operator.
- `GR` - returns a set of integers which consists only from values which exists in more then N sets - arguments of operator.

Example:
```shell
./scalc [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]
2
3

./scalc [ LE 2 a.txt [ GR 1 b.txt c.txt ] ]
1
4

cat a.txt
1
2
3

cat b.txt
2
3
4

cat c.txt
1
2
3
4
5
```

## Development

### How to Build

> Prerequisites: go 1.16, make, docker

With `Go`, binary `./scalc`:

```shell
go build -o ./scalc ./cmd/...
```

Or with `make`, binary `./bin/scalc`:

```shell
make build
```

Or `docker` image `scalc`:

```shell
docker build -t scalc .
```

### How to Run

Built binary:

```shell
./scalc [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]
```

Or with predefined `make` targets:

```shell
make example1

make example2
```

Or with `docker`:

```shell
docker run --rm scalc [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]
```

### How to Execute Unit Tests

```shell
make test
```

### How to Execute Integration Tests

```shell
make test-it
```

### How to Run Linter

```shell
make lint
```
