package Game

func getActionNames(g *Game) map[string]Action {
  return map[string]Action{
    "unknown": g.Unknown,
    "hi": g.Hi,
    "idol": g.Idol,
    "move": g.Move,
    "describe": g.Describe,
    "exit": g.Exit,
  }
}

type Action func(args ...string)

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
