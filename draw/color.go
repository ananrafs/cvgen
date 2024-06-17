package draw

type Color [3]uint8

func (c *Color) GetRGB() (uint8, uint8, uint8) {
	return (*c)[0], (*c)[1], (*c)[2]
}

var (
	Red        Color = Color{255, 0, 0}
	Green      Color = Color{0, 255, 0}
	Blue       Color = Color{0, 0, 255}
	Yellow     Color = Color{255, 255, 0}
	Cyan       Color = Color{0, 255, 255}
	Magenta    Color = Color{255, 0, 255}
	Black      Color = Color{0, 0, 0}
	White      Color = Color{255, 255, 255}
	Orange     Color = Color{255, 165, 0}
	Purple     Color = Color{128, 0, 128}
	Brown      Color = Color{165, 42, 42}
	Pink       Color = Color{255, 192, 203}
	Gray       Color = Color{128, 128, 128}
	LightGray  Color = Color{211, 211, 211}
	DarkGray   Color = Color{169, 169, 169}
	LightBlue  Color = Color{173, 216, 230}
	DarkBlue   Color = Color{0, 0, 139}
	LightGreen Color = Color{144, 238, 144}
	DarkGreen  Color = Color{0, 100, 0}
	LightRed   Color = Color{255, 182, 193}
	DarkRed    Color = Color{139, 0, 0}
)

func WithColor(color Color) DrawOptions {
	return func(do *DrawOption) {
		do.Color = &color
	}
}
