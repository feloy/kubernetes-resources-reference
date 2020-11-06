package kwebsite

import (
	"fmt"

	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
)

// Part of a KWebsite output
// implements the outputs.Part interface
type Part struct {
	kwebsite *KWebsite
	name     string
}

// AddChapter adds a chapter to the Part
func (o Part) AddChapter(i int, name string, gv string, version *kubernetes.APIVersion, description string, importPrefix string) (outputs.Chapter, error) {
	title := name
	if version != nil && version.Stage != kubernetes.StageGA {
		title += " " + version.String()
	}
	chaptername, err := o.kwebsite.addChapter(o.name, name, version.String(), map[string]interface{}{
		"title":        title,
		"description":  description,
		"draft":        false,
		"collapsible":  false,
		"weight":       i + 1,
		"content_type": "api_reference",
		"api_metadata": map[string]interface{}{
			"kind":       name,
			"apiVersion": gv,
			"import":     importPrefix,
		},
	})
	if err != nil {
		return Chapter{}, fmt.Errorf("Error creating chapter %s/%s: %s", o.name, name, err)
	}

	return Chapter{
		kwebsite: o.kwebsite,
		part:     &o,
		name:     chaptername,
	}, nil
}
