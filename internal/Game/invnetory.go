package Game

import "errors"

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

func (inv *Inventory) Describe() {
  for _,i := range inv.items {
    Respond(i.getDescription())
  }
}
