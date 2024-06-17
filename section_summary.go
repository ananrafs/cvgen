package cvgen

import (
	"github.com/ananrafs/cvgen/text"
)

type SummarySection struct {
	Summary string
}

func (p SummarySection) Render(w PDFWriter) {
	CommonRenderTitle("Summary", w)
	w.Write(
		p.Summary,
		text.WithStyle(text.Content),
		text.WithWidth(w.GetMargin(), w.GetWidth()-w.GetMargin()),
	)

}
