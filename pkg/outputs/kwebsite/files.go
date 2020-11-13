package kwebsite

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func (o *KWebsite) addMainIndex() error {
	t := template.Must(template.ParseFiles("./pkg/outputs/kwebsite/templates/main-index.tmpl"))

	filename := filepath.Join(o.Directory, "_index.md")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, nil)
}

type PartIndex struct {
	Title  string
	Weight int
}

func (o *KWebsite) addPartIndex(subdir string, name string, weight int) error {
	dirname := filepath.Join(o.Directory, subdir)
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		return err
	}

	t := template.Must(template.ParseFiles("./pkg/outputs/kwebsite/templates/part-index.tmpl"))
	data := PartIndex{
		Title:  name,
		Weight: weight,
	}

	filename := filepath.Join(o.Directory, subdir, "_index.md")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, data)
}

// escapeName returns a name usable as file name
func escapeName(s string) string {
	result := strings.ToLower(s)
	result = strings.ReplaceAll(result, " ", "-")
	return result
}
