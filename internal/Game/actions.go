package Game

import (
	"slices"
	"strings"
)

func getActionNames(g *Game) map[string]Action {
  return map[string]Action{
    "unknown": g.Unknown,
    "hi": g.Hi,
    "idol": g.Idol,
    "move": g.Move,
    "describe": g.Describe,
    "exit": g.Exit,
    "test": g.Test,
    "search": g.Search,
    "open": g.Open,
  }
}

type Action func(args ...string)

func (a *Game) Open(args ...string) {
  if notEnoughArgs(2, args...) {
    Outputln("Open what?")
  }
  target := args[1]

  switch target {
    case "box":
      Outputln("searching for the box...")
    default:
      Outputln("what is ", target, "?")
  }
}

func (a *Game) Test(args ...string) {
  for _,i := range a.currentView.HiddenNotes {
    println(i.Content)
  }
}

func (a *Game) Search(args ...string) {
  wholeSentence := strings.Join(args, " ")
  
  found := false
  for _,i := range a.currentView.HiddenNotes {
    if strings.Contains(wholeSentence, i.Keyword) {
      Outputln(i.Content)
      found = true
    }
  }
  if !found {
    Outputln("Nothing found...")
  }
}
 
func (a *Game) Hi(args ...string) {
 Outputln("Hello!")
}

func (a *Game) Idol(args ...string) {
  Outputln("Doing nothing!...")
}

func (a *Game) Unknown(args ...string) {
  Outputln("i dunno what \"" + args[0] + "\" is :/")
}

func (a *Game) Move(args ...string) {
  if notEnoughArgs(2, args...) {
    Outputln("Move where? :/")
    return
  }
  var dir Direction
  switch args[1] {
    case "up", "north", "n", "u", "forward":
      dir = Up;
    case "down", "south", "s", "d", "backwards":
      dir = Down;
    case "left", "west", "w", "l":
      dir = Left;
    case "right", "east", "r", "e":
      dir = Right;
    default:
      Outputln("you can't go \"", args[1], "\"")
      return
  }
  if slices.Contains(a.currentView.Neighbors, dir) {
    Outputln("going ", args[1], "...")
    a.ChangeLocation(dir)
  } else {
    Outputln("No path there...")
  }
}

func (a *Game) Describe(args ...string) {
  Outputln(a.currentView.StoryNote)
}

func (a *Game) Exit(args ...string) {
 Outputln("Sending the exit signal...") 
 a.exit = true;
}

func notEnoughArgs(required int, args ...string) bool {
  return len(args) < required
}
