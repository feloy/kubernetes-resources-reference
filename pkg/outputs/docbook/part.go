package docbook

import (
	"github.com/feloy/kubernetes-api-reference/pkg/formats/dbxml"
	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
	x "github.com/shabbyrobe/xmlwriter"
)

// DocbookPart represents a Part of the Docbook output
type DocbookPart struct {
	w *x.Writer
}

// AddChapter adds a chapter to the Docbook part
func (o DocbookPart) AddChapter(i int, name string, gv string, version *kubernetes.APIVersion, description string, importPrefix string) (outputs.Chapter, error) {
	if i > 0 {
		o.w.EndToDepth(chapterDepth, x.ElemNode, "chapter")
	}
	id := name + " " + version.String()

	chapterName := name
	if version != nil && version.Stage != kubernetes.StageGA {
		chapterName += " " + version.String()
	}

	chapter := dbxml.Section("chapter", chapterName)
	return DocbookChapter{id: id, w: o.w}, o.w.StartElem(chapter)
}
