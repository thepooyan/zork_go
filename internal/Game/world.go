package Game

import (
	"fmt"
)

type World struct {
  name string
  cache map[Coordinate]View
}

func (w *World) readView(c Coordinate) View {
  targetFile := fmt.Sprintf("./Worlds/%s/%d_%d.xml", w.name, c.x, c.y)
  view,err := ReadFile(targetFile, c)

  if err != nil {
    panic("error getting the view")
  }
  return view
}

func (w *World) GetView(c Coordinate) View {
  read, ok := w.cache[c]

  if ok {
    return read
  }

  return w.readView(c)
}
