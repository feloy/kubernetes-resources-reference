---
title: "Binding"
description: "Binding ties one object to another; for example, a pod is bound to a node by a scheduler."
draft: false
collapsible: false
weight: 9
---
## Binding
`apiVersion: v1`
`import "k8s.io/api/core/v1"`
### Binding
Binding ties one object to another; for example, a pod is bound to a node by a scheduler. Deprecated in 1.7, please use the bindings subresource of pods instead.
- **apiVersion**: v1
  
- **kind**: Binding
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **target** (<a href="{{< ref "/docs/common-definitions/objectreference-#objectreference" >}}">ObjectReference</a>), required
  The target object that you want to bind to the standard object.
### Operations
#### `create` create a Binding

##### HTTP Request
POST /api/v1/namespaces/{namespace}/bindings

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/cluster/binding-v1#binding" >}}">Binding</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/binding-v1#binding" >}}">Binding</a>): OK
201 (<a href="{{< ref "/docs/cluster/binding-v1#binding" >}}">Binding</a>): Created
202 (<a href="{{< ref "/docs/cluster/binding-v1#binding" >}}">Binding</a>): Accepted
401: Unauthorized
#### `create` create binding of a Pod

##### HTTP Request
POST /api/v1/namespaces/{namespace}/pods/{name}/binding

##### Parameters
  - **{name}** (string), required
    name of the Binding
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/cluster/binding-v1#binding" >}}">Binding</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/binding-v1#binding" >}}">Binding</a>): OK
201 (<a href="{{< ref "/docs/cluster/binding-v1#binding" >}}">Binding</a>): Created
202 (<a href="{{< ref "/docs/cluster/binding-v1#binding" >}}">Binding</a>): Accepted
401: Unauthorized
