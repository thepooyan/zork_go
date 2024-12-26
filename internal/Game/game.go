package Game

import (
	"encoding/gob"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	"os"
	"strings"
)

type Game struct {
	currentWorld World
	currentView  View
	Inventory
	exit      bool
	textInput textinput.Model
}

func initTextInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
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
	}
}

func (g *Game) loop() {
	for !g.exit {
		res := GetUserInput()
		action, args := g.analyzeResponse(res)
		action(args...)
	}
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
		Respond("Error creating file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(g)
	if err != nil {
		Respond("Error saving the game")
		fmt.Println(err)
		return
	}
	Respond("Saved successfully!")
}

func (g *Game) Load(saveName string) {
	file, err := os.Open(fmt.Sprintf("./saves/%s", saveName))
	if err != nil {
		Respond("Error opening the file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&g)
	if err != nil {
		Respond("Error loading the game")
		fmt.Println(err)
		return
	}
	Respond("Loaded successfully!")
}
