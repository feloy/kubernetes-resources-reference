---
title: "SelfSubjectAccessReview"
description: "SelfSubjectAccessReview checks whether or the current user can perform an action."
draft: false
collapsible: false
weight: 2
---
## SelfSubjectAccessReview
`apiVersion: authorization.k8s.io/v1`
`import "k8s.io/api/authorization/v1"`
### SelfSubjectAccessReview
SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means "in all namespaces".  Self is a special case, because users should always be able to check whether they can perform an action
- **apiVersion**: authorization.k8s.io/v1
  
- **kind**: SelfSubjectAccessReview
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/authorization/selfsubjectaccessreview-v1#selfsubjectaccessreviewspec" >}}">SelfSubjectAccessReviewSpec</a>), required
  Spec holds information about the request being evaluated.  user and groups must be empty
- **status** (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreviewstatus" >}}">SubjectAccessReviewStatus</a>)
  Status is filled in by the server and indicates whether the request is allowed or not
### SelfSubjectAccessReviewSpec
SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
- **nonResourceAttributes** (NonResourceAttributes)
  NonResourceAttributes describes information for a non-resource access request
*NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface*
  - **nonResourceAttributes.path** (string)
    Path is the URL path of the request
  - **nonResourceAttributes.verb** (string)
    Verb is the standard HTTP verb
- **resourceAttributes** (ResourceAttributes)
  ResourceAuthorizationAttributes describes information for a resource access request
*ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface*
  - **resourceAttributes.group** (string)
    Group is the API Group of the Resource.  "*" means all.
  - **resourceAttributes.name** (string)
    Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
  - **resourceAttributes.namespace** (string)
    Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
  - **resourceAttributes.resource** (string)
    Resource is one of the existing resource types.  "*" means all.
  - **resourceAttributes.subresource** (string)
    Subresource is one of the existing resource types.  "" means none.
  - **resourceAttributes.verb** (string)
    Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
  - **resourceAttributes.version** (string)
    Version is the API Version of the Resource.  "*" means all.
### Operations
#### `create` create a SelfSubjectAccessReview

##### HTTP Request
POST /apis/authorization.k8s.io/v1/selfsubjectaccessreviews

##### Parameters
  - **body** (<a href="{{< ref "/docs/authorization/selfsubjectaccessreview-v1#selfsubjectaccessreview" >}}">SelfSubjectAccessReview</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/selfsubjectaccessreview-v1#selfsubjectaccessreview" >}}">SelfSubjectAccessReview</a>): OK
201 (<a href="{{< ref "/docs/authorization/selfsubjectaccessreview-v1#selfsubjectaccessreview" >}}">SelfSubjectAccessReview</a>): Created
202 (<a href="{{< ref "/docs/authorization/selfsubjectaccessreview-v1#selfsubjectaccessreview" >}}">SelfSubjectAccessReview</a>): Accepted
401: Unauthorized
