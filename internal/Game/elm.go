package Game

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (g Game) Init() tea.Cmd {
  cmds := []tea.Cmd{g.spinner.Tick, textinput.Blink}
	return tea.Batch(cmds...)
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if g.exit {
		return g, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "g":
			return g, g.spinner.Tick
		case "ctrl+c":
			return g, tea.Quit
		case "enter":
			if g.innerPrompt != nil {
				g.RunInnerPrompt(g.textInput.Value())
				return g, nil
			}
			g.ResponseRecieved(g.textInput.Value())
			return g, nil
		}
	}

	var cmds []tea.Cmd
  {
    var cmd tea.Cmd
    g.spinner, cmd = g.spinner.Update(msg)
    cmds = append(cmds, cmd)
  }
  {
    var cmd tea.Cmd
    g.textInput, cmd = g.textInput.Update(msg)
    cmds = append(cmds, cmd)
  }
	return g, tea.Batch(cmds...)
}

func (g Game) View() string {
	view := ""
	view += g.VirtualOutput.Output
	view += "\n\n"
	view += g.textInput.View()
  view += g.spinner.View()
	return view
}
