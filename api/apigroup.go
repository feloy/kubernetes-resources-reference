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

type ApiGroup string

func (g ApiGroup) String() string {
	return string(g)
}

func (this ApiGroup) LessThan(that ApiGroup) bool {
	// admissionregistration > apiextensions
	if this.String() == "apiextensions" && that.String() == "admissionregistration" {
		return true
	}
	if that.String() == "apiextensions" && this.String() == "admissionregistration" {
		return false
	}

	// policy > extensions
	if this.String() == "extensions" && that.String() == "policy" {
		return true
	}
	if that.String() == "extensions" && this.String() == "policy" {
		return false
	}
	return this.String() < that.String()
}
