package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ryltarrr/nba-cli/commands"
	"github.com/Ryltarrr/nba-cli/parser"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	result      parser.Results
	textInput   textinput.Model
	spinner     spinner.Model
	loading     bool
	showResults bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 10
	ti.Width = 20
	ti.Placeholder = "2022-03-27"

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		result:      parser.Results{},
		textInput:   ti,
		spinner:     s,
		loading:     false,
		showResults: false,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

// TODO: toggle input on command selection
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.Type {

		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyEsc:
			m.showResults = false
			return m, nil

		case tea.KeyEnter:
			m.loading = true
			return m, commands.GetGamesForDateCommand(m.textInput.Value())
		}

	case parser.Results:
		m.result = msg
		m.loading = false
		m.showResults = true
		return m, nil
	}

	var cmdTextInput, cmdSpinner tea.Cmd
	m.textInput, cmdTextInput = m.textInput.Update(msg)
	m.spinner, cmdSpinner = m.spinner.Update(msg)

	return m, tea.Batch(cmdTextInput, cmdSpinner)
}

func (m model) View() string {
	s := "Date of the game:\n"

	if m.loading {
		s += fmt.Sprintf("%s\n", m.spinner.View())
	}

	resultString, err := m.result.StringifyResults()
	if err != nil {
		log.Fatal("error", err)
	}

	if m.showResults {
		s += fmt.Sprintf("\n%s\n", resultString)
	}

	if !m.showResults && !m.loading {
		s += fmt.Sprintf("%s\n", m.textInput.View())
	}

	s += "\nPress ctrl+c to quit.\n"

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
