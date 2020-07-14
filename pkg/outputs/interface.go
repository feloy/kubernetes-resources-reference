package outputs

import (
	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
)

// Output is an interface for output formats
type Output interface {
	Prepare() error
	AddPart(i int, name string) (Part, error)
	Terminate() error
}

// Part is an interface to a part of an output
type Part interface {
	AddChapter(i int, name string, version *kubernetes.APIVersion, description string) (Chapter, error)
}

// Chapter is an interface to a chapter of an output
type Chapter interface {
	SetAPIVersion(s string) error
	SetGoImport(s string) error
	AddSection(i int, name string) (Section, error)
}

// Section is an interface to a section of an output
type Section interface {
	AddContent(s string) error
	AddTypeDefinition(s string) error
	StartPropertyList() error
	AddProperty(name string, property *kubernetes.Property, linkend []string, indent bool) error
	EndProperty() error
	EndPropertyList() error
	AddOperation(operation *kubernetes.ActionInfo) error
}
