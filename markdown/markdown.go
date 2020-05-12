package markdown

import (
	"os"
	"path"

	"github.com/feloy/k8s-api/api"
	"github.com/feloy/k8s-api/structure"
)

// Generate Markdown documents from structure
func Generate(dir string, config *api.Config, document *structure.Section) {
	f, err := os.Create(path.Join(dir, "_index.md"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("---\ntitle: \"Resources\"\n---\n")
	document.AsMarkdown(dir, 1, nil)
}
