package gameList

import (
	"fmt"

	"github.com/Ryltarrr/nba-cli/parser"
	"github.com/Ryltarrr/nba-cli/utils"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Data     parser.Results
	Selected parser.Game
	Spinner  spinner.Model
	Loading  bool
	cursor   int
}

func New() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return Model{
		cursor:  0,
		Spinner: s,
		Loading: false,
	}
}

func (m Model) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.Data.Scoreboard.Games)-1 {
				m.cursor++
			}

		case "enter":
			m.Loading = true

		}

	case parser.Results:
		m.Data = msg
		m.Loading = false
	}

	var cmdSpinner tea.Cmd
	m.Spinner, cmdSpinner = m.Spinner.Update(msg)
	return m, cmdSpinner
}

func (m Model) View() string {
	s := ""
	if m.Loading {
		s += lipgloss.NewStyle().Padding(1).Render(m.Spinner.View()) + "\n"
	} else {
		s += "Results:\n"
		for idx, game := range m.Data.Scoreboard.Games {
			awayTeam := game.AwayTeam
			homeTeam := game.HomeTeam
			awayColor := utils.TeamColors[awayTeam.TeamTricode]
			homeColor := utils.TeamColors[homeTeam.TeamTricode]

			selected := false
			if idx == m.cursor {
				selected = true
			}
			gameStyle := lipgloss.NewStyle().
				MarginBottom(1).
				MarginLeft(1).
				Faint(!selected).
				Border(lipgloss.NormalBorder(), false, false, false, selected)
			gameStr := ""
			gameStr += awayColor.Faint(!selected).Render(awayTeam.TeamTricode)
			gameStr += " @ "
			gameStr += homeColor.Faint(!selected).Render(homeTeam.TeamTricode)

			awayBlock, homeBlock := getScoreBlocks(awayTeam.Score, homeTeam.Score)
			gameStr += fmt.Sprintf("\n%s - %s", awayBlock, homeBlock)
			s += gameStyle.Render(gameStr) + "\n"
		}
	}

	glStyle := lipgloss.NewStyle().MarginLeft(5)
	return glStyle.Render(s)
}

func getScoreBlocks(awayScore int, homeScore int) (string, string) {
	scoreBlock := lipgloss.NewStyle().
		Width(5).
		Align(0.5)

	awayStr := fmt.Sprint(awayScore)
	homeStr := fmt.Sprint(homeScore)

	if awayScore == homeScore {
		return scoreBlock.Bold(false).Render(awayStr), scoreBlock.Bold(false).Render(homeStr)
	} else if awayScore > homeScore {
		return scoreBlock.Bold(true).Render(awayStr), scoreBlock.Bold(false).Render(homeStr)
	} else {
		return scoreBlock.Bold(false).Render(awayStr), scoreBlock.Bold(true).Render(homeStr)
	}
}
