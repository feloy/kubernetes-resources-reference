---
collapsible: false
weight: 3
title: "ValidatingWebhookConfiguration"
description: "ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it."
draft: false
---
## ValidatingWebhookConfiguration
`apiVersion: admissionregistration.k8s.io/v1`
`import "k8s.io/api/admissionregistration/v1"`
### ValidatingWebhookConfiguration
ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.
- **apiVersion**: admissionregistration.k8s.io/v1
  
- **kind**: ValidatingWebhookConfiguration
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
- **webhooks** ([]ValidatingWebhook)
  *Patch strategy: merge on key `name`*
  Webhooks is a list of webhooks and the affected resources and operations.
*ValidatingWebhook describes an admission webhook and the resources and operations it applies to.*
  - **webhooks.admissionReviewVersions** ([]string), required
    AdmissionReviewVersions is an ordered list of preferred `AdmissionReview` versions the Webhook expects. API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail and be subject to the failure policy.
  - **webhooks.clientConfig** (WebhookClientConfig), required
    ClientConfig defines how to communicate with the hook. Required
*WebhookClientConfig contains the information to make a TLS connection with the webhook*
  - **webhooks.clientConfig.caBundle** ([]byte)
    `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
  - **webhooks.clientConfig.service** (ServiceReference)
    `service` is a reference to the service for this webhook. Either `service` or `url` must be specified.
    
    If the webhook is running within the cluster, then you should use `service`.
*ServiceReference holds a reference to Service.legacy.k8s.io*
  - **webhooks.clientConfig.service.name** (string), required
    `name` is the name of the service. Required
  - **webhooks.clientConfig.service.namespace** (string), required
    `namespace` is the namespace of the service. Required
  - **webhooks.clientConfig.service.path** (string)
    `path` is an optional URL path which will be sent in any request to this service.
  - **webhooks.clientConfig.service.port** (int32)
    If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
  - **webhooks.clientConfig.url** (string)
    `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.
    
    The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.
    
    Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.
    
    The scheme must be "https"; the URL must begin with "https://".
    
    A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.
    
    Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
  - **webhooks.name** (string), required
    The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where "imagepolicy" is the name of the webhook, and kubernetes.io is the name of the organization. Required.
  - **webhooks.sideEffects** (string), required
    SideEffects states whether this webhook has side effects. Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown). Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission change and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects == Unknown or Some.
  - **webhooks.failurePolicy** (string)
    FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Fail.
  - **webhooks.matchPolicy** (string)
    matchPolicy defines how the "rules" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent".
    
    - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.
    
    - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook.
    
    Defaults to "Equivalent"
  - **webhooks.namespaceSelector** (<a href="{{< ref "/docs/common-definitions/labelselector-#labelselector" >}}">LabelSelector</a>)
    NamespaceSelector decides whether to run the webhook on an object based on whether the namespace for that object matches the selector. If the object itself is a namespace, the matching is performed on object.metadata.labels. If the object is another cluster scoped resource, it never skips the webhook.
    
    For example, to run the webhook on any objects whose namespace is not associated with "runlevel" of "0" or "1";  you will set the selector as follows: "namespaceSelector": {
      "matchExpressions": [
        {
          "key": "runlevel",
          "operator": "NotIn",
          "values": [
            "0",
            "1"
          ]
        }
      ]
    }
    
    If instead you want to only run the webhook on any objects whose namespace is associated with the "environment" of "prod" or "staging"; you will set the selector as follows: "namespaceSelector": {
      "matchExpressions": [
        {
          "key": "environment",
          "operator": "In",
          "values": [
            "prod",
            "staging"
          ]
        }
      ]
    }
    
    See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels for more examples of label selectors.
    
    Default to the empty LabelSelector, which matches everything.
  - **webhooks.objectSelector** (<a href="{{< ref "/docs/common-definitions/labelselector-#labelselector" >}}">LabelSelector</a>)
    ObjectSelector decides whether to run the webhook based on if the object has matching labels. objectSelector is evaluated against both the oldObject and newObject that would be sent to the webhook, and is considered to match if either object matches the selector. A null object (oldObject in the case of create, or newObject in the case of delete) or an object that cannot have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match. Use the object selector only if the webhook is opt-in, because end users may skip the admission webhook by setting the labels. Default to the empty LabelSelector, which matches everything.
  - **webhooks.rules** ([]RuleWithOperations)
    Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
*RuleWithOperations is a tuple of Operations and Resources. It is recommended to make sure that all the tuple expansions are valid.*
  - **webhooks.rules.apiGroups** ([]string)
    APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
  - **webhooks.rules.apiVersions** ([]string)
    APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
  - **webhooks.rules.operations** ([]string)
    Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.
  - **webhooks.rules.resources** ([]string)
    Resources is a list of resources this rule applies to.
    
    For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources.
    
    If wildcard is present, the validation rule will ensure resources do not overlap with each other.
    
    Depending on the enclosing object, subresources might not be allowed. Required.
  - **webhooks.rules.scope** (string)
    scope specifies the scope of this rule. Valid values are "Cluster", "Namespaced", and "*" "Cluster" means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. "Namespaced" means that only namespaced resources will match this rule. "*" means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is "*".
  - **webhooks.timeoutSeconds** (int32)
    TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.
### ValidatingWebhookConfigurationList
ValidatingWebhookConfigurationList is a list of ValidatingWebhookConfiguration.
- **apiVersion**: admissionregistration.k8s.io/v1
  
- **kind**: ValidatingWebhookConfigurationList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- **items** ([]<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>), required
  List of ValidatingWebhookConfiguration.
### Operations
#### `get` read the specified ValidatingWebhookConfiguration

##### HTTP Request
GET /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name}

##### Parameters
  - **{name}** (string), required
    name of the ValidatingWebhookConfiguration
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind ValidatingWebhookConfiguration

##### HTTP Request
GET /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations

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
200 (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfigurationlist" >}}">ValidatingWebhookConfigurationList</a>): OK
401: Unauthorized
#### `create` create a ValidatingWebhookConfiguration

##### HTTP Request
POST /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations

##### Parameters
  - **body** (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>): OK
201 (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>): Created
202 (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>): Accepted
401: Unauthorized
#### `update` replace the specified ValidatingWebhookConfiguration

##### HTTP Request
PUT /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name}

##### Parameters
  - **{name}** (string), required
    name of the ValidatingWebhookConfiguration
  - **body** (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>): OK
201 (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>): Created
401: Unauthorized
#### `patch` partially update the specified ValidatingWebhookConfiguration

##### HTTP Request
PATCH /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name}

##### Parameters
  - **{name}** (string), required
    name of the ValidatingWebhookConfiguration
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
200 (<a href="{{< ref "/docs/extend/validatingwebhookconfiguration-v1#validatingwebhookconfiguration" >}}">ValidatingWebhookConfiguration</a>): OK
401: Unauthorized
#### `delete` delete a ValidatingWebhookConfiguration

##### HTTP Request
DELETE /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name}

##### Parameters
  - **{name}** (string), required
    name of the ValidatingWebhookConfiguration
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
#### `deletecollection` delete collection of ValidatingWebhookConfiguration

##### HTTP Request
DELETE /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations

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
