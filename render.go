package cvgen

func Render(fileName string, sections Sections, g PDFGenerator) error {
	return g.Generate(fileName, sections)
}
