---
draft: false
collapsible: false
weight: 17
title: "PodPreset v1alpha1"
description: "PodPreset is a policy resource that defines additional runtime requirements for a Pod."
---
## PodPreset
`apiVersion: settings.k8s.io/v1alpha1`
`import "k8s.io/api/settings/v1alpha1"`
### PodPreset
PodPreset is a policy resource that defines additional runtime requirements for a Pod.
- **apiVersion**: settings.k8s.io/v1alpha1
  
- **kind**: PodPreset
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **spec** (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpresetspec" >}}">PodPresetSpec</a>)
  
### PodPresetSpec
PodPresetSpec is a description of a pod preset.
- **selector** (<a href="{{< ref "/docs/common-definitions/labelselector-#labelselector" >}}">LabelSelector</a>)
  Selector is a label query over a set of resources, in this case pods. Required.
- **volumes** ([]<a href="{{< ref "/docs/config-and-storage/volume-#volume" >}}">Volume</a>)
  Volumes defines the collection of Volume to inject into the pod.
- **volumeMounts** ([]VolumeMount)
  VolumeMounts defines the collection of VolumeMount to inject into containers.
*VolumeMount describes a mounting of a Volume within a container.*
  - **volumeMounts.mountPath** (string), required
    Path within the container at which the volume should be mounted.  Must not contain ':'.
  - **volumeMounts.name** (string), required
    This must match the Name of a Volume.
  - **volumeMounts.mountPropagation** (string)
    mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.
  - **volumeMounts.readOnly** (boolean)
    Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
  - **volumeMounts.subPath** (string)
    Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
  - **volumeMounts.subPathExpr** (string)
    Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to "" (volume's root). SubPathExpr and SubPath are mutually exclusive.
- **env** ([]EnvVar)
  Env defines the collection of EnvVar to inject into containers.
*EnvVar represents an environment variable present in a Container.*
  - **env.name** (string), required
    Name of the environment variable. Must be a C_IDENTIFIER.
  - **env.value** (string)
    Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
  - **env.valueFrom** (EnvVarSource)
    Source for the environment variable's value. Cannot be used if value is not empty.
*EnvVarSource represents a source for the value of an EnvVar.*
  - **env.valueFrom.configMapKeyRef** (ConfigMapKeySelector)
    Selects a key of a ConfigMap.
*Selects a key from a ConfigMap.*
  - **env.valueFrom.configMapKeyRef.key** (string), required
    The key to select.
  - **env.valueFrom.configMapKeyRef.name** (string)
    Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
  - **env.valueFrom.configMapKeyRef.optional** (boolean)
    Specify whether the ConfigMap or its key must be defined
  - **env.valueFrom.fieldRef** (<a href="{{< ref "/docs/common-definitions/objectfieldselector-#objectfieldselector" >}}">ObjectFieldSelector</a>)
    Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
  - **env.valueFrom.resourceFieldRef** (<a href="{{< ref "/docs/common-definitions/resourcefieldselector-#resourcefieldselector" >}}">ResourceFieldSelector</a>)
    Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
  - **env.valueFrom.secretKeyRef** (SecretKeySelector)
    Selects a key of a secret in the pod's namespace
*SecretKeySelector selects a key of a Secret.*
  - **env.valueFrom.secretKeyRef.key** (string), required
    The key of the secret to select from.  Must be a valid secret key.
  - **env.valueFrom.secretKeyRef.name** (string)
    Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
  - **env.valueFrom.secretKeyRef.optional** (boolean)
    Specify whether the Secret or its key must be defined
- **envFrom** ([]EnvFromSource)
  EnvFrom defines the collection of EnvFromSource to inject into containers.
*EnvFromSource represents the source of a set of ConfigMaps*
  - **envFrom.configMapRef** (ConfigMapEnvSource)
    The ConfigMap to select from
*ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.*
*The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.*
  - **envFrom.configMapRef.name** (string)
    Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
  - **envFrom.configMapRef.optional** (boolean)
    Specify whether the ConfigMap must be defined
  - **envFrom.prefix** (string)
    An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
  - **envFrom.secretRef** (SecretEnvSource)
    The Secret to select from
*SecretEnvSource selects a Secret to populate the environment variables with.*
*The contents of the target Secret's Data field will represent the key-value pairs as environment variables.*
  - **envFrom.secretRef.name** (string)
    Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
  - **envFrom.secretRef.optional** (boolean)
    Specify whether the Secret must be defined
### PodPresetList
PodPresetList is a list of PodPreset objects.
- **apiVersion**: settings.k8s.io/v1alpha1
  
- **kind**: PodPresetList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>), required
  Items is a list of schema objects.
### Operations
#### `get` read the specified PodPreset

##### HTTP Request
GET /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets/{name}

##### Parameters
  - **{name}** (string), required
    name of the PodPreset
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PodPreset

##### HTTP Request
GET /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets

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
200 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpresetlist" >}}">PodPresetList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind PodPreset

##### HTTP Request
GET /apis/settings.k8s.io/v1alpha1/podpresets

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
200 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpresetlist" >}}">PodPresetList</a>): OK
401: Unauthorized
#### `create` create a PodPreset

##### HTTP Request
POST /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>): OK
201 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>): Created
202 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>): Accepted
401: Unauthorized
#### `update` replace the specified PodPreset

##### HTTP Request
PUT /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets/{name}

##### Parameters
  - **{name}** (string), required
    name of the PodPreset
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>): OK
201 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>): Created
401: Unauthorized
#### `patch` partially update the specified PodPreset

##### HTTP Request
PATCH /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets/{name}

##### Parameters
  - **{name}** (string), required
    name of the PodPreset
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
200 (<a href="{{< ref "/docs/workloads/podpreset-v1alpha1#podpreset" >}}">PodPreset</a>): OK
401: Unauthorized
#### `delete` delete a PodPreset

##### HTTP Request
DELETE /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets/{name}

##### Parameters
  - **{name}** (string), required
    name of the PodPreset
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
#### `deletecollection` delete collection of PodPreset

##### HTTP Request
DELETE /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets

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
