---
description: "ResourceFieldSelector represents container resources (cpu, memory) and their output format."
draft: false
collapsible: false
weight: 19
title: "ResourceFieldSelector"
---
## ResourceFieldSelector
`import "k8s.io/api/core/v1"`
### ResourceFieldSelector
ResourceFieldSelector represents container resources (cpu, memory) and their output format
- **resource** (string), required
  Required: resource to select
- **containerName** (string)
  Container name: required for volumes, optional for env vars
- **divisor** (<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
  Specifies the output format of the exposed resources, defaults to "1"
