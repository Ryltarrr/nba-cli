package menu

import (
	"time"

	"github.com/Ryltarrr/nba-cli/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Focused   bool
	TextInput textinput.Model
}

func Init(m Model) tea.Cmd {
	return nil
}

func New() Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 10
	ti.Placeholder = time.Now().Format(utils.DATE_FORMAT)
	// TODO: Use "2022-01-19" to handle long list
	ti.SetValue("2022-01-17")

	return Model{
		Focused:   true,
		TextInput: ti,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmdTextInput tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "esc":
			m.Focused = true
			m.TextInput.Focus()
			return m, textinput.Blink

		case "backspace":
			m.TextInput, cmdTextInput = m.TextInput.Update(msg)
			newVal := autoDelete(m.TextInput.Value())
			m.TextInput.SetValue(newVal)
			return m, cmdTextInput

		case "-":
			lenTi := len(m.TextInput.Value())
			if lenTi == 4 || lenTi == 7 {
				m.TextInput, cmdTextInput = m.TextInput.Update(msg)
			}
			return m, cmdTextInput
		}

		if m.Focused {
			m.TextInput, cmdTextInput = m.TextInput.Update(msg)
			newVal, newPos := autoComplete(m.TextInput.Value())
			m.TextInput.SetValue(newVal)
			m.TextInput.SetCursor(newPos)
			return m, cmdTextInput
		}
	}

	return m, textinput.Blink
}

func autoComplete(s string) (string, int) {
	count := len(s)
	if count == 4 || count == 7 {
		return s + "-", count + 2
	}
	return s, count
}

func autoDelete(s string) string {
	if len(s) > 0 && s[len(s)-1] == '-' {
		return s[:len(s)-2]
	}
	return s
}

func (m Model) View() string {
	s := "Date of the game:\n"
	s += lipgloss.NewStyle().PaddingLeft(1).Render(m.TextInput.View())

	menuBorderColor := lipgloss.Color("#eee")
	if m.Focused {
		menuBorderColor = lipgloss.Color("205")
	}

	menuStyle := lipgloss.NewStyle().
		MarginRight(2).
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(menuBorderColor)

	return menuStyle.Render(s)
}
