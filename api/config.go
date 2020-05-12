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
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

var ConfigDir = flag.String("config-dir", "", "Directory contain api files.")

type Config struct {
	ResourceCategories []ResourceCategory `yaml:"resource_categories,omitempty"`

	ConfigDefinitions map[string]struct{} `yaml:"-"`

	Definitions Definitions `yaml:"-"`
	SpecTitle   string      `yaml:"-"`
	SpecVersion string      `yaml:"-"`
}

// ResourceCategory defines a category of Concepts
type ResourceCategory struct {
	// Name is the display name of this group
	Name string `yaml:",omitempty"`
	// Resources are the collection of Resources in this group
	Resources  Resources  `yaml:",omitempty"`
	Appendixes Appendixes `yaml:"definitions,omitempty"`
}

type Resources []*Resource

type Appendixes []*Appendix

type Resource struct {
	// Name is the display name of this Resource
	Name    string `yaml:",omitempty"`
	Version string `yaml:",omitempty"`
	Group   string `yaml:",omitempty"`

	GoImport *string `yaml:",omitempty"`

	FieldCategories []*FieldCategories `yaml:"field_categories,omitempty"`

	// Definition of the object
	Definition *Definition `yaml:"-"`
}

type Appendix struct {
	Name        string    `yaml:",omitempty"`
	Definitions Resources `yaml:",omitempty"`
}

type FieldCategory struct {
	Name  string `yaml:",omitempty"`
	Depth int    `yaml:",omitempty"`
	// Linkend indicates if the field is a target of links
	Linkend bool `yaml:",omitempty"`
}

type FieldCategories struct {
	Definition *string          `yaml:",omitempty"`
	List       []FieldsCategory `yaml:",omitempty"`
}

type FieldsCategory struct {
	Name   *string         `yaml:",omitempty"`
	Fields []FieldCategory `yaml:",omitempty"`
}

func NewConfig() *Config {
	config := LoadConfigFromYAML()
	specs := LoadOpenApiSpec()

	// Parse spec version
	ParseSpecInfo(specs, config)

	// Initialize all of the operations
	config.Definitions = NewDefinitions(specs)

	// Initialization of resources
	config.visitResources()

	config.CleanUp()

	config.FindConfigDefinitions()

	return config
}

// CleanUp sorts fields
func (c *Config) CleanUp() {
	for _, d := range c.Definitions.All {
		sort.Sort(d.Fields)
	}
}

// LoadConfigFromYAML reads the config yaml file into a struct
func LoadConfigFromYAML() *Config {
	f := filepath.Join(*ConfigDir, "config.yaml")

	config := &Config{}
	contents, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Printf("Failed to read yaml file %s: %v", f, err)
		os.Exit(2)
	} else {
		err = yaml.Unmarshal(contents, config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return config
}

// EscapeDescriptions The OpenAPI spec has escape sequences like \u003c. When the spec is unmarshaled,
// the escape sequences get converted to ordinary characters. For example,
// \u003c gets converted to a regular < character. But we can't use  regular <
// and > characters in our HTML document. This function replaces these regular
// characters with HTML entities: <, >, &, ', and ".
func (c *Config) EscapeDescriptionsDocbook() {
	for _, d := range c.Definitions.All {
		paras := strings.Split(d.Description(), "\n\n")
		var sb strings.Builder
		for _, para := range paras {
			sb.WriteString("<para>")
			sb.WriteString(html.EscapeString(EscapeAsterisks(para)))
			sb.WriteString("</para>")
		}
		d.DescriptionWithEntities = sb.String()

		//		d.DescriptionWithEntities = html.EscapeString(d.Description())

		for _, f := range d.Fields {
			paras := strings.Split(f.Description, "\n\n")
			var sb strings.Builder
			for _, para := range paras {
				sb.WriteString("<para>")
				sb.WriteString(html.EscapeString(EscapeAsterisks(para)))
				sb.WriteString("</para>")
			}
			f.DescriptionWithEntities = sb.String()
		}
	}
}

func (c *Config) EscapeDescriptionsMarkdown() {
	for _, d := range c.Definitions.All {
		d.DescriptionWithEntities = html.EscapeString(d.Description())

		for _, f := range d.Fields {
			f.DescriptionWithEntities = html.EscapeString(f.Description)
		}
	}
}

// For each resource, look up its definition and visit it.
func (c *Config) visitResources() {
	missing := false
	for _, cat := range c.ResourceCategories {
		for _, r := range cat.Resources {
			if d, ok := c.Definitions.GetByVersionKind(r.Group, r.Version, r.Name); ok {
				r.Definition = d
			} else {
				fmt.Printf("Could not find definition for resource: %s %s %s.\n", r.Group, r.Version, r.Name)
				missing = true
				os.Exit(1)
			}
		}
	}
	if missing {
		fmt.Printf("All known definitions: %v\n", c.Definitions.All)
	}
}

// FindConfigDefinitions populates the ConfigDefinitions map
// with Definitions identified in the config file
func (c *Config) FindConfigDefinitions() {
	c.ConfigDefinitions = map[string]struct{}{}
	for _, rescat := range c.ResourceCategories {
		for _, res := range rescat.Resources {
			c.ConfigDefinitions[res.Name] = struct{}{}
		}
		for _, appendix := range rescat.Appendixes {
			for _, def := range appendix.Definitions {
				c.ConfigDefinitions[def.Name] = struct{}{}
			}
		}
	}
}
