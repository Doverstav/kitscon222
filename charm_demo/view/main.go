package view

import (
	"fmt"

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
	PRESENTATION_LIST    View = "PRESENTATION_LIST"
	ADD_NEW_PRESENTATION View = "ADD_NEW_PRESENTATION"
	PRESENTATION         View = "PRESENTATION"
)

type Model struct {
	DB              *badger.DB
	CurrentView     View
	SelectedKitscon commands.Kitscon
	// Shared list (switching between two lists produced weird artifacts)
	ItemList list.Model
	// AddKitsconView
	KitsconTitleInput       textinput.Model
	KitsconDescriptionInput textarea.Model
	// AddPresentationView
	PresentationInputFocus       int
	PresentationTitleInput       textinput.Model
	PresentationPresenterInput   textinput.Model
	PresentationDescriptionInput textarea.Model
	PresentationRatingInput      textinput.Model // TODO Make nice star/number input
	PresentationReviewInput      textarea.Model
}

func (m Model) Init() tea.Cmd {
	return commands.GetKitscons(m.DB)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.ItemList.SetSize(msg.Width-h, msg.Height-v)
		//m.PresentationList.SetSize(msg.Width-h, msg.Height-v)
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
		m.ItemList.SetItems(listItems)
		return m, nil
	case commands.PresentationsMsg:
		presentations := []commands.Presentation(msg)
		listItems := make([]list.Item, len(presentations))
		for i := range presentations {
			listItems[i] = presentations[i]
		}
		m.CurrentView = PRESENTATION_LIST
		m.ItemList.SetItems(listItems)
		return m, nil
	case commands.KitsconAddedMsg:
		m.CurrentView = KITSCON_LIST
		return m, commands.GetKitscons(m.DB)
	case commands.KitsconRemovedMsg:
		return m, commands.GetKitscons(m.DB)
	case commands.PresentationAddedMsg:
		m.CurrentView = PRESENTATION_LIST
		return m, commands.GetPresentations(m.DB, m.SelectedKitscon.Id)
	case commands.PresentationRemovedMsg:
		return m, commands.GetPresentations(m.DB, m.SelectedKitscon.Id)
	case commands.KitsconSelectedMsg:
		selectedKitscon := commands.Kitscon(msg)
		m.SelectedKitscon = selectedKitscon
		m.ItemList.Title = fmt.Sprintf("%s presentations", selectedKitscon.Name)
		return m, commands.GetPresentations(m.DB, selectedKitscon.Id)
	}

	if m.CurrentView == ADD_NEW_KITSCON {
		return AddKitsconUpdate(m, msg)
	} else if m.CurrentView == KITSCON_LIST {
		return KitsConListUpdate(m, msg)
	} else if m.CurrentView == PRESENTATION_LIST {
		return PresentationListUpdate(m, msg)
	} else if m.CurrentView == ADD_NEW_PRESENTATION {
		return AddPresentationUpdate(m, msg)
	}

	return m, nil
}

func (m Model) View() string {
	if m.CurrentView == ADD_NEW_KITSCON {
		return AddKitsconView(m)
	} else if m.CurrentView == KITSCON_LIST {
		return KitsConListView(m)
	} else if m.CurrentView == PRESENTATION_LIST {
		return PresentationListView(m)
	} else if m.CurrentView == ADD_NEW_PRESENTATION {
		return AddPresentationView(m)
	}

	return ""
}
