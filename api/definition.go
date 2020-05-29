/*
Copyright 2016 The Kubernetes Authors.
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

package api

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

type Definition struct {
	// open-api schema for the definition
	schema spec.Schema

	// Display name of the definition (e.g. Deployment)
	Name string

	GoImport string

	Group   ApiGroup
	Version ApiVersion
	Kind    ApiKind

	DescriptionWithEntities string
	GroupFullName           string

	FoundInOperation bool

	// Inline is a list of definitions that should appear inlined with this one in the documentations
	Inline SortDefinitionsByName

	// Fields is a list of fields in this definition
	Fields Fields

	FullName string
	Resource string

	Required []string
}

// Definitions indexes open-api definitions
type Definitions struct {
	All    map[string]*Definition
	ByKind map[string]SortDefinitionsByVersion
}

type DefinitionList []*Definition

// inlineDefinition is a definition that should be inlined when displaying a Concept
// instead of appearing the in "Definitions"
type inlineDefinition struct {
	Name  string
	Match string
}

var _INLINE_DEFINITIONS = []inlineDefinition{
	{Name: "Spec", Match: "${resource}Spec"},
	{Name: "Status", Match: "${resource}Status"},
	{Name: "List", Match: "${resource}List"},
}

func NewDefinitions(specs []*loads.Document) Definitions {
	s := Definitions{
		All:    map[string]*Definition{},
		ByKind: map[string]SortDefinitionsByVersion{},
	}

	LoadDefinitions(specs, &s)
	s.initialize()
	return s
}

func (s *Definitions) initialize() {
	// initialize fields for all definitions
	for _, d := range s.All {
		s.InitializeFields(d)
	}

	for _, d := range s.All {
		s.ByKind[d.Name] = append(s.ByKind[d.Name], d)
	}

	// If there are multiple versions for an object.  Mark all by the newest as old
	// Sort the ByKind index in by version with newer versions coming before older versions.
	for _, l := range s.ByKind {
		if len(l) <= 1 {
			continue
		}
		sort.Sort(l)
	}

	// Note: examples of inline definitions are "Spec", "Status", "List", etc
	for _, d := range s.All {
		for _, name := range s.getInlineDefinitionNames(d.Name) {
			if cr, ok := s.GetByVersionKind(string(d.Group), string(d.Version), name); ok {
				d.Inline = append(d.Inline, cr)
			}
		}
	}
}

func (s *Definitions) getInlineDefinitionNames(parent string) []string {
	names := []string{}
	for _, id := range _INLINE_DEFINITIONS {
		name := strings.Replace(id.Match, "${resource}", parent, -1)
		names = append(names, name)
	}

	if parent == "SelfSubjectRulesReview" {
		names = append(names, "SubjectRulesReviewStatus")
	}
	return names
}

// GetByVersionKind looks up a definition using its primary key (version,kind)
func (s *Definitions) GetByVersionKind(group, version, kind string) (*Definition, bool) {
	key := &Definition{Group: ApiGroup(group), Version: ApiVersion(version), Kind: ApiKind(kind)}
	r, f := s.All[key.Key()]
	return r, f
}

func (s *Definitions) GetForSchema(schema spec.Schema) (*Definition, bool) {
	g, v, k := GetDefinitionVersionKind(schema)
	if len(k) <= 0 {
		return nil, false
	}
	return s.GetByVersionKind(g, v, k)
}

// Initializes the fields for a definition
func (s *Definitions) InitializeFields(d *Definition) {
	for fieldName, property := range d.schema.Properties {

		required := false
		for _, req := range d.Required {
			if req == fieldName {
				required = true
				break
			}
		}

		f := &Field{
			Name:        fieldName,
			Type:        GetTypeName(property),
			Description: EscapeAsterisks(property.Description),
			Required:    required,
		}
		if len(property.Extensions) > 0 {
			if ps, ok := property.Extensions.GetString(patchStrategyKey); ok {
				f.PatchStrategy = ps
			}
			if pmk, ok := property.Extensions.GetString(patchMergeKeyKey); ok {
				f.PatchMergeKey = pmk
			}
		}

		if fd, ok := s.GetForSchema(property); ok {
			f.Definition = fd
		}
		d.Fields = append(d.Fields, f)
	}
}

func (d *Definition) Key() string {
	return fmt.Sprintf("%s.%s.%s", d.Group, d.Version, d.Kind)
}

func (d *Definition) LinkID() string {
	groupName := strings.Replace(strings.ToLower(d.GroupFullName), ".", "-", -1)
	link := fmt.Sprintf("%s-%s-%s", d.Name, d.Version, groupName)
	return strings.ToLower(link)
}

func (d *Definition) Description() string {
	return EscapeAsterisks(d.schema.Description)
}
