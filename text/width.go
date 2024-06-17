package text

type Width struct {
	Start float64
	End   float64
}

func WithWidth(start, end float64) WriterOptions {
	return func(wo *WriterOption) {
		wo.Width = &Width{
			Start: start, End: end,
		}
	}
}
