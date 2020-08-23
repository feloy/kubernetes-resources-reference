package docbook

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/feloy/kubernetes-api-reference/pkg/formats/dbxml"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
	x "github.com/shabbyrobe/xmlwriter"
)

// Docbook generator
type Docbook struct {
	w      *x.Writer
	buffer *bytes.Buffer
}

const (
	partDepth           = 1
	chapterDepth        = 2
	sectionDepth        = 3
	propertiesListDepth = 5
)

// NewDocbook returns a new Docbook generator
func NewDocbook() *Docbook {
	return &Docbook{}
}

// Prepare Docbook document
func (o *Docbook) Prepare() error {
	o.buffer = &bytes.Buffer{}
	o.w = x.Open(o.buffer, x.WithIndent())

	book := dbxml.NewBook()
	bookinfo := dbxml.Bookinfo(
		"Kubernetes API Reference",
		"version 1.19", "2020",
		"The Kubernetes Authors",
		[]string{
			"By the Kubernetes Authors",
			"Edited and published by Philippe Martin",
		},
		[][]x.Writable{
			{
				x.Text("Permission is granted to copy, distribute and/or modify this document under the terms of the Apache License version 2. A copy of the license is included in "),
				dbxml.XrefLinkend("license"),
			},
			{
				x.Text("The tool used to generate this document is available at https://github.com/feloy/kubernetes-resources-reference"),
			},
			{
				x.Text("The OpenAPI spec file used to generate this document is available at https://github.com/kubernetes/kubernetes"),
			},
		})

	ec := &x.ErrCollector{}
	defer ec.Panic()
	ec.Do(
		o.w.StartDoc(x.Doc{}),
		o.w.StartElem(book),
		o.w.StartElem(bookinfo),
		o.w.EndElem("bookinfo"),
	)
	return nil
}

// Terminate Docbook document
func (o *Docbook) Terminate() error {
	ec := &x.ErrCollector{}
	defer ec.Panic()
	ec.Do(
		o.w.EndToDepth(partDepth, x.ElemNode, "part"),
		insertLicense(o.w),
		insertIndex(o.w, "resources", "Resources Index"),
		insertIndex(o.w, "definitions", "Definitions Index"),
		insertIndex(o.w, "fields", "Fields Index"),
		o.w.EndAllFlush(),
	)
	fmt.Println(o.buffer.String())
	return nil
}

func insertLicense(w *x.Writer) error {
	f, err := os.Open("./static/license.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	w.WriteRaw("\n")
	for scanner.Scan() {
		w.WriteRaw(string(scanner.Bytes()))
		w.WriteRaw("\n")
	}

	return nil
}

func insertIndex(w *x.Writer, name string, title string) error {
	ec := &x.ErrCollector{}
	defer ec.Panic()
	index1 := dbxml.Section("index", title)
	ec.Do(
		w.StartElem(index1),
		w.WriteAttr(x.Attr{Name: "type", Value: name}),
		w.EndElem("index"),
	)
	return nil
}

// AddPart adds a part to the docbook output
func (o *Docbook) AddPart(i int, name string) (outputs.Part, error) {
	if i > 0 {
		o.w.EndToDepth(partDepth, x.ElemNode, "part")
	}
	part := dbxml.Section("part", name)
	return DocbookPart{w: o.w}, o.w.StartElem(part)
}

func escapeName(s string) string {
	result := strings.ToLower(s)
	result = strings.ReplaceAll(result, " ", "-")
	return result
}
