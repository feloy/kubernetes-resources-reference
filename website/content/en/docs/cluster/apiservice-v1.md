---
title: "APIService"
description: "APIService represents a server for a particular GroupVersion."
draft: false
collapsible: false
weight: 4
---
## APIService
`apiVersion: apiregistration.k8s.io/v1`
`import "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"`
### APIService
APIService represents a server for a particular GroupVersion. Name must be "version.group".
- **apiVersion**: apiregistration.k8s.io/v1
  
- **kind**: APIService
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservicespec" >}}">APIServiceSpec</a>)
  Spec contains information for locating and communicating with a server
- **status** (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservicestatus" >}}">APIServiceStatus</a>)
  Status contains derived information about an API server
### APIServiceSpec
APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.
- **groupPriorityMinimum** (int32), required
  GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s
- **versionPriority** (int32), required
  VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
- **caBundle** ([]byte)
  *Atomic: will be replaced during a merge*
  CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.
- **group** (string)
  Group is the API group name this server hosts
- **insecureSkipTLSVerify** (boolean)
  InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged.  You should use the CABundle instead.
- **service** (ServiceReference)
  Service is a reference to the service for this API server.  It must communicate on port 443. If the Service is nil, that means the handling for the API groupversion is handled locally on this server. The call will simply delegate to the normal handler chain to be fulfilled.
*ServiceReference holds a reference to Service.legacy.k8s.io*
  - **service.name** (string)
    Name is the name of the service
  - **service.namespace** (string)
    Namespace is the namespace of the service
  - **service.port** (int32)
    If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
- **version** (string)
  Version is the API version this server hosts.  For example, "v1"
### APIServiceStatus
APIServiceStatus contains derived information about an API server
- **conditions** ([]APIServiceCondition)
  *Patch strategy: merge on key `type`*
  *Map: unique values on key type will be kept during a merge*
  Current service state of apiService.
*APIServiceCondition describes the state of an APIService at a particular point*
  - **conditions.status** (string), required
    Status is the status of the condition. Can be True, False, Unknown.
  - **conditions.type** (string), required
    Type is the type of the condition.
  - **conditions.lastTransitionTime** (Time)
    Last time the condition transitioned from one status to another.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.message** (string)
    Human-readable message indicating details about last transition.
  - **conditions.reason** (string)
    Unique, one-word, CamelCase reason for the condition's last transition.
### APIServiceList
APIServiceList is a list of APIService objects.
- **apiVersion**: apiregistration.k8s.io/v1
  
- **kind**: APIServiceList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  
- **items** ([]<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>), required
  
### Operations
#### `get` read the specified APIService

##### HTTP Request
GET /apis/apiregistration.k8s.io/v1/apiservices/{name}

##### Parameters
  - **{name}** (string), required
    name of the APIService
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): OK
401: Unauthorized
#### `get` read status of the specified APIService

##### HTTP Request
GET /apis/apiregistration.k8s.io/v1/apiservices/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the APIService
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind APIService

##### HTTP Request
GET /apis/apiregistration.k8s.io/v1/apiservices

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
200 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservicelist" >}}">APIServiceList</a>): OK
401: Unauthorized
#### `create` create an APIService

##### HTTP Request
POST /apis/apiregistration.k8s.io/v1/apiservices

##### Parameters
  - **body** (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): OK
201 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): Created
202 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): Accepted
401: Unauthorized
#### `update` replace the specified APIService

##### HTTP Request
PUT /apis/apiregistration.k8s.io/v1/apiservices/{name}

##### Parameters
  - **{name}** (string), required
    name of the APIService
  - **body** (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): OK
201 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): Created
401: Unauthorized
#### `update` replace status of the specified APIService

##### HTTP Request
PUT /apis/apiregistration.k8s.io/v1/apiservices/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the APIService
  - **body** (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): OK
201 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): Created
401: Unauthorized
#### `patch` partially update the specified APIService

##### HTTP Request
PATCH /apis/apiregistration.k8s.io/v1/apiservices/{name}

##### Parameters
  - **{name}** (string), required
    name of the APIService
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
200 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified APIService

##### HTTP Request
PATCH /apis/apiregistration.k8s.io/v1/apiservices/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the APIService
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
200 (<a href="{{< ref "/docs/cluster/apiservice-v1#apiservice" >}}">APIService</a>): OK
401: Unauthorized
#### `delete` delete an APIService

##### HTTP Request
DELETE /apis/apiregistration.k8s.io/v1/apiservices/{name}

##### Parameters
  - **{name}** (string), required
    name of the APIService
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
#### `deletecollection` delete collection of APIService

##### HTTP Request
DELETE /apis/apiregistration.k8s.io/v1/apiservices

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
