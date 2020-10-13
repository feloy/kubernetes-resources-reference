---
title: "CronJob v2alpha1"
description: "CronJob represents the configuration of a single cron job."
draft: false
collapsible: false
weight: 13
---
## CronJob
`apiVersion: batch/v2alpha1`
`import "k8s.io/api/batch/v2alpha1"`
### CronJob
CronJob represents the configuration of a single cron job.
- **apiVersion**: batch/v2alpha1
  
- **kind**: CronJob
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **spec** (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjobspec" >}}">CronJobSpec</a>)
  Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **status** (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjobstatus" >}}">CronJobStatus</a>)
  Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
### CronJobSpec
CronJobSpec describes how the job execution will look like and when it will actually run.
- **jobTemplate** (JobTemplateSpec), required
  Specifies the job that will be created when executing a CronJob.
*JobTemplateSpec describes the data a Job should have when created from a template*
  - **jobTemplate.metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
    Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
  - **jobTemplate.spec** (<a href="{{< ref "/docs/workloads/job-v1#jobspec" >}}">JobSpec</a>)
    Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
- **schedule** (string), required
  The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
- **concurrencyPolicy** (string)
  Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
- **failedJobsHistoryLimit** (int32)
  The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
- **startingDeadlineSeconds** (int64)
  Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
- **successfulJobsHistoryLimit** (int32)
  The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
- **suspend** (boolean)
  This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
### CronJobStatus
CronJobStatus represents the current state of a cron job.
- **active** ([]<a href="{{< ref "/docs/common-definitions/objectreference-#objectreference" >}}">ObjectReference</a>)
  A list of pointers to currently running jobs.
- **lastScheduleTime** (Time)
  Information when was the last time the job was successfully scheduled.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
### CronJobList
CronJobList is a collection of cron jobs.
- **apiVersion**: batch/v2alpha1
  
- **kind**: CronJobList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>), required
  items is the list of CronJobs.
### Operations
#### `get` read the specified CronJob

##### HTTP Request
GET /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}

##### Parameters
  - **{name}** (string), required
    name of the CronJob
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): OK
401: Unauthorized
#### `get` read status of the specified CronJob

##### HTTP Request
GET /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the CronJob
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind CronJob

##### HTTP Request
GET /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs

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
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjoblist" >}}">CronJobList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind CronJob

##### HTTP Request
GET /apis/batch/v2alpha1/cronjobs

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
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjoblist" >}}">CronJobList</a>): OK
401: Unauthorized
#### `create` create a CronJob

##### HTTP Request
POST /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): OK
201 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): Created
202 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): Accepted
401: Unauthorized
#### `update` replace the specified CronJob

##### HTTP Request
PUT /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}

##### Parameters
  - **{name}** (string), required
    name of the CronJob
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): OK
201 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): Created
401: Unauthorized
#### `update` replace status of the specified CronJob

##### HTTP Request
PUT /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the CronJob
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): OK
201 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): Created
401: Unauthorized
#### `patch` partially update the specified CronJob

##### HTTP Request
PATCH /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}

##### Parameters
  - **{name}** (string), required
    name of the CronJob
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
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): OK
401: Unauthorized
#### `patch` partially update status of the specified CronJob

##### HTTP Request
PATCH /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status

##### Parameters
  - **{name}** (string), required
    name of the CronJob
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
200 (<a href="{{< ref "/docs/workloads/cronjob-v2alpha1#cronjob" >}}">CronJob</a>): OK
401: Unauthorized
#### `delete` delete a CronJob

##### HTTP Request
DELETE /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}

##### Parameters
  - **{name}** (string), required
    name of the CronJob
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
#### `deletecollection` delete collection of CronJob

##### HTTP Request
DELETE /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs

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
