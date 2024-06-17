package text

type Aligner int

const (
	Left   Aligner = iota
	Center Aligner = iota
	Right  Aligner = iota
)

func WithAlign(align Aligner) WriterOptions {
	return func(wo *WriterOption) {
		wo.Align = align
	}
}
