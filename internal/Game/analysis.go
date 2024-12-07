package Game

import (
	"strings"
)


func (g *Game) analyzeResponse(response string) {

  words := strings.Fields(response);

  viableActions := map[string]func() {
    "hi": g.Hi,
    "idol": g.Idol,
  }

  if action, exists := viableActions[words[0]]; exists {
    action()
  } else {
    Outputln("What?")
  }
}

