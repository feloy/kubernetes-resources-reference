---
weight: 16
title: "PriorityClass"
description: "PriorityClass defines mapping from a priority class name to the priority integer value."
draft: false
collapsible: false
---
## PriorityClass
`apiVersion: scheduling.k8s.io/v1`
`import "k8s.io/api/scheduling/v1"`
### PriorityClass
PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
- **apiVersion**: scheduling.k8s.io/v1
  
- **kind**: PriorityClass
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **value** (int32), required
  The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
- **description** (string)
  description is an arbitrary string that usually provides guidelines on when this priority class should be used.
- **globalDefault** (boolean)
  globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
- **preemptionPolicy** (string)
  PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
### PriorityClassList
PriorityClassList is a collection of priority classes.
- **apiVersion**: scheduling.k8s.io/v1
  
- **kind**: PriorityClassList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>), required
  items is the list of PriorityClasses
### Operations
#### `get` read the specified PriorityClass

##### HTTP Request
GET /apis/scheduling.k8s.io/v1/priorityclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the PriorityClass
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PriorityClass

##### HTTP Request
GET /apis/scheduling.k8s.io/v1/priorityclasses

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
200 (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclasslist" >}}">PriorityClassList</a>): OK
401: Unauthorized
#### `create` create a PriorityClass

##### HTTP Request
POST /apis/scheduling.k8s.io/v1/priorityclasses

##### Parameters
  - **body** (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>): OK
201 (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>): Created
202 (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>): Accepted
401: Unauthorized
#### `update` replace the specified PriorityClass

##### HTTP Request
PUT /apis/scheduling.k8s.io/v1/priorityclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the PriorityClass
  - **body** (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>): OK
201 (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>): Created
401: Unauthorized
#### `patch` partially update the specified PriorityClass

##### HTTP Request
PATCH /apis/scheduling.k8s.io/v1/priorityclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the PriorityClass
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
200 (<a href="{{< ref "/docs/workloads/priorityclass-v1#priorityclass" >}}">PriorityClass</a>): OK
401: Unauthorized
#### `delete` delete a PriorityClass

##### HTTP Request
DELETE /apis/scheduling.k8s.io/v1/priorityclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the PriorityClass
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
#### `deletecollection` delete collection of PriorityClass

##### HTTP Request
DELETE /apis/scheduling.k8s.io/v1/priorityclasses

##### Parameters
  - **body** (<a href="{{< ref "/docs/common-definitions/deleteoptions-#deleteoptions" >}}">DeleteOptions</a>)
    
  - **?continue** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#continue" >}}">continue</a>
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldSelector** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldselector" >}}">fieldSelector</a>
  - **?gracePeriodSeconds** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#graceperiodseconds" >}}">gracePeriodSeconds</a>
  - **?labelSelector** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#labelselector" >}}">labelSelector</a>
  - **?limit** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#limit" >}}">limit</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>
  - **?propagationPolicy** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#propagationpolicy" >}}">propagationPolicy</a>
  - **?resourceVersion** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#resourceversion" >}}">resourceVersion</a>
  - **?resourceVersionMatch** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#resourceversionmatch" >}}">resourceVersionMatch</a>
  - **?timeoutSeconds** (integer)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#timeoutseconds" >}}">timeoutSeconds</a>

##### Response
200 (<a href="{{< ref "/docs/common-definitions/status-#status" >}}">Status</a>): OK
401: Unauthorized
