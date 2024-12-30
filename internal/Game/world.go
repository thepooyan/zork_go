package Game

type World struct {
  name string
  cache map[Coordinate]View
  WorldReader
}

func (w *World) readView(c Coordinate) (View, error) {
  view,err := w.ReadFile(c)
  return view, err
}

func (w *World) GetView(c Coordinate) (View, error) {
  view, ok := w.cache[c]

  if ok {
    return view, nil
  }

  return w.readView(c)
}
