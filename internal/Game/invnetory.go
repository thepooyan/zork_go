package Game

import (
	"errors"
	"strconv"
	"strings"
)

type Inventory struct {
  items []PickableInt
  CarryWeight
}

func (i *Inventory) Add(item PickableInt) bool {
  weight := item.getWeight()
  _,err := i.CarryWeight.Add(weight)
  if err != nil {
    return false
  }
  i.items = append(i.items, item)
  return true
}

func (inv *Inventory) Find(itemName string) ( PickableInt, bool ) {
  for _,i := range inv.items {
    if strings.Contains(i.getDescription(), itemName) {
      return i, true
    }
  }
  return nil, false
}

func (inv *Inventory) Remove(item PickableInt) {
  FilterInPlace(&inv.items, func(i PickableInt) bool { return !i.equals(item) })
  inv.CarryWeight.Drop(item.getWeight())
}

func (inv *Inventory) Describe() {
  if len(inv.items) == 0 {
    Respond("Inventory empty")
    return
  }
  descs := make([]string, 0)
  for _,i := range inv.items {
    descs = append(descs, i.getDescription())
  }
  printBoxedText(descs, 5)
}

type CarryWeight struct {
  Value int
  max int
}

func (c *CarryWeight) test(amount int) bool {
  return c.Value + amount <= c.max
}

func (c *CarryWeight) Add(amount int) (int, error) {
  if c.test(amount) {
     c.Value += amount
     return amount, nil
  }
  return amount, errors.New("More than you can carry!")
}

func (c *CarryWeight) Drop(amount int) {
  c.Value -= amount
}

func (c *CarryWeight) Print() {
  Respond("the weight you're carrying: ", strconv.Itoa(c.Value))
}
