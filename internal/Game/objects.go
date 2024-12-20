package Game

type Object struct {
	description string
}

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

type LockedBox struct {
  *Object
  *Container
  *Pickable
  *Lockable
}

func NewLockedBox(description, id string) LockedBox {
  obj := &Object{description}
  return LockedBox{
    obj,
    &Container{
      obj,
      make([]interface{}, 0),
    },
    &Pickable{
      obj,
      10,
    },
    &Lockable{
      obj,
      id,
      true,
    },
  }
}

type Key struct {
  *Object
  *Pickable
  *Unlockable
}

func NewKey(description,id string) Key {
  obj := &Object{description}
  return Key{
    obj,
    &Pickable{
      obj,
      2,
    },
    &Unlockable{
      obj,
      id,
    },
  }
}
