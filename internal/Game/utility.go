package Game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func Outputln(write ...string) {
  join := strings.Join(write, "")
  fmt.Println(join)
}

func Output(write string) {
  fmt.Print(write)
}

func Input() string {
  reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
  return res
}

func GetUserInput() string {
  Output(" => ")
  response := Input()
  return response
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
