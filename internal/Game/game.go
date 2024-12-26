package Game

import "strings"

type Game struct {
  currentWorld World
  currentView View
  exit bool
  Inventory
}

func NewGame(worldName string) *Game {
  world := World{name: worldName, WorldReader: WorldReader{worldName}}
  view, _ := world.GetView(Coordinate{0,0})
  return &Game{
    currentWorld: world,
    currentView: view,
    exit: false,
    Inventory: Inventory{CarryWeight: CarryWeight{max: 100}},
  }
}

func StartGame() {
  game := NewGame("World1");
  game.Describe()
  game.loop()
}

func (g *Game) loop() {
  for !g.exit {
    res := GetUserInput()
    action, args := g.analyzeResponse(res);
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
    return Coordinate{g.currentView.Coordinates.x, g.currentView.Coordinates.y+1}
  case Down:
    return Coordinate{g.currentView.Coordinates.x, g.currentView.Coordinates.y-1}
  case Left:
    return Coordinate{g.currentView.Coordinates.x-1, g.currentView.Coordinates.y}
  case Right:
    return Coordinate{g.currentView.Coordinates.x+1, g.currentView.Coordinates.y}
  }
  return Coordinate{0,0}
}


func (g *Game) findObjsInGame(objectName string) []ObjectInt {
  result := make([]ObjectInt, 0)

  for _,o := range g.currentView.Objects {
    if strings.Contains(o.getDescription(), objectName) {
      result = append(result, o)
    }
  }

  for _,o := range g.Inventory.items {
    if strings.Contains(o.getDescription(), objectName) {
      result = append(result, o)
    }
  }

  return result
}

func (g *Game) RemoveFromScene(obj ObjectInt) {
  FilterInPlace(&g.currentView.Objects, func(o ObjectInt)bool {
    return !o.equals(obj)
  })
}
