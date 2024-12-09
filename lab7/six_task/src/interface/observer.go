package observer

type Observer interface {
	Update(o Observable)
}
