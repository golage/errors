# Golage errors handling ![https://github.com/golage/errors](https://github.com/golage/errors/workflows/Go/badge.svg) ![https://github.com/golage/errors](https://codecov.io/gh/golage/errors/branch/master/graph/badge.svg)
Simple and useful error handling package in Golang

## Features
- Fundamental
- Support stacktrace
- Wrap existing errors
- Easy and pretty error handling 
- Support serialization (for transports in grpc, api, etc...)
- Support stacktrace errors in [github.com/pkg/errors](github.com/pkg/errors)
- Handling error types with code numbers (you can extends with constants)

## Installation
Get from Github:
```bash
go get github.com/golage/errors
```

## How to use
Import into your code:
```go
import "github.com/golage/errors"
```
Create error instance:
```go
errors.New(errors.NotFound, "somethings")
errors.New(errors.NotFound, "somethings %v", 123)
```
Wrap existing error:
```go
errors.Wrap(err, errors.Internal, "somethings")
errors.Wrap(err, errors.Internal, "somethings %v", 123)
```
Handle error:
```go
switch err, code := errors.Parse(err); code {
case errors.Nil:
case errors.NotFound:
    log.Fatalf("not found: %v", err)
default:
    log.Fatalf("others: %v\n%v", err, err.StackTrace())
}
```
For more see [example](examples/main.go)
