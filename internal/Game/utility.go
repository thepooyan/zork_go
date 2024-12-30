package Game

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type VirtualOutput struct {
  Output string
}

func (v *VirtualOutput) clear() {
  v.Output = ""
}

func (v *VirtualOutput) write(content ...string) {
  l := strings.Join(content, " ")
  v.Output += l
  v.Output += "\n"
}

func (g *Game) Respond(write ...string) {
  g.VirtualOutput.write(write...)
}

func (g *Game) Tell(write ...string) {
  prefix := "| "
  modify := make([]string, 0)

  for _,i := range write {
    modify = append(modify, prefix + GetStyle(story).Render(i))
  }
  g.VirtualOutput.write(modify...)
}

func Input() string {
  reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
  return res
}

func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false // File does not exist
    }
    return err == nil // Return true if no error, false otherwise
}

func AddRandomSmalltalk(description string) string {
  sm := getSmalltalks()
  rand := rand.Intn(len(sm));
  return sm[rand] + " " + description
}

func getSmalltalks() []string {
  file, err := os.Open("./internal/Dict/small-talk.json")
  if err != nil {
    panic(err)
  }
  defer file.Close()
  var st []string
  decoder := json.NewDecoder(file)
  if err := decoder.Decode(&st); err != nil {
    panic(err)
  }
  return st
}

func RemoveXmlIndentation(text string) string {
  fls := strings.Fields(text)
  return strings.Join(fls, " ")
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s // Return the original string if it's empty
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func FilterInPlace[T any](slice *[]T, predicate func(T) bool) {
	newIndex := 0
	for _, item := range *slice {
		if predicate(item) { // Check if the item satisfies the predicate
			(*slice)[newIndex] = item // Place the item at the new index
			newIndex++               // Increment the new index
		}
	}
	// Resize the slice to keep only the filtered elements
	*slice = (*slice)[:newIndex]
}

func printBoxedText(texts []string, padding int) {
	maxLength := 0
	for _, text := range texts {
		if len(text) > maxLength {
			maxLength = len(text)
		}
	}

	textLength := maxLength + (2 * padding)
	border := strings.Repeat("*", textLength+4)

	fmt.Println(border)
	for _, text := range texts {
		paddedText := fmt.Sprintf("%s%s%s", strings.Repeat(" ", padding), text, strings.Repeat(" ", textLength-len(text)-padding))
		fmt.Printf("* %s *\n", paddedText)
	}
	fmt.Println(border)
}
