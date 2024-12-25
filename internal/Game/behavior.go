package Game

import "strconv"

type Object struct {
	description string
}

type ObjectInt interface {
  getDescription() string
}

func (o *Object) getDescription() string {
  return o.description
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
	println("Reading the", r.description, "=>", r.message)
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
	Respond("opening the ", c.description, "...")
  g.currentView.Objects = append(g.currentView.Objects, c.content...)
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
  Unlock(k KeyInt)
}

type Unlocker struct {
	*Object
	id string
}

type KeyInt interface {
  getKeyId() string
}

func (l *Lockable) Unlock(k KeyInt) {
	if l.id == k.getKeyId() {
		println("Unlocked the", l.description)
    l.isLocked = false
	} else {
    println("the key does not match the lock")
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
    Respond("can't open the ", l.getDescription(),". it's Locked")
  } else {
    l.Container.Open(g)
  }
}
