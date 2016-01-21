package object

type IDestroy interface {
	destroy()
}

//Разрушаемый объект - стены, столы, камни - всё, что потенциально можно разрушить.
type DestroyableObject struct {
	Name string
}
