package view

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/doverstav/kitscon222/charm_demo/commands"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var inputStyle = lipgloss.NewStyle().Margin(1, 0)

type View string

const (
	KITSCON_LIST         View = "KITSCON_LIST"
	ADD_NEW_KITSCON      View = "ADD_NEW_KITSCON"
	PRESENTATIONS_LIST   View = "PRESENTATIONS_LIST"
	ADD_NEW_PRESENTATION View = "ADD_NEW_PRESENTATION"
	PRESENTATION         View = "PRESENTATION"
)

type kitsconListItem struct {
	title       string
	description string
}

func (k kitsconListItem) Title() string       { return k.title }
func (k kitsconListItem) Description() string { return k.description }
func (k kitsconListItem) FilterValue() string { return k.title }

type Model struct {
	CurrentView View
	List        list.Model
	Input       textinput.Model
	TextArea    textarea.Model
}

func aMsg() tea.Msg {
	return characterMsg("a")
}

func bMsg() tea.Msg {
	return characterMsg("b")
}

type characterMsg string

func (m Model) Init() tea.Cmd {
	return commands.GetKitscons
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		if msg.String() == "a" {
			m.CurrentView = ADD_NEW_KITSCON
			m.Input.Focus()
			return m, aMsg
		}

		if msg.String() == "b" {
			return m, bMsg
		}
	case commands.KitsconsMsg:
		m.List.SetItems(KitsconMsgToListItem([]commands.Kitscon(msg)))
		return m, nil
	}

	if m.CurrentView == ADD_NEW_KITSCON {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			fmt.Print("test")
			if msg.Type == tea.KeyEnter || msg.Type == tea.KeyTab {
				m.Input.Blur()
				m.TextArea.Focus()

				return m, nil
			}
		}

		var fieldCmd tea.Cmd
		m.Input, fieldCmd = m.Input.Update(msg)

		var areaCmd tea.Cmd
		m.TextArea, areaCmd = m.TextArea.Update(msg)

		return m, tea.Batch(fieldCmd, areaCmd)
	} else {
		var cmd tea.Cmd
		m.List, cmd = m.List.Update(msg)

		return m, cmd
	}
}

func (m Model) View() string {
	if m.CurrentView == ADD_NEW_KITSCON {
		return docStyle.Render(inputStyle.Render(m.Input.View()) + "\n" + m.TextArea.View())
	} else if m.CurrentView == KITSCON_LIST {
		return KitsConListView(m)
	}

	return ""
}

// -------- HELPERS -----------
func KitsconMsgToListItem(kitscons []commands.Kitscon) []list.Item {
	listItems := []list.Item{}

	for _, kitscon := range kitscons {
		listItems = append(listItems, kitsconListItem{title: kitscon.Name, description: kitscon.Description})
	}

	return listItems
}
