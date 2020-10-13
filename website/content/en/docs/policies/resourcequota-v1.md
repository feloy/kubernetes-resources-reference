---
title: "ResourceQuota"
description: "ResourceQuota sets aggregate quota restrictions enforced per namespace."
draft: false
collapsible: false
weight: 2
---
## ResourceQuota
`apiVersion: v1`
`import "k8s.io/api/core/v1"`
### ResourceQuota
ResourceQuota sets aggregate quota restrictions enforced per namespace
- **apiVersion**: v1
  
- **kind**: ResourceQuota
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequotaspec" >}}">ResourceQuotaSpec</a>)
  Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequotastatus" >}}">ResourceQuotaStatus</a>)
  Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### ResourceQuotaSpec
ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
- **hard** (map[string]<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
  hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
- **scopeSelector** (ScopeSelector)
  scopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector (if specified in spec), must be matched.
*A scope selector represents the AND of the selectors represented by the scoped-resource selector requirements.*
  - **scopeSelector.matchExpressions** ([]ScopedResourceSelectorRequirement)
    A list of scope selector requirements by scope of the resources.
*A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator that relates the scope name and values.*
  - **scopeSelector.matchExpressions.operator** (string), required
    Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
  - **scopeSelector.matchExpressions.scopeName** (string), required
    The name of the scope that the selector applies to.
  - **scopeSelector.matchExpressions.values** ([]string)
    An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
- **scopes** ([]string)
  A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
### ResourceQuotaStatus
ResourceQuotaStatus defines the enforced hard limits and observed use.
- **hard** (map[string]<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
  Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
- **used** (map[string]<a href="{{< ref "/docs/common-definitions/quantity-#quantity" >}}">Quantity</a>)
  Used is the current observed total usage of the resource in the namespace.
### ResourceQuotaList
ResourceQuotaList is a list of ResourceQuota items.
- **apiVersion**: v1
  
- **kind**: ResourceQuotaList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- **items** ([]<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>), required
  Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
### Operations
#### `get` read the specified ResourceQuota

##### HTTP Request
GET /api/v1/namespaces/{namespace}/resourcequotas/{name}

##### Parameters
  - **{name}** (string), required
    name of the ResourceQuota
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): OK
401: Unauthorized
#### `get` read status of the specified ResourceQuota

##### HTTP Request
GET /api/v1/namespaces/{namespace}/resourcequotas/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the ResourceQuota
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind ResourceQuota

##### HTTP Request
GET /api/v1/namespaces/{namespace}/resourcequotas

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
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequotalist" >}}">ResourceQuotaList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind ResourceQuota

##### HTTP Request
GET /api/v1/resourcequotas

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
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequotalist" >}}">ResourceQuotaList</a>): OK
401: Unauthorized
#### `create` create a ResourceQuota

##### HTTP Request
POST /api/v1/namespaces/{namespace}/resourcequotas

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): OK
201 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): Created
202 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): Accepted
401: Unauthorized
#### `update` replace the specified ResourceQuota

##### HTTP Request
PUT /api/v1/namespaces/{namespace}/resourcequotas/{name}

##### Parameters
  - **{name}** (string), required
    name of the ResourceQuota
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): OK
201 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): Created
401: Unauthorized
#### `update` replace status of the specified ResourceQuota

##### HTTP Request
PUT /api/v1/namespaces/{namespace}/resourcequotas/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the ResourceQuota
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): OK
201 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): Created
401: Unauthorized
#### `patch` partially update the specified ResourceQuota

##### HTTP Request
PATCH /api/v1/namespaces/{namespace}/resourcequotas/{name}

##### Parameters
  - **{name}** (string), required
    name of the ResourceQuota
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
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified ResourceQuota

##### HTTP Request
PATCH /api/v1/namespaces/{namespace}/resourcequotas/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the ResourceQuota
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
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): OK
401: Unauthorized
#### `delete` delete a ResourceQuota

##### HTTP Request
DELETE /api/v1/namespaces/{namespace}/resourcequotas/{name}

##### Parameters
  - **{name}** (string), required
    name of the ResourceQuota
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
200 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): OK
202 (<a href="{{< ref "/docs/policies/resourcequota-v1#resourcequota" >}}">ResourceQuota</a>): Accepted
401: Unauthorized
#### `deletecollection` delete collection of ResourceQuota

##### HTTP Request
DELETE /api/v1/namespaces/{namespace}/resourcequotas

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
