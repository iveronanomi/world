package item

type IInteraction interface {
	PickUp()
}

type Item struct {
	Name string
}
