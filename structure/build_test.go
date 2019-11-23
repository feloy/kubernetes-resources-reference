/*
Copyright 2019 Philippe Martin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
