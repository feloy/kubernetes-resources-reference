package api

type Field struct {
	Name                    string
	Type                    string
	Description             string
	DescriptionWithEntities string

	Definition *Definition // Optional Definition for complex types

	PatchStrategy string
	PatchMergeKey string
}

type Fields []*Field

func (a Fields) Len() int           { return len(a) }
func (a Fields) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Fields) Less(i, j int) bool { return a[i].Name < a[j].Name }
