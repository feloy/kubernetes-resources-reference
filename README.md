# Kubernetes Resources Reference

This tool creates a DocBook documentation from a Kubernetes API definition (as a swagger file).

The documentation includes the definition of the current resources of the API.

All efforts are done to include as much as possible the definitions of the fields of a resource in a recursive way.

```
# Create the Docbook file
$ make docbook

# Create a PDF file, in USletter format
$ make pdf

# Create a PDF file, in A4 format
$ make pdf FORMAT=A4
```

Get the companion documentation about kubectl: https://github.com/feloy/kubectl-reference

Get a printed copy at https://www.amazon.com/Kubernetes-Resources-Reference-OpenAPI-specification/dp/1709123451
