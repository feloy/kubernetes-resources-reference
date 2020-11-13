package kwebsite

import (
	"fmt"

	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
)

// KWebsite output
// implements the Output interface
type KWebsite struct {
	Directory    string
	TemplatesDir string
}

// NewKWebsite returns a new KWebsite
func NewKWebsite(dir string, templatesDir string) *KWebsite {
	return &KWebsite{Directory: dir, TemplatesDir: templatesDir}
}

// Prepare a new output
func (o *KWebsite) Prepare() error {
	err := o.addMainIndex()
	if err != nil {
		return fmt.Errorf("Error writing index file in %s: %s", o.Directory, err)
	}
	return nil
}

// AddPart adds a part to the output
func (o *KWebsite) AddPart(i int, name string) (outputs.Part, error) {
	partname := escapeName(name)
	err := o.addPartIndex(partname, name, i+1)
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
