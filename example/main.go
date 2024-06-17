package main

import (
	"log"
	"time"

	"github.com/ananrafs/cvgen"
	"github.com/ananrafs/cvgen/draw"
	"github.com/ananrafs/cvgen/text"
	"github.com/ananrafs/gopdf"
)

func main() {

	sections := []cvgen.Section{
		cvgen.PersonalInfoSection{
			Name:    "John Doe",
			Address: "Jalan santai kemana aja ayo ayo",
			Link: []cvgen.Link{
				{
					"08112111",
					"tel:08112111",
				}, {
					"08112111",
					"mailto:jogh@doe.doe",
				},
				{
					"08112111",
					"mailto:jogh@doe.doe",
				},
			},
		},
		cvgen.EducationSection{
			School: "University of Example",
			Degree: "B.Sc. Computer Science",
			Period: cvgen.Period{
				StartTime: time.Now().AddDate(-4, -1, 0),
				EndTime:   time.Now().AddDate(0, 1, 0),
			},
			Location: "Jakarta",
		},
		cvgen.ExperienceSection{
			WorkExperiences: []cvgen.WorkExperience{
				{
					Period: cvgen.Period{
						StartTime: time.Now().AddDate(-4, -1, 0),
						EndTime:   time.Now().AddDate(0, 1, 0),
					},
					Company: "Example Corp",
					Role:    "Software Engineer",
					JobDesc: []string{
						"Doing 1, Doing 2, Doing 3, Doing 4,Doing 5, Doing 6, Doing 7, Doing 8,Doing 9, Doing 10, Doing 11, Doing 12,Doing 13, Doing 14, Doing Nothing",
						"Helping people",
					},
					TechStack: []string{
						"Go", "Docker", "Python", "SQLServer", "Go", "Docker", "Python", "SQLServer", "Go", "Docker", "Python", "SQLServer", "Go", "Docker", "Python", "SQLServer", "Go", "Docker", "Python", "SQLServer",
					},
				},
				{
					Period: cvgen.Period{
						StartTime: time.Now().AddDate(-3, -2, 0),
						EndTime:   time.Now().AddDate(0, -1, 0),
					},
					Company: "Paypal Corp",
					Role:    "Software Engineer",
					JobDesc: []string{
						"Doing PR Review",
						"Fix Prod Issue",
					},
					TechStack: []string{
						"Go", "Docker", "Python", "SQLServer",
					},
				},
			},
		},
	}

	if err := cvgen.Render("cv.pdf", sections,
		NewGoPDFAdapter(gopdf.Config{PageSize: *gopdf.PageSizeA4})); err != nil {
		log.Fatal(err)
	}

}

type GoPDFAdapter struct {
	pdf *gopdf.GoPdf
	cfg *gopdf.Config
}

func NewGoPDFAdapter(cfg gopdf.Config) *GoPDFAdapter {
	pdf := &gopdf.GoPdf{}
	pdf.Start(cfg)
	return &GoPDFAdapter{
		pdf: pdf,
		cfg: &cfg,
	}
}
func (g *GoPDFAdapter) Generate(fileName string, sections cvgen.Sections) error {
	g.pdf.AddPage()
	err := g.pdf.AddTTFFont("Lato", "./ttf/Lato-Regular.ttf")
	if err != nil {

		return err
	}
	err = g.pdf.AddTTFFontWithOption("Lato", "./ttf/Lato-Bold.ttf", gopdf.TtfOption{
		Style: gopdf.Bold,
	})
	if err != nil {
		return err
	}
	err = g.pdf.AddTTFFontWithOption("Lato", "./ttf/Lato-Italic.ttf", gopdf.TtfOption{
		Style: gopdf.Italic,
	})
	if err != nil {
		return err
	}
	// h := 0
	g.pdf.Br(g.GetMargin())
	for _, section := range sections {
		section.Render(g)
		g.NextSection()
	}

	return g.pdf.WritePdf("../out/cv.pdf")
}
func (g *GoPDFAdapter) GetWidth() float64 {
	return g.cfg.PageSize.W
}
func (g *GoPDFAdapter) GetMargin() float64 {
	return float64(20)
}
func (g *GoPDFAdapter) Write(input string, opts ...text.WriterOptions) {
	defer func() {
		g.pdf.Br(5)
	}()
	wOpt := text.WriterOption{}
	for _, opt := range opts {
		opt(&wOpt)
	}

	switch wOpt.Style {
	case text.Content:
		g.pdf.SetFont("Lato", "", 12)
	case text.Subtitle:
		g.pdf.SetFont("Lato", "I", 13)
	case text.Title:
		g.pdf.SetFont("Lato", "B", 14)
	case text.ContentBold:
		g.pdf.SetFont("Lato", "B", 12)
	case text.Heading:
		g.pdf.SetFont("Lato", "B", 18)
	case text.MainHeading:
		g.pdf.SetFont("Lato", "B", 21)
	}

	textAlignMap := map[text.Aligner]int{
		text.Left:   gopdf.Left,
		text.Center: gopdf.Center,
		text.Right:  gopdf.Right,
	}

	currX, currY := g.GetPosition()
	cellWidth := g.GetWidth() - currX
	if wOpt.Width != nil {
		cellWidth = wOpt.Width.End - wOpt.Width.Start
		g.SetPosition(wOpt.Width.Start, currY)
	}

	if wOpt.Link != nil {
		width, _ := g.pdf.MeasureTextWidth(input)
		height, _ := g.pdf.MeasureCellHeightByText(input)
		prevX, prevY := g.GetPosition()

		var start, end float64
		switch wOpt.Align {
		case text.Center:
			start, end = prevX+0.5*cellWidth-0.5*width, prevX+0.5*cellWidth+0.5*width
		case text.Left:
			start, end = prevX, prevX+width
		case text.Right:
			start, end = prevX+cellWidth-width, prevX+cellWidth
		}

		g.pdf.SetTextColor(draw.Blue.GetRGB())
		g.pdf.MultiCellWithOption(&gopdf.Rect{
			W: cellWidth,
		}, input, gopdf.CellOption{Align: textAlignMap[wOpt.Align]})
		g.pdf.AddExternalLink(*wOpt.Link, start, prevY, end-start, height)
		g.pdf.SetTextColor(draw.Black.GetRGB())
		g.DrawLine(start, end, draw.WithColor(draw.Blue))
		return
	}

	g.pdf.MultiCellWithOption(&gopdf.Rect{
		W: cellWidth,
	}, input, gopdf.CellOption{Align: textAlignMap[wOpt.Align], BreakOption: &gopdf.BreakOption{Mode: gopdf.BreakModeIndicatorSensitive, BreakIndicator: ' '}})
}
func (g *GoPDFAdapter) Next() {
	g.pdf.Br(10)
}
func (g *GoPDFAdapter) NextSection() {
	g.pdf.Br(30)
}
func (g *GoPDFAdapter) SetPosition(x float64, y float64) {
	g.pdf.SetX(x)
	g.pdf.SetY(y)
}
func (g *GoPDFAdapter) GetPosition() (float64, float64) {
	return g.pdf.GetX(), g.pdf.GetY()
}
func (g *GoPDFAdapter) DrawLine(x1, x2 float64, opts ...draw.DrawOptions) {
	do := draw.DrawOption{}
	for _, opt := range opts {
		opt(&do)
	}

	if do.Color != nil {
		g.pdf.SetStrokeColor(do.Color.GetRGB())
	} else {
		g.pdf.SetStrokeColor(draw.Black.GetRGB())
	}

	g.pdf.SetLineWidth(2)
	g.pdf.SetLineType("solid")
	_, currY := g.GetPosition()
	g.pdf.Line(x1, currY, x2, currY)
}
