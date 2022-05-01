package displayer

import (
	"fmt"

	"github.com/Ryltarrr/go-nba/parser"
)

func DisplayGameResults(results parser.Results) {
	for _, el := range results.Scoreboard.Games {
		fmt.Printf("%v @ %v\n", el.AwayTeam.TeamTricode, el.HomeTeam.TeamTricode)
		fmt.Printf("%v - %v\n", el.AwayTeam.Score, el.HomeTeam.Score)
	}
}
