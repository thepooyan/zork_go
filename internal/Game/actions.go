package Game

type Action func(args []string)

func (a *Game) Hi(args []string) {
 Outputln("Hello!")
}

func (a *Game) Idol(args []string) {
  Outputln("Doing nothing!...")
}

func (a *Game) Unknown(args []string) {
  Outputln("what?")
}
