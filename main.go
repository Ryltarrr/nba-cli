package main

import (
	"fmt"
	"os"

	"github.com/Ryltarrr/nba-cli/commands"
	"github.com/Ryltarrr/nba-cli/components/gameList"
	"github.com/Ryltarrr/nba-cli/components/menu"
	"github.com/Ryltarrr/nba-cli/parser"
	"github.com/charmbracelet/bubbles/help"
	Key "github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Enter     Key.Binding
	Escape    Key.Binding
	Help      Key.Binding
	Quit      Key.Binding
	Backspace Key.Binding
	Dash      Key.Binding
}

func (k keyMap) ShortHelp() []Key.Binding {
	return []Key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]Key.Binding {
	return [][]Key.Binding{
		{k.Enter, k.Escape},
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Enter: Key.NewBinding(
		Key.WithKeys("enter"),
		Key.WithHelp("↵", "enter"),
	),
	Escape: Key.NewBinding(
		Key.WithKeys("esc"),
		Key.WithHelp("esc", "hide results"),
	),
	Help: Key.NewBinding(
		Key.WithKeys("?"),
		Key.WithHelp("?", "toggle help"),
	),
	Quit: Key.NewBinding(
		Key.WithKeys("ctrl+c"),
		Key.WithHelp("ctrl+c", "quit"),
	),
	Backspace: Key.NewBinding(
		Key.WithKeys(tea.KeyBackspace.String()),
	),
	Dash: Key.NewBinding(
		Key.WithKeys("-"),
	),
}

type model struct {
	gameList gameList.Model
	menu     menu.Model
	keys     keyMap
	help     help.Model
}

func initialModel() model {
	gl := gameList.New()
	m := menu.New()

	return model{
		menu:     m,
		gameList: gl,
		keys:     keys,
		help:     help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.gameList.Spinner.Tick)
}

// TODO: toggle input on command selection
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmdSpinner, cmdGameList, cmdMenu tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {

		case Key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case Key.Matches(msg, m.keys.Escape):
			m.menu.Focused = true
			m.menu.TextInput.Focus()
			return m, textinput.Blink

		case m.menu.Focused && Key.Matches(msg, m.keys.Enter):
			m.gameList.Loading = true
			return m, commands.GetGamesForDateCommand(m.menu.TextInput.Value())

		case Key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}

	case parser.Results:
		m.gameList.Data = msg
		m.gameList.Loading = false
		m.menu.Focused = false
		m.menu.TextInput.Blur()
		return m, nil
	}

	m.gameList.Spinner, cmdSpinner = m.gameList.Spinner.Update(msg)
	m.gameList, cmdGameList = m.gameList.Update(msg)
	m.menu, cmdMenu = m.menu.Update(msg)
	return m, tea.Batch(cmdSpinner, cmdGameList, cmdMenu)
}

func (m model) View() string {
	menuView := m.menu.View()

	gameListView := m.gameList.View()

	helpStyle := lipgloss.NewStyle().MarginTop(1)
	helpView := helpStyle.Render(m.help.View(m.keys))

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
