package main

import (
	"flag"
	"log"
	"time"

	"github.com/Ryltarrr/go-nba/displayer"
	"github.com/Ryltarrr/go-nba/fetcher"
	"github.com/Ryltarrr/go-nba/parser"
)

func main() {
	dt := time.Now()
	date := flag.String("date", dt.Format("2006-01-02"), "the date")
	flag.Parse()
	var fetcher fetcher.Fetcher
	body := fetcher.GetGamesForDate(*date)
	var parser parser.Parser
	results, err := parser.ParseResults(body)
	if err != nil {
		log.Fatalln("Error while parsing results", err)
	}
	displayer.DisplayGameResults(results)
}
