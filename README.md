# Kubernetes Resources Reference

Tool to create documentation of the Kubernetes API.

Get a printed book at:

- US: https://www.amazon.com/dp/B08GV3ZLNS
- UK: https://www.amazon.co.uk/dp/B08GV3ZLNS
- DE: https://www.amazon.de/dp/B08GV3ZLNS
- FR: https://www.amazon.fr/dp/B08GV3ZLNS
- ES: https://www.amazon.es/dp/B08GV3ZLNS
- IT: https://www.amazon.it/dp/B08GV3ZLNS
- JP: https://www.amazon.co.jp/dp/B08GV3ZLNS
- CA: https://www.amazon.ca/dp/B08GV3ZLNS

## Usage

### Web output

```sh
make website
make serve
```

### PDF output

```sh
make pdf FORMAT=A4
make pdf FORMAT=USletter
```

## OpenAPI specification

The source of truth for the Kubernetes API is an OpenAPI specification. A standard OpenAPI specification describes:

- a list of *Definitions*,
- a list of *Paths*, each describing a list of *Operations*.

## Kubernetes extensions

https://github.com/kubernetes/kubernetes/tree/master/api/openapi-spec

Kubernetes API extends OpenAPI using these extensions:

- `x-kubernetes-group-version-kind`:
  - Definitions associated with a Kubernetes *Resource* use this extension to declare the GVK to which the resource belongs.
  - Operations use this extension to declare on which Kubernetes resource they operate.
- `x-kubernetes-action`: OpenAPI Operations (get, post, etc) are mapped to Kubernetes API *actions* (get, list, watch, etc) with this extension.
- `x-kubernetes-patch-strategy`: a comma-separated list of strategic merge patch strategies supported by a field of a Kubernetes resource.
- `x-kubernetes-patch-merge-key`: when a field supports the `merge` strategy, this extension indicates the key used to identify the fields to merge.
- `x-kubernetes-list-type`: atomic, map, set. Applicable to lists. atomic and set apply to lists with scalar elements only. map applies to lists of nested types only. If configured as atomic, the entire list is replaced during merge; a single manager manages the list as a whole at any one time. If granular, different managers can manage entries separately.
- `x-kubernetes-list-map-keys`: Only applicable when x-kubernetes-list-type=map. A slice of strings whose values in combination must uniquely identify list entries.
- `x-kubernetes-unions`
