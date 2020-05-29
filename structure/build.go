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
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/feloy/k8s-api/api"
)

var debug = false

type Builder struct {
	// documented lists the definition names already documented,
	// not to be included in the last section of Other definitions
	documented map[string]struct{}
	config     *api.Config
}

// Build the resulting document structure
func Build(config *api.Config) (result *Section) {
	builder := Builder{
		documented: map[string]struct{}{
			// For Operations only
			"APIGroup":                  struct{}{},
			"APIGroupList":              struct{}{},
			"APIResource":               struct{}{},
			"APIResourceList":           struct{}{},
			"APIVersions":               struct{}{},
			"DeleteOptions":             struct{}{},
			"Preconditions":             struct{}{},
			"DeploymentRollback":        struct{}{},
			"Eviction":                  struct{}{},
			"GroupVersionForDiscovery":  struct{}{},
			"Patch":                     struct{}{},
			"ServerAddressByClientCIDR": struct{}{},
			"Status":                    struct{}{},
			"StatusCause":               struct{}{},
			"StatusDetails":             struct{}{},
			"WatchEvent":                struct{}{},

			// For Old versions only
			"ExternalMetricSource":             struct{}{},
			"ExternalMetricStatus":             struct{}{},
			"MetricSpec":                       struct{}{},
			"MetricStatus":                     struct{}{},
			"MetricTarget":                     struct{}{},
			"MetricValueStatus":                struct{}{},
			"MetricIdentifier":                 struct{}{},
			"ObjectMetricSource":               struct{}{},
			"ObjectMetricStatus":               struct{}{},
			"PodsMetricSource":                 struct{}{},
			"PodsMetricStatus":                 struct{}{},
			"ResourceMetricSource":             struct{}{},
			"ResourceMetricStatus":             struct{}{},
			"HorizontalPodAutoscalerCondition": struct{}{},
			"RollbackConfig":                   struct{}{},
			"RuntimeClassSpec":                 struct{}{},
			"Scale":                            struct{}{},
			"ScaleSpec":                        struct{}{},
			"ScaleStatus":                      struct{}{},

			// For next alpha/beta version only
			"HPAScalingPolicy": struct{}{},
			"HPAScalingRules":  struct{}{},
		},
		config: config,
	}
	result = NewSection("document", nil)
	for _, rescats := range config.ResourceCategories {
		sect1 := NewSection(rescats.Name, nil)
		result.AddSection(sect1)
		for _, resource := range rescats.Resources {
			definition := resource.Definition
			sect2 := NewSection(fmt.Sprintf("%s (%s)", resource.Name, resource.Version), nil)
			sect1.AddSection(sect2)
			var sect3 *Section
			if debug {
				sect3 = NewSection(fmt.Sprintf("%s %s %s", definition.Name, definition.Version, definition.Group), &definition.DescriptionWithEntities)
			} else {
				sect3 = NewSection(definition.Name, &definition.DescriptionWithEntities)
			}
			sect3.SetID(definition.LinkID())
			sect3.SetGVK(definition.Group.String(), definition.Version.String(), definition.Kind.String())
			sect3.SetGoImport(definition.GoImport)
			sect2.AddSection(sect3)
			builder.insertFields(sect3, definition, resource.FieldCategories, 0, nil)
			for _, inline := range definition.Inline {
				var sect3i *Section
				if debug {
					sect3i = NewSection(fmt.Sprintf("%s %s %s", inline.Name, inline.Version, inline.Group), &inline.DescriptionWithEntities)
				} else {
					sect3i = NewSection(inline.Name, &inline.DescriptionWithEntities)
				}
				sect3i.SetID(inline.LinkID())
				sect3i.SetGVK(inline.Group.String(), inline.Version.String(), inline.Kind.String())
				sect3i.SetGoImport(inline.GoImport)
				sect2.AddSection(sect3i)
				builder.insertFields(sect3i, inline, resource.FieldCategories, 0, nil)
			}
		}
		for _, appendix := range rescats.Appendixes {
			sect2 := NewSection(appendix.Name, nil)
			sect2.SetAppendix()
			sect1.AddSection(sect2)
			for _, definition := range appendix.Definitions {
				allDefs := config.Definitions.ByKind[definition.Name]
				for _, def := range allDefs {
					if string(def.Version) == definition.Version && string(def.Group) == definition.Group {
						sect3 := NewSection(definition.Name, &def.DescriptionWithEntities)
						sect3.SetID(def.LinkID())
						sect3.SetGVK(def.Group.String(), def.Version.String(), def.Kind.String())
						sect3.SetGoImport(def.GoImport)
						sect2.AddSection(sect3)
						builder.insertFields(sect3, def, definition.FieldCategories, 0, nil)
					}
				}
			}
		}
	}

	sectOthers := NewSection("Other definitions", nil)
	definitions := config.Definitions.ByKind
	var defsToDocument []string
	for kind := range definitions {
		if _, documented := builder.documented[kind]; !documented {
			defsToDocument = append(defsToDocument, kind)
		}
	}
	sort.Strings(defsToDocument)
	for _, defToDocument := range defsToDocument {
		defSection := NewSection(defToDocument, nil)
		defSection.SetID(definitions[defToDocument][0].LinkID())
		sectOthers.AddSection(defSection)
	}

	if len(sectOthers.SubSections) > 0 {
		result.AddSection(sectOthers)
	}
	return
}

// getCategoriesForDefinition returns the list of categories defined in config for a definition
func (b *Builder) getCategoriesForDefinition(o []*api.FieldCategories, definition string) ([]api.FieldsCategory, bool) {
	for _, fc := range o {
		if *fc.Definition == definition {
			return fc.List, true
		}
	}
	return []api.FieldsCategory{}, false
}

// getDefinitionField returns the field value in API spec for a given definition and field name
func (b *Builder) getDefinitionField(definition *api.Definition, name string) *api.Field {
	for _, field := range definition.Fields {
		if field.Name == name {
			return field
		}
	}
	log.Fatal("field not found")
	return nil
}

// displayFields displays the fields of the definition.
// If categories contains the definition, this config is used
// Elsewhere the fields are displayed, using the depth parameter as config
func (b *Builder) insertFields(s *Section, definition *api.Definition, categories []*api.FieldCategories, depth int, fieldEntry *FieldEntry) {
	if definition == nil {
		return
	}
	b.documented[definition.Name] = struct{}{}
	fieldCategories, found := b.getCategoriesForDefinition(categories, definition.Name)
	if !found {
		for _, field := range definition.Fields {
			b.insertField(s, field, definition.Name, depth, categories, fieldEntry, nil)
		}
	} else {
		currentSection := s
		addedFields := map[string]struct{}{}
		for _, fieldCategory := range fieldCategories {
			if fieldCategory.Name != nil {
				currentSection = NewSection(*fieldCategory.Name, nil)
				s.AddSection(currentSection)
			}
			for _, field := range fieldCategory.Fields {
				realField := b.getDefinitionField(definition, field.Name)
				addedFields[field.Name] = struct{}{}
				b.insertField(currentSection, realField, definition.Name, field.Depth, categories, fieldEntry, &field)
			}
		}
		// Error if some not configured
		for _, field := range definition.Fields {
			if _, found := addedFields[field.Name]; !found {
				fmt.Fprintf(os.Stderr, "Field %s in %s not configured\n", field.Name, s.Name)
			}

		}
	}
}

// Recursively insert a field in the structure
func (b *Builder) insertField(s *Section, field *api.Field, definitionName string, depth int, categories []*api.FieldCategories, fieldEntry *FieldEntry, fieldCategory *api.FieldCategory) {

	typ := field.Type
	if debug && field.Definition != nil {
		typ = fmt.Sprintf("%s %s %s", field.Definition.Name, field.Definition.Version, field.Definition.Group)
	}
	newField := NewFieldEntry(field.Name, definitionName, typ, &field.DescriptionWithEntities, field.Required)
	if field.Definition != nil {
		newField.SetTypeID(field.Definition.LinkID())
	}

	if fieldCategory != nil && fieldCategory.Linkend {
		newField.SetLinkend()
	}

	if fieldEntry != nil {
		fieldEntry.AddFieldEntry(newField)
	} else {
		s.AddFieldEntry(newField)
	}

	if _, documented := b.config.ConfigDefinitions[field.Type]; !documented {
		if depth != 0 {
			b.insertFields(s, field.Definition, categories, depth-1, newField)
		}
	}
}
