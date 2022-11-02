package view

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dgraph-io/badger/v3"
	"github.com/doverstav/kitscon222/charm_demo/commands"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type View string

const (
	KITSCON_LIST         View = "KITSCON_LIST"
	ADD_NEW_KITSCON      View = "ADD_NEW_KITSCON"
	PRESENTATIONS_LIST   View = "PRESENTATIONS_LIST"
	ADD_NEW_PRESENTATION View = "ADD_NEW_PRESENTATION"
	PRESENTATION         View = "PRESENTATION"
)

type presentationListItem struct {
	title  string
	rating int
}

func (p presentationListItem) Title() string       { return p.title }
func (p presentationListItem) Description() string { return strings.Repeat("⭐", p.rating) }
func (p presentationListItem) FilterValue() string { return p.title }

type Model struct {
	DB                      *badger.DB
	CurrentView             View
	List                    list.Model
	KitsconTitleInput       textinput.Model
	KitsconDescriptionInput textarea.Model
	SelectedKitscon         commands.Kitscon
}

func (m Model) Init() tea.Cmd {
	return commands.GetKitscons(m.DB)
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
	case commands.KitsconsMsg:
		kitscons := []commands.Kitscon(msg)
		listItems := make([]list.Item, len(kitscons))
		for i := range kitscons {
			listItems[i] = kitscons[i]
		}
		m.List.SetItems(listItems)
		return m, nil
	case commands.PresentationsMsg:
		m.List.SetItems(PresentationMsgToListItem([]commands.Presentation(msg)))
	case commands.KitsconAddedMsg:
		m.CurrentView = KITSCON_LIST
		return m, commands.GetKitscons(m.DB)
	}

	if m.CurrentView == ADD_NEW_KITSCON {
		return AddKitsconUpdate(m, msg)
	} else if m.CurrentView == KITSCON_LIST {
		return KitsConListUpdate(m, msg)
	}

	return m, nil
}

func (m Model) View() string {
	if m.CurrentView == ADD_NEW_KITSCON {
		return AddKitsconView(m)
	} else if m.CurrentView == KITSCON_LIST {
		return KitsConListView(m)
	}

	return ""
}

// -------- HELPERS -----------

func PresentationMsgToListItem(presentations []commands.Presentation) []list.Item {
	listItems := []list.Item{}

	for _, presentation := range presentations {
		listItems = append(listItems, presentationListItem{title: presentation.Title, rating: presentation.Rating})
	}

	return listItems
}
