---
title: "ObjectFieldSelector"
description: "ObjectFieldSelector selects an APIVersioned field of an object."
draft: false
collapsible: false
weight: 12
---
## ObjectFieldSelector
`import "k8s.io/api/core/v1"`
### ObjectFieldSelector
ObjectFieldSelector selects an APIVersioned field of an object.
- **fieldPath** (string), required
  Path of the field to select in the specified API version.
- **apiVersion** (string)
  Version of the schema the FieldPath is written in terms of, defaults to "v1".
