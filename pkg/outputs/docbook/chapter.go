package docbook

import (
	"github.com/feloy/kubernetes-api-reference/pkg/formats/dbxml"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
	x "github.com/shabbyrobe/xmlwriter"
)

// DocbookChapter represents a Docbook chapter
type DocbookChapter struct {
	id string
	w  *x.Writer
}

// SetAPIVersion writes the APIVersion for a chapter
func (o DocbookChapter) SetAPIVersion(s string) error {
	o.w.StartElem(x.Elem{Name: "para"})
	o.w.StartElem(dbxml.ElemWithText("varname", "apiVersion: "+s))
	o.w.EndElem("varname")
	o.w.EndElem("para")
	return nil
}

// SetGoImport writes the Go import for a chapter
func (o DocbookChapter) SetGoImport(s string) error {
	o.w.StartElem(x.Elem{Name: "para"})
	o.w.StartElem(dbxml.ElemWithText("varname", "import \""+s+"\""))
	o.w.EndElem("varname")
	o.w.EndElem("para")
	return nil
}

// AddSection adds a section to the Docbook chapter
func (o DocbookChapter) AddSection(i int, name string, apiVersion *string) (outputs.Section, error) {
	if i > 0 {
		o.w.EndToDepth(sectionDepth, x.ElemNode, "sect1")
	}
	section := dbxml.Section("sect1", name)
	err := o.w.StartElem(section)
	if err != nil {
		return DocbookSection{}, err
	}
	o.w.WriteAttr(x.Attr{
		Name:  "id",
		Value: escapeName(o.id + "." + name),
	})
	err = o.w.WriteAttr(x.Attr{
		Name:  "renderas",
		Value: "sect2",
	})
	if apiVersion != nil {
		o.w.StartElem(dbxml.IndexTerm("resources", name, *apiVersion))
		o.w.EndElem("indexterm")
	}
	if err != nil {
		return DocbookSection{}, err
	}
	return DocbookSection{w: o.w}, nil
}
