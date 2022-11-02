package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dgraph-io/badger/v3"
)

func GetKitscons(db *badger.DB) tea.Cmd {
	// Dummy implementation
	return func() tea.Msg {
		kitsconList := []Kitscon{
			{Id: "1", Name: "KitsCon 22.2", Description: "Strandbad edition"},
			{Id: "2", Name: "KitsCon 22.1", Description: "Jorgen edition"},
			{Id: "3", Name: "KitsCon 21.2", Description: "Vann edition"},
			{Id: "4", Name: "KitsCon 21.1", Description: "Mölle edition"},
		}

		return KitsconsMsg(kitsconList)
	}
}

type Kitscon struct {
	Id          string
	Name        string
	Description string
}

type KitsconsMsg []Kitscon
