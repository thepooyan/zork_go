package Game

type Game struct {
  currentWorld World
  currentView View
  exit bool
}

func NewGame(worldName string) *Game {
  world := World{name: worldName, WorldReader: WorldReader{worldName}}
  return &Game{
    currentWorld: world,
    currentView: world.GetView(Coordinate{0,0}),
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
