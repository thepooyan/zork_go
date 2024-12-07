package Game

import (
	"strings"
)


func (g *Game) analyzeResponse(response string) (Action, []string) {

  words := strings.Fields(response);

  ActionNames := getActionNames(g);
  ActionAliases := getActionAliases();

  if action, exists := ActionNames[words[0]]; exists {
    return action, words
  } else if name, exists := ActionAliases[words[0]]; exists {
    return ActionNames[name], words
  } else {
    return g.Unknown, words
  }
}

func getActionAliases() map[string]string {
  return map[string]string{
    "hello": "hi",
    "hey": "hi",
    "nothing": "idol",
    "stand": "idol",
  }
}

func getActionNames(g *Game) map[string]Action {
  return map[string]Action{
    "hi": g.Hi,
    "idol": g.Idol,
  }
}
