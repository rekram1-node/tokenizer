package main

import (
	"fmt"

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
