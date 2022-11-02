package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/doverstav/kitscon222/charm_demo/view"
)

func main() {
	initalView := view.KITSCON_LIST

	input := textinput.New()
	input.Placeholder = "New KitsCon"
	input.CharLimit = 156
	input.Width = 20

	areainput := textarea.New()
	areainput.Placeholder = "Your thought about the presentation"

	if err := tea.NewProgram(view.Model{List: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0), CurrentView: initalView, Input: input, TextArea: areainput}).Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
