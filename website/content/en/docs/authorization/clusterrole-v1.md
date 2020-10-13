---
title: "ClusterRole"
description: "ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding."
draft: false
collapsible: false
weight: 5
---
## ClusterRole
`apiVersion: rbac.authorization.k8s.io/v1`
`import "k8s.io/api/rbac/v1"`
### ClusterRole
ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.
- **apiVersion**: rbac.authorization.k8s.io/v1
  
- **kind**: ClusterRole
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata.
- **aggregationRule** (AggregationRule)
  AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
*AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole*
  - **aggregationRule.clusterRoleSelectors** ([]<a href="{{< ref "/docs/common-definitions/labelselector-#labelselector" >}}">LabelSelector</a>)
    ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
- **rules** ([]PolicyRule)
  Rules holds all the PolicyRules for this ClusterRole
*PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.*
  - **rules.apiGroups** ([]string)
    APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
  - **rules.resources** ([]string)
    Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
  - **rules.verbs** ([]string), required
    Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
  - **rules.resourceNames** ([]string)
    ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
  - **rules.nonResourceURLs** ([]string)
    NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
### ClusterRoleList
ClusterRoleList is a collection of ClusterRoles
- **apiVersion**: rbac.authorization.k8s.io/v1
  
- **kind**: ClusterRoleList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard object's metadata.
- **items** ([]<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>), required
  Items is a list of ClusterRoles
### Operations
#### `get` read the specified ClusterRole

##### HTTP Request
GET /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}

##### Parameters
  - **{name}** (string), required
    name of the ClusterRole
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind ClusterRole

##### HTTP Request
GET /apis/rbac.authorization.k8s.io/v1/clusterroles

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
200 (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrolelist" >}}">ClusterRoleList</a>): OK
401: Unauthorized
#### `create` create a ClusterRole

##### HTTP Request
POST /apis/rbac.authorization.k8s.io/v1/clusterroles

##### Parameters
  - **body** (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>): OK
201 (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>): Created
202 (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>): Accepted
401: Unauthorized
#### `update` replace the specified ClusterRole

##### HTTP Request
PUT /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}

##### Parameters
  - **{name}** (string), required
    name of the ClusterRole
  - **body** (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>): OK
201 (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>): Created
401: Unauthorized
#### `patch` partially update the specified ClusterRole

##### HTTP Request
PATCH /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}

##### Parameters
  - **{name}** (string), required
    name of the ClusterRole
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
200 (<a href="{{< ref "/docs/authorization/clusterrole-v1#clusterrole" >}}">ClusterRole</a>): OK
401: Unauthorized
#### `delete` delete a ClusterRole

##### HTTP Request
DELETE /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}

##### Parameters
  - **{name}** (string), required
    name of the ClusterRole
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
#### `deletecollection` delete collection of ClusterRole

##### HTTP Request
DELETE /apis/rbac.authorization.k8s.io/v1/clusterroles

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
