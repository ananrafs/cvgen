package cvgen

import (
	"log"

	"github.com/ananrafs/cvgen/text"
)

// ExperienceSection represents the experience section
type ExperienceSection struct {
	WorkExperiences []WorkExperience
}

type WorkExperience struct {
	Period  Period
	Company string
	Role    string
	JobDesc []string
}

// Render renders the experience section as a string
func (e ExperienceSection) Render(w PDFWriter) {
	CommonRenderTitle("Work Experience", w)

	widthRatio, err := getWidth(w.GetWidth()-(2*w.GetMargin()), 70, 30)
	if err != nil {
		log.Fatal(err)
	}

	for _, we := range e.WorkExperiences {
		j := 0
		for {
			margin := w.GetMargin()
			_, currY := w.GetPosition()
			for _, width := range widthRatio {
				switch j {
				case 0:
					w.SetPosition(margin, currY)
					w.Write(
						we.Company,
						text.WithStyle(text.Title),
						text.WithWidth(margin, w.GetWidth()-w.GetMargin()),
					)
				case 1:
					w.SetPosition(margin, currY)
					w.Write(
						we.Period.String(),
						text.WithAlign(text.Right),
						text.WithStyle(text.ContentBold),
						text.WithWidth(margin, w.GetWidth()-w.GetMargin()),
					)
				case 2:
					w.Write(
						we.Role,
						text.WithStyle(text.Subtitle),
						text.WithWidth(margin, w.GetWidth()-w.GetMargin()),
					)
				}

				margin += width
				j++
			}
			if j > 2 {
				break
			}
		}

		for _, jd := range we.JobDesc {
			_, currY := w.GetPosition()
			w.SetPosition(2*w.GetMargin(), currY)
			w.Write("-")
			w.SetPosition(0, currY)
			w.Write(jd,
				text.WithWidth(2*w.GetMargin()+10, w.GetWidth()-w.GetMargin()),
			)
		}

		w.Next()
	}
}
