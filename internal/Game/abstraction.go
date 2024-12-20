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
  StoryNote string
  People []Person
  Objects []interface{}
  HiddenNotes []Note
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
