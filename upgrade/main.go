package main

import (
	"fmt"
	"log"

	"github.com/feloy/k8s-api/api"
	"gopkg.in/yaml.v2"
)

func main() {
	*api.ConfigDir = ".."
	config := api.NewConfig()

	for _, rescat := range config.ResourceCategories {
		for _, res := range rescat.Resources {
			d := res.Definition
			inlines := d.Inline
			for _, def := range append(inlines, d) {
				found := false
				for _, fieldcat := range res.FieldCategories {
					if *fieldcat.Definition == def.Name {
						found = true
						break
					}
				}
				if found {
					continue
				}
				fields := []api.FieldCategory{}
				for _, defField := range def.Fields {
					fields = append(fields, api.FieldCategory{
						Name: defField.Name,
					})
				}
				list := []api.FieldsCategory{
					{
						Fields: fields,
					},
				}
				// Add field category with the fields of the definition
				res.FieldCategories = append(res.FieldCategories, &api.FieldCategories{
					Definition: &def.Name,
					List:       list,
				})
			}
		}
	}

	bytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(bytes))
}
