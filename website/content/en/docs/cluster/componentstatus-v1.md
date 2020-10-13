---
description: "ComponentStatus (and ComponentStatusList) holds the cluster validation info."
draft: false
collapsible: false
weight: 10
title: "ComponentStatus"
---
## ComponentStatus
`apiVersion: v1`
`import "k8s.io/api/core/v1"`
### ComponentStatus
ComponentStatus (and ComponentStatusList) holds the cluster validation info. Deprecated: This API is deprecated in v1.19+
- **apiVersion**: v1
  
- **kind**: ComponentStatus
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **conditions** ([]ComponentCondition)
  *Patch strategy: merge on key `type`*
  List of component conditions observed
*Information about the condition of a component.*
  - **conditions.status** (string), required
    Status of the condition for a component. Valid values for "Healthy": "True", "False", or "Unknown".
  - **conditions.type** (string), required
    Type of condition for a component. Valid value: "Healthy"
  - **conditions.error** (string)
    Condition error code for a component. For example, a health check error code.
  - **conditions.message** (string)
    Message about the condition for a component. For example, information about a health check.
### ComponentStatusList
Status of all the conditions for the component as a list of ComponentStatus objects. Deprecated: This API is deprecated in v1.19+
- **apiVersion**: v1
  
- **kind**: ComponentStatusList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- **items** ([]<a href="{{< ref "/docs/cluster/componentstatus-v1#componentstatus" >}}">ComponentStatus</a>), required
  List of ComponentStatus objects.
### Operations
#### `get` read the specified ComponentStatus

##### HTTP Request
GET /api/v1/componentstatuses/{name}

##### Parameters
  - **{name}** (string), required
    name of the ComponentStatus
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/componentstatus-v1#componentstatus" >}}">ComponentStatus</a>): OK
401: Unauthorized
#### `list` list objects of kind ComponentStatus

##### HTTP Request
GET /api/v1/componentstatuses

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
200 (<a href="{{< ref "/docs/cluster/componentstatus-v1#componentstatuslist" >}}">ComponentStatusList</a>): OK
401: Unauthorized
