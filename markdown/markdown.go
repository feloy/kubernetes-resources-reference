package markdown

import (
	"github.com/feloy/k8s-api/api"
	"github.com/feloy/k8s-api/structure"
)

// Generate Markdown documents from structure
func Generate(dir string, config *api.Config, document *structure.Section) {
	document.AsMarkdown(dir, 1, nil)
}
