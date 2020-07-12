package config

import (
	"fmt"

	"github.com/feloy/kubernetes-api-reference/pkg/kubernetes"
)

// LinkEnds maps definition key to a link-end
type LinkEnds map[kubernetes.Key][]string

// Add a new map between key and linkend
func (o LinkEnds) Add(key kubernetes.Key, linkend []string) {
	o[key] = linkend
}

func (o LinkEnds) Debug() {
	for k, v := range o {
		fmt.Printf("%s: %v\n", k, v)
	}
}
