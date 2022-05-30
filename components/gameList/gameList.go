package gameList

import (
	"fmt"

	"github.com/Ryltarrr/nba-cli/parser"
	"github.com/Ryltarrr/nba-cli/utils"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const useHighPerformanceRenderer = false
const numLinesPerGameResult = 3

type Model struct {
	data     parser.Results
	Spinner  spinner.Model
	Loading  bool
	cursor   int
	ready    bool
	viewport viewport.Model
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
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {

	case tea.KeyMsg:
		numGames := len(m.data.Scoreboard.Games)
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.viewport.Width < numGames*numLinesPerGameResult {
					return m, nil
				}
			}

		case "down", "j":
			if m.cursor < len(m.data.Scoreboard.Games)-1 {
				m.cursor++
				if m.viewport.Width < numGames*numLinesPerGameResult {
					return m, nil
				}
			}

		}

	case tea.WindowSizeMsg:
		verticalMarginHeight := 1
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = 0
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

	case parser.Results:
		m.data = msg
		m.Loading = false
		m.viewport.SetContent(m.getContent())
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	m.Spinner, cmd = m.Spinner.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.Loading {
		return lipgloss.NewStyle().Padding(1).Render(m.Spinner.View()) + "\n"
	} else {
		glStyle := lipgloss.NewStyle().MarginLeft(5)
		m.viewport.SetContent(m.getContent())
		return glStyle.Render(m.viewport.View())
	}
}

func (m Model) getContent() string {
	s := "Results:\n"
	for idx, game := range m.data.Scoreboard.Games {
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
	return s
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
