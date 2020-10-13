---
title: "PriorityLevelConfiguration v1alpha1"
description: "PriorityLevelConfiguration represents the configuration of a priority level."
draft: false
collapsible: false
weight: 8
---
## PriorityLevelConfiguration
`apiVersion: flowcontrol.apiserver.k8s.io/v1alpha1`
`import "k8s.io/api/flowcontrol/v1alpha1"`
### PriorityLevelConfiguration
PriorityLevelConfiguration represents the configuration of a priority level.
- **apiVersion**: flowcontrol.apiserver.k8s.io/v1alpha1
  
- **kind**: PriorityLevelConfiguration
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfigurationspec" >}}">PriorityLevelConfigurationSpec</a>)
  `spec` is the specification of the desired behavior of a "request-priority". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfigurationstatus" >}}">PriorityLevelConfigurationStatus</a>)
  `status` is the current status of a "request-priority". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### PriorityLevelConfigurationSpec
PriorityLevelConfigurationSpec specifies the configuration of a priority level.
- **type** (string), required
  `type` indicates whether this priority level is subject to limitation on request execution.  A value of `"Exempt"` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `"Limited"` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.
- **limited** (LimitedPriorityLevelConfiguration)
  `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `"Limited"`.
*LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:*
* * How are requests for this priority level limited?*
* * What should be done with requests that exceed the limit?*
  - **limited.assuredConcurrencyShares** (int32)
    `assuredConcurrencyShares` (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.  ACS must be a positive number. The server's concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) --- the number of requests that may be executing at a time --- for each such priority level:
    
                ACV(l) = ceil( SCL * ACS(l) / ( sum[priority levels k] ACS(k) ) )
    
    bigger numbers of ACS mean more reserved concurrent requests (at the expense of every other PL). This field has a default value of 30.
  - **limited.limitResponse** (LimitResponse)
    `limitResponse` indicates what to do with requests that can not be executed right now
*LimitResponse defines how to handle requests that can not be executed right now.*
  - **limited.limitResponse.type** (string), required
    `type` is "Queue" or "Reject". "Queue" means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. "Reject" means that requests that can not be executed upon arrival are rejected. Required.
  - **limited.limitResponse.queuing** (QueuingConfiguration)
    `queuing` holds the configuration parameters for queuing. This field may be non-empty only if `type` is `"Queue"`.
*QueuingConfiguration holds the configuration parameters for queuing*
  - **limited.limitResponse.queuing.handSize** (int32)
    `handSize` is a small positive number that configures the shuffle sharding of requests into queues.  When enqueuing a request at this priority level the request's flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here.  The request is put into one of the shortest queues in that hand. `handSize` must be no larger than `queues`, and should be significantly smaller (so that a few heavy flows do not saturate most of the queues).  See the user-facing documentation for more extensive guidance on setting this field.  This field has a default value of 8.
  - **limited.limitResponse.queuing.queueLengthLimit** (int32)
    `queueLengthLimit` is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.
  - **limited.limitResponse.queuing.queues** (int32)
    `queues` is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.
### PriorityLevelConfigurationStatus
PriorityLevelConfigurationStatus represents the current state of a "request-priority".
- **conditions** ([]PriorityLevelConfigurationCondition)
  *Map: unique values on key type will be kept during a merge*
  `conditions` is the current state of "request-priority".
*PriorityLevelConfigurationCondition defines the condition of priority level.*
  - **conditions.lastTransitionTime** (Time)
    `lastTransitionTime` is the last time the condition transitioned from one status to another.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.message** (string)
    `message` is a human-readable message indicating details about last transition.
  - **conditions.reason** (string)
    `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
  - **conditions.status** (string)
    `status` is the status of the condition. Can be True, False, Unknown. Required.
  - **conditions.type** (string)
    `type` is the type of the condition. Required.
### PriorityLevelConfigurationList
PriorityLevelConfigurationList is a list of PriorityLevelConfiguration objects.
- **apiVersion**: flowcontrol.apiserver.k8s.io/v1alpha1
  
- **kind**: PriorityLevelConfigurationList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>), required
  `items` is a list of request-priorities.
### Operations
#### `get` read the specified PriorityLevelConfiguration

##### HTTP Request
GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}

##### Parameters
  - **{name}** (string), required
    name of the PriorityLevelConfiguration
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): OK
401: Unauthorized
#### `get` read status of the specified PriorityLevelConfiguration

##### HTTP Request
GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PriorityLevelConfiguration
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PriorityLevelConfiguration

##### HTTP Request
GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations

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
200 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfigurationlist" >}}">PriorityLevelConfigurationList</a>): OK
401: Unauthorized
#### `create` create a PriorityLevelConfiguration

##### HTTP Request
POST /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations

##### Parameters
  - **body** (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): OK
201 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): Created
202 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): Accepted
401: Unauthorized
#### `update` replace the specified PriorityLevelConfiguration

##### HTTP Request
PUT /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}

##### Parameters
  - **{name}** (string), required
    name of the PriorityLevelConfiguration
  - **body** (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): OK
201 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): Created
401: Unauthorized
#### `update` replace status of the specified PriorityLevelConfiguration

##### HTTP Request
PUT /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PriorityLevelConfiguration
  - **body** (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): OK
201 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): Created
401: Unauthorized
#### `patch` partially update the specified PriorityLevelConfiguration

##### HTTP Request
PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}

##### Parameters
  - **{name}** (string), required
    name of the PriorityLevelConfiguration
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
200 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified PriorityLevelConfiguration

##### HTTP Request
PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the PriorityLevelConfiguration
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
200 (<a href="{{< ref "/docs/cluster/prioritylevelconfiguration-v1alpha1#prioritylevelconfiguration" >}}">PriorityLevelConfiguration</a>): OK
401: Unauthorized
#### `delete` delete a PriorityLevelConfiguration

##### HTTP Request
DELETE /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}

##### Parameters
  - **{name}** (string), required
    name of the PriorityLevelConfiguration
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
#### `deletecollection` delete collection of PriorityLevelConfiguration

##### HTTP Request
DELETE /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations

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
