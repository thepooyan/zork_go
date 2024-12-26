package Game

import (
	"strconv"
)

type Object struct {
	description string
  id string
}

type ObjectInt interface {
  getDescription() string
  equals(i ObjectInt) bool
  getId() string
}

func (o *Object) getDescription() string {
  return o.description
}

func (o *Object) getId() string {
  return o.id
}

func (o *Object) equals(i ObjectInt) bool {
  return i.getId() == o.id
}

// _____________________

type Pickable struct {
	*Object
	weight int
}

type PickableInt interface  {
  ObjectInt
  getWeight() int
  Pickup(g *Game)
}

func (p *Pickable) Pickup(g *Game) {
  g.Inventory.Add(p)
  g.RemoveFromScene(p)
	Respond("picked up ", strconv.Itoa(p.weight) , " grams of ", p.description)
}

func (p *Pickable) getWeight() int {
  return p.weight
}

//__________________

type Readable struct {
	*Object
	message string
}

type ReadableInt interface {
  Read()
}

func (r *Readable) Read() {
  Respond("Reading the ", r.description, ":")
  printBoxedText([]string{r.message}, 10)
}

//_________________

type Container struct {
	*Object
	content []ObjectInt
}

type ContainerInt interface {
  ObjectInt
  Open(g *Game)
  Add(item ObjectInt)
}

func (c *Container) Open(g *Game) {
  c.ListStuff()
  g.currentView.Objects = append(g.currentView.Objects, c.content...)
  c.EmptyStuff()
}

func (c *Container) ListStuff() {
  if len(c.content) == 0 {
    Respond("the ", c.getDescription()," is empty!")
    return
  }
  Respond("inside the ", c.getDescription(), ":")
  inside := make([]string, 0)
  for _,i := range c.content {
    inside = append(inside, i.getDescription())
  }
  printBoxedText(inside, 5)
}

func (c *Container) EmptyStuff() {
  c.content = make([]ObjectInt, 0)
}

func (c *Container) Add(item ObjectInt) {
	c.content = append(c.content, item)
}

//______________________

type Lockable struct {
  *Object
	id       string
	isLocked bool
}

type LockInt interface {
  ObjectInt
  Unlock(k KeyInt) bool
}

type Unlocker struct {
	*Object
	id string
}

type KeyInt interface {
  getKeyId() string
}

func (l *Lockable) Unlock(k KeyInt) bool {
	if l.id == k.getKeyId() {
		Respond("unlocked the ", l.description)
    l.isLocked = false
    return true
	} else {
    Respond("the key does not match the lock")
    return false
  }
}

func (u *Unlocker) getKeyId() string {
  return u.id
}

//---------------------------------

type LockedContainer struct {
  *Object
  *Lockable
  *Container
}

func (l *LockedContainer) Open(g *Game) {
  if (l.Lockable.isLocked) {
    Respond("the ", l.getDescription()," is Locked. you have to unlock it first")
  } else {
    l.Container.Open(g)
  }
}

func (l *LockedContainer) Unlock(k KeyInt) bool {
  result := l.Lockable.Unlock(k)
  Respond("now you can open it")
  // l.Container.Open(g *Game)
  return result
}
