package Game

import (
	"encoding/gob"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type Game struct {
	currentWorld World
	currentView  View
	Inventory
	exit      bool
	textInput textinput.Model
  spinner spinner.Model
  VirtualOutput
}

func initTextInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "What do you do? \U0001F600"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
  return ti
}

func NewGame(worldName string) *Game {
	world := World{name: worldName, WorldReader: WorldReader{worldName}}
	view, _ := world.GetView(Coordinate{0, 0})
	return &Game{
		currentWorld: world,
		currentView:  view,
		exit:         false,
		Inventory:    Inventory{CarryWeight: CarryWeight{max: 100}},
    textInput: initTextInput(),
    VirtualOutput: VirtualOutput{},
	}
}

func (g *Game) prepareInitialPrompt() {
  g.Describe()
}

func (g *Game) calculateNextPrompt(response string) {
  style := lipgloss.NewStyle().Foreground(lipgloss.Color("#eb246c")).Bold(true)
  g.VirtualOutput.write(style.Render("> ") + response)
  action, args := g.analyzeResponse(response)
  action(args...)
}

func (g *Game) ChangeLocation(d Direction) {
	view, err := g.currentWorld.GetView(g.DirectionToCoordinate(d))
	if err == nil {
		g.currentView = view
	}
	g.Describe()
}

func (g *Game) DirectionToCoordinate(d Direction) Coordinate {
	switch d {
	case Up:
		return Coordinate{g.currentView.Coordinates.x, g.currentView.Coordinates.y + 1}
	case Down:
		return Coordinate{g.currentView.Coordinates.x, g.currentView.Coordinates.y - 1}
	case Left:
		return Coordinate{g.currentView.Coordinates.x - 1, g.currentView.Coordinates.y}
	case Right:
		return Coordinate{g.currentView.Coordinates.x + 1, g.currentView.Coordinates.y}
	}
	return Coordinate{0, 0}
}

func (g *Game) findObjsInGame(objectName string) []ObjectInt {
	result := make([]ObjectInt, 0)

	for _, o := range g.currentView.Objects {
		if strings.Contains(o.getDescription(), objectName) {
			result = append(result, o)
		}
	}

	for _, o := range g.Inventory.items {
		if strings.Contains(o.getDescription(), objectName) {
			result = append(result, o)
		}
	}

	return result
}

func (g *Game) RemoveFromScene(obj ObjectInt) {
	FilterInPlace(&g.currentView.Objects, func(o ObjectInt) bool {
		return !o.equals(obj)
	})
}

func (g *Game) Save(saveName string) {
	file, err := os.Create(fmt.Sprintf("./saves/%s", saveName))
	if err != nil {
		g.Respond("Error creating file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(g)
	if err != nil {
		g.Respond("Error saving the game")
		fmt.Println(err)
		return
	}
	g.Respond("Saved successfully!")
}

func (g *Game) Load(saveName string) {
	file, err := os.Open(fmt.Sprintf("./saves/%s", saveName))
	if err != nil {
		g.Respond("Error opening the file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&g)
	if err != nil {
		g.Respond("Error loading the game")
		fmt.Println(err)
		return
	}
	g.Respond("Loaded successfully!")
}

func (g *Game) GetAnotherPrompt(callback func(response string)) {

}
