package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var client = http.Client{}

func GetGamesForDate(date string) []byte {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://stats.nba.com/stats/scoreboardv3?GameDate=%s&LeagueID=00", date),
		nil,
	)
	if err != nil {
		log.Fatalln("Error while creating request")
	}
	req.Header.Add("Referer", "https://www.nba.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "en-US")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error while fetching games of", date)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error while reading data", err)
	}
	res.Body.Close()

	return body
}
