---
title: "StatefulSet"
description: "StatefulSet represents a set of pods with consistent identities."
draft: false
collapsible: false
weight: 8
---
## StatefulSet
`apiVersion: apps/v1`
`import "k8s.io/api/apps/v1"`
### StatefulSet
StatefulSet represents a set of pods with consistent identities. Identities are defined as:
 - Network: A single stable DNS and hostname.
 - Storage: As many VolumeClaims as requested.
The StatefulSet guarantees that a given network identity will always map to the same storage identity.
- **apiVersion**: apps/v1
  
- **kind**: StatefulSet
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulsetspec" >}}">StatefulSetSpec</a>)
  Spec defines the desired identities of pods in this set.
- **status** (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulsetstatus" >}}">StatefulSetStatus</a>)
  Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
### StatefulSetSpec
A StatefulSetSpec is the specification of a StatefulSet.
- **serviceName** (string), required
  serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
- **selector** (<a href="{{< ref "/docs/common-definitions/labelselector-#labelselector" >}}">LabelSelector</a>), required
  selector is a label query over pods that should match the replica count. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
- **template** (<a href="{{< ref "/docs/workloads/podtemplate-v1#podtemplatespec" >}}">PodTemplateSpec</a>), required
  template is the object that describes the pod that will be created if insufficient replicas are detected. Each pod stamped out by the StatefulSet will fulfill this Template, but have a unique identity from the rest of the StatefulSet.
- **replicas** (int32)
  replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
- **updateStrategy** (StatefulSetUpdateStrategy)
  updateStrategy indicates the StatefulSetUpdateStrategy that will be employed to update Pods in the StatefulSet when a revision is made to Template.
*StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.*
  - **updateStrategy.type** (string)
    Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
  - **updateStrategy.rollingUpdate** (RollingUpdateStatefulSetStrategy)
    RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
*RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.*
  - **updateStrategy.rollingUpdate.partition** (int32)
    Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.
- **podManagementPolicy** (string)
  podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
- **revisionHistoryLimit** (int32)
  revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
- **volumeClaimTemplates** ([]<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>)
  volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
### StatefulSetStatus
StatefulSetStatus represents the current state of a StatefulSet.
- **replicas** (int32), required
  replicas is the number of Pods created by the StatefulSet controller.
- **readyReplicas** (int32)
  readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
- **currentReplicas** (int32)
  currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
- **updatedReplicas** (int32)
  updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
- **collisionCount** (int32)
  collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
- **conditions** ([]StatefulSetCondition)
  *Patch strategy: merge on key `type`*
  Represents the latest available observations of a statefulset's current state.
*StatefulSetCondition describes the state of a statefulset at a certain point.*
  - **conditions.status** (string), required
    Status of the condition, one of True, False, Unknown.
  - **conditions.type** (string), required
    Type of statefulset condition.
  - **conditions.lastTransitionTime** (Time)
    Last time the condition transitioned from one status to another.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.message** (string)
    A human readable message indicating details about the transition.
  - **conditions.reason** (string)
    The reason for the condition's last transition.
- **currentRevision** (string)
  currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
- **updateRevision** (string)
  updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
- **observedGeneration** (int64)
  observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
### StatefulSetList
StatefulSetList is a collection of StatefulSets.
- **apiVersion**: apps/v1
  
- **kind**: StatefulSetList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  
- **items** ([]<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>), required
  
### Operations
#### `get` read the specified StatefulSet

##### HTTP Request
GET /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}

##### Parameters
  - **{name}** (string), required
    name of the StatefulSet
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): OK
401: Unauthorized
#### `get` read status of the specified StatefulSet

##### HTTP Request
GET /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the StatefulSet
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind StatefulSet

##### HTTP Request
GET /apis/apps/v1/namespaces/{namespace}/statefulsets

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
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulsetlist" >}}">StatefulSetList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind StatefulSet

##### HTTP Request
GET /apis/apps/v1/statefulsets

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
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulsetlist" >}}">StatefulSetList</a>): OK
401: Unauthorized
#### `create` create a StatefulSet

##### HTTP Request
POST /apis/apps/v1/namespaces/{namespace}/statefulsets

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): OK
201 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): Created
202 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): Accepted
401: Unauthorized
#### `update` replace the specified StatefulSet

##### HTTP Request
PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}

##### Parameters
  - **{name}** (string), required
    name of the StatefulSet
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): OK
201 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): Created
401: Unauthorized
#### `update` replace status of the specified StatefulSet

##### HTTP Request
PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the StatefulSet
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): OK
201 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): Created
401: Unauthorized
#### `patch` partially update the specified StatefulSet

##### HTTP Request
PATCH /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}

##### Parameters
  - **{name}** (string), required
    name of the StatefulSet
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
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified StatefulSet

##### HTTP Request
PATCH /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the StatefulSet
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
200 (<a href="{{< ref "/docs/workloads/statefulset-v1#statefulset" >}}">StatefulSet</a>): OK
401: Unauthorized
#### `delete` delete a StatefulSet

##### HTTP Request
DELETE /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}

##### Parameters
  - **{name}** (string), required
    name of the StatefulSet
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
#### `deletecollection` delete collection of StatefulSet

##### HTTP Request
DELETE /apis/apps/v1/namespaces/{namespace}/statefulsets

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
