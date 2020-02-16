package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/russross/blackfriday"
	doku "github.com/seankhliao/blackfriday-doku"
)

const (
	es = blackfriday.Tables | blackfriday.FencedCode | blackfriday.Autolink | blackfriday.Strikethrough | blackfriday.SpaceHeadings | blackfriday.BackslashLineBreak
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("read file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", blackfriday.Run(b, blackfriday.WithRenderer(doku.NewRenderer()), blackfriday.WithExtensions(es)))
}
