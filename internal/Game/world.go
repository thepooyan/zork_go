package Game

type World struct {
  name string
  cache map[Coordinate]View
  WorldReader
}

func (w *World) readView(c Coordinate) View {
  view,err := w.ReadFile(c)

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
