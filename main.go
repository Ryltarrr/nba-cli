package main

import (
	"fmt"
	"os"

	"github.com/Ryltarrr/nba-cli/components/gameList"
	"github.com/Ryltarrr/nba-cli/components/help"
	"github.com/Ryltarrr/nba-cli/components/menu"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	gameList gameList.Model
	menu     menu.Model
	help     help.Model
}

func initialModel() model {
	return model{
		menu:     menu.New(),
		gameList: gameList.New(),
		help:     help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.menu.Init(), m.gameList.Init())
}

// TODO: toggle input on command selection
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			cmds = append(cmds, tea.Quit)
		}
	}

	m.help, cmd = m.help.Update(msg)
	cmds = append(cmds, cmd)
	m.gameList, cmd = m.gameList.Update(msg)
	cmds = append(cmds, cmd)
	m.menu, cmd = m.menu.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	menuView := m.menu.View()
	gameListView := m.gameList.View()
	helpView := m.help.View()

	leftColumn := lipgloss.JoinVertical(lipgloss.Top, menuView, helpView)
	return lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, gameListView)
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
