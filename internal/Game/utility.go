package Game

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func Outputln(write ...string) {
  join := strings.Join(write, "")
  fmt.Println(join)
}

func Output(write string) {
  fmt.Print(write)
}

func Respond(write ...string) {
  folan := append([]string{"\n "}, write...)
  folan = append(folan, "\n")
  Outputln(folan...)
}

func Describe(write ...string) {
  prefix := "-     "
  modify := make([]string, 0)

  for _,i := range write {
    modify = append(modify, prefix+i)
  }
  Outputln(modify...)
}

func Input() string {
  reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
  return res
}

func GetUserInput() string {
  Output("\n\n => ")
  response := Input()
  Output("\n")
  return strings.TrimSuffix(response, "\n")
}

func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false // File does not exist
    }
    return err == nil // Return true if no error, false otherwise
}

type CarryWeight struct {
  value int
  max int
}

func (c *CarryWeight) test(amount int) bool {
  return c.value + amount <= c.max
}

func (c *CarryWeight) Add(amount int) (int, error) {
  if c.test(amount) {
     c.value += amount
     return amount, nil
  }
  return amount, errors.New("More than you can carry!")
}

func (c *CarryWeight) Drop(amount int) {
  c.value -= amount
}

type Inventory struct {
  items []interface{}
  CarryWeight
}

func (i *Inventory) Add(item PickableInt) bool {
  weight := item.getWeight()
  _,err := i.CarryWeight.Add(weight)
  if err != nil {
    return false
  }

  return true
}

func ListStuff(stuff []ObjectInt) {
  for _,i := range stuff {
    Outputln(i.getDescription())
  }
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
