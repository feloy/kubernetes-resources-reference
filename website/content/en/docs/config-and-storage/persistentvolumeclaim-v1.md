---
collapsible: false
weight: 4
title: "PersistentVolumeClaim"
description: "PersistentVolumeClaim is a user's request for and claim to a persistent volume."
draft: false
---
## PersistentVolumeClaim
`apiVersion: v1`
`import "k8s.io/api/core/v1"`
### PersistentVolumeClaim
PersistentVolumeClaim is a user's request for and claim to a persistent volume
- **apiVersion**: v1
  
- **kind**: PersistentVolumeClaim
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaimspec" >}}">PersistentVolumeClaimSpec</a>)
  Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
- **status** (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaimstatus" >}}">PersistentVolumeClaimStatus</a>)
  Status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
### PersistentVolumeClaimSpec
PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
- **accessModes** ([]string)
  AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
- **selector** (<a href="{{< ref "/docs/common-definitions/labelselector-#labelselector" >}}">LabelSelector</a>)
  A label query over volumes to consider for binding.
- **resources** (ResourceRequirements)
  Resources represents the minimum resources the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
*ResourceRequirements describes the compute resource requirements.*
  - **resources.limits** (map[string]<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
    Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
  - **resources.requests** (map[string]<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
    Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
- **volumeName** (string)
  VolumeName is the binding reference to the PersistentVolume backing this claim.
- **storageClassName** (string)
  Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
- **volumeMode** (string)
  volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
#### Alpha level

- **dataSource** (<a href="{{< ref "/docs/common-definitions/typedlocalobjectreference-#typedlocalobjectreference" >}}">TypedLocalObjectReference</a>)
  This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot - Beta) * An existing PVC (PersistentVolumeClaim) * An existing custom resource/object that implements data population (Alpha) In order to use VolumeSnapshot object types, the appropriate feature gate must be enabled (VolumeSnapshotDataSource or AnyVolumeDataSource) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the specified data source is not supported, the volume will not be created and the failure will be reported as an event. In the future, we plan to support more data source types and the behavior of the provisioner may change.
### PersistentVolumeClaimStatus
PersistentVolumeClaimStatus is the current status of a persistent volume claim.
- **accessModes** ([]string)
  AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
- **capacity** (map[string]<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
  Represents the actual resources of the underlying volume.
- **conditions** ([]PersistentVolumeClaimCondition)
  *Patch strategy: merge on key `type`*
  Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.
*PersistentVolumeClaimCondition contails details about state of pvc*
  - **conditions.status** (string), required
    
  - **conditions.type** (string), required
    
  - **conditions.lastProbeTime** (Time)
    Last time we probed the condition.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.lastTransitionTime** (Time)
    Last time the condition transitioned from one status to another.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.message** (string)
    Human-readable message indicating details about last transition.
  - **conditions.reason** (string)
    Unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports "ResizeStarted" that means the underlying persistent volume is being resized.
- **phase** (string)
  Phase represents the current phase of PersistentVolumeClaim.
### PersistentVolumeClaimList
PersistentVolumeClaimList is a list of PersistentVolumeClaim items.
- **apiVersion**: v1
  
- **kind**: PersistentVolumeClaimList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- **items** ([]<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>), required
  A list of persistent volume claims. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
### Operations
#### `get` read the specified PersistentVolumeClaim

##### HTTP Request
GET /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolumeClaim
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): OK
401: Unauthorized
#### `get` read status of the specified PersistentVolumeClaim

##### HTTP Request
GET /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolumeClaim
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PersistentVolumeClaim

##### HTTP Request
GET /api/v1/namespaces/{namespace}/persistentvolumeclaims

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
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaimlist" >}}">PersistentVolumeClaimList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PersistentVolumeClaim

##### HTTP Request
GET /api/v1/persistentvolumeclaims

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
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaimlist" >}}">PersistentVolumeClaimList</a>): OK
401: Unauthorized
#### `create` create a PersistentVolumeClaim

##### HTTP Request
POST /api/v1/namespaces/{namespace}/persistentvolumeclaims

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): Created
202 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): Accepted
401: Unauthorized
#### `update` replace the specified PersistentVolumeClaim

##### HTTP Request
PUT /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolumeClaim
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): Created
401: Unauthorized
#### `update` replace status of the specified PersistentVolumeClaim

##### HTTP Request
PUT /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolumeClaim
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): Created
401: Unauthorized
#### `patch` partially update the specified PersistentVolumeClaim

##### HTTP Request
PATCH /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolumeClaim
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
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified PersistentVolumeClaim

##### HTTP Request
PATCH /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolumeClaim
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
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): OK
401: Unauthorized
#### `delete` delete a PersistentVolumeClaim

##### HTTP Request
DELETE /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}

##### Parameters
  - **{name}** (string), required
    name of the PersistentVolumeClaim
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
200 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): OK
202 (<a href="{{< ref "/docs/config-and-storage/persistentvolumeclaim-v1#persistentvolumeclaim" >}}">PersistentVolumeClaim</a>): Accepted
401: Unauthorized
#### `deletecollection` delete collection of PersistentVolumeClaim

##### HTTP Request
DELETE /api/v1/namespaces/{namespace}/persistentvolumeclaims

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
