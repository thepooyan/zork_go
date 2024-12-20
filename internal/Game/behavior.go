package Game

type Pickable struct {
	*Object
	weight int
}

func (p *Pickable) Pickup() {
	println("picked up ", p.weight, " grams of ", p.description)
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
	content []interface{}
}

func (c *Container) Open() []interface{} {
	println("Opening the ", c.description, "...")

	return c.content
}

func (c *Container) Add(item interface{}) {
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
