---
description: "LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace."
draft: false
collapsible: false
weight: 1
title: "LocalSubjectAccessReview"
---
## LocalSubjectAccessReview
`apiVersion: authorization.k8s.io/v1`
`import "k8s.io/api/authorization/v1"`
### LocalSubjectAccessReview
LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace. Having a namespace scoped resource makes it much easier to grant namespace scoped policy that includes permissions checking.
- **apiVersion**: authorization.k8s.io/v1
  
- **kind**: LocalSubjectAccessReview
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreviewspec" >}}">SubjectAccessReviewSpec</a>), required
  Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
- **status** (<a href="{{< ref "/docs/authorization/subjectaccessreview-v1#subjectaccessreviewstatus" >}}">SubjectAccessReviewStatus</a>)
  Status is filled in by the server and indicates whether the request is allowed or not
### Operations
#### `create` create a LocalSubjectAccessReview

##### HTTP Request
POST /apis/authorization.k8s.io/v1/namespaces/{namespace}/localsubjectaccessreviews

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/authorization/localsubjectaccessreview-v1#localsubjectaccessreview" >}}">LocalSubjectAccessReview</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authorization/localsubjectaccessreview-v1#localsubjectaccessreview" >}}">LocalSubjectAccessReview</a>): OK
201 (<a href="{{< ref "/docs/authorization/localsubjectaccessreview-v1#localsubjectaccessreview" >}}">LocalSubjectAccessReview</a>): Created
202 (<a href="{{< ref "/docs/authorization/localsubjectaccessreview-v1#localsubjectaccessreview" >}}">LocalSubjectAccessReview</a>): Accepted
401: Unauthorized
