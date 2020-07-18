package docbook

import (
	"fmt"
	"sort"
	"strings"

	"github.com/feloy/kubernetes-api-reference/pkg/formats/dbxml"
	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
	x "github.com/shabbyrobe/xmlwriter"
)

type DocbookSection struct {
	w *x.Writer
}

// AddContent adds content to the output
func (o DocbookSection) AddContent(s string) error {
	for _, part := range strings.Split(s, "\n") {
		o.w.StartElem(dbxml.ElemWithText("para", part))
		o.w.EndElem("para")
	}
	return nil
}

// AddTypeDefinition adds the definition of a type to the output
func (o DocbookSection) AddTypeDefinition(s string) error {
	for _, part := range strings.Split(s, "\n") {
		o.w.StartElem(x.Elem{Name: "para"})
		o.w.StartElem(dbxml.ElemWithText("emphasis", part))
		o.w.EndElem("emphasis")
		o.w.EndElem("para")
	}
	return nil
}

// StartPropertyList starts the list of properties
func (o DocbookSection) StartPropertyList() error {
	return o.w.StartElem(x.Elem{Name: "variablelist"})
}

// AddProperty adds a property to the list of properties
func (o DocbookSection) AddProperty(name string, property *kubernetes.Property, linkend []string, indent bool) error {
	if property.HardCodedValue != nil {
		o.w.StartElem(x.Elem{Name: "varlistentry"})
		o.w.StartElem(x.Elem{Name: "term"})
		o.w.StartElem(dbxml.ElemWithText("varname", name+": "+*property.HardCodedValue))
		o.w.EndElem("varname")
		o.w.EndElem("term")
		o.w.StartElem(x.Elem{Name: "listitem"})
		o.w.EndElem("listitem")
		o.w.EndElem("varlistentry")
		return nil
	}

	o.w.StartElem(x.Elem{Name: "varlistentry"})
	o.w.StartElem(x.Elem{Name: "term"})
	o.w.StartElem(dbxml.ElemWithText("varname", name))
	o.w.EndElem("varname")
	o.w.WriteText(" (")
	o.w.StartElem(x.Elem{Name: "emphasis"})
	if len(linkend) > 0 {
		o.w.StartElem(x.Elem{Name: "link"})
		o.w.WriteAttr(x.Attr{Name: "linkend", Value: escapeName(linkend[1] + "." + linkend[2])})
	}
	o.w.WriteText(property.Type)
	if len(linkend) > 0 {
		o.w.EndElem("link")
	}
	o.w.EndElem("emphasis")
	o.w.WriteText(")")
	if property.Required {
		o.w.WriteText(", required")
	}
	o.w.EndElem("term")
	o.w.StartElem(x.Elem{Name: "listitem"})

	var patches []x.Writable
	if property.MergeStrategyKey != nil && property.RetainKeysStrategy {
		patches = []x.Writable{x.Text("Patch strategies: retainKeys, merge on key "), dbxml.ElemWithText("varname", *property.MergeStrategyKey)}
	} else if property.MergeStrategyKey != nil {
		patches = []x.Writable{x.Text("Patch strategy: merge on key "), dbxml.ElemWithText("varname", *property.MergeStrategyKey)}
	} else if property.RetainKeysStrategy {
		patches = []x.Writable{x.Text("Patch strategy: retainKeys")}
	}

	if len(patches) > 0 {
		o.w.StartElem(x.Elem{Name: "para"})
		o.w.StartElem(dbxml.ElemWithContent("emphasis", patches))
		o.w.EndElem("emphasis")
		o.w.EndElem("para")
	}

	var listType []x.Writable
	if property.ListType != nil {
		if *property.ListType == "atomic" {
			listType = []x.Writable{x.Text("Atomic: will be replaced during a merge")}
		} else if *property.ListType == "set" {
			listType = []x.Writable{x.Text("Set: unique values will be kept during a merge")}
		} else if *property.ListType == "map" {
			if len(property.ListMapKeys) == 1 {
				listType = []x.Writable{
					x.Text("Map: unique values on key "),
					dbxml.ElemWithText("varname", property.ListMapKeys[0]),
					x.Text(" will be kept during a merge"),
				}
			} else {
				listType = []x.Writable{
					x.Text("Map: unique values on keys "),
					dbxml.ElemWithText("varname", strings.Join(property.ListMapKeys, ", ")),
					x.Text(" will be kept during a merge"),
				}
			}
		}
	}
	if len(listType) > 0 {
		o.w.StartElem(x.Elem{Name: "para"})
		o.w.StartElem(dbxml.ElemWithContent("emphasis", listType))
		o.w.EndElem("emphasis")
		o.w.EndElem("para")
	}

	parts := strings.Split(property.Description, "\n")
	for _, part := range parts {
		o.w.StartElem(dbxml.ElemWithText("para", part))
		o.w.EndElem("para")
	}
	return nil
}

// EndProperty ends a property
func (o DocbookSection) EndProperty() error {
	o.w.EndElem("listitem")
	o.w.EndElem("varlistentry")
	return nil
}

// EndPropertyList ends the list of properties
func (o DocbookSection) EndPropertyList() error {
	return o.w.EndElem("variablelist")
}

// AddOperation adds an operation
func (o DocbookSection) AddOperation(operation *kubernetes.ActionInfo, linkends kubernetes.LinkEnds) error {
	sentences := strings.Split(operation.Operation.Description, ".")

	httpRequest := fmt.Sprintf("%s %s", operation.HTTPMethod, operation.Path.String())

	if len(sentences) > 1 {
		fmt.Printf("SHOULD NOT HAPPEN, sentences: %d\n", len(sentences))
	}

	subsection := dbxml.Section("sect2", fmt.Sprintf("(%s) %s", operation.Action.Verb(), sentences[0]))
	err := o.w.StartElem(subsection)
	if err != nil {
		return err
	}
	o.w.WriteAttr(x.Attr{
		Name:  "renderas",
		Value: "sect3",
	})

	o.w.StartElem(x.Elem{Name: "para"})
	o.w.StartElem(dbxml.ElemWithText("varname", httpRequest))
	o.w.EndElem("varname")
	o.w.EndElem("para")

	paramsSection := dbxml.ElemWithText("bridgehead", "Parameters")
	err = o.w.StartElem(paramsSection)
	if err != nil {
		return err
	}
	o.w.WriteAttr(x.Attr{
		Name:  "renderas",
		Value: "sect4",
	})
	o.w.EndElem("bridgehead")

	// Params

	o.w.StartElem(x.Elem{Name: "variablelist"})

	for _, param := range operation.Parameters {

		typ := param.Type
		var linkend []string
		if param.Schema != nil {
			var typeKey *kubernetes.Key
			typ, typeKey = kubernetes.GetTypeNameAndKey(*param.Schema)
			linkend, _ = linkends[*typeKey]
		}

		desc := param.Description
		o.w.StartElem(x.Elem{Name: "varlistentry"})

		o.w.StartElem(x.Elem{Name: "term"})
		if len(desc) > 0 && kubernetes.ParameterInAnnex(param) {
			o.w.StartElem(x.Elem{Name: "link"})
			o.w.WriteAttr(x.Attr{Name: "linkend", Value: escapeName("common-parameters-." + param.Name)})
		}

		// parameter name
		o.w.StartElem(x.Elem{Name: "varname"})
		o.w.WriteText(paramName(param.Name, param.In))
		o.w.EndElem("varname")

		// type
		o.w.WriteText(" (")
		o.w.StartElem(x.Elem{Name: "emphasis"})
		if len(linkend) > 0 {
			o.w.StartElem(x.Elem{Name: "link"})
			o.w.WriteAttr(x.Attr{Name: "linkend", Value: escapeName(linkend[1] + "." + linkend[2])})
		}
		o.w.WriteText(typ)
		if len(linkend) > 0 {
			o.w.EndElem("link")
		}
		o.w.EndElem("emphasis")
		o.w.WriteText(")")

		// required?
		if param.Required {
			o.w.WriteText(", required")
		}

		if len(desc) > 0 && kubernetes.ParameterInAnnex(param) {
			o.w.EndElem("link")
			desc = ""
		}
		o.w.EndElem("term")

		// Description
		o.w.StartElem(x.Elem{Name: "listitem"})
		if len(desc) > 0 {
			o.w.StartElem(dbxml.ElemWithText("para", desc))
			o.w.EndElem("para")
		}

		o.w.EndElem("listitem")
		o.w.EndElem("varlistentry")
	}

	o.w.EndElem("variablelist")

	responseSection := dbxml.ElemWithText("bridgehead", "Response")
	err = o.w.StartElem(responseSection)
	if err != nil {
		return err
	}
	o.w.WriteAttr(x.Attr{
		Name:  "renderas",
		Value: "sect4",
	})
	o.w.EndElem("bridgehead")

	// Response
	codes := make([]int, len(operation.Operation.Responses.StatusCodeResponses))
	i := 0
	for code := range operation.Operation.Responses.StatusCodeResponses {
		codes[i] = code
		i++
	}
	sort.Ints(codes)

	o.w.StartElem(x.Elem{Name: "variablelist"})

	for _, code := range codes {
		response := operation.Operation.Responses.StatusCodeResponses[code]

		typ := ""
		var linkend []string
		if response.Schema != nil {
			var typeKey *kubernetes.Key
			typ, typeKey = kubernetes.GetTypeNameAndKey(*response.Schema)
			if typeKey != nil {
				linkend, _ = linkends[*typeKey]
			}
		}
		o.w.StartElem(x.Elem{Name: "varlistentry"})
		o.w.StartElem(x.Elem{Name: "term"})
		o.w.StartElem(dbxml.ElemWithText("varname", fmt.Sprintf("%d", code)))
		o.w.EndElem("varname")

		if len(typ) > 0 {
			o.w.WriteText(" (")
			o.w.StartElem(x.Elem{Name: "emphasis"})
			if len(linkend) > 0 {
				o.w.StartElem(x.Elem{Name: "link"})
				o.w.WriteAttr(x.Attr{Name: "linkend", Value: escapeName(linkend[1] + "." + linkend[2])})
			}
			o.w.WriteText(typ)
			if len(linkend) > 0 {
				o.w.EndElem("link")
			}
			o.w.EndElem("emphasis")
			o.w.WriteText(")")
		}

		o.w.WriteText(": " + response.Description)
		o.w.EndElem("term")
		o.w.EndElem("varlistentry")
	}
	o.w.EndElem("variablelist")

	o.w.EndElem("sect2")
	return nil
}

func paramName(s string, in string) string {
	switch in {
	case "path":
		return fmt.Sprintf("{%s}", s)
	case "query":
		return fmt.Sprintf("?%s", s)
	default:
		return fmt.Sprintf("%s", s)
	}
}
