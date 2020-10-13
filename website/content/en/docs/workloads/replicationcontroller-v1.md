---
title: "ReplicationController"
description: "ReplicationController represents the configuration of a replication controller."
draft: false
collapsible: false
weight: 5
---
## ReplicationController
`apiVersion: v1`
`import "k8s.io/api/core/v1"`
### ReplicationController
ReplicationController represents the configuration of a replication controller.
- **apiVersion**: v1
  
- **kind**: ReplicationController
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  If the Labels of a ReplicationController are empty, they are defaulted to be the same as the Pod(s) that the replication controller manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontrollerspec" >}}">ReplicationControllerSpec</a>)
  Spec defines the specification of the desired behavior of the replication controller. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontrollerstatus" >}}">ReplicationControllerStatus</a>)
  Status is the most recently observed status of the replication controller. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### ReplicationControllerSpec
ReplicationControllerSpec is the specification of a replication controller.
- **selector** (map[string]string)
  Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
- **template** (<a href="{{< ref "/docs/workloads/podtemplate-v1#podtemplatespec" >}}">PodTemplateSpec</a>)
  Template is the object that describes the pod that will be created if insufficient replicas are detected. This takes precedence over a TemplateRef. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
- **replicas** (int32)
  Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
- **minReadySeconds** (int32)
  Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
### ReplicationControllerStatus
ReplicationControllerStatus represents the current status of a replication controller.
- **replicas** (int32), required
  Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
- **availableReplicas** (int32)
  The number of available replicas (ready for at least minReadySeconds) for this replication controller.
- **readyReplicas** (int32)
  The number of ready replicas for this replication controller.
- **fullyLabeledReplicas** (int32)
  The number of pods that have labels matching the labels of the pod template of the replication controller.
- **conditions** ([]ReplicationControllerCondition)
  *Patch strategy: merge on key `type`*
  Represents the latest available observations of a replication controller's current state.
*ReplicationControllerCondition describes the state of a replication controller at a certain point.*
  - **conditions.status** (string), required
    Status of the condition, one of True, False, Unknown.
  - **conditions.type** (string), required
    Type of replication controller condition.
  - **conditions.lastTransitionTime** (Time)
    The last time the condition transitioned from one status to another.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.message** (string)
    A human readable message indicating details about the transition.
  - **conditions.reason** (string)
    The reason for the condition's last transition.
- **observedGeneration** (int64)
  ObservedGeneration reflects the generation of the most recently observed replication controller.
### ReplicationControllerList
ReplicationControllerList is a collection of replication controllers.
- **apiVersion**: v1
  
- **kind**: ReplicationControllerList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- **items** ([]<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>), required
  List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
### Operations
#### `get` read the specified ReplicationController

##### HTTP Request
GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}

##### Parameters
  - **{name}** (string), required
    name of the ReplicationController
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): OK
401: Unauthorized
#### `get` read status of the specified ReplicationController

##### HTTP Request
GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the ReplicationController
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind ReplicationController

##### HTTP Request
GET /api/v1/namespaces/{namespace}/replicationcontrollers

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
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontrollerlist" >}}">ReplicationControllerList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind ReplicationController

##### HTTP Request
GET /api/v1/replicationcontrollers

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
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontrollerlist" >}}">ReplicationControllerList</a>): OK
401: Unauthorized
#### `create` create a ReplicationController

##### HTTP Request
POST /api/v1/namespaces/{namespace}/replicationcontrollers

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): OK
201 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): Created
202 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): Accepted
401: Unauthorized
#### `update` replace the specified ReplicationController

##### HTTP Request
PUT /api/v1/namespaces/{namespace}/replicationcontrollers/{name}

##### Parameters
  - **{name}** (string), required
    name of the ReplicationController
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): OK
201 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): Created
401: Unauthorized
#### `update` replace status of the specified ReplicationController

##### HTTP Request
PUT /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the ReplicationController
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): OK
201 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): Created
401: Unauthorized
#### `patch` partially update the specified ReplicationController

##### HTTP Request
PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}

##### Parameters
  - **{name}** (string), required
    name of the ReplicationController
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
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified ReplicationController

##### HTTP Request
PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the ReplicationController
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
200 (<a href="{{< ref "/docs/workloads/replicationcontroller-v1#replicationcontroller" >}}">ReplicationController</a>): OK
401: Unauthorized
#### `delete` delete a ReplicationController

##### HTTP Request
DELETE /api/v1/namespaces/{namespace}/replicationcontrollers/{name}

##### Parameters
  - **{name}** (string), required
    name of the ReplicationController
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
#### `deletecollection` delete collection of ReplicationController

##### HTTP Request
DELETE /api/v1/namespaces/{namespace}/replicationcontrollers

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
