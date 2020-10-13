---
collapsible: false
weight: 6
title: "ClusterRoleBinding"
description: "ClusterRoleBinding references a ClusterRole, but not contain it."
draft: false
---
## ClusterRoleBinding
`apiVersion: rbac.authorization.k8s.io/v1`
`import "k8s.io/api/rbac/v1"`
### ClusterRoleBinding
ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.
- **apiVersion**: rbac.authorization.k8s.io/v1
  
- **kind**: ClusterRoleBinding
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata.
- **roleRef** (RoleRef), required
  RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
*RoleRef contains information that points to the role being used*
  - **roleRef.apiGroup** (string), required
    APIGroup is the group for the resource being referenced
  - **roleRef.kind** (string), required
    Kind is the type of resource being referenced
  - **roleRef.name** (string), required
    Name is the name of resource being referenced
- **subjects** ([]Subject)
  Subjects holds references to the objects the role applies to.
*Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.*
  - **subjects.kind** (string), required
    Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
  - **subjects.name** (string), required
    Name of the object being referenced.
  - **subjects.apiGroup** (string)
    APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
  - **subjects.namespace** (string)
    Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
### ClusterRoleBindingList
ClusterRoleBindingList is a collection of ClusterRoleBindings
- **apiVersion**: rbac.authorization.k8s.io/v1
  
- **kind**: ClusterRoleBindingList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard object's metadata.
- **items** ([]<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>), required
  Items is a list of ClusterRoleBindings
### Operations
#### `get` read the specified ClusterRoleBinding

##### HTTP Request
GET /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name}

##### Parameters
  - **{name}** (string), required
    name of the ClusterRoleBinding
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind ClusterRoleBinding

##### HTTP Request
GET /apis/rbac.authorization.k8s.io/v1/clusterrolebindings

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
200 (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebindinglist" >}}">ClusterRoleBindingList</a>): OK
401: Unauthorized
#### `create` create a ClusterRoleBinding

##### HTTP Request
POST /apis/rbac.authorization.k8s.io/v1/clusterrolebindings

##### Parameters
  - **body** (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>): OK
201 (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>): Created
202 (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>): Accepted
401: Unauthorized
#### `update` replace the specified ClusterRoleBinding

##### HTTP Request
PUT /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name}

##### Parameters
  - **{name}** (string), required
    name of the ClusterRoleBinding
  - **body** (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>): OK
201 (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>): Created
401: Unauthorized
#### `patch` partially update the specified ClusterRoleBinding

##### HTTP Request
PATCH /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name}

##### Parameters
  - **{name}** (string), required
    name of the ClusterRoleBinding
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
200 (<a href="{{< ref "/docs/authorization/clusterrolebinding-v1#clusterrolebinding" >}}">ClusterRoleBinding</a>): OK
401: Unauthorized
#### `delete` delete a ClusterRoleBinding

##### HTTP Request
DELETE /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name}

##### Parameters
  - **{name}** (string), required
    name of the ClusterRoleBinding
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
#### `deletecollection` delete collection of ClusterRoleBinding

##### HTTP Request
DELETE /apis/rbac.authorization.k8s.io/v1/clusterrolebindings

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
