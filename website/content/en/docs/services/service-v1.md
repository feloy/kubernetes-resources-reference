---
collapsible: false
weight: 1
title: "Service"
description: "Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy."
draft: false
---
## Service
`apiVersion: v1`
`import "k8s.io/api/core/v1"`
### Service
Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.
- **apiVersion**: v1
  
- **kind**: Service
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/services/service-v1#servicespec" >}}">ServiceSpec</a>)
  Spec defines the behavior of a service. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/services/service-v1#servicestatus" >}}">ServiceStatus</a>)
  Most recently observed status of the service. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### ServiceSpec
ServiceSpec describes the attributes that a user creates on a service.
- **selector** (map[string]string)
  Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/
- **ports** ([]ServicePort)
  *Patch strategy: merge on key `port`*
  *Map: unique values on keys `port, protocol` will be kept during a merge*
  The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
*ServicePort contains information on service's port.*
  - **ports.port** (int32), required
    The port that will be exposed by this service.
  - **ports.targetPort** (IntOrString)
    Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
*IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.*
  - **ports.protocol** (string)
    The IP protocol for this port. Supports "TCP", "UDP", and "SCTP". Default is TCP.
  - **ports.name** (string)
    The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.
  - **ports.nodePort** (int32)
    The port on each node on which this service is exposed when type=NodePort or LoadBalancer. Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the ServiceType of this Service requires one. More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
  - **ports.appProtocol** (string)
    The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. This is a beta field that is guarded by the ServiceAppProtocol feature gate and enabled by default.
- **type** (string)
  type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. "ExternalName" maps to the specified externalName. "ClusterIP" allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object. If clusterIP is "None", no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a stable IP. "NodePort" builds on ClusterIP and allocates a port on every node which routes to the clusterIP. "LoadBalancer" builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the clusterIP. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
- **ipFamily** (string)
  ipFamily specifies whether this Service has a preference for a particular IP family (e.g. IPv4 vs. IPv6) when the IPv6DualStack feature gate is enabled. In a dual-stack cluster, you can specify ipFamily when creating a ClusterIP Service to determine whether the controller will allocate an IPv4 or IPv6 IP for it, and you can specify ipFamily when creating a headless Service to determine whether it will have IPv4 or IPv6 Endpoints. In either case, if you do not specify an ipFamily explicitly, it will default to the cluster's primary IP family. This field is part of an alpha feature, and you should not make any assumptions about its semantics other than those described above. In particular, you should not assume that it can (or cannot) be changed after creation time; that it can only have the values "IPv4" and "IPv6"; or that its current value on a given Service correctly reflects the current state of that Service. (For ClusterIP Services, look at clusterIP to see if the Service is IPv4 or IPv6. For headless Services, look at the endpoints, which may be dual-stack in the future. For ExternalName Services, ipFamily has no meaning, but it may be set to an irrelevant value anyway.)
- **clusterIP** (string)
  clusterIP is the IP address of the service and is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. This field can not be changed through updates. Valid values are "None", empty string (""), or a valid IP address. "None" can be specified for headless services when proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
- **externalIPs** ([]string)
  externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.
- **sessionAffinity** (string)
  Supports "ClientIP" and "None". Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
- **loadBalancerIP** (string)
  Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.
- **loadBalancerSourceRanges** ([]string)
  If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature." More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
- **externalName** (string)
  externalName is the external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName.
- **externalTrafficPolicy** (string)
  externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. "Local" preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. "Cluster" obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.
- **healthCheckNodePort** (int32)
  healthCheckNodePort specifies the healthcheck nodePort for the service. If not specified, HealthCheckNodePort is created by the service api backend with the allocated nodePort. Will use user-specified nodePort value if specified by the client. Only effects when Type is set to LoadBalancer and ExternalTrafficPolicy is set to Local.
- **publishNotReadyAddresses** (boolean)
  publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered "ready" even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior.
- **sessionAffinityConfig** (SessionAffinityConfig)
  sessionAffinityConfig contains the configurations of session affinity.
*SessionAffinityConfig represents the configurations of session affinity.*
  - **sessionAffinityConfig.clientIP** (ClientIPConfig)
    clientIP contains the configurations of Client IP based session affinity.
*ClientIPConfig represents the configurations of Client IP based session affinity.*
  - **sessionAffinityConfig.clientIP.timeoutSeconds** (int32)
    timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP". Default value is 10800(for 3 hours).
- **topologyKeys** ([]string)
  topologyKeys is a preference-order list of topology keys which implementations of services should use to preferentially sort endpoints when accessing this Service, it can not be used at the same time as externalTrafficPolicy=Local. Topology keys must be valid label keys and at most 16 keys may be specified. Endpoints are chosen based on the first topology key with available backends. If this field is specified and all entries have no backends that match the topology of the client, the service has no backends for that client and connections should fail. The special value "*" may be used to mean "any topology". This catch-all value, if used, only makes sense as the last value in the list. If this is not specified or empty, no topology constraints will be applied.
### ServiceStatus
ServiceStatus represents the current status of a service.
- **loadBalancer** (LoadBalancerStatus)
  LoadBalancer contains the current status of the load-balancer, if one is present.
*LoadBalancerStatus represents the status of a load-balancer.*
  - **loadBalancer.ingress** ([]LoadBalancerIngress)
    Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
*LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.*
  - **loadBalancer.ingress.hostname** (string)
    Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
  - **loadBalancer.ingress.ip** (string)
    IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
### ServiceList
ServiceList holds a list of services.
- **apiVersion**: v1
  
- **kind**: ServiceList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- **items** ([]<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>), required
  List of services
### Operations
#### `get` read the specified Service

##### HTTP Request
GET /api/v1/namespaces/{namespace}/services/{name}

##### Parameters
  - **{name}** (string), required
    name of the Service
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): OK
401: Unauthorized
#### `get` read status of the specified Service

##### HTTP Request
GET /api/v1/namespaces/{namespace}/services/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Service
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind Service

##### HTTP Request
GET /api/v1/namespaces/{namespace}/services

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
200 (<a href="{{< ref "/docs/services/service-v1#servicelist" >}}">ServiceList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind Service

##### HTTP Request
GET /api/v1/services

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
200 (<a href="{{< ref "/docs/services/service-v1#servicelist" >}}">ServiceList</a>): OK
401: Unauthorized
#### `create` create a Service

##### HTTP Request
POST /api/v1/namespaces/{namespace}/services

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): OK
201 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): Created
202 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): Accepted
401: Unauthorized
#### `update` replace the specified Service

##### HTTP Request
PUT /api/v1/namespaces/{namespace}/services/{name}

##### Parameters
  - **{name}** (string), required
    name of the Service
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): OK
201 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): Created
401: Unauthorized
#### `update` replace status of the specified Service

##### HTTP Request
PUT /api/v1/namespaces/{namespace}/services/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Service
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): OK
201 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): Created
401: Unauthorized
#### `patch` partially update the specified Service

##### HTTP Request
PATCH /api/v1/namespaces/{namespace}/services/{name}

##### Parameters
  - **{name}** (string), required
    name of the Service
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
200 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified Service

##### HTTP Request
PATCH /api/v1/namespaces/{namespace}/services/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the Service
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
200 (<a href="{{< ref "/docs/services/service-v1#service" >}}">Service</a>): OK
401: Unauthorized
#### `delete` delete a Service

##### HTTP Request
DELETE /api/v1/namespaces/{namespace}/services/{name}

##### Parameters
  - **{name}** (string), required
    name of the Service
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
