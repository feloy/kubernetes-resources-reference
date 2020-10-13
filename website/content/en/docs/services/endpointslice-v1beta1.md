---
title: "EndpointSlice v1beta1"
description: "EndpointSlice represents a subset of the endpoints that implement a service."
draft: false
collapsible: false
weight: 3
---
## EndpointSlice
`apiVersion: discovery.k8s.io/v1beta1`
`import "k8s.io/api/discovery/v1beta1"`
### EndpointSlice
EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
- **apiVersion**: discovery.k8s.io/v1beta1
  
- **kind**: EndpointSlice
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata.
- **addressType** (string), required
  addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
- **endpoints** ([]Endpoint), required
  *Atomic: will be replaced during a merge*
  endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
*Endpoint represents a single logical "backend" implementing a service.*
  - **endpoints.addresses** ([]string), required
    *Set: unique values will be kept during a merge*
    addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100.
  - **endpoints.conditions** (EndpointConditions)
    conditions contains information about the current status of the endpoint.
*EndpointConditions represents the current condition of an endpoint.*
  - **endpoints.conditions.ready** (boolean)
    ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready.
  - **endpoints.hostname** (string)
    hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must pass DNS Label (RFC 1123) validation.
  - **endpoints.targetRef** (<a href="{{< ref "/docs/common-definitions/objectreference-#objectreference" >}}">ObjectReference</a>)
    targetRef is a reference to a Kubernetes object that represents this endpoint.
  - **endpoints.topology** (map[string]string)
    topology contains arbitrary topology information associated with the endpoint. These key/value pairs must conform with the label format. https://kubernetes.io/docs/concepts/overview/working-with-objects/labels Topology may include a maximum of 16 key/value pairs. This includes, but is not limited to the following well known keys: * kubernetes.io/hostname: the value indicates the hostname of the node
      where the endpoint is located. This should match the corresponding
      node label.
    * topology.kubernetes.io/zone: the value indicates the zone where the
      endpoint is located. This should match the corresponding node label.
    * topology.kubernetes.io/region: the value indicates the region where the
      endpoint is located. This should match the corresponding node label.
- **ports** ([]EndpointPort)
  *Atomic: will be replaced during a merge*
  ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
*EndpointPort represents a Port used by an EndpointSlice*
  - **ports.port** (int32)
    The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.
  - **ports.protocol** (string)
    The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
  - **ports.name** (string)
    The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
  - **ports.appProtocol** (string)
    The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
### EndpointSliceList
EndpointSliceList represents a list of endpoint slices
- **apiVersion**: discovery.k8s.io/v1beta1
  
- **kind**: EndpointSliceList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata.
- **items** ([]<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>), required
  List of endpoint slices
### Operations
#### `get` read the specified EndpointSlice

##### HTTP Request
GET /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}

##### Parameters
  - **{name}** (string), required
    name of the EndpointSlice
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind EndpointSlice

##### HTTP Request
GET /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices

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
200 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslicelist" >}}">EndpointSliceList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind EndpointSlice

##### HTTP Request
GET /apis/discovery.k8s.io/v1beta1/endpointslices

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
200 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslicelist" >}}">EndpointSliceList</a>): OK
401: Unauthorized
#### `create` create an EndpointSlice

##### HTTP Request
POST /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>): OK
201 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>): Created
202 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>): Accepted
401: Unauthorized
#### `update` replace the specified EndpointSlice

##### HTTP Request
PUT /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}

##### Parameters
  - **{name}** (string), required
    name of the EndpointSlice
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>): OK
201 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>): Created
401: Unauthorized
#### `patch` partially update the specified EndpointSlice

##### HTTP Request
PATCH /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}

##### Parameters
  - **{name}** (string), required
    name of the EndpointSlice
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
200 (<a href="{{< ref "/docs/services/endpointslice-v1beta1#endpointslice" >}}">EndpointSlice</a>): OK
401: Unauthorized
#### `delete` delete an EndpointSlice

##### HTTP Request
DELETE /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}

##### Parameters
  - **{name}** (string), required
    name of the EndpointSlice
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
200 (<a href="{{< ref "/docs/common-definitions/status-#status" >}}">Status</a>): OK
202 (<a href="{{< ref "/docs/common-definitions/status-#status" >}}">Status</a>): Accepted
401: Unauthorized
#### `deletecollection` delete collection of EndpointSlice

##### HTTP Request
DELETE /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices

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
