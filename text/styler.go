package text

type styler int

const (
	Content     styler = iota
	ContentBold styler = iota
	Subtitle    styler = iota
	Title       styler = iota
	Heading     styler = iota
	MainHeading styler = iota
)

func WithStyle(style styler) WriterOptions {
	return func(wo *WriterOption) {
		wo.Style = style
	}
}
