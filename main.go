/*
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

package main

import (
	"os"

	"github.com/feloy/k8s-api/api"
	"github.com/feloy/k8s-api/docbook"
	"github.com/feloy/k8s-api/markdown"
	"github.com/feloy/k8s-api/structure"
)

func main() {
	*api.ConfigDir = "."
	config := api.NewConfig()

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	structure := structure.Build(config)

	switch os.Args[1] {
	case "docbook":
		docbook.Generate(os.Stdout, config, structure)
	case "md":
		markdown.Generate("md", config, structure)
	}

}
