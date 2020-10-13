---
title: "Namespace"
description: "Namespace provides a scope for Names."
draft: false
collapsible: false
weight: 2
---
## Namespace
`apiVersion: v1`
`import "k8s.io/api/core/v1"`
### Namespace
Namespace provides a scope for Names. Use of multiple namespaces is optional.
- **apiVersion**: v1
  
- **kind**: Namespace
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/cluster/namespace-v1#namespacespec" >}}">NamespaceSpec</a>)
  Spec defines the behavior of the Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/cluster/namespace-v1#namespacestatus" >}}">NamespaceStatus</a>)
  Status describes the current status of a Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### NamespaceSpec
NamespaceSpec describes the attributes on a Namespace.
- **finalizers** ([]string)
  Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
### NamespaceStatus
NamespaceStatus is information about the current status of a Namespace.
- **conditions** ([]NamespaceCondition)
  *Patch strategy: merge on key `type`*
  Represents the latest available observations of a namespace's current state.
*NamespaceCondition contains details about state of namespace.*
  - **conditions.status** (string), required
    Status of the condition, one of True, False, Unknown.
  - **conditions.type** (string), required
    Type of namespace controller condition.
  - **conditions.lastTransitionTime** (Time)
    
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.message** (string)
    
  - **conditions.reason** (string)
    
- **phase** (string)
  Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
### NamespaceList
NamespaceList is a list of Namespaces.
- **apiVersion**: v1
  
- **kind**: NamespaceList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- **items** ([]<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>), required
  Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
### Operations
#### `get` read the specified Namespace

##### HTTP Request
GET /api/v1/namespaces/{name}

##### Parameters
  - **{name}** (string), required
    name of the Namespace
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): OK
401: Unauthorized
#### `get` read status of the specified Namespace

##### HTTP Request
GET /api/v1/namespaces/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Namespace
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind Namespace

##### HTTP Request
GET /api/v1/namespaces

##### Parameters
  - **?allowWatchBookmarks** (boolean)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#allowwatchbookmarks" >}}">allowWatchBookmarks</a>
  - **?continue** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#continue" >}}">continue</a>
  - **?fieldSelector** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldselector" >}}">fieldSelector</a>
  - **?labelSelector** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#labelselector" >}}">labelSelector</a>
  - **?limit** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#limit" >}}">limit</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>
  - **?resourceVersion** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#resourceversion" >}}">resourceVersion</a>
  - **?resourceVersionMatch** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#resourceversionmatch" >}}">resourceVersionMatch</a>
  - **?timeoutSeconds** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#timeoutseconds" >}}">timeoutSeconds</a>
  - **?watch** (boolean)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#watch" >}}">watch</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespacelist" >}}">NamespaceList</a>): OK
401: Unauthorized
#### `create` create a Namespace

##### HTTP Request
POST /api/v1/namespaces

##### Parameters
  - **body** (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): OK
201 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): Created
202 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): Accepted
401: Unauthorized
#### `update` replace the specified Namespace

##### HTTP Request
PUT /api/v1/namespaces/{name}

##### Parameters
  - **{name}** (string), required
    name of the Namespace
  - **body** (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): OK
201 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): Created
401: Unauthorized
#### `update` replace finalize of the specified Namespace

##### HTTP Request
PUT /api/v1/namespaces/{name}/finalize

##### Parameters
  - **{name}** (string), required
    name of the Namespace
  - **body** (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): OK
201 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): Created
401: Unauthorized
#### `update` replace status of the specified Namespace

##### HTTP Request
PUT /api/v1/namespaces/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Namespace
  - **body** (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): OK
201 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): Created
401: Unauthorized
#### `patch` partially update the specified Namespace

##### HTTP Request
PATCH /api/v1/namespaces/{name}

##### Parameters
  - **{name}** (string), required
    name of the Namespace
  - **body** (<a href="{{< ref "/docs/common-definitions/patch-#patch" >}}">Patch</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?force** (boolean)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#force" >}}">force</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified Namespace

##### HTTP Request
PATCH /api/v1/namespaces/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Namespace
  - **body** (<a href="{{< ref "/docs/common-definitions/patch-#patch" >}}">Patch</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?force** (boolean)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#force" >}}">force</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/namespace-v1#namespace" >}}">Namespace</a>): OK
401: Unauthorized
#### `delete` delete a Namespace

##### HTTP Request
DELETE /api/v1/namespaces/{name}

##### Parameters
  - **{name}** (string), required
    name of the Namespace
  - **body** (<a href="{{< ref "/docs/common-definitions/deleteoptions-#deleteoptions" >}}">DeleteOptions</a>)
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?gracePeriodSeconds** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#graceperiodseconds" >}}">gracePeriodSeconds</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>
  - **?propagationPolicy** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#propagationpolicy" >}}">propagationPolicy</a>

##### Response
200 (<a href="{{< ref "/docs/common-definitions/status-#status" >}}">Status</a>): OK
202 (<a href="{{< ref "/docs/common-definitions/status-#status" >}}">Status</a>): Accepted
401: Unauthorized
