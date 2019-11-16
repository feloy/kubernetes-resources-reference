package main

import (
	"os"

	"github.com/feloy/k8s-api/api"
	"github.com/feloy/k8s-api/docbook"
	"github.com/feloy/k8s-api/structure"
)

func main() {
	*api.ConfigDir = "."
	config := api.NewConfig()

	//	fmt.Printf("%+v\n", config.ConfigDefinitions)
	//	os.Exit(0)
	structure := structure.Build(config)
	//	structure.AsMarkdown(os.Stdout)

	docbook.Generate(os.Stdout, config, structure)
}
