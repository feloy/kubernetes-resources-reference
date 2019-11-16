package structure

import (
	"fmt"
	"io"
	"strings"
)

// Section at any level in the resulting document
type Section struct {
	Name        string
	Description *string
	ID          *string
	SubSections []*Section
	FieldList   []*FieldEntry
	IsAppendix  bool
}

// NewSection creates a new section with the given name
func NewSection(name string, description *string) *Section {
	return &Section{Name: name, Description: description}
}

// SetID sets the ID of the section
func (o *Section) SetID(id string) {
	o.ID = &id
}

// SetAppendix indicates that the section is an appendix
func (o *Section) SetAppendix() {
	o.IsAppendix = true
}

// AddSection adds a new subsection
func (o *Section) AddSection(s *Section) {
	o.SubSections = append(o.SubSections, s)
}

// AddFieldEntry adds a field entry in the list of fields of the section
func (o *Section) AddFieldEntry(f *FieldEntry) {
	o.FieldList = append(o.FieldList, f)
}

// AsMarkdown dumps the section as markdown
func (o *Section) AsMarkdown(w io.Writer, depths ...int) {
	depth := 0
	if len(depths) > 0 {
		depth = depths[0]
	}
	if depth > 0 {
		for i := 0; i < depth; i++ {
			fmt.Fprintf(w, "#")
		}
		fmt.Fprintf(w, " %s\n", o.Name)
	}

	for _, field := range o.FieldList {
		field.AsMarkdown(w)
	}

	for _, section := range o.SubSections {
		section.AsMarkdown(w, depth+1)
	}
}

// AsDocbook dumps the section as Docbook
func (o *Section) AsDocbook(w io.Writer, depths ...int) {
	// Create an index entry for sections at this level
	indexLevel := 3

	normalTags := []string{"", "part", "chapter", "sect1 renderas=\"sect2\"", "bridgehead renderas=\"sect4\""}
	appendixTags := []string{"", "", "appendix"}

	tags := normalTags
	if o.IsAppendix {
		tags = appendixTags
	}

	depth := 0
	if len(depths) > 0 {
		depth = depths[0]
	}

	isBridgeHead := false
	if depth > 0 {
		ID := ""
		if o.ID != nil {
			ID = fmt.Sprintf(" id=\"%s\"", *o.ID)
		}
		if strings.HasPrefix(tags[depth], "bridgehead") {
			isBridgeHead = true
		}
		if isBridgeHead {
			fmt.Fprintf(w, "<%s%s>%s</bridgehead>\n", tags[depth], ID, o.Name)
		} else {
			if depth == indexLevel {
				fmt.Fprintf(w, "<%s%s><title><indexterm type=\"resources\"><primary>%s</primary></indexterm>%s</title>\n", tags[depth], ID, o.Name, o.Name)
			} else {
				fmt.Fprintf(w, "<%s%s><title>%s</title>\n", tags[depth], ID, o.Name)
			}
		}
	}

	if o.Description != nil {
		fmt.Fprintf(w, "<para>%s</para>", *o.Description)
	}

	if len(o.FieldList) > 0 {
		fmt.Fprintf(w, "<variablelist>\n")
		for _, field := range o.FieldList {
			field.AsDocbook(w)
		}
		fmt.Fprintf(w, "</variablelist>\n")
	}

	for _, section := range o.SubSections {
		section.AsDocbook(w, depth+1)
	}

	if depth > 0 && !isBridgeHead {
		tagParts := strings.Split(tags[depth], " ")
		fmt.Fprintf(w, "</%s>\n", tagParts[0])
	}
}
