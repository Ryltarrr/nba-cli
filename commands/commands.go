package commands

import (
	"log"
	"time"

	"github.com/Ryltarrr/nba-cli/fetcher"
	"github.com/Ryltarrr/nba-cli/parser"
	"github.com/Ryltarrr/nba-cli/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func GetGamesForDateCommand(date string) tea.Cmd {
	return func() tea.Msg {
		_, err := time.Parse(utils.DATE_FORMAT, date)

		if date == "" || err != nil {
			dt := time.Now()
			date = dt.Format(utils.DATE_FORMAT)
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
