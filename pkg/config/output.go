package config

import (
	"sort"
	"strings"

	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
	"github.com/go-openapi/spec"
)

// OutputDocument outputs contents using output
func (o *TOC) OutputDocument(output outputs.Output) error {
	err := output.Prepare()
	if err != nil {
		return err
	}

	for p, tocPart := range o.Parts {
		err = o.OutputPart(p, tocPart, output)
		if err != nil {
			return err
		}
	}

	o.OutputCommonParameters(len(o.Parts), output)

	err = output.Terminate()
	if err != nil {
		return err
	}
	return nil
}

// OutputPart outputs a Part
func (o *TOC) OutputPart(i int, part *Part, output outputs.Output) error {
	outputPart, err := output.AddPart(i, part.Name)
	if err != nil {
		return err
	}

	for c, tocChapter := range part.Chapters {
		err = o.OutputChapter(c, tocChapter, outputPart)
		if err != nil {
			return err
		}
	}
	return nil
}

// OutputChapter outputs a chapter of the part
func (o *TOC) OutputChapter(i int, chapter *Chapter, outputPart outputs.Part) error {
	description := ""
	if len(chapter.Sections) > 0 {
		description = getEscapedFirstPhrase(chapter.Sections[0].Definition.Description)
	}
	outputChapter, err := outputPart.AddChapter(i, chapter.Name, chapter.Version, description)
	if err != nil {
		return err
	}

	if chapter.Group != nil && chapter.Version != nil {
		err = outputChapter.SetAPIVersion(GetGV(*chapter.Group, *chapter.Version))
		if err != nil {
			return err
		}
	}
	err = outputChapter.SetGoImport(chapter.Key.GoImportPrefix())
	if err != nil {
		return err
	}

	for s, tocSection := range chapter.Sections {
		err = o.OutputSection(s, tocSection, outputChapter)
		if err != nil {
			return err
		}
	}

	if chapter.Group != nil && chapter.Version != nil {
		gvkString := chapter.Group.String() + "." + chapter.Version.String() + "." + chapter.Name
		actions := o.Actions.Get(gvkString)
		if actions != nil {
			o.OutputOperations(len(chapter.Sections), outputChapter, &actions)
		}
	}
	return nil
}

// OutputSection outputs a section of the chapter
func (o *TOC) OutputSection(i int, section *Section, outputChapter outputs.Chapter) error {
	outputSection, err := outputChapter.AddSection(i, section.Name)
	if err != nil {
		return err
	}
	err = outputSection.AddContent(section.Definition.Description)
	if err != nil {
		return err
	}

	outputSection.StartPropertyList()

	err = o.OutputProperties(section.Name, section.Definition, outputSection, []string{}, section.APIVersion)
	if err != nil {
		return err
	}

	return outputSection.EndPropertyList()
}

// OutputProperties outputs the properties of a definition
func (o *TOC) OutputProperties(defname string, definition spec.Schema, outputSection outputs.Section, prefix []string, apiVersion *string) error {
	requiredProperties := definition.Required
	ordered := orderedPropertyKeys(requiredProperties, definition.Properties, apiVersion != nil)

	for _, name := range ordered {

		if apiVersion != nil && (name == "apiVersion" || name == "kind") {
			var property *kubernetes.Property
			if name == "apiVersion" {
				property = kubernetes.NewHardCodedValueProperty(name, *apiVersion)
			} else if name == "kind" {
				property = kubernetes.NewHardCodedValueProperty(name, defname)
			}
			err := outputSection.AddProperty(name, property, []string{}, false)
			if err != nil {
				return err
			}
			continue
		}

		details := definition.Properties[name]
		property, err := kubernetes.NewProperty(name, details, requiredProperties)
		if err != nil {
			return err
		}
		var linkend []string
		if property.TypeKey != nil {
			linkend = o.LinkEnds[*property.TypeKey]
		}
		completeName := prefix
		completeName = append(completeName, name)
		err = outputSection.AddProperty(strings.Join(completeName, "."), property, linkend, len(prefix) > 0)
		if err != nil {
			return err
		}
		if property.TypeKey != nil && len(linkend) == 0 {
			// The type is documented inline
			if target, found := (*o.Definitions)[property.TypeKey.String()]; found {
				o.setDocumentedDefinition(property.TypeKey, defname+"/"+strings.Join(completeName, "."))

				err = outputSection.AddTypeDefinition(target.Description)
				if err != nil {
					return err
				}

				sublist := false
				if len(prefix) == 0 {
					sublist = true
					outputSection.StartPropertyList()
				} else {
					err = outputSection.EndProperty()
				}

				o.OutputProperties(defname, target, outputSection, append(prefix, name), nil)
				if sublist {
					outputSection.EndPropertyList()
				}
			}
		}
		err = outputSection.EndProperty()
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *TOC) setDocumentedDefinition(key *kubernetes.Key, from string) {
	if _, found := o.DocumentedDefinitions[*key]; found {
		o.DocumentedDefinitions[*key] = append(o.DocumentedDefinitions[*key], from)
	} else {
		o.DocumentedDefinitions[*key] = []string{from}
	}
}

// OutputOperations outputs the Operations chapter
func (o *TOC) OutputOperations(i int, outputChapter outputs.Chapter, operations *kubernetes.ActionInfoList) error {
	operationsSection, err := outputChapter.AddSection(i, "Operations")
	if err != nil {
		return err
	}
	for i, operation := range *operations {
		o.OutputOperation(i, operationsSection, &operation)
		_ = operation
	}
	return nil
}

// OutputOperation outputs details of an Operation
func (o *TOC) OutputOperation(i int, outputSection outputs.Section, operation *kubernetes.ActionInfo) error {
	outputSection.AddOperation(operation, o.LinkEnds)
	return nil
}

// OutputCommonParameters outputs the parameters in common
func (o *TOC) OutputCommonParameters(i int, output outputs.Output) error {
	outputPart, err := output.AddPart(i, "Common Parameters")
	if err != nil {
		return err
	}

	outputChapter, err := outputPart.AddChapter(0, "Common Parameters", nil, "")

	params := make([]string, len(kubernetes.ParametersAnnex))
	j := 0
	for k := range kubernetes.ParametersAnnex {
		params[j] = k
		j++
	}
	sort.Strings(params)
	for i, param := range params {
		if len(kubernetes.ResourcesDescriptions[param][0].Description) == 0 {
			continue
		}
		outputSection, err := outputChapter.AddSection(i, param)
		if err != nil {
			return err
		}
		err = outputSection.AddContent(kubernetes.ResourcesDescriptions[param][0].Description)
	}
	return nil
}

// orderedPropertyKeys returns the keys of m alphabetically ordered
// keys in required will be placed first
func orderedPropertyKeys(required []string, m map[string]spec.Schema, isResource bool) []string {
	sort.Strings(required)

	if isResource {
		mkeys := make(map[string]struct{})
		for k := range m {
			mkeys[k] = struct{}{}
		}
		for _, special := range []string{"metadata", "kind", "apiVersion"} {
			if !isRequired(special, required) {
				if _, ok := mkeys[special]; ok {
					required = append([]string{special}, required...)
				}
			}
		}
	}

	keys := make([]string, len(m)-len(required))
	i := 0
	for k := range m {
		if !isRequired(k, required) {
			keys[i] = k
			i++
		}
	}
	sort.Strings(keys)
	return append(required, keys...)
}

// isRequired returns true if k is in the required array
func isRequired(k string, required []string) bool {
	for _, r := range required {
		if r == k {
			return true
		}
	}
	return false
}
