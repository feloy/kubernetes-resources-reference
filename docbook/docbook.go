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
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/feloy/k8s-api/api"
	"github.com/feloy/k8s-api/structure"
)

// Generate Docbook document from structure
func Generate(w io.Writer, config *api.Config, document *structure.Section) {
	fmt.Fprintf(w, "<?xml version=\"1.0\"?>\n<!DOCTYPE book PUBLIC \"-//OASIS//DTD DocBook XML V4.5//EN\" \"http://www.oasis-open.org/docbook/xml/4.5/docbookx.dtd\">\n")
	fmt.Fprintf(w, "<book>\n")
	fmt.Fprintf(w, "<bookinfo>\n")
	fmt.Fprintf(w, "<title>%s Resources</title><subtitle>%s</subtitle>\n", config.SpecTitle, config.SpecVersion)
	fmt.Fprintf(w, "<releaseinfo>By the Kubernetes Authors</releaseinfo>")
	fmt.Fprintf(w, "<releaseinfo>Edited and published by Philippe Martin</releaseinfo>")
	fmt.Fprintf(w, "<copyright><year>2020</year><holder>The Kubernetes Authors</holder></copyright>")
	fmt.Fprintf(w, "<legalnotice><para>Permission is granted to copy, distribute and/or modify this document under the terms of the Apache License version 2. A copy of the license is included in <xref linkend=\"license\"></xref>.</para></legalnotice>")
	fmt.Fprintf(w, "<legalnotice><para>The tool used to generate this document is available at https://github.com/feloy/kubernetes-resources-reference</para></legalnotice>")
	fmt.Fprintf(w, "<legalnotice><para>The OpenAPI spec file used to generate this document is available at https://github.com/kubernetes/kubernetes</para></legalnotice>")
	fmt.Fprintf(w, "</bookinfo>\n")
	document.AsDocbook(w)
	addLicense(w)
	fmt.Fprintf(w, "<index type=\"resources\"><title>Resources Index</title></index>")
	fmt.Fprintf(w, "<index type=\"fields\"><title>Fields Index</title></index>")
	fmt.Fprintf(w, "</book>\n")
}

func addLicense(w io.Writer) {
	f, err := os.Open("./static/license.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		w.Write(scanner.Bytes())
		w.Write([]byte("\n"))
	}
}
