# Golang Basic

This repository for personal learning material. Learning basic programming with Go Language.

## Dev Environment

* Go version go1.14.3 windows/amd64

## Run

* Build `go build filename.go`
* Or Run without build `go run filename.go`
* Run `.\filename.go`

## Test

* Run `go test` or `go test -v`
* Benchmark: on powershell (windows) `go test -bench="."`
* Test coverage: `go test -cover`
* Run spesific column (Table Driven Test), example: `go test -run FunctionName/Name`
* DO `errcheck .` for check error handling scenario we have not tested.
