package Game

import "encoding/xml"

type Object struct {
	name string
}

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
	XMLName     xml.Name     `xml:"view"`
	StoryNote   string       `xml:"story_note"`
	People      People       `xml:"people"`
	Objects     Objects      `xml:"objects"`
	HiddenNotes []HiddenNote `xml:"hidden_notes>note"`
}

type People struct {
	Persons []Person `xml:"person"`
}

type Person struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description,attr"`
	Text        string `xml:",chardata"`
}

type Objects struct {
	Boxes       []Box       `xml:"box"`
	LockedBoxes []LockedBox `xml:"lockedBox"`
	Keys        []Key       `xml:"key"`
}

type Box struct {
	Description string `xml:"description,attr"`
}

type LockedBox struct {
	ID          string  `xml:"id,attr"`
	Description string  `xml:"description,attr"`
	Letter      *Letter `xml:"letter"`
}

type Letter struct {
	Description string `xml:"description,attr"`
	Text        string `xml:",chardata"`
}

type Key struct {
	ID          string `xml:"id,attr"`
	Description string `xml:"description,attr"`
}

type HiddenNote struct {
	Keyword string `xml:"keyword,attr"`
	Text    string `xml:",chardata"`
}
