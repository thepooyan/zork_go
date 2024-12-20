package Game

type Game struct {
  currentWorld World
  currentView View
  exit bool
}

func NewGame(worldName string) *Game {
  world := World{name: worldName, WorldReader: WorldReader{worldName}}
  view, _ := world.GetView(Coordinate{0,0})
  return &Game{
    currentWorld: world,
    currentView: view,
    exit: false,
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
