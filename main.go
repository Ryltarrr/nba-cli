package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Ryltarrr/nba-cli/commands"
	"github.com/Ryltarrr/nba-cli/displayer"
	"github.com/Ryltarrr/nba-cli/parser"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Enter     key.Binding
	Escape    key.Binding
	Help      key.Binding
	Quit      key.Binding
	Backspace key.Binding
	Dash      key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter, k.Escape},
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↵", "enter"),
	),
	Escape: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "hide results"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Backspace: key.NewBinding(
		key.WithKeys(tea.KeyBackspace.String()),
	),
	Dash: key.NewBinding(
		key.WithKeys("-"),
	),
}

type model struct {
	displayer   displayer.Model
	textInput   textinput.Model
	spinner     spinner.Model
	loading     bool
	showResults bool
	keys        keyMap
	help        help.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 10
	ti.Width = 20
	ti.Placeholder = time.Now().Format(commands.DATE_FORMAT)
	ti.SetValue("2022-01-17")

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	d := displayer.New()

	return model{
		textInput:   ti,
		displayer:   d,
		spinner:     s,
		loading:     false,
		showResults: false,
		keys:        keys,
		help:        help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

// TODO: toggle input on command selection
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmdTextInput, cmdSpinner tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Escape):
			m.showResults = false
			m.textInput.Focus()
			return m, textinput.Blink

		case key.Matches(msg, m.keys.Enter):
			m.loading = true
			return m, commands.GetGamesForDateCommand(m.textInput.Value())

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil

		case key.Matches(msg, m.keys.Backspace):
			m.textInput, cmdTextInput = m.textInput.Update(msg)
			return m, cmdTextInput

		case !key.Matches(msg, m.keys.Dash):
			m.textInput, cmdTextInput = m.textInput.Update(msg)
			newVal, newPos := autoComplete(m.textInput.Value())
			m.textInput.SetValue(newVal)
			m.textInput.SetCursor(newPos)
			return m, cmdTextInput

		}

	case parser.Results:
		m.displayer.Data = msg
		m.loading = false
		m.showResults = true
		m.textInput.Blur()
		return m, nil
	}

	m.spinner, cmdSpinner = m.spinner.Update(msg)
	return m, cmdSpinner
}

func autoComplete(s string) (string, int) {
	count := len(s)
	if count == 4 || count == 7 {
		return s + "-", count + 2
	}
	return s, count
}

func (m model) View() string {
	s := ""
	if !m.showResults {
		s += "Date of the game:\n"
	}
	padding := lipgloss.NewStyle().Padding(1)

	if m.loading {
		s += fmt.Sprintf("%s\n", padding.Render(m.spinner.View()))
	}

	if m.showResults {
		s += m.displayer.View()
	}

	if !m.showResults && !m.loading {
		s += fmt.Sprintf("%s\n", padding.Render(m.textInput.View()))
	}

	helpMargin := lipgloss.NewStyle().
		MarginTop(3)
	helpView := m.help.View(m.keys)
	s += fmt.Sprintf("%s\n", helpMargin.Render(helpView))

	return s
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
