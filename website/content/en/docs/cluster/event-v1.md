---
title: "Event"
description: "Event is a report of an event somewhere in the cluster."
draft: false
collapsible: false
weight: 3
---
## Event
`apiVersion: events.k8s.io/v1`
`import "k8s.io/api/events/v1"`
### Event
Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
- **apiVersion**: events.k8s.io/v1
  
- **kind**: Event
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/objectmeta-#objectmeta" >}}">ObjectMeta</a>)
  
- **eventTime** (MicroTime), required
  eventTime is the time when this Event was first observed. It is required.
*MicroTime is version of Time with microsecond level precision.*
- **action** (string)
  action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
- **deprecatedCount** (int32)
  deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
- **deprecatedFirstTimestamp** (Time)
  deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
- **deprecatedLastTimestamp** (Time)
  deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
*Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.*
- **deprecatedSource** (EventSource)
  deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
*EventSource contains information for an event.*
  - **deprecatedSource.component** (string)
    Component from which the event is generated.
  - **deprecatedSource.host** (string)
    Node name on which the event is generated.
- **note** (string)
  note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
- **reason** (string)
  reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
- **regarding** (<a href="{{< ref "/docs/common-definitions/objectreference-#objectreference" >}}">ObjectReference</a>)
  regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
- **related** (<a href="{{< ref "/docs/common-definitions/objectreference-#objectreference" >}}">ObjectReference</a>)
  related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
- **reportingController** (string)
  reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
- **reportingInstance** (string)
  reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
- **series** (EventSeries)
  series is data about the Event series this event represents or nil if it's a singleton Event.
*EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in "k8s.io/client-go/tools/events/event_broadcaster.go" shows how this struct is updated on heartbeats and can guide customized reporter implementations.*
  - **series.count** (int32), required
    count is the number of occurrences in this series up to the last heartbeat time.
  - **series.lastObservedTime** (MicroTime), required
    lastObservedTime is the time when last Event from the series was seen before last heartbeat.
*MicroTime is version of Time with microsecond level precision.*
- **type** (string)
  type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
### EventList
EventList is a list of Event objects.
- **apiVersion**: events.k8s.io/v1
  
- **kind**: EventList
  
- **metadata** (<a href="{{< ref "/docs/common-definitions/listmeta-#listmeta" >}}">ListMeta</a>)
  Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
- **items** ([]<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>), required
  items is a list of schema objects.
### Operations
#### `get` read the specified Event

##### HTTP Request
GET /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}

##### Parameters
  - **{name}** (string), required
    name of the Event
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind Event

##### HTTP Request
GET /apis/events.k8s.io/v1/namespaces/{namespace}/events

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
200 (<a href="{{< ref "/docs/cluster/event-v1#eventlist" >}}">EventList</a>): OK
401: Unauthorized
#### `list` list or watch objects of kind Event

##### HTTP Request
GET /apis/events.k8s.io/v1/events

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
200 (<a href="{{< ref "/docs/cluster/event-v1#eventlist" >}}">EventList</a>): OK
401: Unauthorized
#### `create` create an Event

##### HTTP Request
POST /apis/events.k8s.io/v1/namespaces/{namespace}/events

##### Parameters
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>): OK
201 (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>): Created
202 (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>): Accepted
401: Unauthorized
#### `update` replace the specified Event

##### HTTP Request
PUT /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}

##### Parameters
  - **{name}** (string), required
    name of the Event
  - **{namespace}** (string), required
    <a href="{{< ref "/docs/common-parameters/common-parameters-#namespace" >}}">namespace</a>
  - **body** (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>), required
    
  - **?dryRun** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#dryrun" >}}">dryRun</a>
  - **?fieldManager** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#fieldmanager" >}}">fieldManager</a>
  - **?pretty** (string)
    <a href="{{< ref "/docs/common-parameters/common-parameters-#pretty" >}}">pretty</a>

##### Response
200 (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>): OK
201 (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>): Created
401: Unauthorized
#### `patch` partially update the specified Event

##### HTTP Request
PATCH /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}

##### Parameters
  - **{name}** (string), required
    name of the Event
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
200 (<a href="{{< ref "/docs/cluster/event-v1#event" >}}">Event</a>): OK
401: Unauthorized
#### `delete` delete an Event

##### HTTP Request
DELETE /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}

##### Parameters
  - **{name}** (string), required
    name of the Event
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
#### `deletecollection` delete collection of Event

##### HTTP Request
DELETE /apis/events.k8s.io/v1/namespaces/{namespace}/events

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
