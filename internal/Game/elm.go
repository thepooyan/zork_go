package Game

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (g Game) Init() tea.Cmd {
	g.prepareInitialPrompt()
	return textinput.Blink
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

  if g.exit {
    return g, tea.Quit
  }

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return g, tea.Quit
		case "enter":
			if g.innerPrompt != nil {
				g.innerPrompt(g.textInput.Value())
				g.innerPrompt = nil
				return g, nil
			}
			g.ResponseRecieved(g.textInput.Value())
			g.textInput.Reset()
			return g, nil
		}
	}

	var cmd tea.Cmd
	g.textInput, cmd = g.textInput.Update(msg)
	return g, cmd
}

func (g Game) View() string {
	view := ""
	view += g.VirtualOutput.Output + "\n"
	view += "\n"

	view += g.textInput.View()

	return view
}
