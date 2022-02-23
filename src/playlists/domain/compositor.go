package domain

type Compositor struct {
	Name string
}

func NewCompositor(name string) *Compositor {
	return &Compositor{Name: name}
}
