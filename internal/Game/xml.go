package Game

import (
	"github.com/beevik/etree"
)

func ReadFile(filename string) (View, error) {
  doc := etree.NewDocument()
  err := doc.ReadFromFile(filename)
  if err != nil {
    return View{}, err
  }

  root := doc.SelectElement("view")

  StoryNote := root.SelectElement("story_note").Text()

  PeopleNode := root.SelectElement("people")
  People := make([]Person, 0)

  for _,p := range  PeopleNode.SelectElements("person") {
    guy := Person{
      p.SelectAttrValue("name", ""),
      p.SelectAttrValue("description", ""),
      p.Text(),
    }
    People = append(People, guy)
  }

  ObjectsNode := root.SelectElement("objects")
  Objects := make([]interface{}, 0)

  for _,o := range ObjectsNode.ChildElements() {
    switch o.Tag {
      case "letter":
        l := NewLetter(o.Text(), o.SelectAttrValue("description",""))
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
        println("unknown object while parsing", filename,". ", o.Tag)
    }
  }


  return View{
    StoryNote,
    People,
    Objects,
    make([]Note, 0),
  }, nil
}
