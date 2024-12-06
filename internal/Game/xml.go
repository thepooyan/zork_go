package Game

import (
	"encoding/xml"
	"io"
	"os"
)

type XML_view struct {
	XMLName     xml.Name     `xml:"view"`
	XML_storyNote   string       `xml:"story_note"`
	XML_people      XML_people       `xml:"people"`
	XML_objects     XML_objects     `xml:"objects"`
	XML_hiddenNotes []XML_hiddenNote `xml:"hidden_notes>note"`
}

type XML_people struct {
	XML_persons []XML_person `xml:"person"`
}

type XML_person struct {
	XML_name        string `xml:"name,attr"`
	XML_description string `xml:"description,attr"`
	XML_text        string `xml:",chardata"`
}

type XML_objects struct {
	XML_boxes       []XML_box       `xml:"box"`
	XML_lockedBoxes []XML_lockedBox `xml:"lockedBox"`
	XML_keys        []XML_key       `xml:"key"`
}

type XML_box struct {
	XML_description string `xml:"description,attr"`
}

type XML_lockedBox struct {
	XML_id          string     `xml:"id,attr"`
	XML_description string     `xml:"description,attr"`
	XML_letter      *XML_letter `xml:"letter"`
}

type XML_letter struct {
	XML_description string `xml:"description,attr"`
	XML_text        string `xml:",chardata"`
}

type XML_key struct {
	XML_id          string `xml:"id,attr"`
	XML_description string `xml:"description,attr"`
}

type XML_hiddenNote struct {
	XML_keyword string `xml:"keyword,attr"`
	XML_text    string `xml:",chardata"`
}

func ReadFile(filename string) (View, error) {
  file, err := os.Open(filename)

  if err != nil {
    return View{}, err
  }
  defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
    return View{}, err
	}

	// Unmarshal the XML into the struct
	var view XML_view
	err = xml.Unmarshal(content, &view)
	if err != nil {
    return View{}, err
	}

  return View {
    storyNote: view.XML_storyNote,
  }, nil
}
