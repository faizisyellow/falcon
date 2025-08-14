package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ModelResult struct {
	Cursor  int
	Choices []string
	Next    int
}

// options based the questions index.
var options = map[int][]string{
	0: {"Chi"},
	1: {"Mysql"},
}

var questions = []string{
	// router
	"What kind of router would you like to use ?",
	// database
	"What kind of rational database would you like to use ?",
}

// Style definitions
var (
	header = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000ff")).Bold(true).Render
	list   = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Render
)

func (m ModelResult) Init() tea.Cmd {
	return nil
}

func (m ModelResult) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":

			return m, tea.Quit

		case "enter":
			// Send the Choices on the channel and exit.
			m.Choices = append(m.Choices, options[m.Next][m.Cursor])

			// if the question is the last one, quit
			if len(questions)-1 == m.Next {
				return m, tea.Quit
			}

			m.Next++
			m.Cursor = 0

			return m, nil

		case "down", "j":
			m.Cursor++
			if m.Cursor >= len(options[m.Next]) {
				m.Cursor = 0
			}

		case "up", "k":
			m.Cursor--
			if m.Cursor < 0 {
				m.Cursor = len(options[m.Next]) - 1
			}
		}
	}

	return m, nil
}

func (m ModelResult) View() string {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString(header(questions[m.Next]))
	s.WriteString("\n\n")

	for i := 0; i < len(options[m.Next]); i++ {
		if m.Cursor == i {
			s.WriteString(list(lipgloss.JoinHorizontal(lipgloss.Left, "(•) ", options[m.Next][i])))
		} else {
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, "( ) ", options[m.Next][i]))
		}
		s.WriteString("\n")
	}
	s.WriteString("\n(press ctrl + c to quit)\n\n")

	return s.String()
}
