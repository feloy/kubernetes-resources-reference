package kwebsite

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/feloy/kubernetes-api-reference/pkg/formats/markdown"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
)

// Chapter of a KWebsite output
// implements the outputs.Chapter interface
type Chapter struct {
	kwebsite *KWebsite
	part     *Part
	name     string
	data     *ChapterData
}

type ChapterData struct {
	ApiVersion  string
	Version     string
	Import      string
	Kind        string
	Metadata    ChapterMetadata
	ChapterName string
	Sections    []SectionData
}

type ChapterMetadata struct {
	Description string
	Title       string
	Weight      int
}

type SectionData struct {
	Name            string
	Description     string
	Fields          []FieldData
	FieldCategories []FieldCategoryData
}

type FieldCategoryData struct {
	Name   string
	Fields []FieldData
}

type FieldData struct {
	Name           string
	Value          string
	Description    string
	TypeDefinition string
	Indent         int
}

// SetAPIVersion writes the APIVersion for a chapter
func (o Chapter) SetAPIVersion(s string) error {
	err := o.kwebsite.addContent(o.part.name, o.name, markdown.Code("apiVersion: "+s)+"\n")
	if err != nil {
		return fmt.Errorf("Error adding GV for chapter %s/%s: %s", o.part.name, o.name, err)
	}
	return nil
}

// SetGoImport writes the Go import for a chapter
func (o Chapter) SetGoImport(s string) error {
	err := o.kwebsite.addContent(o.part.name, o.name, markdown.Code("import \""+s+"\"")+"\n")
	if err != nil {
		return fmt.Errorf("Error adding Go Import for chapter %s/%s: %s", o.part.name, o.name, err)
	}
	return nil
}

// AddSection adds a section to the chapter
func (o Chapter) AddSection(i int, name string, apiVersion *string) (outputs.Section, error) {
	o.data.Sections = append(o.data.Sections, SectionData{
		Name: name,
	})

	err := o.kwebsite.addSection(o.part.name, o.name, name)
	if err != nil {
		return Section{}, err
	}
	return Section{
		kwebsite: o.kwebsite,
		part:     o.part,
		chapter:  &o,
	}, nil
}

func (o Chapter) Write() error {
	chaptername := escapeName(o.data.ChapterName + "-" + o.data.Version)
	filename := filepath.Join(o.kwebsite.Directory, o.part.name, chaptername) + "-new.md"
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	t := template.Must(template.New("chapter.tmpl").Funcs(sprig.TxtFuncMap()).ParseFiles("./pkg/outputs/kwebsite/templates/chapter.tmpl"))
	return t.Execute(f, o.data)
}
