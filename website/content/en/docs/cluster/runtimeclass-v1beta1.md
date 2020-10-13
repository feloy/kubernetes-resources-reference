---
collapsible: false
weight: 6
title: "RuntimeClass v1beta1"
description: "RuntimeClass defines a class of container runtime supported in the cluster."
draft: false
---
## RuntimeClass
`apiVersion: node.k8s.io/v1beta1`
`import "k8s.io/api/node/v1beta1"`
### RuntimeClass
RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
- **apiVersion**: node.k8s.io/v1beta1
  
- **kind**: RuntimeClass
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **handler** (string), required
  Handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must conform to the DNS Label (RFC 1123) requirements, and is immutable.
- **overhead** (Overhead)
  Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
*Overhead structure represents the resource overhead associated with running a pod.*
  - **overhead.podFixed** (map[string]<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
    PodFixed represents the fixed resource overhead associated with running a pod.
- **scheduling** (Scheduling)
  Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
*Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.*
  - **scheduling.nodeSelector** (map[string]string)
    nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
  - **scheduling.tolerations** ([]Toleration)
    *Atomic: will be replaced during a merge*
    tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
*The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.*
  - **scheduling.tolerations.key** (string)
    Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
  - **scheduling.tolerations.operator** (string)
    Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
  - **scheduling.tolerations.value** (string)
    Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
  - **scheduling.tolerations.effect** (string)
    Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
  - **scheduling.tolerations.tolerationSeconds** (int64)
    TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
### RuntimeClassList
RuntimeClassList is a list of RuntimeClass objects.
- **apiVersion**: node.k8s.io/v1beta1
  
- **kind**: RuntimeClassList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>), required
  Items is a list of schema objects.
### Operations
#### `get` read the specified RuntimeClass

##### HTTP Request
GET /apis/node.k8s.io/v1beta1/runtimeclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the RuntimeClass
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind RuntimeClass

##### HTTP Request
GET /apis/node.k8s.io/v1beta1/runtimeclasses

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
200 (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclasslist" >}}">RuntimeClassList</a>): OK
401: Unauthorized
#### `create` create a RuntimeClass

##### HTTP Request
POST /apis/node.k8s.io/v1beta1/runtimeclasses

##### Parameters
  - **body** (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>): OK
201 (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>): Created
202 (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>): Accepted
401: Unauthorized
#### `update` replace the specified RuntimeClass

##### HTTP Request
PUT /apis/node.k8s.io/v1beta1/runtimeclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the RuntimeClass
  - **body** (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>): OK
201 (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>): Created
401: Unauthorized
#### `patch` partially update the specified RuntimeClass

##### HTTP Request
PATCH /apis/node.k8s.io/v1beta1/runtimeclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the RuntimeClass
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
200 (<a href="{{< ref "/docs/cluster/runtimeclass-v1beta1#runtimeclass" >}}">RuntimeClass</a>): OK
401: Unauthorized
#### `delete` delete a RuntimeClass

##### HTTP Request
DELETE /apis/node.k8s.io/v1beta1/runtimeclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the RuntimeClass
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
#### `deletecollection` delete collection of RuntimeClass

##### HTTP Request
DELETE /apis/node.k8s.io/v1beta1/runtimeclasses

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
