/*
Copyright 2016 The Kubernetes Authors.
Copyright 2019 Philippe Martin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
