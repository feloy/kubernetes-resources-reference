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

package docbook

import (
	"fmt"
	"io"

	"github.com/feloy/k8s-api/api"
	"github.com/feloy/k8s-api/structure"
)

// Generate Docbook document from structure
func Generate(w io.Writer, config *api.Config, document *structure.Section) {
	fmt.Fprintf(w, "<?xml version=\"1.0\"?>\n<!DOCTYPE book PUBLIC \"-//OASIS//DTD DocBook XML V4.5//EN\" \"http://www.oasis-open.org/docbook/xml/4.5/docbookx.dtd\">\n")
	fmt.Fprintf(w, "<book>\n")
	fmt.Fprintf(w, "<bookinfo>\n")
	fmt.Fprintf(w, "<title>%s Resources</title><subtitle>%s</subtitle>\n", config.SpecTitle, config.SpecVersion)
	fmt.Fprintf(w, "</bookinfo>\n")
	document.AsDocbook(w)
	fmt.Printf("<index type=\"resources\"><title>Index of Resources</title></index>")
	fmt.Printf("<index type=\"fields\"><title>Index of Fields</title></index>")
	fmt.Fprintf(w, "</book>\n")
}
