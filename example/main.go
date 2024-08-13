package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ananrafs/cvgen"
	"github.com/ananrafs/cvgen/draw"
	"github.com/ananrafs/cvgen/text"
	"github.com/ananrafs/gopdf"
)

func main() {

	defaultSections := []cvgen.Section{
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
		cvgen.SummarySection{
			Summary: "Results-driven Automation Engineer with 5+ years of experience in designing, developing, and implementing automated systems and processes. Proficient in various automation tools and programming languages, with a proven track record of improving efficiency and reducing operational costs.",
		},
		cvgen.ExpertiseSection{
			Expertise: map[string][]string{
				"Programming Language": {"Golang", "C#", "Javascript", "Java"},
				"Tools":                {"Jenkins", "Postman", "git", "docker"},
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
				},
			},
		},
	}
	// if err := CreateDummy(defaultSections, "dummy.json"); err != nil {
	// 	log.Fatal(err)
	// }

	sections, err := GetSection("cv.json")
	if err != nil {
		fmt.Println(err)
		sections = defaultSections
	}

	if err := cvgen.Render("cv.pdf", sections,
		NewGoPDFAdapter(gopdf.Config{PageSize: *gopdf.PageSizeA4})); err != nil {
		log.Fatal(err)
	}
}

func GetSection(filename string) (cvgen.Sections, error) {
	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a variable of the appropriate type to hold the unmarshaled data
	var sections cvgen.Sections

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON data
	err = decoder.Decode(&sections)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func CreateDummy(s cvgen.Sections, fileName string) error {
	// Marshal the struct to JSON.
	jsonData, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return err
	}

	// Create and open a file for writing.
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the JSON data to the file.
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
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

	return g.pdf.WritePdf("./out/cv.pdf")
}
func (g *GoPDFAdapter) GetWidth() float64 {
	return g.cfg.PageSize.W
}
func (g *GoPDFAdapter) GetMargin() float64 {
	return float64(20)
}
func (g *GoPDFAdapter) Write(input string, opts ...text.WriterOptions) {
	defer func() {
		g.pdf.Br(2)
	}()
	wOpt := text.WriterOption{}
	for _, opt := range opts {
		opt(&wOpt)
	}

	var lineSpacing float64 = 2
	switch wOpt.Style {
	case text.Content:
		g.pdf.SetFont("Lato", "", 12)
		lineSpacing = 1
	case text.Subtitle:
		g.pdf.SetFont("Lato", "I", 13)
	case text.Title:
		g.pdf.SetFont("Lato", "B", 13)
		lineSpacing = 1
	case text.ContentBold:
		g.pdf.SetFont("Lato", "B", 12)
	case text.Heading:
		g.pdf.SetFont("Lato", "B", 17)
		lineSpacing = 1
	case text.MainHeading:
		g.pdf.SetFont("Lato", "B", 20)
		lineSpacing = 1
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

		g.pdf.MultiCellWithOption(&gopdf.Rect{
			W: cellWidth,
		}, input,
			gopdf.CellOption{
				Align:       textAlignMap[wOpt.Align],
				BreakOption: &gopdf.BreakOption{Mode: gopdf.BreakModeIndicatorSensitive, BreakIndicator: ' '},
				LineSpacing: &lineSpacing,
			})
		g.pdf.AddExternalLink(*wOpt.Link, start, prevY, end-start, height)
		g.pdf.SetTextColor(draw.Black.GetRGB())
		x, y := g.GetPosition()
		g.SetPosition(x, y-1)
		g.DrawLine(start, end)
		return
	}

	g.pdf.MultiCellWithOption(&gopdf.Rect{
		W: cellWidth,
	}, input,
		gopdf.CellOption{
			Align:       textAlignMap[wOpt.Align],
			BreakOption: &gopdf.BreakOption{Mode: gopdf.BreakModeIndicatorSensitive, BreakIndicator: ' '},
			LineSpacing: &lineSpacing,
		})
}
func (g *GoPDFAdapter) Next() {
	g.pdf.Br(8)
}
func (g *GoPDFAdapter) NextSection() {
	g.pdf.Br(12)
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
