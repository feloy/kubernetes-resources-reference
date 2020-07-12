package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"

	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs/docbook"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs/hugo"
	"github.com/go-openapi/spec"
	"gopkg.in/yaml.v2"
)

// TOC is the table of contents of the documentation
type TOC struct {
	Parts                 []*Part              `yaml:"parts"`
	SkippedResources      []kubernetes.APIKind `yaml:"skippedResources"`
	Definitions           *spec.Definitions
	LinkEnds              LinkEnds
	DocumentedDefinitions map[kubernetes.Key][]string
}

// Part contains chapters
type Part struct {
	Name     string     `yaml:"name"`
	Chapters []*Chapter `yaml:"chapters"`
}

// Section contains a definition of a Kind for a given Group/Version
type Section struct {
	Name       string
	APIVersion *string
	Definition spec.Schema
}

// NewSection returns a Section
func NewSection(name string, definition *spec.Schema, apiVersion *string) *Section {
	return &Section{
		Name:       name,
		APIVersion: apiVersion,
		Definition: *definition,
	}
}

// LoadTOC loads a config file containing the TOC definition
func LoadTOC(filename string) (*TOC, error) {
	var result TOC

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}

	result.DocumentedDefinitions = map[kubernetes.Key][]string{}
	return &result, nil
}

// PopulateAssociates adds sections to the chapters found in the spec
func (o *TOC) PopulateAssociates(thespec *kubernetes.Spec) error {
	o.LinkEnds = make(LinkEnds)

	for _, part := range o.Parts {
		for _, chapter := range part.Chapters {
			err := chapter.populate(part, o, thespec)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AddOtherResources adds not documented and not replaced resources to a new Part
func (o *TOC) AddOtherResources(spec *kubernetes.Spec) {
	part := &Part{}
	part.Name = "Other Resources"
	part.Chapters = []*Chapter{}

	for k, resource := range *spec.Resources {
		for _, v := range resource {
			if v.ReplacedBy == nil && !v.Documented && !o.skippedResource(k) {
				part.Chapters = append(part.Chapters, &Chapter{
					Name:    v.Kind.String(),
					Group:   &v.Group,
					Version: &v.Version,
					Key:     v.Key,
				})
			}
		}
	}
	sort.Slice(part.Chapters, func(i, j int) bool {
		return part.Chapters[i].Name < part.Chapters[j].Name
	})
	if len(part.Chapters) > 0 {
		o.Parts = append(o.Parts, part)
	}
}

// ToMarkdown writes a Markdown representation of the TOC
func (o *TOC) ToMarkdown(w io.Writer) {
	for _, part := range o.Parts {
		fmt.Fprintf(w, "\n## %s\n", part.Name)
		for _, chapter := range part.Chapters {
			fmt.Fprintf(w, "### %s\n", chapter.Name)
			for _, section := range chapter.Sections {
				fmt.Fprintf(w, "#### %s\n", section.Name)
			}
		}
	}
}

// GetGV returns the group/version for a resource and version (used for apiVersion:)
func GetGV(group kubernetes.APIGroup, version kubernetes.APIVersion) string {
	if group == "" {
		return version.String()
	}
	return fmt.Sprintf("%s/%s", group, version.String())
}

// ToHugo outputs documentation in Markdown format for Hugo in dir directory
func (o *TOC) ToHugo(dir string) error {
	// Test that dir is empty
	fileinfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("Unable to open directory %s", dir)
	}
	if len(fileinfos) > 0 {
		return fmt.Errorf("Directory %s must be empty", dir)
	}

	hugo := hugo.NewHugo(dir)

	o.OutputDocument(hugo)
	return nil
}

// ToDocbook outputs documentation in Docbook format
func (o *TOC) ToDocbook(w io.Writer) error {
	docbook := docbook.NewDocbook()
	o.OutputDocument(docbook)
	return nil
}

// OutputDocumentedDefinitions outputs the list of definitions
// and on which properties they are defined
func (o *TOC) OutputDocumentedDefinitions() {
	for k, v := range o.DocumentedDefinitions {
		if len(v) != 1 {
			fmt.Printf("%s:\n", k)
			for _, s := range v {
				fmt.Printf(" - %s\n", s)
			}
		}
	}
	for k, v := range o.DocumentedDefinitions {
		if len(v) == 1 {
			fmt.Printf("%s:\n", k)
			for _, s := range v {
				fmt.Printf(" - %s\n", s)
			}
		}
	}
	for k := range *o.Definitions {
		if _, found := o.DocumentedDefinitions[kubernetes.Key(k)]; !found {
			fmt.Printf("%s\n", k)
		}
	}
}

func (o *TOC) skippedResource(k kubernetes.APIKind) bool {
	fmt.Printf("%s\n", k)
	for _, skip := range o.SkippedResources {
		if skip == k {
			return true
		}
	}
	return false
}
