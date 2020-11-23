oexec
=====
[![GitHub release](https://img.shields.io/github/release/amalfra/oexec.svg)](https://github.com/amalfra/oexec/releases)
[![Build Status](https://travis-ci.org/amalfra/oexec.svg?branch=master)](https://travis-ci.org/amalfra/oexec)
[![GoDoc](https://godoc.org/github.com/amalfra/oexec?status.svg)](https://godoc.org/github.com/amalfra/oexec)
[![Go Report Card](https://goreportcard.com/badge/github.com/amalfra/oexec)](https://goreportcard.com/report/github.com/amalfra/oexec)

A go package to execute shell commands in specified order. Currently supports executing list of shell commands in following orders:
* Series
* Parallel

## Installation
You can download the package using
```sh
go get github.com/amalfra/oexec
```
## Usage
Next, import the package
``` go
import (
  "github.com/amalfra/oexec"
)
```
You can execute list of shell commands in following orders:
* Series
* Parallel

The result of each command will be returned as ```oexec.Output``` struct which has fields
* Stdout - a ```byte``` array containing stdout produced by the command. Will be nil if command status is non zero
* Stderr - an ```error``` object containing stderr returned by the command. Will be nil if command status is zero  

### executing in series
To execute commands in series call the function as
``` go
oexec.Series("ls -l", "pwd")
```
You can pass any number of commands to execute as parameters. All the passed commands will get executed in series and results will be return once all of them are completed. The results are returned as an array of ```oexec.Output``` struct, having the result corresponding to a command at same position of argument as in Series function call(position starts are zero)

### executing in parallel
> Note: Precisely the commands will be executed concurrently unless configured otherwise via ```GOMAXPROCS``` environment variable and depending on underlying hardware

To execute commands in parallel call the function as
``` go
oexec.Parallel("ls -l", "pwd")
```
You can pass any number of commands to execute as parameters. All the passed commands will get executed in parallel and results will be return once all of them are completed. The results are returned as an array of ```oexec.Output``` struct, having the result corresponding to a command at same position of argument as in Series function call(position starts are zero)

## Development

Questions, problems or suggestions? Please post them on the [issue tracker](https://github.com/amalfra/oexec/issues).

You can contribute changes by forking the project and submitting a pull request. You can ensure the tests are passing by running ```make test```. Feel free to contribute :heart_eyes:

## UNDER MIT LICENSE

The MIT License (MIT)

Copyright (c) 2017 Amal Francis

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
