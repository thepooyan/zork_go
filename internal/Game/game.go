package Game

type Game struct {
  currentWorld World
  currentView View
}

func NewGame(worldName string) *Game {
  world := World{name: worldName}
  return &Game{
    currentWorld: world,
    currentView: world.GetView(Coordinate{0,0}),
  }
}

func StartGame() {
  game := NewGame("World1");
  game.describe()
  game.loop()
}

func (g *Game) loop() {
  for {
    res := GetUserInput()
    action, args := g.analyzeResponse(res);
    action(args)
  }
}

func (g *Game) describe() {
  Outputln(g.currentView.storyNote)
}
