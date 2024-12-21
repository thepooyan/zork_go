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
    Respond("Open what?")
    return
  }

  target := strings.Join(args[1:], " ")
  o := a.findObjsInGame(target)

  switch len(o) {
  case 1:
    obj, ok := o[0].(ContainerInt)
    if !ok {
      Respond("can't open the ", o[0].getDescription())
      return
    }
    Respond("opening the ", o[0].getDescription(),"...")
    stuff := obj.Open()
    ListStuff(stuff)
  case 0:
    Respond("can't find any \"", target, "\"s")
  default:
    str := make([]string, 0)
    for _,i := range o {
      str = append(str, i.getDescription())
    }
    Respond(strings.Join(str, " or "), "?")
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
      Respond(i.Content)
      found = true
    }
  }
  if !found {
    Respond("Nothing found...")
  }
}
 
func (a *Game) Hi(args ...string) {
 Respond("Hello!")
}

func (a *Game) Idol(args ...string) {
  Respond("Doing nothing!...")
}

func (a *Game) Unknown(args ...string) {
  Respond("i dunno what \"" + args[0] + "\" is :/")
}

func (a *Game) Move(args ...string) {
  if notEnoughArgs(2, args...) {
    Respond("Move where? :/")
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
      Respond("you can't go \"", args[1], "\"")
      return
  }
  if slices.Contains(a.currentView.Neighbors, dir) {
    Respond("going ", args[1], "...")
    a.ChangeLocation(dir)
  } else {
    Respond("No path there...")
  }
}

func (a *Game) Describe(args ...string) {
  Describe(a.prefix + a.currentView.StoryNote)

  for _,o := range a.currentView.Objects {
    Describe(a.prefix + AddRandomSmalltalk(o.getDescription()))
  }
  for _,o := range a.currentView.People {
    Describe(a.prefix + AddRandomSmalltalk(o.Description))
  }
}

func (a *Game) Exit(args ...string) {
 Respond("Sending the exit signal...") 
 a.exit = true;
}

func notEnoughArgs(required int, args ...string) bool {
  return len(args) < required
}
