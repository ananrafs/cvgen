package cvgen

import (
	"fmt"
	"log"

	"github.com/ananrafs/cvgen/text"
)

// EducationSection represents the education section
type EducationSection struct {
	School   string
	Location string
	Degree   string
	Period   Period
}

// Render renders the education section as a string
func (e EducationSection) Render(w PDFWriter) {
	CommonRenderTitle("Education", w)
	widthRatio, err := getWidth(w.GetWidth()-(2*w.GetMargin()), 70, 30)
	if err != nil {
		log.Fatal(err)
	}

	j := 0
	for {
		margin := w.GetMargin()
		_, currY := w.GetPosition()
		for _, width := range widthRatio {
			w.SetPosition(margin, currY)
			switch j {
			case 0:
				w.Write(
					fmt.Sprintf("%s, %s", e.School, e.Location),
					text.WithStyle(text.Title),
					text.WithWidth(margin, w.GetWidth()-w.GetMargin()),
				)
			case 1:
				w.Write(
					e.Period.String(),
					text.WithAlign(text.Right),
					text.WithStyle(text.ContentBold),
					text.WithWidth(margin, w.GetWidth()-w.GetMargin()),
				)
			case 2:
				w.Write(
					e.Degree,
					text.WithStyle(text.Subtitle),
					text.WithWidth(margin, w.GetWidth()-w.GetMargin()),
				)
			}

			margin += width
			j++
		}

		if j >= 3 {
			break
		}
	}

}
