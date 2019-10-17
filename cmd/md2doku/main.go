package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/russross/blackfriday"
	doku "github.com/seankhliao/blackfriday-doku"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	es = blackfriday.Tables | blackfriday.FencedCode | blackfriday.Autolink | blackfriday.Strikethrough | blackfriday.SpaceHeadings | blackfriday.BackslashLineBreak
)

func main() {
	initLog()

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal().Err(err).Msg("read file")
	}

	b = blackfriday.Run(b, blackfriday.WithRenderer(doku.NewRenderer()), blackfriday.WithExtensions(es))

	fmt.Printf("\n%s\n", b)

}

func initLog() {
	logfmt := os.Getenv("LOGFMT")
	if logfmt != "json" {
		logfmt = "text"
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: !terminal.IsTerminal(int(os.Stdout.Fd()))})
	}

	level, _ := zerolog.ParseLevel(os.Getenv("LOGLVL"))
	if level == zerolog.NoLevel {
		level = zerolog.InfoLevel
	}
	log.Info().Str("FMT", logfmt).Str("LVL", level.String()).Msg("log initialized")
	zerolog.SetGlobalLevel(level)
}
