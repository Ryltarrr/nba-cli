package displayer

import (
	"fmt"

	"github.com/Ryltarrr/nba-cli/parser"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

const DRAW string = "DRAW"

type Model struct {
	Data parser.Results
}

func New() Model {
	return Model{}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	s := ""
	for _, game := range m.Data.Scoreboard.Games {
		awayTeam := game.AwayTeam
		homeTeam := game.HomeTeam
		awayColor := teamColors[awayTeam.TeamTricode].SprintFunc()
		homeColor := teamColors[homeTeam.TeamTricode].SprintFunc()

		// prints teams with their colors
		s += fmt.Sprintf("%s @ %s\n",
			awayColor(leftRightPad(awayTeam.TeamTricode)),
			homeColor(leftRightPad(homeTeam.TeamTricode)),
		)

		// prints scores, the winner is bold
		winnerTricode := getWinnerTricode(awayTeam, homeTeam)
		winnerScoreColor := color.New(color.Bold)
		if winnerTricode == DRAW {
			s += fmt.Sprintf("%4d  -  %d \n", awayTeam.Score, homeTeam.Score)
		} else if winnerTricode == awayTeam.TeamTricode {
			s += winnerScoreColor.Sprintf("%4d", awayTeam.Score)
			s += fmt.Sprintf("  -  %d \n", homeTeam.Score)
		} else {
			s += fmt.Sprintf("%4d  -  ", awayTeam.Score)
			s += winnerScoreColor.Sprintf("%d\n", homeTeam.Score)
		}
		s += "\n"
	}
	return s
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
