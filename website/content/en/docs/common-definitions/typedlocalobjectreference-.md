---
title: "TypedLocalObjectReference"
description: "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace."
draft: false
collapsible: false
weight: 22
---
## TypedLocalObjectReference
`import "k8s.io/api/core/v1"`
### TypedLocalObjectReference
TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
- **kind** (string), required
  Kind is the type of resource being referenced
- **name** (string), required
  Name is the name of resource being referenced
- **apiGroup** (string)
  APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
