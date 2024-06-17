package cvgen

import (
	"encoding/json"

	"github.com/ananrafs/cvgen/draw"
	"github.com/ananrafs/cvgen/text"
)

type Section interface {
	Render(w PDFWriter)
}

type Sections []Section

func (s *Sections) UnmarshalJSON(data []byte) error {
	var m []json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	for _, raw := range m {
		var identifier = struct {
			Type string `json:"__type"`
		}{}
		if err := json.Unmarshal(raw, &identifier); err != nil {
			return err
		}
		var instance Section
		switch identifier.Type {
		case "edu":
			instance = new(EducationSection)
		case "perso":
			instance = new(PersonalInfoSection)
		case "ts":
			instance = new(ExpertiseSection)
		case "sum":
			instance = new(SummarySection)
		case "exp":
			instance = new(ExperienceSection)
		}
		if err := json.Unmarshal(raw, instance); err != nil {
			return err
		}
		*s = append(*s, instance)
	}
	return nil
}

type PDFWriter interface {
	GetWidth() float64
	GetMargin() float64
	Write(string, ...text.WriterOptions)
	Next()
	NextSection()
	SetPosition(x float64, y float64)
	GetPosition() (x float64, y float64)
	DrawLine(x1, x2 float64, opts ...draw.DrawOptions)
}

type PDFGenerator interface {
	Generate(string, Sections) error
}

func CommonRenderTitle(title string, w PDFWriter) {
	w.Write(title,
		text.WithStyle(text.Heading),
		text.WithWidth(w.GetMargin(), w.GetWidth()-w.GetMargin()),
	)
	w.DrawLine(w.GetMargin()-5, w.GetWidth()-w.GetMargin()+5)
	w.Next()
}
