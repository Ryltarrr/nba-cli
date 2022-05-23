package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Ryltarrr/go-nba/fetcher"
	"github.com/Ryltarrr/go-nba/parser"
	tea "github.com/charmbracelet/bubbletea"
)

type choice struct {
	text  string
	value string
}

type model struct {
	choices  []choice
	cursor   int
	selected int
	result   parser.Results
}

func initialModel() model {
	return model{
		choices: []choice{
			{text: "Games for date", value: "gamesForDate"},
			{text: "Games for date", value: "gamesForDate"},
			{text: "Games for date", value: "gamesForDate"},
			{text: "Games for date", value: "gamesForDate"},
			{text: "Games for date", value: "gamesForDate"},
		},
		result: parser.Results{},
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			m.selected = m.cursor
			return m, getGameForDate("")
		}

	case parser.Results:
		m.result = msg
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What to fetch\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.text)
	}

	resultString, err := m.result.StringifyResults()
	if err != nil {
		log.Fatal("error", err)
	}

	// the results
	s += fmt.Sprintf("\n%s\n", resultString)

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

type errMsg struct{ err error }

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }

func getGameForDate(url string) tea.Cmd {
	return func() tea.Msg {
		dt := time.Now()
		date := flag.String("date", dt.Format("2006-01-02"), "the date")
		flag.Parse()
		var fetcher fetcher.Fetcher
		body := fetcher.GetGamesForDate(*date)
		var parser parser.Parser
		results, err := parser.ParseResults(body)
		if err != nil {
			log.Fatalln("Error while parsing results", err)
		}
		return results
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
