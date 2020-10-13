---
collapsible: false
weight: 4
title: "PodDisruptionBudget v1beta1"
description: "PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods."
draft: false
---
## PodDisruptionBudget
`apiVersion: policy/v1beta1`
`import "k8s.io/api/policy/v1beta1"`
### PodDisruptionBudget
PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
- **apiVersion**: policy/v1beta1
  
- **kind**: PodDisruptionBudget
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudgetspec" >}}">PodDisruptionBudgetSpec</a>)
  Specification of the desired behavior of the PodDisruptionBudget.
- **status** (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudgetstatus" >}}">PodDisruptionBudgetStatus</a>)
  Most recently observed status of the PodDisruptionBudget.
### PodDisruptionBudgetSpec
PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
- **maxUnavailable** (IntOrString)
  An eviction is allowed if at most "maxUnavailable" pods selected by "selector" are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with "minAvailable".
*IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.*
- **minAvailable** (IntOrString)
  An eviction is allowed if at least "minAvailable" pods selected by "selector" will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying "100%!"(MISSING).
*IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.*
- **selector** (<a href="{{< ref "/docs/common-definitions/labelselector-#labelselector" >}}">LabelSelector</a>)
  Label query over pods whose evictions are managed by the disruption budget.
### PodDisruptionBudgetStatus
PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
- **currentHealthy** (int32), required
  current number of healthy pods
- **desiredHealthy** (int32), required
  minimum desired number of healthy pods
- **disruptionsAllowed** (int32), required
  Number of pod disruptions that are currently allowed.
- **expectedPods** (int32), required
  total number of pods counted by this disruption budget
- **disruptedPods** (map[string]Time)
  DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
- **observedGeneration** (int64)
  Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB's object generation.
### PodDisruptionBudgetList
PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
- **apiVersion**: policy/v1beta1
  
- **kind**: PodDisruptionBudgetList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  
- **items** ([]<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>), required
  
### Operations
#### `get` read the specified PodDisruptionBudget

##### HTTP Request
GET /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}

##### Parameters
  - **{name}** (string), required
    name of the PodDisruptionBudget
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): OK
401: Unauthorized
#### `get` read status of the specified PodDisruptionBudget

##### HTTP Request
GET /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PodDisruptionBudget
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PodDisruptionBudget

##### HTTP Request
GET /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
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
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudgetlist" >}}">PodDisruptionBudgetList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PodDisruptionBudget

##### HTTP Request
GET /apis/policy/v1beta1/poddisruptionbudgets

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
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudgetlist" >}}">PodDisruptionBudgetList</a>): OK
401: Unauthorized
#### `create` create a PodDisruptionBudget

##### HTTP Request
POST /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): OK
201 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): Created
202 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): Accepted
401: Unauthorized
#### `update` replace the specified PodDisruptionBudget

##### HTTP Request
PUT /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}

##### Parameters
  - **{name}** (string), required
    name of the PodDisruptionBudget
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): OK
201 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): Created
401: Unauthorized
#### `update` replace status of the specified PodDisruptionBudget

##### HTTP Request
PUT /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PodDisruptionBudget
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): OK
201 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): Created
401: Unauthorized
#### `patch` partially update the specified PodDisruptionBudget

##### HTTP Request
PATCH /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}

##### Parameters
  - **{name}** (string), required
    name of the PodDisruptionBudget
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
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
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified PodDisruptionBudget

##### HTTP Request
PATCH /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PodDisruptionBudget
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
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
200 (<a href="{{< ref "/docs/policies/poddisruptionbudget-v1beta1#poddisruptionbudget" >}}">PodDisruptionBudget</a>): OK
401: Unauthorized
#### `delete` delete a PodDisruptionBudget

##### HTTP Request
DELETE /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}

##### Parameters
  - **{name}** (string), required
    name of the PodDisruptionBudget
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
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
#### `deletecollection` delete collection of PodDisruptionBudget

##### HTTP Request
DELETE /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
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
