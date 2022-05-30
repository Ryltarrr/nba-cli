package help

import (
	"github.com/charmbracelet/bubbles/help"
	Key "github.com/charmbracelet/bubbles/key"
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

type Model struct {
	help help.Model
	keys keyMap
}

func New() Model {
	return Model{
		help: help.New(),
		keys: keys,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case Key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, nil
}

func (m Model) View() string {
	helpStyle := lipgloss.NewStyle().MarginTop(1)
	return helpStyle.Render(m.help.View(m.keys))
}
