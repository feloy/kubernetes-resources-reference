---
draft: false
collapsible: false
weight: 6
title: "StorageClass"
description: "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned."
---
## StorageClass
`apiVersion: storage.k8s.io/v1`
`import "k8s.io/api/storage/v1"`
### StorageClass
StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.

StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
- **apiVersion**: storage.k8s.io/v1
  
- **kind**: StorageClass
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **provisioner** (string), required
  Provisioner indicates the type of the provisioner.
- **allowVolumeExpansion** (boolean)
  AllowVolumeExpansion shows whether the storage class allow volume expand
- **allowedTopologies** ([]TopologySelectorTerm)
  Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
*A topology selector term represents the result of label queries. A null or empty topology selector term matches no objects. The requirements of them are ANDed. It provides a subset of functionality as NodeSelectorTerm. This is an alpha feature and may change in the future.*
  - **allowedTopologies.matchLabelExpressions** ([]TopologySelectorLabelRequirement)
    A list of topology selector requirements by labels.
*A topology selector requirement is a selector that matches given label. This is an alpha feature and may change in the future.*
  - **allowedTopologies.matchLabelExpressions.key** (string), required
    The label key that the selector applies to.
  - **allowedTopologies.matchLabelExpressions.values** ([]string), required
    An array of string values. One value must match the label to be selected. Each entry in Values is ORed.
- **mountOptions** ([]string)
  Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
- **parameters** (map[string]string)
  Parameters holds the parameters for the provisioner that should create volumes of this storage class.
- **reclaimPolicy** (string)
  Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
- **volumeBindingMode** (string)
  VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
### StorageClassList
StorageClassList is a collection of storage classes.
- **apiVersion**: storage.k8s.io/v1
  
- **kind**: StorageClassList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>), required
  Items is the list of StorageClasses
### Operations
#### `get` read the specified StorageClass

##### HTTP Request
GET /apis/storage.k8s.io/v1/storageclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the StorageClass
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind StorageClass

##### HTTP Request
GET /apis/storage.k8s.io/v1/storageclasses

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
200 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclasslist" >}}">StorageClassList</a>): OK
401: Unauthorized
#### `create` create a StorageClass

##### HTTP Request
POST /apis/storage.k8s.io/v1/storageclasses

##### Parameters
  - **body** (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): Created
202 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): Accepted
401: Unauthorized
#### `update` replace the specified StorageClass

##### HTTP Request
PUT /apis/storage.k8s.io/v1/storageclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the StorageClass
  - **body** (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): Created
401: Unauthorized
#### `patch` partially update the specified StorageClass

##### HTTP Request
PATCH /apis/storage.k8s.io/v1/storageclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the StorageClass
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
200 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): OK
401: Unauthorized
#### `delete` delete a StorageClass

##### HTTP Request
DELETE /apis/storage.k8s.io/v1/storageclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the StorageClass
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
200 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): OK
202 (<a href="{{< ref "/docs/config-and-storage/storageclass-v1#storageclass" >}}">StorageClass</a>): Accepted
401: Unauthorized
#### `deletecollection` delete collection of StorageClass

##### HTTP Request
DELETE /apis/storage.k8s.io/v1/storageclasses

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
