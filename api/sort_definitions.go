package api

type SortDefinitionsByName []*Definition

type SortDefinitionsByVersion []*Definition

func (a SortDefinitionsByVersion) Len() int      { return len(a) }
func (a SortDefinitionsByVersion) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortDefinitionsByVersion) Less(i, j int) bool {
	if a[i].Group.String() == a[j].Group.String() {
		return a[i].Version.LessThan(a[j].Version)
	}
	if a[i].Version.String() == a[j].Version.String() {
		return a[i].Group.LessThan(a[j].Group)
	}

	return a[i].Version.LessThan(a[j].Version)
}
