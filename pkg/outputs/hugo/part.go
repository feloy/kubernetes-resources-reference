package hugo

import (
	"fmt"

	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
)

// Part of a Hugo output
// implements the outputs.Part interface
type Part struct {
	hugo *Hugo
	name string
}

// AddChapter adds a chapter to the Part
func (o Part) AddChapter(i int, name string, version *kubernetes.APIVersion, description string) (outputs.Chapter, error) {
	title := name
	if version != nil && version.Stage != kubernetes.StageGA {
		title += " " + version.String()
	}
	chaptername, err := o.hugo.addChapter(o.name, name, version.String(), map[string]interface{}{
		"title":       title,
		"description": description,
		"draft":       false,
		"collapsible": false,
		"weight":      i + 1,
	})
	if err != nil {
		return Chapter{}, fmt.Errorf("Error creating chapter %s/%s: %s", o.name, name, err)
	}

	return Chapter{
		hugo: o.hugo,
		part: &o,
		name: chaptername,
	}, nil
}
