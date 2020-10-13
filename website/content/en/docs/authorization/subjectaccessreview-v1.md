---
weight: 4
title: "SubjectAccessReview"
description: "SubjectAccessReview checks whether or not a user or group can perform an action."
draft: false
collapsible: false
---
## SubjectAccessReview
`apiVersion: authorization.k8s.io/v1`
`import "k8s.io/api/authorization/v1"`
### SubjectAccessReview
SubjectAccessReview checks whether or not a user or group can perform an action.
- **apiVersion**: authorization.k8s.io/v1
  
- **kind**: SubjectAccessReview
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreviewspec" >}}">SubjectAccessReviewSpec</a>), required
  Spec holds information about the request being evaluated
- **status** (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreviewstatus" >}}">SubjectAccessReviewStatus</a>)
  Status is filled in by the server and indicates whether the request is allowed or not
### SubjectAccessReviewSpec
SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
- **extra** (map[string][]string)
  Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
- **groups** ([]string)
  Groups is the groups you're testing for.
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
- **uid** (string)
  UID information about the requesting user.
- **user** (string)
  User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
### SubjectAccessReviewStatus
SubjectAccessReviewStatus
- **allowed** (boolean), required
  Allowed is required. True if the action would be allowed, false otherwise.
- **denied** (boolean)
  Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
- **evaluationError** (string)
  EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
- **reason** (string)
  Reason is optional.  It indicates why a request was allowed or denied.
### Operations
#### `create` create a SubjectAccessReview

##### HTTP Request
POST /apis/authorization.k8s.io/v1/subjectaccessreviews

##### Parameters
  - **body** (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreview" >}}">SubjectAccessReview</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreview" >}}">SubjectAccessReview</a>): OK
201 (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreview" >}}">SubjectAccessReview</a>): Created
202 (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreview" >}}">SubjectAccessReview</a>): Accepted
401: Unauthorized
