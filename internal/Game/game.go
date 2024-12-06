package Game

import "fmt"

type Game struct {
  currentWorld World
  currentView View
}

func StartGame() {
  fmt.Println("starting game...")
  gameInstance := Game{
    currentWorld: World{},
  }
  gameInstance.currentView = gameInstance.currentWorld.GetView(Coordinate{0,0})
  gameInstance.loop()
}

func (g *Game) loop() {
  for {
    res := getUserInput(g.describe())
    analyzeResponse(res);
  }
}

func (g *Game) describe() string {
  return g.currentView.storyNote
}

func getUserInput(prompt string) string {
  Outputln(prompt)
  Output(" => ")
  response := Input()
  return response
}
