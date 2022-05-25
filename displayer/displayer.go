package displayer

import (
	"fmt"

	"github.com/Ryltarrr/nba-cli/parser"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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
	s := "Results:\n"
	for _, game := range m.Data.Scoreboard.Games {
		awayTeam := game.AwayTeam
		homeTeam := game.HomeTeam
		awayColor := teamColors[awayTeam.TeamTricode]
		homeColor := teamColors[homeTeam.TeamTricode]

		s += awayColor.Render(awayTeam.TeamTricode)
		s += " @ "
		s += homeColor.Render(homeTeam.TeamTricode)

		awayBlock, homeBlock := getScoreBlocks(fmt.Sprint(awayTeam.Score), fmt.Sprint(homeTeam.Score))
		s += fmt.Sprintf("\n%s - %s\n\n", awayBlock, homeBlock)
	}
	return s
}

func getScoreBlocks(awayScore string, homeScore string) (string, string) {
	scoreBlock := lipgloss.NewStyle().
		Width(5).
		Align(0.5)

	if awayScore == homeScore {
		return scoreBlock.Bold(false).Render(awayScore), scoreBlock.Bold(false).Render(homeScore)
	} else if awayScore > homeScore {
		return scoreBlock.Bold(true).Render(awayScore), scoreBlock.Bold(false).Render(homeScore)
	} else {
		return scoreBlock.Bold(false).Render(awayScore), scoreBlock.Bold(true).Render(homeScore)
	}
}
