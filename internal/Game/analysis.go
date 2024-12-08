package Game

import (
	"encoding/json"
	"os"
	"strings"
)

func (g *Game) analyzeResponse(response string) (Action, []string) {

	words := strings.Fields(response)

	ActionNames := getActionNames(g)
	ActionAliases := getActionAliases()

	if action, exists := ActionNames[words[0]]; exists {
		return action, words
	} else if name, exists := ActionAliases[words[0]]; exists {
		return ActionNames[name], words
	} else {
		return g.Unknown, words
	}
}

func getActionAliases() map[string]string {
	file, err := os.Open("./internal/Dict/Dict.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode JSON into a map
	var commands map[string][]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&commands); err != nil {
		panic(err)
	}

  result := make(map[string]string)
  for key, aliases := range commands {
    for _, alias := range aliases {
      result[alias] = key;
    }
  }
  return result
}
