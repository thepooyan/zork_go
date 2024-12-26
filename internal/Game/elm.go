package Game

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (g Game) Init() tea.Cmd {
	return nil
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return g, tea.Quit
		}
	}
	return g, nil
}

func (g Game) View() string {
  view := "** dummy ** \n"


	return view
}
