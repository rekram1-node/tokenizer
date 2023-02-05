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

Using the benchmark test, you can see that even with 1 million words we still only have a meager 39 allocations per operation. I believe this number can be beat as well but I am still looking into the feasibility of some operations

```text
task: [bench] go test -bench=. ./tokenizer -run=^# -count=10 -benchmem | tee preformance.txt
goos: darwin
goarch: arm64
pkg: github.com/rekram1-node/tokenizer/tokenizer
BenchmarkTokenize/word_count:_10-8                893482              1302 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                909124              1309 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                908605              1298 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                905206              1270 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                900795              1254 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                912558              1267 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                928588              1255 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                885170              1264 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                894373              1253 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_10-8                917389              1258 ns/op             496 B/op          5 allocs/op
BenchmarkTokenize/word_count:_100-8                81775             13954 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                85268             13914 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                85263             13912 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                85602             13973 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                85400             13908 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                85245             13941 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                85260             13896 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                85585             14014 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                85687             13925 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_100-8                84966             13921 ns/op            4080 B/op          8 allocs/op
BenchmarkTokenize/word_count:_1000-8                8143            143447 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                7996            144030 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                8072            144614 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                8059            143350 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                8186            143046 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                8046            144263 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                8026            143982 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                8038            144174 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                8079            144182 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_1000-8                8030            142584 ns/op           50416 B/op         12 allocs/op
BenchmarkTokenize/word_count:_10000-8                795           1447904 ns/op          685299 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                784           1460662 ns/op          685299 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                824           1459105 ns/op          685300 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                818           1462055 ns/op          685299 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                805           1454443 ns/op          685299 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                818           1458763 ns/op          685300 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                819           1453174 ns/op          685299 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                811           1452126 ns/op          685298 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                823           1450348 ns/op          685299 B/op         19 allocs/op
BenchmarkTokenize/word_count:_10000-8                817           1457028 ns/op          685298 B/op         19 allocs/op
BenchmarkTokenize/word_count:_100000-8                70          15327946 ns/op         8942875 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                75          15442076 ns/op         8942871 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                76          15366360 ns/op         8942882 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                75          15329089 ns/op         8942887 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                74          15337678 ns/op         8942874 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                74          15398484 ns/op         8942891 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                74          15302757 ns/op         8942877 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                75          15362398 ns/op         8942871 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                74          15387052 ns/op         8942882 B/op         29 allocs/op
BenchmarkTokenize/word_count:_100000-8                75          15310445 ns/op         8942881 B/op         29 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         171073882 ns/op        88036674 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         169618688 ns/op        88036672 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         169880493 ns/op        88036592 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         169492944 ns/op        88036672 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         170651521 ns/op        88036640 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         169312208 ns/op        88036624 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         169648021 ns/op        88036640 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         170238188 ns/op        88036624 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         170449750 ns/op        88036656 B/op         39 allocs/op
BenchmarkTokenize/word_count:_1000000-8                6         169968132 ns/op        88036592 B/op         39 allocs/op
PASS
ok      github.com/rekram1-node/tokenizer/tokenizer     73.772s
```