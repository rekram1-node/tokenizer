# Tokenizer

[![Go Report](https://goreportcard.com/badge/github.com/rekram1-node/tokenizer)](https://goreportcard.com/report/github.com/rekram1-node/tokenizer) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/rekram1-node/tokenizer/blob/main/LICENSE) ![Build Status](https://github.com/rekram1-node/tokenizer/actions/workflows/main.yml/badge.svg)

NLP Tokenization Libary designed for English. Fast, Lean, Customizable

## Features

* [Blazingly Fast](#benchmarks)
* [Low to zero allocation](#benchmarks)
* [Convert Text to Tokens](#usage)

## Installation

```bash
go get -u github.com/rs/zerolog/log
```

## Usage

```go
package main

import (
    "fmt"
    "strings"

    "github.com/rekram1-node/tokenizer/tokenizer"
)

func main() {
    s := `
    The world is a wonderful place...
    There are many places in this world!!!	

    123

    This place is wonderful.
    `
    s = strings.TrimSpace(s)
    t := tokenizer.New()
    tokens := t.TokenizeByWord(s)
    fmt.Println(tokens)
}

// Output: [The world is a wonderful place There are many places in this world 123 This place is wonderful]
```

## Benchmarks

See [benchmark test](https://github.com/rekram1-node/tokenizer/blob/main/tokenizer/benchmark_test.go)

```text
goos: darwin
goarch: arm64
pkg: github.com/rekram1-node/tokenizer/tokenizer
BenchmarkTokenize-8     1000000000               0.001590 ns/op        0 B/op          0 allocs/op
PASS
ok      github.com/rekram1-node/tokenizer/tokenizer     0.185s
```