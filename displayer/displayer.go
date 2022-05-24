package displayer

import (
	"fmt"

	"github.com/Ryltarrr/nba-cli/parser"
	"github.com/fatih/color"
)

const DRAW string = "DRAW"

func DisplayGameResults(results parser.Results) {
	for _, game := range results.Scoreboard.Games {
		awayTeam := game.AwayTeam
		homeTeam := game.HomeTeam
		awayColor := teamColors[awayTeam.TeamTricode].SprintFunc()
		homeColor := teamColors[homeTeam.TeamTricode].SprintFunc()

		// prints teams with their colors
		fmt.Printf("%s @ %s\n",
			awayColor(leftRightPad(awayTeam.TeamTricode)),
			homeColor(leftRightPad(homeTeam.TeamTricode)),
		)

		// prints scores, the winner is bold
		winnerTricode := getWinnerTricode(awayTeam, homeTeam)
		winnerScoreColor := color.New(color.Bold)
		if winnerTricode == DRAW {
			fmt.Printf("%4d  -  %d \n", awayTeam.Score, homeTeam.Score)
		} else if winnerTricode == awayTeam.TeamTricode {
			winnerScoreColor.Printf("%4d", awayTeam.Score)
			fmt.Printf("  -  %d \n", homeTeam.Score)
		} else {
			fmt.Printf("%4d  -  ", awayTeam.Score)
			winnerScoreColor.Printf("%d\n", homeTeam.Score)
		}
		fmt.Println()
	}
}

func getWinnerTricode(awayTeam parser.Team, homeTeam parser.Team) string {
	if awayTeam.Score == homeTeam.Score {
		return DRAW
	} else if awayTeam.Score > homeTeam.Score {
		return awayTeam.TeamTricode
	} else {
		return homeTeam.TeamTricode
	}
}

func leftRightPad(tricode string) string {
	return " " + tricode + " "
}
