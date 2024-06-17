package text

func WithLink(link string) WriterOptions {
	return func(wo *WriterOption) {
		wo.Link = &link
	}
}
