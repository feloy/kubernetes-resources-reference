package openapi_test

import (
	"testing"

	"github.com/feloy/kubernetes-api-reference/pkg/openapi"
)

func TestLoadOpenAPISpecV118(t *testing.T) {
	spec, err := openapi.LoadOpenAPISpec("../../api/v1.18/swagger.json")
	if err != nil {
		t.Errorf("Failed to load spec")
	}
	if len(spec.Definitions) != 600 {
		t.Errorf("Spec should contain %d definition but contains %d", 600, len(spec.Definitions))
	}
}
