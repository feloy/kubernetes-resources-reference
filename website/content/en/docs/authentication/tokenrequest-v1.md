---
draft: false
collapsible: false
weight: 2
title: "TokenRequest"
description: "TokenRequest requests a token for a given service account."
---
## TokenRequest
`apiVersion: authentication.k8s.io/v1`
`import "k8s.io/api/authentication/v1"`
### TokenRequest
TokenRequest requests a token for a given service account.
- **apiVersion**: authentication.k8s.io/v1
  
- **kind**: TokenRequest
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/authentication/tokenrequest-v1#tokenrequestspec" >}}">TokenRequestSpec</a>), required
  
- **status** (<a href="{{< ref "/docs/authentication/tokenrequest-v1#tokenrequeststatus" >}}">TokenRequestStatus</a>)
  
### TokenRequestSpec
TokenRequestSpec contains client provided parameters of a token request.
- **audiences** ([]string), required
  Audiences are the intendend audiences of the token. A recipient of a token must identitfy themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
- **boundObjectRef** (BoundObjectReference)
  BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other audiences may not. Keep ExpirationSeconds small if you want prompt revocation.
*BoundObjectReference is a reference to an object that a token is bound to.*
  - **boundObjectRef.apiVersion** (string)
    API version of the referent.
  - **boundObjectRef.kind** (string)
    Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
  - **boundObjectRef.name** (string)
    Name of the referent.
  - **boundObjectRef.uid** (string)
    UID of the referent.
- **expirationSeconds** (int64)
  ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
### TokenRequestStatus
TokenRequestStatus is the result of a token request.
- **expirationTimestamp** (Time), required
  ExpirationTimestamp is the time of expiration of the returned token.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
- **token** (string), required
  Token is the opaque bearer token.
### Operations
#### `create` create token of a ServiceAccount

##### HTTP Request
POST /api/v1/namespaces/{namespace}/serviceaccounts/{name}/token

##### Parameters
  - **{name}** (string), required
    name of the TokenRequest
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/authentication/tokenrequest-v1#tokenrequest" >}}">TokenRequest</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authentication/tokenrequest-v1#tokenrequest" >}}">TokenRequest</a>): OK
201 (<a href="{{< ref "/docs/authentication/tokenrequest-v1#tokenrequest" >}}">TokenRequest</a>): Created
202 (<a href="{{< ref "/docs/authentication/tokenrequest-v1#tokenrequest" >}}">TokenRequest</a>): Accepted
401: Unauthorized
