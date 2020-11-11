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
		"api_metadata": map[string]interface{}{
			"apiVersion": gv,
			"import":     importPrefix,
			"kind":       name,
		},
		"collapsible":  false,
		"content_type": "api_reference",
		"description":  description,
		"draft":        false,
		"title":        title,
		"weight":       i + 1,
	})
	if err != nil {
		return Chapter{}, fmt.Errorf("Error creating chapter %s/%s: %s", o.name, name, err)
	}

	data := ChapterData{
		ApiVersion: gv,
		Version:    version.String(),
		Import:     importPrefix,
		Kind:       name,
		Metadata: ChapterMetadata{
			Description: description,
			Title:       title,
			Weight:      i + 1,
		},
		ChapterName: name,
	}

	return Chapter{
		kwebsite: o.kwebsite,
		part:     &o,
		name:     chaptername,
		data:     &data,
	}, nil
}
