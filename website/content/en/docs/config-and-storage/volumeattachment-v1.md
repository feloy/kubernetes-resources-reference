---
description: "VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node."
draft: false
collapsible: false
weight: 7
title: "VolumeAttachment"
---
## VolumeAttachment
`apiVersion: storage.k8s.io/v1`
`import "k8s.io/api/storage/v1"`
### VolumeAttachment
VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.

VolumeAttachment objects are non-namespaced.
- **apiVersion**: storage.k8s.io/v1
  
- **kind**: VolumeAttachment
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachmentspec" >}}">VolumeAttachmentSpec</a>), required
  Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
- **status** (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachmentstatus" >}}">VolumeAttachmentStatus</a>)
  Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
### VolumeAttachmentSpec
VolumeAttachmentSpec is the specification of a VolumeAttachment request.
- **attacher** (string), required
  Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
- **nodeName** (string), required
  The node that the volume should be attached to.
- **source** (VolumeAttachmentSource), required
  Source represents the volume that should be attached.
*VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.*
  - **source.inlineVolumeSpec** (<a href="{{< ref "/docs/config-and-storage/persistentvolume-v1#persistentvolumespec" >}}">PersistentVolumeSpec</a>)
    inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
  - **source.persistentVolumeName** (string)
    Name of the persistent volume to attach.
### VolumeAttachmentStatus
VolumeAttachmentStatus is the status of a VolumeAttachment request.
- **attached** (boolean), required
  Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
- **attachError** (VolumeError)
  The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
*VolumeError captures an error encountered during a volume operation.*
  - **attachError.message** (string)
    String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
  - **attachError.time** (Time)
    Time the error was encountered.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
- **attachmentMetadata** (map[string]string)
  Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
- **detachError** (VolumeError)
  The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
*VolumeError captures an error encountered during a volume operation.*
  - **detachError.message** (string)
    String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
  - **detachError.time** (Time)
    Time the error was encountered.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
### VolumeAttachmentList
VolumeAttachmentList is a collection of VolumeAttachment objects.
- **apiVersion**: storage.k8s.io/v1
  
- **kind**: VolumeAttachmentList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>), required
  Items is the list of VolumeAttachments
### Operations
#### `get` read the specified VolumeAttachment

##### HTTP Request
GET /apis/storage.k8s.io/v1/volumeattachments/{name}

##### Parameters
  - **{name}** (string), required
    name of the VolumeAttachment
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): OK
401: Unauthorized
#### `get` read status of the specified VolumeAttachment

##### HTTP Request
GET /apis/storage.k8s.io/v1/volumeattachments/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the VolumeAttachment
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind VolumeAttachment

##### HTTP Request
GET /apis/storage.k8s.io/v1/volumeattachments

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
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachmentlist" >}}">VolumeAttachmentList</a>): OK
401: Unauthorized
#### `create` create a VolumeAttachment

##### HTTP Request
POST /apis/storage.k8s.io/v1/volumeattachments

##### Parameters
  - **body** (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): Created
202 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): Accepted
401: Unauthorized
#### `update` replace the specified VolumeAttachment

##### HTTP Request
PUT /apis/storage.k8s.io/v1/volumeattachments/{name}

##### Parameters
  - **{name}** (string), required
    name of the VolumeAttachment
  - **body** (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): Created
401: Unauthorized
#### `update` replace status of the specified VolumeAttachment

##### HTTP Request
PUT /apis/storage.k8s.io/v1/volumeattachments/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the VolumeAttachment
  - **body** (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): OK
201 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): Created
401: Unauthorized
#### `patch` partially update the specified VolumeAttachment

##### HTTP Request
PATCH /apis/storage.k8s.io/v1/volumeattachments/{name}

##### Parameters
  - **{name}** (string), required
    name of the VolumeAttachment
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
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified VolumeAttachment

##### HTTP Request
PATCH /apis/storage.k8s.io/v1/volumeattachments/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the VolumeAttachment
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
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): OK
401: Unauthorized
#### `delete` delete a VolumeAttachment

##### HTTP Request
DELETE /apis/storage.k8s.io/v1/volumeattachments/{name}

##### Parameters
  - **{name}** (string), required
    name of the VolumeAttachment
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
200 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): OK
202 (<a href="{{< ref "/docs/config-and-storage/volumeattachment-v1#volumeattachment" >}}">VolumeAttachment</a>): Accepted
401: Unauthorized
#### `deletecollection` delete collection of VolumeAttachment

##### HTTP Request
DELETE /apis/storage.k8s.io/v1/volumeattachments

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
