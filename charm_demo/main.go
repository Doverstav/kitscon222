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

/* TODO:
- Add nicer keybinds with bubbles components ##MUST
- Add nice help message on all screen ##MUST
- Support "going back" from screens, all the way to root screen ##MUST
- Improve "rating" input to be bounded to 0 - 5 stars ##NICE TO HAVE
- Add some nice styling to "Add"-screens ##NICE TO HAVE
- Can models be setup in a nicer way? I.e. move ##IF TIME PERMITS
	state/model into the corresponding view, not
	having it all in the "main" view
- Think about edit functionality ##IF TIME PERMITS
*/

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

	kitsconTitleInput := textinput.New()
	kitsconTitleInput.Placeholder = "New KitsCon"
	kitsconTitleInput.CharLimit = 156
	kitsconTitleInput.Width = 20

	kitsconDescInput := textarea.New()
	kitsconDescInput.Placeholder = "KitsCon description"

	presentationTitleInput := textinput.New()
	presentationTitleInput.Placeholder = "New presentation"
	presentationTitleInput.CharLimit = 156
	presentationTitleInput.Width = 20

	presentationPresenterInput := textinput.New()
	presentationPresenterInput.Placeholder = "Who presented"
	presentationPresenterInput.CharLimit = 156
	presentationPresenterInput.Width = 20

	presentationDescInput := textarea.New()
	presentationDescInput.Placeholder = "Presentation description"

	presentationRatingInput := textinput.New()
	presentationRatingInput.Placeholder = "Your rating"
	presentationRatingInput.CharLimit = 156
	presentationRatingInput.Width = 20

	presentationReviewInput := textarea.New()
	presentationReviewInput.Placeholder = "Your thoughts on the presentation"

	if err := tea.NewProgram(view.Model{
		DB:       db,
		ItemList: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		//PresentationList:             list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		CurrentView:                  initalView,
		KitsconTitleInput:            kitsconTitleInput,
		KitsconDescriptionInput:      kitsconDescInput,
		PresentationTitleInput:       presentationTitleInput,
		PresentationPresenterInput:   presentationPresenterInput,
		PresentationDescriptionInput: presentationDescInput,
		PresentationRatingInput:      presentationRatingInput,
		PresentationReviewInput:      presentationReviewInput,
	}).Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
