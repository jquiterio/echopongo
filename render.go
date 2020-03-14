package echopongo

import (
	"path"

	"github.com/jquiterio/pongo2"
	"github.com/labstack/echo/v4"
)

// RenderOptions : render configuration for echopongo
type RenderOptions struct {
	TemplateDir string
	ContentType string
}

// Pongo2Render : Render template with context and options
type Pongo2Render struct {
	Options  *RenderOptions
	Template *pongo2.Template
	Content  pongo2.Context
}

//New creates new Pongo2Render
func New(opts RenderOptions) *Pongo2Render {
	return &Pongo2Render{
		Options: &opts,
	}
}

func Default() *Pongo2Render {
	return New(RenderOptions{
		TemplateDir: "templates",
		ContentType: "text/html; charset=utf-8",
	})
}

// func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
// 	if viewContext, isMap := data.(map[string]interface{}); isMap {
// 		viewContext["reverse"] = c.Echo().Reverse
// 	}
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func (p2r Pongo2Render) Instance(name string, data interface{}) echo.Render {
	var template *pongo2.Template
	filename := path.Join(p2r.Options.TemplateDir, name)
	template = pongo2.Must(pongo2.FromFile(filename))
	return Pongo2Render{
		Template: template,
		Context:  data.(pongo2.Context),
		Options:  p2r.Options,
	}
}

func (pr Pongo2Render) Render(w httpp.ResponseWriter) error {
	header := w.Header()
	if len(header["Content-Type"]) == 0 {
		header["Content-Type"] = []string{pr.Options.ContentType}
	}
	//pr.WriteContentType(w)
	err := pr.Template.ExecuteWriter(p.Context, w)
	return err
}
