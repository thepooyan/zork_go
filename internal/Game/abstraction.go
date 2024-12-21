package Game

type Coordinate struct {
	x, y int
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type View struct {
  Coordinates Coordinate
  StoryNote string
  People []Person
  Objects []ObjectInt
  HiddenNotes []Note
  Neighbors []Direction
}

type Person struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description,attr"`
	Prompt      string `xml:",chardata"`
}

type Note struct {
  Keyword string `xml:"keyword,attr"`
  Content string `xml:",chardata"`
}
