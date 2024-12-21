package Game

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
	println("picked up ", p.weight, " grams of ", p.description)
}

func (p *Pickable) getWeight() int {
  return p.weight
}

//__________________

type Readable struct {
	*Object
	message string
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
  Open() []ObjectInt
  Add(item ObjectInt)
}

func (c *Container) Open() []ObjectInt {
	Respond("opening the ", c.description, "...")

	return c.content
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

type Unlockable struct {
	*Object
	id string
}

func (l *Lockable) Unlock(k Unlockable) {
	if l.id == k.id {
		println("Unlocked the", l.description)
    l.isLocked = false
	} else {
    println("the key does not match the lock")
  }
}
