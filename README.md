# Kubernetes Resources Reference

This tool creates a DocBook documentation from a Kubernetes API definition (as a swagger file).

The documentation includes the definition of the current resources of the API.

All efforts are done to include as much as possible the definitions of the fields of a resource in a recursive way.

Get the companion documentation about kubectl: https://github.com/feloy/kubectl-reference

## PDF Output

```
# Create the Docbook file
$ make docbook

# Create a PDF file, in USletter format
$ make pdf

# Create a PDF file, in A4 format
$ make pdf FORMAT=A4
```

Get a printed copy at https://kdp.amazon.com/amazon-dp-action/us/dualbookshelf.marketplacelink/B086G11WY1

## Website Output

You will need to install [hugo](https://gohugo.io/) and the [go language tools](https://golang.org/).

```
make website
```

Visit the website at https://www.k8sref.io
