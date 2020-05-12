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
	"log"
	"os"
	"path"
	"strings"
)

type GVK struct {
	Group   string
	Version string
	Kind    string
}

// Section at any level in the resulting document
type Section struct {
	Name        string
	GVK         *GVK
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

func (o *Section) SetGVK(group, version, kind string) {
	o.GVK = &GVK{
		Group:   group,
		Version: version,
		Kind:    kind,
	}
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
func (o *Section) AsMarkdown(dir string, weight int, w *bufio.Writer, depths ...int) {
	depth := 0
	if len(depths) > 0 {
		depth = depths[0]
	}

	if depth == 1 {
		dir = path.Join(dir, fmt.Sprintf("part%d", weight+1))
		os.MkdirAll(dir, 0755)
		filename := path.Join(dir, "_index.md")
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Unable to create file %s: %s", filename, err)
		}
		defer f.Close()
		f.WriteString(fmt.Sprintf("---\ntitle: \"%s\"\ndraft: false\ncollapsible: true\nweight: %d\n---\n", o.Name, weight+1))
	}

	if depth == 2 {
		filename := path.Join(dir, o.Name+".md")
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Unable to create file %s: %s", filename, err)
		}
		defer f.Close()
		w = bufio.NewWriter(f)
		defer w.Flush()
	}

	if depth >= 2 {
		if depth == 2 {
			fmt.Fprintf(w, "---\ntitle: \"%s\"\ndraft: false\ncollapsible: false\nweight: %d\n---\n", o.Name, weight+1)
		}
		for i := 0; i < depth; i++ {
			fmt.Fprintf(w, "#")
		}
		fmt.Fprintf(w, " %s\n", o.Name)
	}

	if depth > 2 {
		if o.GVK != nil {
			fmt.Fprintf(w, "\n`%s.%s.%s`\n\n", o.GVK.Group, o.GVK.Version, o.GVK.Kind)
		}

		if o.Description != nil {
			fmt.Fprintf(w, "%s\n\n", *o.Description)
		}

		for _, field := range o.FieldList {
			field.AsMarkdown(w)
		}
	}

	for i, section := range o.SubSections {
		if depth == 0 {
			section.AsMarkdown(dir, i, w, depth+1)
		} else {
			section.AsMarkdown(dir, i, w, depth+1)
		}
	}
}

// AsDocbook dumps the section as Docbook
func (o *Section) AsDocbook(w io.Writer, depths ...int) {
	// Create an index entry for sections at this level
	indexLevel := 3

	normalTags := []string{"", "part", "chapter", "sect1 renderas=\"sect2\"", "bridgehead renderas=\"sect4\""}
	appendixTags := []string{"", "", "chapter"}

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
				fmt.Fprintf(w, "<%s%s><title><indexterm type=\"resources\"><primary>%s</primary><secondary>%s.%s</secondary></indexterm>%s</title>\n", tags[depth], ID, o.Name, o.GVK.Group, o.GVK.Version, o.Name)
			} else {
				fmt.Fprintf(w, "<%s%s><title>%s</title>\n", tags[depth], ID, o.Name)
			}
		}
	}

	if o.GVK != nil {
		fmt.Fprintf(w, "<subtitle>%s.%s.%s</subtitle>", o.GVK.Group, o.GVK.Version, o.GVK.Kind)
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
