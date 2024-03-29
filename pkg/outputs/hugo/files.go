package hugo

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/feloy/kubernetes-api-reference/pkg/formats/markdown"
)

// addIndex adds an _index.md file to a Hugo directory
func (o *Hugo) addIndex(subdir string, metadata map[string]interface{}) error {

	filename := filepath.Join(o.Directory, subdir, "_index.md")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return writeMetadata(f, metadata)
}

// addPart adds a directory in the Hugo content
func (o *Hugo) addPart(name string) (string, error) {
	subdir := escapeName(name)
	dirname := filepath.Join(o.Directory, subdir)
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		return "", err
	}
	return subdir, nil
}

// addChapter adds a chapter to the part
func (o *Hugo) addChapter(partname string, name string, version string, metadata map[string]interface{}) (string, error) {
	chaptername := escapeName(name + "-" + version)
	filename := filepath.Join(o.Directory, partname, chaptername) + ".md"
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	writeMetadata(f, metadata)
	fmt.Fprintf(f, markdown.Chapter(name))
	return chaptername, nil
}

func (o *Hugo) getChapterFile(partname string, chaptername string) (*os.File, error) {
	filename := filepath.Join(o.Directory, partname, chaptername) + ".md"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("error opening file")
		return nil, err
	}
	return f, err
}

// addSection adds a section to the chapter
func (o *Hugo) addSection(partname string, chaptername string, name string) error {
	f, err := o.getChapterFile(partname, chaptername)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, markdown.Section(name))
	return nil
}

// addSection adds a section to the chapter
func (o *Hugo) addSubsection(partname string, chaptername string, name string) error {
	f, err := o.getChapterFile(partname, chaptername)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, markdown.Subsection(name))
	return nil
}

// addContent adds content to the chapter in part
func (o *Hugo) addContent(partname string, chaptername string, content string) error {
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
func (o *Hugo) addListEntry(partname string, chaptername string, title string, content string, indentLevel int) error {
	f, err := o.getChapterFile(partname, chaptername)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, markdown.ListEntry(title, content, indentLevel, false))
	if err != nil {
		fmt.Print("error printing in file")
	}
	return err
}

// writeMetadata writes metadata at the beginning of a Markdown file
func writeMetadata(f io.Writer, metadata map[string]interface{}) error {
	_, err := fmt.Fprint(f, "---\n")
	if err != nil {
		return err
	}
	for k, v := range metadata {
		switch v.(type) {
		case string:
			_, err = fmt.Fprintf(f, "%s: \"%v\"\n", k, v)
			if err != nil {
				return err
			}
		default:
			_, err = fmt.Fprintf(f, "%s: %v\n", k, v)
			if err != nil {
				return err
			}
		}
	}
	_, err = fmt.Fprint(f, "---\n")
	if err != nil {
		return err
	}
	return nil
}

// escapeName returns a name usable as file name
func escapeName(s string) string {
	result := strings.ToLower(s)
	result = strings.ReplaceAll(result, " ", "-")
	return result
}
