package Game;

type Object struct {
  name string
}

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
