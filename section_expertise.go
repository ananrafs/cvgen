package cvgen

import (
	"log"
	"strings"

	"github.com/ananrafs/cvgen/text"
)

type ExpertiseSection struct {
	Expertise map[string][]string
}

// Render renders the education section as a string
func (e ExpertiseSection) Render(w PDFWriter) {
	CommonRenderTitle("Technical Skills", w)

	widthRatio, err := getWidth(w.GetWidth()-(2*w.GetMargin()), 30, 70)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range e.Expertise {
		_, y := w.GetPosition()
		w.Write(
			k,
			text.WithStyle(text.ContentBold),
			text.WithWidth(w.GetMargin(), widthRatio[0]),
		)
		w.SetPosition(widthRatio[0], y)
		w.Write(
			strings.Join(v, ", "),
			text.WithStyle(text.Content),
			text.WithWidth(w.GetMargin()+widthRatio[0], w.GetWidth()-w.GetMargin()),
		)
	}
}
