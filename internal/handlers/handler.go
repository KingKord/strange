package handlers

type Handler interface {
	root()
}
type handlers struct {
}

func NewHandlers() Handler {
	return handlers{}
}

func (h handlers) root() {

}
