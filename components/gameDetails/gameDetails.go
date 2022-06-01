package gamedetails

import (
	"fmt"

	"github.com/Ryltarrr/nba-cli/parser"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	game parser.Game
}

func New() Model {
	return Model{
		game: parser.Game{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	s := ""
	s += m.game.GameStatusText + "\n"
	homePeriods, awayPeriods := "", ""
	for i := 0; i < m.game.Period; i++ {
		homePeriods += fmt.Sprintf("%v ", m.game.HomeTeam.Periods[i].Score)
		awayPeriods += fmt.Sprintf("%v ", m.game.AwayTeam.Periods[i].Score)
	}

	homePeriods += fmt.Sprintf(" | %v", m.game.HomeTeam.Score)
	awayPeriods += fmt.Sprintf(" | %v", m.game.AwayTeam.Score)
	s += homePeriods + "\n" + awayPeriods
	s += "\n"
	return s
}

func (m *Model) SetGame(g parser.Game) {
	m.game = g
}
