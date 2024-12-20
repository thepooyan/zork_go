package Game

import (
	"fmt"

	"github.com/beevik/etree"
)

type WorldReader struct {
	WorldName string
}

func (w *WorldReader) GetFileName(c Coordinate) string {
	return fmt.Sprintf("./Worlds/%s/%d_%d.xml", w.WorldName, c.x, c.y)
}

func (w *WorldReader) ReadFile(c Coordinate) (View, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromFile(w.GetFileName(c))
	if err != nil {
		return View{}, err
	}

	root := doc.SelectElement("view")

	StoryNote := root.SelectElement("story_note").Text()

	PeopleNode := root.SelectElement("people")
	People := make([]Person, 0)

	if PeopleNode != nil {

		for _, p := range PeopleNode.SelectElements("person") {
			guy := Person{
				p.SelectAttrValue("name", ""),
				p.SelectAttrValue("description", ""),
				p.Text(),
			}
			People = append(People, guy)
		}
	}

	ObjectsNode := root.SelectElement("objects")
	Objects := make([]interface{}, 0)

	for _, o := range ObjectsNode.ChildElements() {
		switch o.Tag {
		case "letter":
			l := NewLetter(o.Text(), o.SelectAttrValue("description", ""))
			Objects = append(Objects, l)
		case "box":
			b := NewBox(o.SelectAttrValue("description", ""))
			Objects = append(Objects, b)
		case "lockedBox":
			b := NewLockedBox(o.SelectAttrValue("description", ""), o.SelectAttrValue("id", ""))
			Objects = append(Objects, b)
		case "key":
			b := NewKey(o.SelectAttrValue("description", ""), o.SelectAttrValue("id", ""))
			Objects = append(Objects, b)
		default:
			println("unknown object while parsing", w.GetFileName(c), ". ", o.Tag)
		}
	}

	notesNode := root.SelectElement("hidden_notes").SelectElements("note")
	Notes := make([]Note, 0)

	for _, n := range notesNode {

		newNote := Note{
			Keyword: n.SelectAttrValue("keyword", ""),
			Content: n.Text(),
		}
		Notes = append(Notes, newNote)
	}

	Neighbors := make([]Direction, 0)

	if fileExists(w.GetFileName(Coordinate{c.x + 1, c.y})) {
		Neighbors = append(Neighbors, Right)
	}
	if fileExists(w.GetFileName(Coordinate{c.x - 1, c.y})) {
		Neighbors = append(Neighbors, Left)
	}
	if fileExists(w.GetFileName(Coordinate{c.x, c.y + 1})) {
		Neighbors = append(Neighbors, Up)
	}
	if fileExists(w.GetFileName(Coordinate{c.x, c.y - 1})) {
		Neighbors = append(Neighbors, Down)
	}

	return View{
		c,
		StoryNote,
		People,
		Objects,
		Notes,
		Neighbors,
	}, nil
}
