package kubernetes_test

import (
	"testing"

	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
)

func TestSpecV118(t *testing.T) {
	spec, err := kubernetes.NewSpec("../../api/v1.18/swagger.json")
	if err != nil {
		t.Errorf("NewSpec should not return an errors but returns %s", err)
	}
	if len(*spec.Resources) != 114 {
		t.Errorf("Spec should contain %d resources but contains %d", 114, len(*spec.Resources))
	}
}

func TestGetResourceV118(t *testing.T) {
	spec, err := kubernetes.NewSpec("../../api/v1.18/swagger.json")
	if err != nil {
		t.Errorf("NewSpec should not return an errors but returns %s", err)
	}
	v1 := newAPIVersionAssert(t, "v1")
	_, res := spec.GetResource("", *v1, "Pod", false)
	if res.Description != "Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts." {
		t.Error("Error getting definition of Pod")
	}
}
