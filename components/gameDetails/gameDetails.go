package gamedetails

import (
	"fmt"

	"github.com/Ryltarrr/nba-cli/parser"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	return lipgloss.JoinVertical(lipgloss.Left, m.getPeriods(), m.getLeaders())
}

func (m *Model) SetGame(g parser.Game) {
	m.game = g
}

func (m Model) getPeriods() string {
	headerStr := lipgloss.NewStyle().Underline(true).Bold(true).Render("Score by quarter:")

	homePeriods, awayPeriods := "", ""
	for i := 0; i < m.game.Period; i++ {
		homeStyle := lipgloss.NewStyle().Bold(false)
		awayStyle := homeStyle.Copy()
		homeScore := m.game.HomeTeam.Periods[i].Score
		awayScore := m.game.AwayTeam.Periods[i].Score

		if homeScore > awayScore {
			homeStyle = homeStyle.Bold(true)
		} else {
			awayStyle = awayStyle.Bold(true)
		}

		homePeriods += homeStyle.Render(fmt.Sprintf("%v ", homeScore))
		awayPeriods += awayStyle.Render(fmt.Sprintf("%v ", awayScore))
	}

	return headerStr + "\n" + homePeriods + "\n" + awayPeriods + "\n"
}

func (m Model) getLeaders() string {
	headerStr := lipgloss.NewStyle().Underline(true).Bold(true).Render("Best players:")
	homeLeaderView := leaderView(m.game.GameLeaders.HomeLeaders)
	awayLeaderView := leaderView(m.game.GameLeaders.AwayLeaders)

	return lipgloss.JoinVertical(lipgloss.Top, headerStr, homeLeaderView, "", awayLeaderView)
}

func leaderView(player parser.Player) string {
	bold := lipgloss.NewStyle().Bold(true)

	s := bold.Render("#"+player.JerseyNum) + " " + player.Name + "\n"
	s += fmt.Sprintf("%v pts, %v rbd, %v ast", player.Points, player.Rebounds, player.Assists)

	return s
}
