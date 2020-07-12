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

// ListEntry returns a list entry
func ListEntry(title string, content string, indentLevel int) string {
	titleIndent := strings.Repeat("  ", indentLevel) + "- "
	descIndent := strings.Repeat("  ", indentLevel) + "  "

	// Indent all lines
	parts := strings.Split(content, "\n")
	for i := range parts {
		parts[i] = descIndent + parts[i]
	}
	return fmt.Sprintf("%s%s\n%s\n", titleIndent, title, strings.Join(parts, "\n"))
}
