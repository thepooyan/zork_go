package main


type Object struct {
  name string
}

type Pickable struct {
  *Object
  weight int
}

func (p *Pickable) Pickup() {
  println("picked up ", p.weight, " grams of ", p.name);
}

//__________________

type Readable struct {
  *Object
  message string
}

func (r *Readable) Read() {
  println("Reading the", r.name,"=>", r.message)
}

//_________________

type Container struct {
  *Object
  content []interface{}
}

func (c *Container) Open() []interface{} {
  println("Opening the ", c.name, "...")

  return c.content
}

func (c *Container) Add(item interface{}) {
  c.content = append(c.content, item)
}

//______________________

type Letter struct {
  *Object
  *Readable
  *Pickable
}

func NewLetter(message, description string) Letter {
  Object := Object{description}
  return Letter{
    &Object,
    &Readable{
      &Object,
      message,
    },
    &Pickable{
      &Object,
      5,
    },
  }
}


type Box struct {
  *Object
  *Pickable
  *Container
}

func NewBox(description string) Box {
  Object := Object{description}
  return Box{
    &Object,
    &Pickable{
      &Object,
      20,
    },
    &Container{
      &Object,
      make([]interface{}, 0),
    },
  }
}

func findLetterInBox(b Box) {
  c := b.Open()
  for _,item := range c {
    switch v := item.(type) {
    case Letter:
      println("found letter")
      v.Read()
    default:
      println("found nothing")
    }
  }
}

func main() {
  l := NewLetter("hi!", "old letter")
  b := NewBox("rusty box")
  b.Add(l)
  findLetterInBox(b)
}
