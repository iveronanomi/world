package item

type IInteraction interface {
	PickUp()
}

//Предмет - то, что можно поднять с земли и поместить в инвентарь.
type Item struct {
	Name string
}
