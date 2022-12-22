package main

import (
	"fmt"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().Background(lipgloss.Color("21")).MarginLeft(5)

type model struct {
	url    string
	status int
}

type statusMsg int

func checkURL(url string) tea.Cmd {
	return func() tea.Msg {
		c := &http.Client{Timeout: 10 * time.Second}
		res, _ := c.Get(url)

		return statusMsg(res.StatusCode)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		} else if msg.Type == tea.KeyEnter {
			return m, checkURL(m.url)
		}

		m.url += msg.String()

		return m, nil
	case statusMsg:
		m.status = int(msg)

		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	return style.Render(fmt.Sprintf("URL to check: %s\nStatus: %d", m.url, m.status))
}

func main() {
	tea.NewProgram(model{}).Start()
}
