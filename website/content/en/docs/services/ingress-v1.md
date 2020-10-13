---
title: "Ingress"
description: "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend."
draft: false
collapsible: false
weight: 4
---
## Ingress
`apiVersion: networking.k8s.io/v1`
`import "k8s.io/api/networking/v1"`
### Ingress
Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.
- **apiVersion**: networking.k8s.io/v1
  
- **kind**: Ingress
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/services/ingress-v1#ingressspec" >}}">IngressSpec</a>)
  Spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/services/ingress-v1#ingressstatus" >}}">IngressStatus</a>)
  Status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### IngressSpec
IngressSpec describes the Ingress the user wishes to exist.
- **defaultBackend** (IngressBackend)
  DefaultBackend is the backend that should handle requests that don't match any rule. If Rules are not specified, DefaultBackend must be specified. If DefaultBackend is not set, the handling of requests that do not match any of the rules will be up to the Ingress controller.
*IngressBackend describes all endpoints for a given service and port.*
  - **defaultBackend.resource** (<a href="{{< ref "/docs/common-definitions/typedlocalobjectreference-#typedlocalobjectreference" >}}">TypedLocalObjectReference</a>)
    Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with "Service".
  - **defaultBackend.service** (IngressServiceBackend)
    Service references a Service as a Backend. This is a mutually exclusive setting with "Resource".
*IngressServiceBackend references a Kubernetes Service as a Backend.*
  - **defaultBackend.service.name** (string), required
    Name is the referenced service. The service must exist in the same namespace as the Ingress object.
  - **defaultBackend.service.port** (ServiceBackendPort)
    Port of the referenced service. A port name or port number is required for a IngressServiceBackend.
*ServiceBackendPort is the service port being referenced.*
  - **defaultBackend.service.port.name** (string)
    Name is the name of the port on the Service. This is a mutually exclusive setting with "Number".
  - **defaultBackend.service.port.number** (int32)
    Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with "Name".
- **ingressClassName** (string)
  IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.
- **rules** ([]IngressRule)
  *Atomic: will be replaced during a merge*
  A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
*IngressRule represents the rules mapping the paths under a specified host to the related backend services. Incoming requests are first evaluated for a host match, then routed to the backend associated with the matching IngressRuleValue.*
  - **rules.host** (string)
    Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the "host" part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to
       the IP in the Spec of the parent Ingress.
    2. The `:` delimiter is not respected because ports are not allowed.
    	  Currently the port of an Ingress is implicitly :80 for http and
    	  :443 for https.
    Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.
    
    Host can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.bar.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. "*.foo.com"). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == "*"). Requests will be matched against the Host field in the following way: 1. If Host is precise, the request matches this rule if the http host header is equal to Host. 2. If Host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.
  - **rules.http** (HTTPIngressRuleValue)
    
*HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.*
  - **rules.http.paths** ([]HTTPIngressPath), required
    *Atomic: will be replaced during a merge*
    A collection of paths that map requests to backends.
*HTTPIngressPath associates a path with a backend. Incoming urls matching the path are forwarded to the backend.*
  - **rules.http.paths.backend** (IngressBackend), required
    Backend defines the referenced service endpoint to which the traffic will be forwarded to.
*IngressBackend describes all endpoints for a given service and port.*
  - **rules.http.paths.backend.resource** (<a href="{{< ref "/docs/common-definitions/typedlocalobjectreference-#typedlocalobjectreference" >}}">TypedLocalObjectReference</a>)
    Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with "Service".
  - **rules.http.paths.backend.service** (IngressServiceBackend)
    Service references a Service as a Backend. This is a mutually exclusive setting with "Resource".
*IngressServiceBackend references a Kubernetes Service as a Backend.*
  - **rules.http.paths.backend.service.name** (string), required
    Name is the referenced service. The service must exist in the same namespace as the Ingress object.
  - **rules.http.paths.backend.service.port** (ServiceBackendPort)
    Port of the referenced service. A port name or port number is required for a IngressServiceBackend.
*ServiceBackendPort is the service port being referenced.*
  - **rules.http.paths.backend.service.port.name** (string)
    Name is the name of the port on the Service. This is a mutually exclusive setting with "Number".
  - **rules.http.paths.backend.service.port.number** (int32)
    Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with "Name".
  - **rules.http.paths.path** (string)
    Path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. When unspecified, all paths from incoming requests are matched.
  - **rules.http.paths.pathType** (string)
    PathType determines the interpretation of the Path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is
      done on a path element by element basis. A path element refers is the
      list of labels in the path split by the '/' separator. A request is a
      match for path p if every p is an element-wise prefix of p of the
      request path. Note that if the last element of the path is a substring
      of the last element in request path, it is not a match (e.g. /foo/bar
      matches /foo/bar/baz, but does not match /foo/barbaz).
    * ImplementationSpecific: Interpretation of the Path matching is up to
      the IngressClass. Implementations can treat this as a separate PathType
      or treat it identically to Prefix or Exact path types.
    Implementations are required to support all path types.
- **tls** ([]IngressTLS)
  *Atomic: will be replaced during a merge*
  TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
*IngressTLS describes the transport layer security associated with an Ingress.*
  - **tls.hosts** ([]string)
    *Atomic: will be replaced during a merge*
    Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
  - **tls.secretName** (string)
    SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
### IngressStatus
IngressStatus describe the current state of the Ingress.
- **loadBalancer** (LoadBalancerStatus)
  LoadBalancer contains the current status of the load-balancer.
*LoadBalancerStatus represents the status of a load-balancer.*
  - **loadBalancer.ingress** ([]LoadBalancerIngress)
    Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
*LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.*
  - **loadBalancer.ingress.hostname** (string)
    Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
  - **loadBalancer.ingress.ip** (string)
    IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
### IngressList
IngressList is a collection of Ingress.
- **apiVersion**: networking.k8s.io/v1
  
- **kind**: IngressList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>), required
  Items is the list of Ingress.
### Operations
#### `get` read the specified Ingress

##### HTTP Request
GET /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}

##### Parameters
  - **{name}** (string), required
    name of the Ingress
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): OK
401: Unauthorized
#### `get` read status of the specified Ingress

##### HTTP Request
GET /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Ingress
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind Ingress

##### HTTP Request
GET /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses

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
200 (<a href="{{< ref "/docs/services/ingress-v1#ingresslist" >}}">IngressList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind Ingress

##### HTTP Request
GET /apis/networking.k8s.io/v1/ingresses

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
200 (<a href="{{< ref "/docs/services/ingress-v1#ingresslist" >}}">IngressList</a>): OK
401: Unauthorized
#### `create` create an Ingress

##### HTTP Request
POST /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): OK
201 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): Created
202 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): Accepted
401: Unauthorized
#### `update` replace the specified Ingress

##### HTTP Request
PUT /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}

##### Parameters
  - **{name}** (string), required
    name of the Ingress
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): OK
201 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): Created
401: Unauthorized
#### `update` replace status of the specified Ingress

##### HTTP Request
PUT /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Ingress
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): OK
201 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): Created
401: Unauthorized
#### `patch` partially update the specified Ingress

##### HTTP Request
PATCH /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}

##### Parameters
  - **{name}** (string), required
    name of the Ingress
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
200 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified Ingress

##### HTTP Request
PATCH /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Ingress
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
200 (<a href="{{< ref "/docs/services/ingress-v1#ingress" >}}">Ingress</a>): OK
401: Unauthorized
#### `delete` delete an Ingress

##### HTTP Request
DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}

##### Parameters
  - **{name}** (string), required
    name of the Ingress
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
#### `deletecollection` delete collection of Ingress

##### HTTP Request
DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses

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
