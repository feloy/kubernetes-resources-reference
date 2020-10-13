---
title: "IngressClass"
description: "IngressClass represents the class of the Ingress, referenced by the Ingress Spec."
draft: false
collapsible: false
weight: 5
---
## IngressClass
`apiVersion: networking.k8s.io/v1`
`import "k8s.io/api/networking/v1"`
### IngressClass
IngressClass represents the class of the Ingress, referenced by the Ingress Spec. The `ingressclass.kubernetes.io/is-default-class` annotation can be used to indicate that an IngressClass should be considered default. When a single IngressClass resource has this annotation set to true, new Ingress resources without a class specified will be assigned this default class.
- **apiVersion**: networking.k8s.io/v1
  
- **kind**: IngressClass
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclassspec" >}}">IngressClassSpec</a>)
  Spec is the desired state of the IngressClass. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### IngressClassSpec
IngressClassSpec provides information about the class of an Ingress.
- **controller** (string)
  Controller refers to the name of the controller that should handle this class. This allows for different "flavors" that are controlled by the same controller. For example, you may have different Parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
- **parameters** (<a href="{{< ref "/docs/common-definitions/typedlocalobjectreference-#typedlocalobjectreference" >}}">TypedLocalObjectReference</a>)
  Parameters is a link to a custom resource containing additional configuration for the controller. This is optional if the controller does not require extra parameters.
### IngressClassList
IngressClassList is a collection of IngressClasses.
- **apiVersion**: networking.k8s.io/v1
  
- **kind**: IngressClassList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata.
- **items** ([]<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>), required
  Items is the list of IngressClasses.
### Operations
#### `get` read the specified IngressClass

##### HTTP Request
GET /apis/networking.k8s.io/v1/ingressclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the IngressClass
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind IngressClass

##### HTTP Request
GET /apis/networking.k8s.io/v1/ingressclasses

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
200 (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclasslist" >}}">IngressClassList</a>): OK
401: Unauthorized
#### `create` create an IngressClass

##### HTTP Request
POST /apis/networking.k8s.io/v1/ingressclasses

##### Parameters
  - **body** (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>): OK
201 (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>): Created
202 (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>): Accepted
401: Unauthorized
#### `update` replace the specified IngressClass

##### HTTP Request
PUT /apis/networking.k8s.io/v1/ingressclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the IngressClass
  - **body** (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>): OK
201 (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>): Created
401: Unauthorized
#### `patch` partially update the specified IngressClass

##### HTTP Request
PATCH /apis/networking.k8s.io/v1/ingressclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the IngressClass
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
200 (<a href="{{< ref "/docs/services/ingressclass-v1#ingressclass" >}}">IngressClass</a>): OK
401: Unauthorized
#### `delete` delete an IngressClass

##### HTTP Request
DELETE /apis/networking.k8s.io/v1/ingressclasses/{name}

##### Parameters
  - **{name}** (string), required
    name of the IngressClass
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
#### `deletecollection` delete collection of IngressClass

##### HTTP Request
DELETE /apis/networking.k8s.io/v1/ingressclasses

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
