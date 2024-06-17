package text

type WriterOptions func(*WriterOption)

type WriterOption struct {
	Align Aligner
	Style styler
	Link  *string
	Width *Width
}
