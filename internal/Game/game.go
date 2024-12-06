package Game

import "fmt"

type Game struct {
  currentWorld World
  currentView View
}

func StartGame() {
  fmt.Println("starting game...")
  gameInstance := &Game{
    currentWorld: World{name: "World1"},
  }
  gameInstance.currentView = gameInstance.currentWorld.GetView(Coordinate{0,0})
  gameInstance.loop()
}

func (g *Game) loop() {
  for {
    res := GetUserInput(g.describe())
    analyzeResponse(res);
  }
}

func (g *Game) describe() string {
  return g.currentView.storyNote
}
