package cvgen

type ExpertiseSection struct {
	Expertise map[string][]string
}

// Render renders the education section as a string
func (e ExpertiseSection) Render(w PDFWriter) {
	CommonRenderTitle("Expertise", w)

}
