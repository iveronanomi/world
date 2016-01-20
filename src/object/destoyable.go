package object

type IDestroy interface {
	destroy()
}

type DestroyableObject struct {
	Name string
}
