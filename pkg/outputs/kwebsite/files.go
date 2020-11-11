package kwebsite

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/feloy/kubernetes-api-reference/pkg/formats/markdown"
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

// addPart adds a directory in the KWebsite content
func (o *KWebsite) addPart(name string) (string, error) {
	subdir := escapeName(name)
	dirname := filepath.Join(o.Directory, subdir)
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		return "", err
	}
	return subdir, nil
}

// addChapter adds a chapter to the part
func (o *KWebsite) addChapter(partname string, name string, version string, metadata map[string]interface{}) (string, error) {
	chaptername := escapeName(name + "-" + version)
	filename := filepath.Join(o.Directory, partname, chaptername) + ".md"
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	writeMetadata(f, metadata, 0)
	fmt.Fprintf(f, markdown.Chapter(name))
	return chaptername, nil
}

func (o *KWebsite) getChapterFile(partname string, chaptername string) (*os.File, error) {
	filename := filepath.Join(o.Directory, partname, chaptername) + ".md"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("error opening file")
		return nil, err
	}
	return f, err
}

// addSection adds a section to the chapter
func (o *KWebsite) addSection(partname string, chaptername string, name string) error {
	f, err := o.getChapterFile(partname, chaptername)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, markdown.Section(name))
	return nil
}

// addSection adds a section to the chapter
func (o *KWebsite) addSubsection(partname string, chaptername string, name string) error {
	f, err := o.getChapterFile(partname, chaptername)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, markdown.Subsection(name))
	return nil
}

// addContent adds content to the chapter in part
func (o *KWebsite) addContent(partname string, chaptername string, content string) error {
	f, err := o.getChapterFile(partname, chaptername)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%s\n", content)
	if err != nil {
		fmt.Print("error printing in file")
	}
	return err
}

// addListEntry adds a list entry to the chapter in part
func (o *KWebsite) addListEntry(partname string, chaptername string, title string, content string, indentLevel int) error {
	f, err := o.getChapterFile(partname, chaptername)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, markdown.ListEntry(title, content, indentLevel, true))
	if err != nil {
		fmt.Print("error printing in file")
	}
	return err
}

// writeMetadata writes metadata at the beginning of a Markdown file
func writeMetadata(f io.Writer, metadata map[string]interface{}, indent int) error {
	if indent == 0 {
		_, err := fmt.Fprint(f, "---\n")
		if err != nil {
			return err
		}
	}

	keys := make([]string, 0, len(metadata))
	for k := range metadata {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := metadata[k]
		switch v.(type) {
		case string:
			_, err := fmt.Fprintf(f, "%s%s: \"%v\"\n", strings.Repeat(" ", indent*2), k, v)
			if err != nil {
				return err
			}
		case map[string]interface{}:
			_, err := fmt.Fprintf(f, "%s:\n", k)
			if err != nil {
				return err
			}
			err = writeMetadata(f, v.(map[string]interface{}), indent+1)
			if err != nil {
				return err
			}
		default:
			_, err := fmt.Fprintf(f, "%s%s: %v\n", strings.Repeat(" ", indent*2), k, v)
			if err != nil {
				return err
			}
		}
	}
	if indent == 0 {
		_, err := fmt.Fprint(f, "---\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// escapeName returns a name usable as file name
func escapeName(s string) string {
	result := strings.ToLower(s)
	result = strings.ReplaceAll(result, " ", "-")
	return result
}
