# Demo doc

`go.mod` should ideally be setup before so we don't have to wait for `go get`

Build same app in both frameworks. An app that takes a URL and return the status code. To expand, we can take some optional parameter/flag that tells us how many times we want to make our request.

## Cobra

### Live coding
Duration: appr. 10 min

- go mod init cobra_demo (setup go repo)
- cobra-cli init (create empty cobra command)
- go run main.go
- cobra-cli add check (create single command)
- take arg and print it
- add exactargs(1)
- check URL and return status
```golang
    url := args[0]

    c := &http.Client{Timeout: 10 * time.Second}
    res, err := c.Get(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(res.Status)
```
- show flags
```golang
// access
times, _ := strconv.Atoi(cmd.Flag("times").Value.String())
// create
checkCmd.Flags().IntP("times", "t", 1, "How many times should the URL be checked")
```
- show shorthand
- required flag
```golang
checkCmd.MarkFlagRequired("times")
```
- show command aliases
- show input
```golang
review := ""
survey.AskOne(&survey.Input{
    Message: "Please rate this app and write a review!",
}, &review)

fmt.Printf("Review submitted: %s", review)
```
- go build and just run the .exe [optional step]

## Charm

### Live code
Duration: appr. 12 min


- go mod init charm_demo
- write model
- write init
- write update
- write view
```golang
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	url   string
	status int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "hello!"
}

func main() {
	err := tea.NewProgram(model{}).Start()
	if err != nil {
		fmt.Printf("Broke! %v", err)
		os.Exit(1)
	}
}
```
- simple input to URL
- create a msg first
```golang
type model struct {
	url    string
	status int
}

type StatusMsg int

func checkURL() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, _ := c.Get("http://kits.se")

	return StatusMsg(res.StatusCode)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Omitted
	case StatusMsg:
		m.status = int(msg)
	}
    // Omitted
}

func (m model) View() string {
	return fmt.Sprintf("URL to check: %s\nStatus: %d", m.url, m.status)
}
```
- create command
```golang
func checkURL(url string) tea.Cmd {
	return func() tea.Msg {
		c := &http.Client{Timeout: 10 * time.Second}
		res, _ := c.Get(url)

		return StatusMsg(res.StatusCode)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Omitted
		else if msg.Type == tea.KeyEnter {
			return m, checkURL(m.url)
		}
        // Omitted
	}

	return m, nil
}
```
- show lipgloss
```golang
var style = lipgloss.NewStyle().Background(lipgloss.Color("21"))

func (m model) View() string {
	return style.Render(fmt.Sprintf("URL to check: %s\nStatus: %d", m.url, m.status))
}
```
- backspace implementation [optional]
```golang
else if msg.Type == tea.KeyBackspace {
    m.url = m.url[:len(m.url)-1]

    return m, nil
}
```