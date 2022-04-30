package main

import (
	"flag"
	"time"

	"github.com/Ryltarrr/go-nba/fetcher"
)

func main() {
	dt := time.Now()
	date := flag.String("date", dt.Format("2006-01-02"), "the date")
	var fetcher fetcher.Fetcher
	fetcher.GetGamesForDate(*date)
}
