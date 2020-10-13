---
title: "CSINode"
description: "CSINode holds information about all CSI drivers installed on a node."
draft: false
collapsible: false
weight: 9
---
## CSINode
`apiVersion: storage.k8s.io/v1`
`import "k8s.io/api/storage/v1"`
### CSINode
CSINode holds information about all CSI drivers installed on a node. CSI drivers do not need to create the CSINode object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node object.
- **apiVersion**: storage.k8s.io/v1
  
- **kind**: CSINode
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  metadata.name must be the Kubernetes node name.
- **spec** (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinodespec" >}}">CSINodeSpec</a>), required
  spec is the specification of CSINode
### CSINodeSpec
CSINodeSpec holds information about the specification of all CSI drivers installed on a node
- **drivers** ([]CSINodeDriver), required
  *Patch strategy: merge on key `name`*
  drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty.
*CSINodeDriver holds information about the specification of one CSI driver installed on a node*
  - **drivers.name** (string), required
    This is the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI GetPluginName() call for that driver.
  - **drivers.nodeID** (string), required
    nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as "node1", but the storage system may refer to the same node as "nodeA". When Kubernetes issues a command to the storage system to attach a volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will understand, e.g. "nodeA" instead of "node1". This field is required.
  - **drivers.allocatable** (VolumeNodeResources)
    allocatable represents the volume resources of a node that are available for scheduling. This field is beta.
*VolumeNodeResources is a set of resource limits for scheduling of volumes.*
  - **drivers.allocatable.count** (int32)
    Maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is not specified, then the supported number of volumes on this node is unbounded.
  - **drivers.topologyKeys** ([]string)
    topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set of topology keys that it understands (e.g. "company.com/zone", "company.com/region"). When a driver is initialized on a node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology keys. This can be empty if driver does not support topology.
### CSINodeList
CSINodeList is a collection of CSINode objects.
- **apiVersion**: storage.k8s.io/v1
  
- **kind**: CSINodeList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>), required
  items is the list of CSINode
### Operations
#### `get` read the specified CSINode

##### HTTP Request
GET /apis/storage.k8s.io/v1/csinodes/{name}

##### Parameters
  - **{name}** (string), required
    name of the CSINode
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind CSINode

##### HTTP Request
GET /apis/storage.k8s.io/v1/csinodes

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
200 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinodelist" >}}">CSINodeList</a>): OK
401: Unauthorized
#### `create` create a CSINode

##### HTTP Request
POST /apis/storage.k8s.io/v1/csinodes

##### Parameters
  - **body** (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): Created
202 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): Accepted
401: Unauthorized
#### `update` replace the specified CSINode

##### HTTP Request
PUT /apis/storage.k8s.io/v1/csinodes/{name}

##### Parameters
  - **{name}** (string), required
    name of the CSINode
  - **body** (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): Created
401: Unauthorized
#### `patch` partially update the specified CSINode

##### HTTP Request
PATCH /apis/storage.k8s.io/v1/csinodes/{name}

##### Parameters
  - **{name}** (string), required
    name of the CSINode
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
200 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): OK
401: Unauthorized
#### `delete` delete a CSINode

##### HTTP Request
DELETE /apis/storage.k8s.io/v1/csinodes/{name}

##### Parameters
  - **{name}** (string), required
    name of the CSINode
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
200 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): OK
202 (<a href="{{< ref "/docs/config-and-storage/csinode-v1#csinode" >}}">CSINode</a>): Accepted
401: Unauthorized
#### `deletecollection` delete collection of CSINode

##### HTTP Request
DELETE /apis/storage.k8s.io/v1/csinodes

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
