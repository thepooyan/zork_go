package Game

import (
	"fmt"
)

type View struct {
  storyNote string
  coordinates Coordinate
  neighbors []Direction
  objects []Object
  hiddenNote map[string]string
}

type World struct {
  name string
  cache map[Coordinate]View
}

func (w *World) ReadView(c Coordinate) View {
  targetFile := fmt.Sprintf("./Worlds/%s/%d_%d.xml", w.name, c.x, c.y)
  view,_ := ReadFile(targetFile)
  return view
}

func (w *World) GetView(c Coordinate) View {
  read, ok := w.cache[c]

  if ok {
    return read
  }

  return w.ReadView(c)
}
