package Game

func getActionNames(g *Game) map[string]Action {
  return map[string]Action{
    "hi": g.Hi,
    "idol": g.Idol,
    "unknown": g.Unknown,
    "move": g.Move,
  }
}

type Action func(args []string)

func (a *Game) Hi(args []string) {
 Outputln("Hello!")
}

func (a *Game) Idol(args []string) {
  Outputln("Doing nothing!...")
}

func (a *Game) Unknown(args []string) {
  Outputln("i dunno what \"" + args[0] + "\" is :/")
}

func (a *Game) Move(args []string) {
  if notEnoughArgs(2, args) {
    Outputln("Move where? :/")
    return
  }
}

func notEnoughArgs(required int, args []string) bool {
  return len(args) < required
}
