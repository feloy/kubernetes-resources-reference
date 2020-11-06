# Kubernetes Resources Reference

Tool to create documentation of the Kubernetes API.

Visit the website at https://k8sref.io

Get a printed book at:

- US: https://www.amazon.com/dp/B086G11WY1
- UK: https://www.amazon.co.uk/dp/B086G11WY1
- DE: https://www.amazon.de/dp/B086G11WY1
- FR: https://www.amazon.fr/dp/B086G11WY1
- ES: https://www.amazon.es/dp/B086G11WY1
- IT: https://www.amazon.it/dp/B086G11WY1
- JP: https://www.amazon.co.jp/dp/B086G11WY1
- CA: https://www.amazon.ca/dp/B086G11WY1

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

## Changes

### v1.20

- `PodPreset`: removed
- `ServiceSpec`: `ipFamily` replaced by `ipFamilies` and `ipFamilyPolicy`
- `ServiceSpec`: `clusterIPs` added
- `HorizontalPodAutoscalerSpec`: `metrics.containerResource` added
- `HorizontalPodAutoscalerStatus`: `currentMetrics.containerResource` added
- `EndpointSlice`: `endpoints.conditions.serving`, `endpoints.conditions.terminating` added

## New version

To create documentation for a new version, for example v1.20:

```
$ mkdir api/v1.20
$ wget https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/openapi-spec/swagger.json -O api/v1.20/swagger.json
$ mkdir config/v1.20
$ cp config/v1.19/* config/v1.20/
```

Update Makefile to replace v1.19 references with v1.20.

Examine the differences between swagger definitions, and make changes on `config/v1.20/{toc,fields}.yaml` accordingly:

```
$ diff -u api/v1.19/swagger.json api/v1.20/swagger.json
```