package commands

import (
	"time"

	"github.com/Ryltarrr/go-nba/fetcher"
	"github.com/Ryltarrr/go-nba/parser"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func GetGamesForDateCommand(date string) tea.Cmd {
	return func() tea.Msg {
		dt := time.Now()

		if date == "" {
			date = dt.Format("2006-01-02")
		}
		body := fetcher.GetGamesForDate(date)
		results, err := parser.ParseResults(body)
		if err != nil {
			return errMsg{err}
		}
		return results
	}
}
