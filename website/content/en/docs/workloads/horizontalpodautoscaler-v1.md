---
title: "HorizontalPodAutoscaler"
description: "configuration of a horizontal pod autoscaler."
draft: false
collapsible: false
weight: 14
---
## HorizontalPodAutoscaler
`apiVersion: autoscaling/v1`
`import "k8s.io/api/autoscaling/v1"`
### HorizontalPodAutoscaler
configuration of a horizontal pod autoscaler.
- **apiVersion**: autoscaling/v1
  
- **kind**: HorizontalPodAutoscaler
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscalerspec" >}}">HorizontalPodAutoscalerSpec</a>)
  behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
- **status** (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscalerstatus" >}}">HorizontalPodAutoscalerStatus</a>)
  current information about the autoscaler.
### HorizontalPodAutoscalerSpec
specification of a horizontal pod autoscaler.
- **maxReplicas** (int32), required
  upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
- **scaleTargetRef** (CrossVersionObjectReference), required
  reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
*CrossVersionObjectReference contains enough information to let you identify the referred resource.*
  - **scaleTargetRef.kind** (string), required
    Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
  - **scaleTargetRef.name** (string), required
    Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
  - **scaleTargetRef.apiVersion** (string)
    API version of the referent
- **minReplicas** (int32)
  minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
- **targetCPUUtilizationPercentage** (int32)
  target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
### HorizontalPodAutoscalerStatus
current status of a horizontal pod autoscaler
- **currentReplicas** (int32), required
  current number of replicas of pods managed by this autoscaler.
- **desiredReplicas** (int32), required
  desired number of replicas of pods managed by this autoscaler.
- **currentCPUUtilizationPercentage** (int32)
  current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70%!o(MISSING)f its requested CPU.
- **lastScaleTime** (Time)
  last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
- **observedGeneration** (int64)
  most recent generation observed by this autoscaler.
### HorizontalPodAutoscalerList
list of horizontal pod autoscaler objects.
- **apiVersion**: autoscaling/v1
  
- **kind**: HorizontalPodAutoscalerList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata.
- **items** ([]<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>), required
  list of horizontal pod autoscaler objects.
### Operations
#### `get` read the specified HorizontalPodAutoscaler

##### HTTP Request
GET /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}

##### Parameters
  - **{name}** (string), required
    name of the HorizontalPodAutoscaler
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): OK
401: Unauthorized
#### `get` read status of the specified HorizontalPodAutoscaler

##### HTTP Request
GET /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the HorizontalPodAutoscaler
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind HorizontalPodAutoscaler

##### HTTP Request
GET /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers

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
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscalerlist" >}}">HorizontalPodAutoscalerList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind HorizontalPodAutoscaler

##### HTTP Request
GET /apis/autoscaling/v1/horizontalpodautoscalers

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
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscalerlist" >}}">HorizontalPodAutoscalerList</a>): OK
401: Unauthorized
#### `create` create a HorizontalPodAutoscaler

##### HTTP Request
POST /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): OK
201 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): Created
202 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): Accepted
401: Unauthorized
#### `update` replace the specified HorizontalPodAutoscaler

##### HTTP Request
PUT /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}

##### Parameters
  - **{name}** (string), required
    name of the HorizontalPodAutoscaler
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): OK
201 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): Created
401: Unauthorized
#### `update` replace status of the specified HorizontalPodAutoscaler

##### HTTP Request
PUT /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the HorizontalPodAutoscaler
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): OK
201 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): Created
401: Unauthorized
#### `patch` partially update the specified HorizontalPodAutoscaler

##### HTTP Request
PATCH /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}

##### Parameters
  - **{name}** (string), required
    name of the HorizontalPodAutoscaler
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
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified HorizontalPodAutoscaler

##### HTTP Request
PATCH /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the HorizontalPodAutoscaler
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
200 (<a href="{{< ref "/docs/workloads/horizontalpodautoscaler-v1#horizontalpodautoscaler" >}}">HorizontalPodAutoscaler</a>): OK
401: Unauthorized
#### `delete` delete a HorizontalPodAutoscaler

##### HTTP Request
DELETE /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}

##### Parameters
  - **{name}** (string), required
    name of the HorizontalPodAutoscaler
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
#### `deletecollection` delete collection of HorizontalPodAutoscaler

##### HTTP Request
DELETE /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers

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
