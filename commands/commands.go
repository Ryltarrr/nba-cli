package commands

import (
	"log"
	"time"

	"github.com/Ryltarrr/nba-cli/fetcher"
	"github.com/Ryltarrr/nba-cli/parser"
	tea "github.com/charmbracelet/bubbletea"
)

const DATE_FORMAT = "2006-01-02"

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func GetGamesForDateCommand(date string) tea.Cmd {
	return func() tea.Msg {
		_, err := time.Parse(DATE_FORMAT, date)

		if date == "" || err != nil {
			dt := time.Now()
			date = dt.Format(DATE_FORMAT)
		}

		log.Printf("Fetching results for %s", date)
		body := fetcher.GetGamesForDate(date)
		results, err := parser.ParseResults(body)
		if err != nil {
			return errMsg{err}
		}
		return results
	}
}
