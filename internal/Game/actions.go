package Game

import (
	"os"
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
    "read": g.Read,
    "unlock": g.Unlock,
    "pickup": g.Pickup,
    "inventory": g.GetInventory,
    "weight": g.CarryWeight,
    "iloveyou": g.Love,
    "lock": g.Lock,
    "drop": g.Drop,
    "save": g.SaveGame,
  }
}

type Action func(args ...string)

func (a *Game) LoadGame(args ...string) {
  saves,err := os.ReadDir("./saves")
  if err != nil {
    Respond("Error opening the saves directory")
    return
  }
  if len(saves) == 0 {
    Respond("You have no save files")
    return
  }
  Respond("Which save to load?")
  for _,s := range saves {
    Respond("- ", s.Name())
  }
  input := GetUserInput()
  a.Load(input)
}

func (a *Game) SaveGame(args ...string) {
  Respond("Saving the game as:")
  name := GetUserInput()
  a.Save(name)
}

func (a *Game) Drop(args ...string) {
  target := strings.Join(args[1:], " ")
  item, ok := a.Inventory.Find(target)
  if ok {
    a.Inventory.Remove(item)
    a.currentView.Objects = append(a.currentView.Objects, item)
    Respond("Dropped ", item.getDescription())
  } else {
    Respond("You have no '", target, "' in your inventory!")
  }
}

func (a *Game) Lock(args ...string) {
  action := a.NewSingleAction("lock", (*LockInt)(nil), args...)
  if action != nil {
    action()
  }
}

func (a *Game) Love(args ...string) {
  Respond("iloveyou too ^^")
}


func (a *Game) CarryWeight(args ...string) {
  a.Inventory.CarryWeight.Print()
}

func (a *Game) GetInventory(args ...string) {
  a.Inventory.CarryWeight.Print()
  a.Inventory.Describe()
}

func (a *Game) Pickup(args ...string) {
  action := a.NewSingleAction("pickup", (*PickableInt)(nil), args...)
  if action != nil {
    action(a)
  }
}

func (a *Game) Unlock(args ...string) {
  action := a.NewTwoStepAction("unlock", (*LockInt)(nil), (*KeyInt)(nil), args...)
  if action != nil {
    action(a)
  }
}

func (a *Game) Read(args ...string) {
  action := a.NewSingleAction("read", (*ReadableInt)(nil), args...)
  if action != nil {
    action()
  }
}

func (a *Game) Open(args ...string) {
  action := a.NewSingleAction("open", (*ContainerInt)(nil), args...)
  if action != nil {
    action(a)
  }
}

func (a *Game) Test(args ...string) {
  for _,i := range a.currentView.HiddenNotes {
    Respond(i.Content)
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
  Describe(a.currentView.StoryNote)

  for _,o := range a.currentView.Objects {
    Describe(AddRandomSmalltalk(o.getDescription()))
  }
  for _,o := range a.currentView.People {
    Describe(AddRandomSmalltalk(o.Description))
  }
}

func (a *Game) Exit(args ...string) {
 Respond("Sending the exit signal...") 
 a.exit = true;
}

func notEnoughArgs(required int, args ...string) bool {
  return len(args) < required
}
