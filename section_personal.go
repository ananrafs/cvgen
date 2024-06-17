package cvgen

import (
	"log"

	"github.com/ananrafs/cvgen/text"
)

type PersonalInfoSection struct {
	Name    string
	Address string
	Link    []Link
}

type Link struct {
	Text     string
	Redirect string
}

// Render renders the personal information section as a string
func (p PersonalInfoSection) Render(w PDFWriter) {
	w.Write(
		p.Name,
		text.WithAlign(text.Center),
		text.WithStyle(text.MainHeading),
	)
	w.Write(
		p.Address,
		text.WithAlign(text.Center),
		text.WithStyle(text.Content),
	)

	widthRatio, err := getWidth(w.GetWidth(), 49, 2, 49)
	if err != nil {
		log.Fatal(err)
	}

	j := 0
	for {
		if j >= len(p.Link) {
			break
		}

		x := float64(0)
		_, currY := w.GetPosition()
		for i, width := range widthRatio {
			var (
				start = x
				end   float64
				align text.WriterOptions
			)
			switch i {
			case 0:
				end = width
				align = text.WithAlign(text.Right)
			case 1:
				x += width
				continue
			case 2:
				end = w.GetWidth()
				align = text.WithAlign(text.Left)
			}
			w.SetPosition(start, currY)
			w.Write(
				p.Link[j].Text,
				text.WithLink(p.Link[j].Redirect),
				text.WithWidth(start, end),
				align,
			)
			x += width
			j++
		}

		if len(p.Link)-j == 1 {
			w.Write(
				p.Link[j].Text,
				text.WithLink(p.Link[j].Redirect),
				text.WithAlign(text.Center),
			)
			j++
		}

	}

}
