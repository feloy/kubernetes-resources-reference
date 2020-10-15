package markdown

import (
	"fmt"
	"strings"
)

// Code returns 's' as code
func Code(s string) string {
	return fmt.Sprintf("`%s`", s)
}

// Italic returns text s in italic
func Italic(s string) string {
	return fmt.Sprintf("*%s*", s)
}

// Chapter returns a Level 2 mark
func Chapter(name string) string {
	return fmt.Sprintf("## %s\n", name)
}

// Section returns a Level 3 mark
func Section(name string) string {
	return fmt.Sprintf("### %s\n", name)
}

// Subsection returns a Level 4 mark
func Subsection(name string) string {
	return fmt.Sprintf("#### %s\n", name)
}

// ListEntry returns a list entry
func ListEntry(title string, content string, indentLevel int, nl bool) string {
	titleIndent := strings.Repeat("  ", indentLevel) + "- "
	descIndent := strings.Repeat("  ", indentLevel) + "  "

	// Indent all lines
	parts := strings.Split(content, "\n")
	for i := range parts {
		parts[i] = descIndent + parts[i]
	}

	separator := "\n"
	if nl {
		separator = "\n\n"
	}
	return fmt.Sprintf("%s%s%s%s\n", titleIndent, title, separator, strings.Join(parts, "\n"))
}
