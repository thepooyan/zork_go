package Game

type Coordinate struct {
  x,y int
}

type Direction int;

const (
  Left Direction = iota
  Right 
  Up
  Down
)

type View struct {
  storyNote string
  coordinates Coordinate
  neighbors []Direction
  objects []Object
  hiddenNote map[string]string
}

type World struct {
  cache map[Coordinate]View
}

func ReadView(c Coordinate) View {
  // read xml of c
  // make the view 
  return View{}
}

func (w *World) GetView(c Coordinate) View {
  read, ok := w.cache[c]

  if ok {
    return read
  }

  return ReadView(c)
}
