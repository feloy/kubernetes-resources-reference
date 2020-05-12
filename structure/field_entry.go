/*
Copyright 2019 Philippe Martin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package structure

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// FieldEntry contains a Definition field information
type FieldEntry struct {
	Name         string
	SectionName  string
	Type         string
	TypeID       *string
	Description  *string
	Linkend      bool
	SubFieldList []*FieldEntry
}

// NewFieldEntry creates a new field entry for a list of fields
func NewFieldEntry(name string, sectionName string, typ string, description *string) *FieldEntry {
	return &FieldEntry{Name: name, SectionName: sectionName, Type: typ, Description: description}
}

// SetTypeID sets the ID of the type for linking
func (o *FieldEntry) SetTypeID(typeID string) {
	o.TypeID = &typeID
}

// SetLinkend sets linkend to true for the field entry
func (o *FieldEntry) SetLinkend() {
	o.Linkend = true
}

// AddFieldEntry adds a field entry in the list of fields of a section
func (o *FieldEntry) AddFieldEntry(f *FieldEntry) {
	o.SubFieldList = append(o.SubFieldList, f)
}

func indentString(w *bufio.Writer, s string, indent int) {
	w.WriteString(strings.Repeat(" ", indent))
	w.WriteString(strings.ReplaceAll(s, "\n", "\n"+strings.Repeat(" ", indent)))
}

// AsMarkdown dumps the field entry as Markdown
func (o *FieldEntry) AsMarkdown(w *bufio.Writer, prefixes ...string) {
	depth := len(prefixes)
	if depth > 0 {
		w.WriteString("  ")
	}
	w.WriteString(fmt.Sprintf("- **%s** (%s)\n", strings.Join(append(prefixes, o.Name), "."), o.Type))

	indent := 2
	if depth > 0 {
		indent += 2
	}
	indentString(w, *o.Description, indent)
	w.WriteString("\n\n")

	for _, subfield := range o.SubFieldList {
		subfield.AsMarkdown(w, append(prefixes, o.Name)...)
	}
}

func isIndexable(name string) bool {
	noIndex := []string{"apiVersion", "kind", "metadata", "spec", "status"}
	for _, v := range noIndex {
		if v == name {
			return false
		}
	}
	return true
}

// AsDocbook dumps the field entry as Docbook
func (o *FieldEntry) AsDocbook(w io.Writer, prefixes ...string) {

	// Sublists are inlined after the first level
	depth := len(prefixes)
	inline := depth > 0

	fmt.Fprintf(w, "<varlistentry>\n")
	typ := ""
	if o.Type != "" {
		if o.TypeID != nil && len(o.SubFieldList) == 0 {
			typ = fmt.Sprintf(" (<emphasis><link linkend=\"%s\">%s</link></emphasis>)", *o.TypeID, o.Type)
		} else if o.TypeID != nil && o.Linkend {
			typ = fmt.Sprintf(" (<emphasis id=\"%s\">%s</emphasis>)", *o.TypeID, o.Type)
		} else {
			typ = fmt.Sprintf(" (<emphasis>%s</emphasis>)", o.Type)
		}
	}

	if !isIndexable(o.Name) {
		fmt.Fprintf(w, "<term><varname>%s</varname>%s</term>\n", strings.Join(append(prefixes, o.Name), "."), typ)
	} else {
		fmt.Fprintf(w, "<term><varname><indexterm type=\"fields\"><primary>%s</primary><secondary>%s</secondary></indexterm>%s</varname>%s</term>\n",
			o.Name, o.SectionName,
			strings.Join(append(prefixes, o.Name), "."), typ)
	}
	fmt.Fprintf(w, "<listitem>")
	if o.Description != nil {
		fmt.Fprintf(w, "%s", *o.Description)
	}

	if inline {
		fmt.Fprintf(w, "</listitem>\n")
		fmt.Fprintf(w, "</varlistentry>\n")
	}

	if len(o.SubFieldList) > 0 {
		if !inline {
			fmt.Fprintf(w, "<variablelist>\n")
		}
		for _, subfield := range o.SubFieldList {
			subfield.AsDocbook(w, append(prefixes, o.Name)...)
		}
		if !inline {
			fmt.Fprintf(w, "</variablelist>\n")
		}
	}

	if !inline {
		fmt.Fprintf(w, "</listitem>\n")
		fmt.Fprintf(w, "</varlistentry>\n")
	}

}
