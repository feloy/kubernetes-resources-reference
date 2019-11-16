package structure_test

import (
	"bytes"
	"testing"

	"github.com/feloy/k8s-api/api"
	"github.com/feloy/k8s-api/structure"
	"gotest.tools/assert"
)

func TestBuild(t *testing.T) {
	*api.ConfigDir = "test_files/1"
	config := api.NewConfig()
	document := structure.Build(config)
	var buf bytes.Buffer
	document.AsMarkdown(&buf)
	assert.Equal(t, `# Category 1
## Definition1
### Definition1
- prop1 {[]string}
### Definition1Spec
- name {string}
#### Spec fields
- aspecprop1 {[]Definition1Sub}
  - aspecprop1.subprop1 {[]string}
  - aspecprop1.subprop2 {string}
- aspecprop2 {[]string}
`, buf.String())
}
