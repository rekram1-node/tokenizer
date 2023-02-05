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
