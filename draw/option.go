package draw

type DrawOptions func(*DrawOption)

type DrawOption struct {
	Color *Color
}
