---
title: "CertificateSigningRequest"
description: "CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued."
draft: false
collapsible: false
weight: 4
---
## CertificateSigningRequest
`apiVersion: certificates.k8s.io/v1`
`import "k8s.io/api/certificates/v1"`
### CertificateSigningRequest
CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.

Kubelets use this API to obtain:
 1. client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client-kubelet" signerName).
 2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the "kubernetes.io/kubelet-serving" signerName).

This API can be used to request client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client" signerName), or to obtain certificates from custom non-Kubernetes signers.
- **apiVersion**: certificates.k8s.io/v1
  
- **kind**: CertificateSigningRequest
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequestspec" >}}">CertificateSigningRequestSpec</a>), required
  spec contains the certificate request, and is immutable after creation. Only the request, signerName, and usages fields can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.
- **status** (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequeststatus" >}}">CertificateSigningRequestStatus</a>)
  status contains information about whether the request is approved or denied, and the certificate issued by the signer, or the failure condition indicating signer failure.
### CertificateSigningRequestSpec
CertificateSigningRequestSpec contains the certificate request.
- **request** ([]byte), required
  *Atomic: will be replaced during a merge*
  request contains an x509 certificate signing request encoded in a "CERTIFICATE REQUEST" PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.
- **signerName** (string), required
  signerName indicates the requested signer, and is a qualified name.
  
  List/watch requests for CertificateSigningRequests can filter on this field using a "spec.signerName=NAME" fieldSelector.
  
  Well-known Kubernetes signers are:
   1. "kubernetes.io/kube-apiserver-client": issues client certificates that can be used to authenticate to kube-apiserver.
    Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the "csrsigning" controller in kube-controller-manager.
   2. "kubernetes.io/kube-apiserver-client-kubelet": issues client certificates that kubelets use to authenticate to kube-apiserver.
    Requests for this signer can be auto-approved by the "csrapproving" controller in kube-controller-manager, and can be issued by the "csrsigning" controller in kube-controller-manager.
   3. "kubernetes.io/kubelet-serving" issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.
    Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the "csrsigning" controller in kube-controller-manager.
  
  More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers
  
  Custom signerNames can also be specified. The signer defines:
   1. Trust distribution: how trust (CA bundles) are distributed.
   2. Permitted subjects: and behavior when a disallowed subject is requested.
   3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested.
   4. Required, permitted, or forbidden key usages / extended key usages.
   5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin.
   6. Whether or not requests for CA certificates are allowed.
- **extra** (map[string][]string)
  extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
- **groups** ([]string)
  *Atomic: will be replaced during a merge*
  groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
- **uid** (string)
  uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
- **usages** ([]string)
  *Atomic: will be replaced during a merge*
  usages specifies a set of key usages requested in the issued certificate.
  
  Requests for TLS client certificates typically request: "digital signature", "key encipherment", "client auth".
  
  Requests for TLS serving certificates typically request: "key encipherment", "digital signature", "server auth".
  
  Valid values are:
   "signing", "digital signature", "content commitment",
   "key encipherment", "key agreement", "data encipherment",
   "cert sign", "crl sign", "encipher only", "decipher only", "any",
   "server auth", "client auth",
   "code signing", "email protection", "s/mime",
   "ipsec end system", "ipsec tunnel", "ipsec user",
   "timestamping", "ocsp signing", "microsoft sgc", "netscape sgc"
- **username** (string)
  username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
### CertificateSigningRequestStatus
CertificateSigningRequestStatus contains conditions used to indicate approved/denied/failed status of the request, and the issued certificate.
- **certificate** ([]byte)
  *Atomic: will be replaced during a merge*
  certificate is populated with an issued certificate by the signer after an Approved condition is present. This field is set via the /status subresource. Once populated, this field is immutable.
  
  If the certificate signing request is denied, a condition of type "Denied" is added and this field remains empty. If the signer cannot issue the certificate, a condition of type "Failed" is added and this field remains empty.
  
  Validation requirements:
   1. certificate must contain one or more PEM blocks.
   2. All PEM blocks must have the "CERTIFICATE" label, contain no headers, and the encoded data
    must be a BER-encoded ASN.1 Certificate structure as described in section 4 of RFC5280.
   3. Non-PEM content may appear before or after the "CERTIFICATE" PEM blocks and is unvalidated,
    to allow for explanatory text as described in section 5.2 of RFC7468.
  
  If more than one PEM block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes.
  
  The certificate is encoded in PEM format.
  
  When serialized as JSON or YAML, the data is additionally base64-encoded, so it consists of:
  
      base64(
      -----BEGIN CERTIFICATE-----
      ...
      -----END CERTIFICATE-----
      )
- **conditions** ([]CertificateSigningRequestCondition)
  *Map: unique values on key type will be kept during a merge*
  conditions applied to the request. Known conditions are "Approved", "Denied", and "Failed".
*CertificateSigningRequestCondition describes a condition of a CertificateSigningRequest object*
  - **conditions.status** (string), required
    status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown".
  - **conditions.type** (string), required
    type of the condition. Known conditions are "Approved", "Denied", and "Failed".
    
    An "Approved" condition is added via the /approval subresource, indicating the request was approved and should be issued by the signer.
    
    A "Denied" condition is added via the /approval subresource, indicating the request was denied and should not be issued by the signer.
    
    A "Failed" condition is added via the /status subresource, indicating the signer failed to issue the certificate.
    
    Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once added.
    
    Only one condition of a given type is allowed.
  - **conditions.lastTransitionTime** (Time)
    lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.lastUpdateTime** (Time)
    lastUpdateTime is the time of the last update to this condition
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
  - **conditions.message** (string)
    message contains a human readable message with details about the request state
  - **conditions.reason** (string)
    reason indicates a brief reason for the request state
### CertificateSigningRequestList
CertificateSigningRequestList is a collection of CertificateSigningRequest objects
- **apiVersion**: certificates.k8s.io/v1
  
- **kind**: CertificateSigningRequestList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  
- **items** ([]<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>), required
  items is a collection of CertificateSigningRequest objects
### Operations
#### `get` read the specified CertificateSigningRequest

##### HTTP Request
GET /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
401: Unauthorized
#### `get` read approval of the specified CertificateSigningRequest

##### HTTP Request
GET /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}/approval

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
401: Unauthorized
#### `get` read status of the specified CertificateSigningRequest

##### HTTP Request
GET /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind CertificateSigningRequest

##### HTTP Request
GET /apis/certificates.k8s.io/v1/certificatesigningrequests

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
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequestlist" >}}">CertificateSigningRequestList</a>): OK
401: Unauthorized
#### `create` create a CertificateSigningRequest

##### HTTP Request
POST /apis/certificates.k8s.io/v1/certificatesigningrequests

##### Parameters
  - **body** (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
201 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): Created
202 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): Accepted
401: Unauthorized
#### `update` replace the specified CertificateSigningRequest

##### HTTP Request
PUT /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
  - **body** (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
201 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): Created
401: Unauthorized
#### `update` replace approval of the specified CertificateSigningRequest

##### HTTP Request
PUT /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}/approval

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
  - **body** (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
201 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): Created
401: Unauthorized
#### `update` replace status of the specified CertificateSigningRequest

##### HTTP Request
PUT /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
  - **body** (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
201 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): Created
401: Unauthorized
#### `patch` partially update the specified CertificateSigningRequest

##### HTTP Request
PATCH /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
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
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
401: Unauthorized
#### `patch` partially update approval of the specified CertificateSigningRequest

##### HTTP Request
PATCH /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}/approval

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
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
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified CertificateSigningRequest

##### HTTP Request
PATCH /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
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
200 (<a href="{{< ref "/docs/authentication/certificatesigningrequest-v1#certificatesigningrequest" >}}">CertificateSigningRequest</a>): OK
401: Unauthorized
#### `delete` delete a CertificateSigningRequest

##### HTTP Request
DELETE /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}

##### Parameters
  - **{name}** (string), required
    name of the CertificateSigningRequest
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
#### `deletecollection` delete collection of CertificateSigningRequest

##### HTTP Request
DELETE /apis/certificates.k8s.io/v1/certificatesigningrequests

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
