---
title: "FlowSchema v1alpha1"
description: "FlowSchema defines the schema of a group of flows."
draft: false
collapsible: false
weight: 7
---
## FlowSchema
`apiVersion: flowcontrol.apiserver.k8s.io/v1alpha1`
`import "k8s.io/api/flowcontrol/v1alpha1"`
### FlowSchema
FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a "flow distinguisher".
- **apiVersion**: flowcontrol.apiserver.k8s.io/v1alpha1
  
- **kind**: FlowSchema
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschemaspec" >}}">FlowSchemaSpec</a>)
  `spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschemastatus" >}}">FlowSchemaStatus</a>)
  `status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### FlowSchemaSpec
FlowSchemaSpec describes how the FlowSchema's specification looks like.
- **priorityLevelConfiguration** (PriorityLevelConfigurationReference), required
  `priorityLevelConfiguration` should reference a PriorityLevelConfiguration in the cluster. If the reference cannot be resolved, the FlowSchema will be ignored and marked as invalid in its status. Required.
*PriorityLevelConfigurationReference contains information that points to the "request-priority" being used.*
  - **priorityLevelConfiguration.name** (string), required
    `name` is the name of the priority level configuration being referenced Required.
- **distinguisherMethod** (FlowDistinguisherMethod)
  `distinguisherMethod` defines how to compute the flow distinguisher for requests that match this schema. `nil` specifies that the distinguisher is disabled and thus will always be the empty string.
*FlowDistinguisherMethod specifies the method of a flow distinguisher.*
  - **distinguisherMethod.type** (string), required
    `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace". Required.
- **matchingPrecedence** (int32)
  `matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.
- **rules** ([]PolicyRulesWithSubjects)
  *Atomic: will be replaced during a merge*
  `rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.
*PolicyRulesWithSubjects prescribes a test that applies to a request to an apiserver. The test considers the subject making the request, the verb being requested, and the resource to be acted upon. This PolicyRulesWithSubjects matches a request if and only if both (a) at least one member of subjects matches the request and (b) at least one member of resourceRules or nonResourceRules matches the request.*
  - **rules.subjects** ([]Subject), required
    *Atomic: will be replaced during a merge*
    subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.
*Subject matches the originator of a request, as identified by the request authentication system. There are three ways of matching an originator; by user, group, or service account.*
  - **rules.subjects.kind** (string), required
    Required
  - **rules.subjects.group** (GroupSubject)
    
*GroupSubject holds detailed information for group-kind subject.*
  - **rules.subjects.group.name** (string), required
    name is the user group that matches, or "*" to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.
  - **rules.subjects.serviceAccount** (ServiceAccountSubject)
    
*ServiceAccountSubject holds detailed information for service-account-kind subject.*
  - **rules.subjects.serviceAccount.name** (string), required
    `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
  - **rules.subjects.serviceAccount.namespace** (string), required
    `namespace` is the namespace of matching ServiceAccount objects. Required.
  - **rules.subjects.user** (UserSubject)
    
*UserSubject holds detailed information for user-kind subject.*
  - **rules.subjects.user.name** (string), required
    `name` is the username that matches, or "*" to match all usernames. Required.
  - **rules.nonResourceRules** ([]NonResourcePolicyRule)
    *Atomic: will be replaced during a merge*
    `nonResourceRules` is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.
*NonResourcePolicyRule is a predicate that matches non-resource requests according to their verb and the target non-resource URL. A NonResourcePolicyRule matches a request if and only if both (a) at least one member of verbs matches the request and (b) at least one member of nonResourceURLs matches the request.*
  - **rules.nonResourceRules.nonResourceURLs** ([]string), required
    *Set: unique values will be kept during a merge*
    `nonResourceURLs` is a set of url prefixes that a user should have access to and may not be empty. For example:
      - "/healthz" is legal
      - "/hea*" is illegal
      - "/hea" is legal but matches nothing
      - "/hea/*" also matches nothing
      - "/healthz/*" matches all per-component health checks.
    "*" matches all non-resource urls. if it is present, it must be the only entry. Required.
  - **rules.nonResourceRules.verbs** ([]string), required
    *Set: unique values will be kept during a merge*
    `verbs` is a list of matching verbs and may not be empty. "*" matches all verbs. If it is present, it must be the only entry. Required.
  - **rules.resourceRules** ([]ResourcePolicyRule)
    *Atomic: will be replaced during a merge*
    `resourceRules` is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of `resourceRules` and `nonResourceRules` has to be non-empty.
*ResourcePolicyRule is a predicate that matches some resource requests, testing the request's verb and the target resource. A ResourcePolicyRule matches a resource request if and only if: (a) at least one member of verbs matches the request, (b) at least one member of apiGroups matches the request, (c) at least one member of resources matches the request, and (d) least one member of namespaces matches the request.*
  - **rules.resourceRules.apiGroups** ([]string), required
    *Set: unique values will be kept during a merge*
    `apiGroups` is a list of matching API groups and may not be empty. "*" matches all API groups and, if present, must be the only entry. Required.
  - **rules.resourceRules.resources** ([]string), required
    *Set: unique values will be kept during a merge*
    `resources` is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ "services", "nodes/status" ].  This list may not be empty. "*" matches all resources and, if present, must be the only entry. Required.
  - **rules.resourceRules.verbs** ([]string), required
    *Set: unique values will be kept during a merge*
    `verbs` is a list of matching verbs and may not be empty. "*" matches all verbs and, if present, must be the only entry. Required.
  - **rules.resourceRules.clusterScope** (boolean)
    `clusterScope` indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the `namespaces` field must contain a non-empty list.
  - **rules.resourceRules.namespaces** ([]string)
    *Set: unique values will be kept during a merge*
    `namespaces` is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains "*".  Note that "*" matches any specified namespace but does not match a request that _does not specify_ a namespace (see the `clusterScope` field for that). This list may be empty, but only if `clusterScope` is true.
### FlowSchemaStatus
FlowSchemaStatus represents the current state of a FlowSchema.
- **conditions** ([]FlowSchemaCondition)
  *Map: unique values on key type will be kept during a merge*
  `conditions` is a list of the current states of FlowSchema.
*FlowSchemaCondition describes conditions for a FlowSchema.*
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
### FlowSchemaList
FlowSchemaList is a list of FlowSchema objects.
- **apiVersion**: flowcontrol.apiserver.k8s.io/v1alpha1
  
- **kind**: FlowSchemaList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>), required
  `items` is a list of FlowSchemas.
### Operations
#### `get` read the specified FlowSchema

##### HTTP Request
GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}

##### Parameters
  - **{name}** (string), required
    name of the FlowSchema
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): OK
401: Unauthorized
#### `get` read status of the specified FlowSchema

##### HTTP Request
GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the FlowSchema
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind FlowSchema

##### HTTP Request
GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas

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
200 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschemalist" >}}">FlowSchemaList</a>): OK
401: Unauthorized
#### `create` create a FlowSchema

##### HTTP Request
POST /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas

##### Parameters
  - **body** (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): OK
201 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): Created
202 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): Accepted
401: Unauthorized
#### `update` replace the specified FlowSchema

##### HTTP Request
PUT /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}

##### Parameters
  - **{name}** (string), required
    name of the FlowSchema
  - **body** (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): OK
201 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): Created
401: Unauthorized
#### `update` replace status of the specified FlowSchema

##### HTTP Request
PUT /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the FlowSchema
  - **body** (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): OK
201 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): Created
401: Unauthorized
#### `patch` partially update the specified FlowSchema

##### HTTP Request
PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}

##### Parameters
  - **{name}** (string), required
    name of the FlowSchema
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
200 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified FlowSchema

##### HTTP Request
PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the FlowSchema
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
200 (<a href="{{< ref "/docs/cluster/flowschema-v1alpha1#flowschema" >}}">FlowSchema</a>): OK
401: Unauthorized
#### `delete` delete a FlowSchema

##### HTTP Request
DELETE /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}

##### Parameters
  - **{name}** (string), required
    name of the FlowSchema
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
#### `deletecollection` delete collection of FlowSchema

##### HTTP Request
DELETE /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas

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
