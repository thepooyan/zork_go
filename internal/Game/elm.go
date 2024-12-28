package Game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

func (g Game) Init() tea.Cmd {
	return textinput.Blink
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
      case "ctrl+c", "q":
        return g, tea.Quit
      case "enter":
        g.calculateNextPrompt(g.textInput.Value())
        g.textInput.SetValue("")
        return g, nil
		}
	}

  var cmd tea.Cmd
  g.textInput, cmd = g.textInput.Update(msg)
	return g, cmd
}

func (g Game) View() string {
  view := ""
  view += g.VirtualOutput.output + "\n"
  view += "\n"

  view += g.textInput.View()

	return view
}
