---
title: "NodeSelectorRequirement"
description: "A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values."
draft: false
collapsible: false
weight: 11
---
## NodeSelectorRequirement
`import "k8s.io/api/core/v1"`
### NodeSelectorRequirement
A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
- **key** (string), required
  The label key that the selector applies to.
- **operator** (string), required
  Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
- **values** ([]string)
  An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
