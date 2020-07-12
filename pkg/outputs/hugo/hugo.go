package hugo

import (
	"fmt"

	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
)

// Hugo output
// implements the Output interface
type Hugo struct {
	Directory string
}

// NewHugo returns a new Hugo
func NewHugo(dir string) *Hugo {
	return &Hugo{Directory: dir}
}

// Prepare a new output
func (o *Hugo) Prepare() error {
	err := o.addIndex("", map[string]interface{}{
		"title": "Resources",
	})
	if err != nil {
		return fmt.Errorf("Error writing index file in %s: %s", o.Directory, err)
	}
	return nil
}

// AddPart adds a part to the output
func (o *Hugo) AddPart(i int, name string) (outputs.Part, error) {
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
		hugo: o,
		name: partname,
	}, nil
}

// Terminate hugo document
func (o *Hugo) Terminate() error {
	return nil
}
