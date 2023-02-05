# Tokenizer

[![Go Report](https://goreportcard.com/badge/github.com/rekram1-node/tokenizer)](https://goreportcard.com/report/github.com/rekram1-node/tokenizer) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/rekram1-node/tokenizer/blob/main/LICENSE) ![Build Status](https://github.com/rekram1-node/tokenizer/actions/workflows/main.yml/badge.svg)

Natural Language Processing (NLP) Tokenization Libary designed for English. Fast, Lean, Customizable. Tokenizes text, replaces abbreviations, replaces contractions, lowercases words, optionally you can remove stop words as well (must specify, see [usage](#usage)). This library is a work in progress but all features mentioned in README should be working as advertised.

Tokenizing text is one of the first steps in NLP, this preformant library should help Go users get started with their NLP tasks. 

## Features

* [Convert Text to Tokens](#usage)
* [Replace Contractions](#usage)
* [Replace Abbreviations](#usage)
* [Remove Stop Words](#usage) - defaults to keeping them
* [Blazingly Fast](#benchmarks)
* [Low Allocation](#benchmarks)
* Practically zero dependency (only dependency is [testify](https://github.com/stretchr/testify) for unit testing)

Coming soon:
- Streamed Reading
- Sentence Tokenization

## Installation

```bash
go get -u github.com/rekram1-node/tokenizer
```

## Usage

For detailed examples see [examples](https://github.com/rekram1-node/tokenizer/examples)

### Default

```go
package main

import (
    "fmt"
    "strings"

    "github.com/rekram1-node/tokenizer/tokenizer"
)

func main() {
    myStr := "This is my long string! I can replace contractions like can't or they've! I can replace abbreviations such as: demonstr. or jan."
	t := tokenizer.New()
	tokens := t.TokenizeString(myStr)
	fmt.Println(tokens)
	// Output: [this is my long string i can replace contractions like cannot or they have i can replace abbreviations such as demonstration or jan.]

	/*
		Note: you can remove stop words too!!!
	*/
	t.SetStopWordRemoval(true)
	myOtherStr := "This is another string to demonstrate stop words removal, words like: and or but the are are all stop word examples"
	tokens = t.TokenizeString(myOtherStr)
	fmt.Println(tokens)
	// Output: [string demonstrate stop words removal words stop word examples]
}

// Output: [the world is a wonderful place there are many places in this world 123 this place is wonderful]
```

### Custom Settings

```go
package main

import (
	"fmt"
	"log"

	"github.com/rekram1-node/tokenizer/languages"
	"github.com/rekram1-node/tokenizer/tokenizer"
)

func main() {
	// add a string containing all the "separators" you want
	// important note: including the "." would degrade the ability to replace abbreviations
	customSeparators := "\t\n\r ,:?\"!;()"

	// specify your settings here
	settings := &tokenizer.Settings{
		KeepSeparators:  false,
		RemoveStopWords: true,
		// you can have your own language configuration, see the language struct
		/*
			type Lanuage struct {
				StopWords     map[string]uint8
				Contractions  map[string]string
				Abbreviations map[string]string
			}

			you can create your own using languages.NewLanguage(yourStopWords, yourContractions, yourAbbreviations)
		*/
		Lanuage: languages.English,
	}
	// custom settings return an error incase of a misconfigured/missing setting
	t, err := settings.Custom(customSeparators)
	if err != nil {
		log.Fatal(err)
	}

	myStr := "This is my long string! I can replace contractions like can't or they've! I can replace abbreviations such as: demonstr. or jan. This is another string to demonstrate stop words removal, words like: and or but the are are all stop word examples"
	tokens := t.TokenizeString(myStr)
	fmt.Println(tokens)
	// Output: [long string replace contractions replace abbreviations demonstration january string demonstrate stop words removal words stop word examples]
}
```

## Benchmarks

See [benchmark test](https://github.com/rekram1-node/tokenizer/blob/main/tokenizer/benchmark_test.go)

Using the benchmark test, you can see that even with 1 million words we still only have a meager 39 allocations per operation.

```text
task: [bench] go test -bench=. ./tokenizer -run=^# -count=10 -benchmem | tee preformance.txt
goos: darwin
goarch: arm64
pkg: github.com/rekram1-node/tokenizer/tokenizer
BenchmarkTokenize/word_count:_10-8               1000000              1154 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1153 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1155 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1159 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1152 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1154 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1122 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1115 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1110 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8               1000000              1107 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_100-8               108454             11099 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               108898             10994 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               108961             10968 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               108924             11003 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               109113             10976 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               107854             11127 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               107070             11025 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               108572             11086 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               104055             11044 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8               108446             11075 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            109167 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            109177 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            109985 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            110084 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            109397 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            109133 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            109116 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            109918 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            110418 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8               10000            109553 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_10000-8               1066           1114218 ns/op          685302 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1074           1117825 ns/op          685298 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1074           1116659 ns/op          685299 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1078           1112335 ns/op          685301 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1081           1107204 ns/op          685300 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1076           1111695 ns/op          685300 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1082           1122813 ns/op          685300 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1068           1114729 ns/op          685300 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1063           1165390 ns/op          685299 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8               1030           1113300 ns/op          685300 B/op         19 allocs/op
BenchmarkTokenize/word_count:_100000-8                88          12112446 ns/op         8942866 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                97          12080584 ns/op         8942883 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                97          12054590 ns/op         8942882 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                97          12119115 ns/op         8942878 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                93          12124667 ns/op         8942873 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                97          12046165 ns/op         8942870 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                98          12015797 ns/op         8942871 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                98          12108281 ns/op         8942880 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                96          12039198 ns/op         8942886 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                98          12025329 ns/op         8942878 B/op         29 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         136513245 ns/op        88036641 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         137814833 ns/op        88036665 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         135903448 ns/op        88036664 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         136073953 ns/op        88036616 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         136058141 ns/op        88036664 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         136534161 ns/op        88036628 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         136569016 ns/op        88036908 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         136063990 ns/op        88036616 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         136175146 ns/op        88036628 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                8         137739057 ns/op        88036616 B/op         39 allocs/op
PASS
ok      github.com/rekram1-node/tokenizer/tokenizer     75.054s
```