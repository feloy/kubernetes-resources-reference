---
collapsible: false
weight: 10
title: "DaemonSet"
description: "DaemonSet represents the configuration of a daemon set."
draft: false
---
## DaemonSet
`apiVersion: apps/v1`
`import "k8s.io/api/apps/v1"`
### DaemonSet
DaemonSet represents the configuration of a daemon set.
- **apiVersion**: apps/v1
  
- **kind**: DaemonSet
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonsetspec" >}}">DaemonSetSpec</a>)
  The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonsetstatus" >}}">DaemonSetStatus</a>)
  The current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### DaemonSetSpec
DaemonSetSpec is the specification of a daemon set.
- **selector** (<a href="{{< ref "/docs/common-definitions/labelselector-#labelselector" >}}">LabelSelector</a>), required
  A label query over pods that are managed by the daemon set. Must match in order to be controlled. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
- **template** (<a href="{{< ref "/docs/workloads/podtemplate-v1#podtemplatespec" >}}">PodTemplateSpec</a>), required
  An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
- **minReadySeconds** (int32)
  The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
- **updateStrategy** (DaemonSetUpdateStrategy)
  An update strategy to replace existing DaemonSet pods with new pods.
*DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.*
  - **updateStrategy.type** (string)
    Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
  - **updateStrategy.rollingUpdate** (RollingUpdateDaemonSet)
    Rolling update config params. Present only if type = "RollingUpdate".
*Spec to control the desired behavior of daemon set rolling update.*
  - **updateStrategy.rollingUpdate.maxUnavailable** (IntOrString)
    The maximum number of DaemonSet pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total number of DaemonSet pods at the start of the update (ex: 10%!)(MISSING). Absolute number is calculated from percentage by rounding up. This cannot be 0. Default value is 1. Example: when this is set to 30%!,(MISSING) at most 30%!o(MISSING)f the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their pods stopped for an update at any given time. The update starts by stopping at most 30%!o(MISSING)f those DaemonSet pods and then brings up new DaemonSet pods in their place. Once the new pods are available, it then proceeds onto other DaemonSet pods, thus ensuring that at least 70%!o(MISSING)f original number of DaemonSet pods are available at all times during the update.
*IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.*
- **revisionHistoryLimit** (int32)
  The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
### DaemonSetStatus
DaemonSetStatus represents the current status of a daemon set.
- **numberReady** (int32), required
  The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
- **numberAvailable** (int32)
  The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)
- **numberUnavailable** (int32)
  The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)
- **numberMisscheduled** (int32), required
  The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
- **desiredNumberScheduled** (int32), required
  The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
- **currentNumberScheduled** (int32), required
  The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
- **updatedNumberScheduled** (int32)
  The total number of nodes that are running updated daemon pod
- **collisionCount** (int32)
  Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
- **conditions** ([]DaemonSetCondition)
  *Patch strategy: merge on key `type`*
  Represents the latest available observations of a DaemonSet's current state.
*DaemonSetCondition describes the state of a DaemonSet at a certain point.*
  - **conditions.status** (string), required
    Status of the condition, one of True, False, Unknown.
  - **conditions.type** (string), required
    Type of DaemonSet condition.
  - **conditions.lastTransitionTime** (Time)
    Last time the condition transitioned from one status to another.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.message** (string)
    A human readable message indicating details about the transition.
  - **conditions.reason** (string)
    The reason for the condition's last transition.
- **observedGeneration** (int64)
  The most recent generation observed by the daemon set controller.
### DaemonSetList
DaemonSetList is a collection of daemon sets.
- **apiVersion**: apps/v1
  
- **kind**: DaemonSetList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>), required
  A list of daemon sets.
### Operations
#### `get` read the specified DaemonSet

##### HTTP Request
GET /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}

##### Parameters
  - **{name}** (string), required
    name of the DaemonSet
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): OK
401: Unauthorized
#### `get` read status of the specified DaemonSet

##### HTTP Request
GET /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the DaemonSet
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind DaemonSet

##### HTTP Request
GET /apis/apps/v1/namespaces/{namespace}/daemonsets

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
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonsetlist" >}}">DaemonSetList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind DaemonSet

##### HTTP Request
GET /apis/apps/v1/daemonsets

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
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonsetlist" >}}">DaemonSetList</a>): OK
401: Unauthorized
#### `create` create a DaemonSet

##### HTTP Request
POST /apis/apps/v1/namespaces/{namespace}/daemonsets

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): OK
201 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): Created
202 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): Accepted
401: Unauthorized
#### `update` replace the specified DaemonSet

##### HTTP Request
PUT /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}

##### Parameters
  - **{name}** (string), required
    name of the DaemonSet
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): OK
201 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): Created
401: Unauthorized
#### `update` replace status of the specified DaemonSet

##### HTTP Request
PUT /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the DaemonSet
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): OK
201 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): Created
401: Unauthorized
#### `patch` partially update the specified DaemonSet

##### HTTP Request
PATCH /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}

##### Parameters
  - **{name}** (string), required
    name of the DaemonSet
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
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified DaemonSet

##### HTTP Request
PATCH /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the DaemonSet
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
200 (<a href="{{< ref "/docs/workloads/daemonset-v1#daemonset" >}}">DaemonSet</a>): OK
401: Unauthorized
#### `delete` delete a DaemonSet

##### HTTP Request
DELETE /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}

##### Parameters
  - **{name}** (string), required
    name of the DaemonSet
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
#### `deletecollection` delete collection of DaemonSet

##### HTTP Request
DELETE /apis/apps/v1/namespaces/{namespace}/daemonsets

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
