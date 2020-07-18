package hugo

import (
	"fmt"
	"sort"
	"strings"

	"github.com/feloy/kubernetes-api-reference/pkg/formats/markdown"
	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
)

// Section of a Hugo output
// implements the outputs.Section interface
type Section struct {
	hugo    *Hugo
	part    *Part
	chapter *Chapter
}

// AddContent adds content to a section
func (o Section) AddContent(s string) error {
	return o.hugo.addContent(o.part.name, o.chapter.name, s)
}

// AddTypeDefinition adds the definition of a type to a section
func (o Section) AddTypeDefinition(s string) error {
	return o.hugo.addContent(o.part.name, o.chapter.name, markdown.Italic(s))
}

// StartPropertyList starts the list of properties
func (o Section) StartPropertyList() error {
	return nil
}

// AddProperty adds a property to the section
func (o Section) AddProperty(name string, property *kubernetes.Property, linkend []string, indent bool) error {
	if property.HardCodedValue != nil {
		return o.hugo.addListEntry(o.part.name, o.chapter.name, "**"+name+"**: "+*property.HardCodedValue, "", 0)
	}

	indentLevel := 0
	if indent {
		indentLevel++
	}
	required := ""
	if property.Required {
		required = ", required"
	}

	typ := property.Type
	if property.TypeKey != nil && len(linkend) > 0 {
		typ = o.hugo.LinkEnd(linkend, property.Type)
	}
	title := fmt.Sprintf("**%s** (%s)%s", name, typ, required)

	description := property.Description

	listType := ""
	if property.ListType != nil {
		if *property.ListType == "atomic" {
			listType = "Atomic: will be replaced during a merge"
		} else if *property.ListType == "set" {
			listType = "Set: unique values will be kept during a merge"
		} else if *property.ListType == "map" {
			if len(property.ListMapKeys) == 1 {
				listType = "Map: unique values on key " + property.ListMapKeys[0] + " will be kept during a merge"
			} else {
				listType = "Map: unique values on keys `" + strings.Join(property.ListMapKeys, ", ") + "` will be kept during a merge"
			}
		}
	}
	if len(listType) > 0 {
		description = "*" + listType + "*\n" + description
	}

	var patches string
	if property.MergeStrategyKey != nil && property.RetainKeysStrategy {
		patches = fmt.Sprintf("Patch strategies: retainKeys, merge on key `%s`", *property.MergeStrategyKey)
	} else if property.MergeStrategyKey != nil {
		patches = fmt.Sprintf("Patch strategy: merge on key `%s`", *property.MergeStrategyKey)
	} else if property.RetainKeysStrategy {
		patches = "Patch strategy: retainKeys"
	}

	if len(patches) > 0 {
		description = "*" + patches + "*\n" + description
	}
	return o.hugo.addListEntry(o.part.name, o.chapter.name, title, description, indentLevel)
}

// EndProperty ends a property
func (o Section) EndProperty() error {
	return nil
}

// EndPropertyList ends the list of properties
func (o Section) EndPropertyList() error {
	return nil
}

// AddOperation adds an operation
func (o Section) AddOperation(operation *kubernetes.ActionInfo, linkends kubernetes.LinkEnds) error {
	sentences := strings.Split(operation.Operation.Description, ".")

	if len(sentences) > 1 {
		fmt.Printf("SHOULD NOT HAPPEN, sentences: %d\n", len(sentences))
	}

	httpRequest := fmt.Sprintf("%s %s", operation.HTTPMethod, operation.Path.String())

	err := o.hugo.addSubsection(o.part.name, o.chapter.name, fmt.Sprintf("`%s` %s", operation.Action.Verb(), sentences[0]))
	if err != nil {
		return err
	}

	err = o.hugo.addContent(o.part.name, o.chapter.name, "\n##### HTTP Request")
	err = o.hugo.addContent(o.part.name, o.chapter.name, httpRequest)
	if err != nil {
		return err
	}

	err = o.hugo.addContent(o.part.name, o.chapter.name, "\n##### Parameters")
	for _, param := range operation.Parameters {

		required := ""
		if param.Required {
			required = ", required"
		}

		typ := param.Type
		if param.Schema != nil {
			t, typeKey := kubernetes.GetTypeNameAndKey(*param.Schema)
			linkend, found := linkends[*typeKey]
			if found {
				typ = o.hugo.LinkEnd(linkend, t)
			} else {
				typ = t
				fmt.Printf("SHOULD NOT HAPPEN: %s\n", typeKey)
			}
		}

		desc := param.Description
		if len(desc) > 0 && kubernetes.ParameterInAnnex(param) {
			desc = o.hugo.LinkEnd([]string{"common-parameters", "common-parameters-"}, param.Name)
		}
		o.hugo.addListEntry(o.part.name, o.chapter.name, paramName(param.Name, param.In)+" ("+typ+")"+required, desc, 1)
	}

	codes := make([]int, len(operation.Operation.Responses.StatusCodeResponses))
	i := 0
	for code := range operation.Operation.Responses.StatusCodeResponses {
		codes[i] = code
		i++
	}
	sort.Ints(codes)
	err = o.hugo.addContent(o.part.name, o.chapter.name, "\n##### Response")
	for _, code := range codes {
		response := operation.Operation.Responses.StatusCodeResponses[code]

		typ := ""
		if response.Schema != nil {
			t, typeKey := kubernetes.GetTypeNameAndKey(*response.Schema)
			if typeKey != nil {
				linkend, found := linkends[*typeKey]
				if found {
					typ = o.hugo.LinkEnd(linkend, t)
					typ = " (" + typ + ")"
				} else {
					typ = " (" + t + ")"
					fmt.Printf("SHOULD NOT HAPPEN: %s\n", typeKey)
				}
			} else {
				typ = " (" + t + ")"
			}
		}

		err = o.hugo.addContent(o.part.name, o.chapter.name, fmt.Sprintf("%d%s: %s", code, typ, response.Description))
	}
	return nil
}

func paramName(s string, in string) string {
	switch in {
	case "path":
		return fmt.Sprintf("**{%s}**", s)
	case "query":
		return fmt.Sprintf("**?%s**", s)
	default:
		return fmt.Sprintf("**%s**", s)
	}
}
