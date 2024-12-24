package Game

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
      make([]ObjectInt, 0),
    },
  }
}

type LockedBox struct {
  *Object
  *Pickable
  *LockedContainer
}

func NewLockedBox(description, id string) LockedBox {
  obj := &Object{description}
  return LockedBox{
    obj,
    &Pickable{
      obj,
      10,
    },
    &LockedContainer{
      obj,
      &Lockable{
        obj,
        id,
        true,
      },
      &Container{
        obj,
        make([]ObjectInt, 0),
      },
    },
  }
}

type Key struct {
  *Object
  *Pickable
  *Unlocker
}

func NewKey(description,id string) Key {
  obj := &Object{description}
  return Key{
    obj,
    &Pickable{
      obj,
      2,
    },
    &Unlocker{
      obj,
      id,
    },
  }
}
