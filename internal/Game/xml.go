package Game

import (
	"encoding/xml"
	"io"
	"os"
)

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
	var view View
	err = xml.Unmarshal(content, &view)
	if err != nil {
    return View{}, err
	}

  return View {
    StoryNote: view.StoryNote,
  }, nil
}
