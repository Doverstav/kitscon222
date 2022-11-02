package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/doverstav/kitscon222/charm_demo/view"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	// Get app dir
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	appDir := expandedDir + "/kitscon-cli"
	fmt.Println(appDir)

	// Setup db
	db, err := badger.Open(badger.DefaultOptions(appDir).WithLoggingLevel(badger.ERROR))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initalView := view.KITSCON_LIST

	input := textinput.New()
	input.Placeholder = "New KitsCon"
	input.CharLimit = 156
	input.Width = 20

	areainput := textarea.New()
	areainput.Placeholder = "Your thought about the presentation"

	if err := tea.NewProgram(view.Model{
		DB:          db,
		List:        list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		CurrentView: initalView,
		Input:       input,
		TextArea:    areainput,
	}).Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
