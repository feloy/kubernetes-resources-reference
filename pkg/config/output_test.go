package config_test

import (
	"testing"

	"github.com/feloy/kubernetes-api-reference/pkg/config"
	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
	"github.com/feloy/kubernetes-api-reference/pkg/outputs"
)

type FakeOutput struct{}

func (o FakeOutput) Prepare() error                                   { return nil }
func (o FakeOutput) AddPart(i int, name string) (outputs.Part, error) { return FakePart{}, nil }
func (o FakeOutput) Terminate() error                                 { return nil }

type FakePart struct{}

func (o FakePart) AddChapter(i int, name string, version *kubernetes.APIVersion, description string) (outputs.Chapter, error) {
	return FakeChapter{}, nil
}

type FakeChapter struct{}

func (o FakeChapter) SetAPIVersion(s string) error { return nil }
func (o FakeChapter) SetGoImport(s string) error   { return nil }
func (o FakeChapter) AddSection(i int, name string) (outputs.Section, error) {
	return FakeSection{}, nil
}

type FakeSection struct{}

func (o FakeSection) AddContent(s string) error        { return nil }
func (o FakeSection) AddTypeDefinition(s string) error { return nil }
func (o FakeSection) AddProperty(name string, property *kubernetes.Property, linkend []string, indent bool) error {
	return nil
}
func (o FakeSection) EndProperty() error       { return nil }
func (o FakeSection) StartPropertyList() error { return nil }
func (o FakeSection) EndPropertyList() error   { return nil }
func (o FakeSection) AddOperation(operation *kubernetes.ActionInfo, linkends kubernetes.LinkEnds) error {
	return nil
}

func TestOutputDocumentV116(t *testing.T) {
	outputDocumentVersion(t, "v1.16")
}

func TestOutputDocumentV117(t *testing.T) {
	outputDocumentVersion(t, "v1.17")
}

func TestOutputDocumentV118(t *testing.T) {
	outputDocumentVersion(t, "v1.18")
}

func TestOutputDocumentV119(t *testing.T) {
	outputDocumentVersion(t, "v1.19")
}

func outputDocumentVersion(t *testing.T, version string) {
	spec, err := kubernetes.NewSpec("../../api/" + version + "/swagger.json")
	if err != nil {
		t.Errorf("Error loding swagger file")
	}

	toc, err := config.LoadTOC("../../config/" + version + "/toc.yaml")
	if err != nil {
		t.Errorf("LoadTOC should not fail")
	}

	err = toc.PopulateAssociates(spec)
	if err != nil {
		t.Errorf("%s", err)
	}
	toc.Definitions = &spec.Swagger.Definitions
	toc.OutputDocument(FakeOutput{})

}
