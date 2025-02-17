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

	StoryNote := RemoveXmlIndentation(root.SelectElement("story_note").Text())

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
	Objects := parseChildNodes(ObjectsNode)

	notesNode := root.SelectElement("hidden_notes")
	Notes := make([]Note, 0)

	if notesNode != nil {
		for _, n := range notesNode.SelectElements("note") {
			newNote := Note{
				Keyword: n.SelectAttrValue("keyword", ""),
				Content: n.Text(),
			}
			Notes = append(Notes, newNote)
		}
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

func parseChildNodes(e *etree.Element) []ObjectInt {
	Objects := make([]ObjectInt, 0)

	if e != nil {
		for _, o := range e.ChildElements() {
			switch o.Tag {
			case "letter":
				l := NewLetter(RemoveXmlIndentation(o.Text()), o.SelectAttrValue("description", ""))
				Objects = append(Objects, l)
			case "box":
				b := NewBox(o.SelectAttrValue("description", ""))
        innerStuff := parseChildNodes(o)
        for _,i := range innerStuff {
          b.Add(i)
        }
				Objects = append(Objects, b)
			case "lockedBox":
				b := NewLockedBox(o.SelectAttrValue("description", ""), o.SelectAttrValue("id", ""))
        innerStuff := parseChildNodes(o)
        for _,i := range innerStuff {
          b.Add(i)
        }
				Objects = append(Objects, b)
			case "key":
				b := NewKey(o.SelectAttrValue("description", ""), o.SelectAttrValue("id", ""))
				Objects = append(Objects, b)
			}
		}
	}
	return Objects
}
