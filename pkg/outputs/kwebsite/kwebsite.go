package kwebsite

import (
	"fmt"

	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
)

// KWebsite output
// implements the Output interface
type KWebsite struct {
	Directory string
}

// NewKWebsite returns a new KWebsite
func NewKWebsite(dir string) *KWebsite {
	return &KWebsite{Directory: dir}
}

// Prepare a new output
func (o *KWebsite) Prepare() error {
	err := o.addIndex("", map[string]interface{}{
		"title": "Resources",
	})
	if err != nil {
		return fmt.Errorf("Error writing index file in %s: %s", o.Directory, err)
	}
	return nil
}

// AddPart adds a part to the output
func (o *KWebsite) AddPart(i int, name string) (outputs.Part, error) {
	partname, err := o.addPart(name)
	if err != nil {
		return Part{}, fmt.Errorf("Error creating part %s: %s", name, err)
	}
	err = o.addIndex(partname, map[string]interface{}{
		"title":       name,
		"draft":       false,
		"collapsible": true,
		"weight":      i + 1,
	})
	if err != nil {
		return Part{}, fmt.Errorf("Error writing index file for part %s: %s", name, err)
	}
	return Part{
		kwebsite: o,
		name:     partname,
	}, nil
}

// Terminate kwebsite document
func (o *KWebsite) Terminate() error {
	return nil
}
