package dbxml

import (
	x "github.com/shabbyrobe/xmlwriter"
)

// NewBook returns a new `book` element
func NewBook() x.Elem {
	return x.Elem{
		Name: "book",
	}
}

// Bookinfo returns a `bookinfo` structure
func Bookinfo(title string, subtitle string, year string, holder string, releaseinfos []string, legalnotices [][]x.Writable) x.Elem {
	result := x.Elem{
		Name: "bookinfo",
		Content: []x.Writable{
			ElemWithText("title", title),
			ElemWithText("subtitle", subtitle),
			x.Elem{
				Name: "copyright",
				Content: []x.Writable{
					ElemWithText("year", year),
					ElemWithText("holder", holder),
				},
			},
		},
	}

	for _, releaseinfo := range releaseinfos {
		result.Content = append(result.Content, ElemWithText("releaseinfo", releaseinfo))
	}
	for _, legalnotice := range legalnotices {
		result.Content = append(result.Content, x.Elem{
			Name: "legalnotice",
			Content: []x.Writable{
				ElemWithContent("para", legalnotice),
			},
		})
	}
	return result
}

// XrefLinkend returns an xref element with a linkend attribute
func XrefLinkend(linkend string) x.Writable {
	return elemWithAttributes("xref", []x.Attr{
		{Name: "linkend", Value: linkend},
	})
}

// Section returns an Elem <nodename><title>...</title></nodename>
func Section(nodename string, title string) x.Elem {
	return x.Elem{
		Name: nodename,
		Content: []x.Writable{
			ElemWithText("title", title),
		},
	}
}

func ElemWithText(nodename string, content string) x.Elem {
	return x.Elem{
		Name: nodename,
		Content: []x.Writable{
			x.Text(content),
		},
	}
}

func ElemWithContent(nodename string, content []x.Writable) x.Elem {
	return x.Elem{
		Name:    nodename,
		Content: content,
	}
}

func elemWithAttributes(nodename string, attrs []x.Attr) x.Writable {
	return x.Elem{
		Name:  "xref",
		Attrs: attrs,
	}
}
