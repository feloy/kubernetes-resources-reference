package hugo

import (
	"fmt"
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

	link := ""
	var title string
	if property.TypeKey != nil {
		link = property.Type
		if len(linkend) > 0 {
			link = o.hugo.LinkEnd(linkend, property.Type)
		}
		title = fmt.Sprintf("**%s** (%s)%s", name, link, required)
	} else {
		title = fmt.Sprintf("**%s** (%s%s)%s", name, property.Type, link, required)
	}

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
func (o Section) AddOperation(operation *kubernetes.ActionInfo) error {
	sentences := strings.Split(operation.Operation.Description, ".")

	description := fmt.Sprintf("%s %s", operation.HTTPMethod, operation.Path.String())
	if len(sentences) > 1 {
		description = strings.Join(sentences[1:], ".") + "\n" + description
	}
	err := o.hugo.addListEntry(o.part.name, o.chapter.name, fmt.Sprintf("**%s** (`%s`)", sentences[0], operation.Action.Verb()), description, 0)
	if err != nil {
		return err
	}

	/*	for _, pathParam := range operation.Operation.Parameters {
		fmt.Printf("IN: %s (%s)\n", pathParam.Name, pathParam.In)
	}*/
	return nil
}
