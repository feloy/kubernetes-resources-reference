---
collapsible: false
weight: 10
title: "NodeAffinity"
description: "Node affinity is a group of node affinity scheduling rules."
draft: false
---
## NodeAffinity
`import "k8s.io/api/core/v1"`
### NodeAffinity
Node affinity is a group of node affinity scheduling rules.
- **preferredDuringSchedulingIgnoredDuringExecution** ([]PreferredSchedulingTerm)
  The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
*An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).*
  - **preferredDuringSchedulingIgnoredDuringExecution.preference** (NodeSelectorTerm), required
    A node selector term, associated with the corresponding weight.
*A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.*
  - **preferredDuringSchedulingIgnoredDuringExecution.preference.matchExpressions** ([]<a href="{{< ref "/docs/common-definitions/nodeselectorrequirement-#nodeselectorrequirement" >}}">NodeSelectorRequirement</a>)
    A list of node selector requirements by node's labels.
  - **preferredDuringSchedulingIgnoredDuringExecution.preference.matchFields** ([]<a href="{{< ref "/docs/common-definitions/nodeselectorrequirement-#nodeselectorrequirement" >}}">NodeSelectorRequirement</a>)
    A list of node selector requirements by node's fields.
  - **preferredDuringSchedulingIgnoredDuringExecution.weight** (int32), required
    Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
- **requiredDuringSchedulingIgnoredDuringExecution** (NodeSelector)
  If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.
*A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.*
  - **requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms** ([]NodeSelectorTerm), required
    Required. A list of node selector terms. The terms are ORed.
*A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.*
  - **requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchExpressions** ([]<a href="{{< ref "/docs/common-definitions/nodeselectorrequirement-#nodeselectorrequirement" >}}">NodeSelectorRequirement</a>)
    A list of node selector requirements by node's labels.
  - **requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchFields** ([]<a href="{{< ref "/docs/common-definitions/nodeselectorrequirement-#nodeselectorrequirement" >}}">NodeSelectorRequirement</a>)
    A list of node selector requirements by node's fields.
